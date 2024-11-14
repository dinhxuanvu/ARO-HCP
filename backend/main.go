package main

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/go-logr/logr"
	ocmsdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/spf13/cobra"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
	"k8s.io/klog/v2"

	"github.com/Azure/ARO-HCP/internal/database"
)

const (
	leaderElectionLockName      = "backend-leader"
	leaderElectionLeaseDuration = 15 * time.Second
	leaderElectionRenewDeadline = 10 * time.Second
	leaderElectionRetryPeriod   = 2 * time.Second
)

var (
	argKubeconfig         string
	argNamespace          string
	argLocation           string
	argCosmosName         string
	argCosmosURL          string
	argClustersServiceURL string
	argInsecure           bool

	processName = filepath.Base(os.Args[0])

	rootCmd = &cobra.Command{
		Use:   processName,
		Args:  cobra.NoArgs,
		Short: "ARO HCP Backend",
		Long: fmt.Sprintf(`ARO HCP Backend

	The command runs the ARO HCP Backend. It executes background processing that
	communicates with Clusters Service and CosmosDB.

	# Run ARO HCP Backend locally to connect to a local Clusters Service at http://localhost:8000
	%s --cosmos-name ${DB_NAME} --cosmos-url ${DB_URL} --location ${LOCATION} \
		--clusters-service-url "http://localhost:8000"
`, processName),
		Version:       "unknown", // overridden by build info below
		RunE:          Run,
		SilenceErrors: true, // errors are printed after Execute
	}
)

func init() {
	rootCmd.SetErrPrefix(rootCmd.Short + " error:")

	rootCmd.Flags().StringVar(&argKubeconfig, "kubeconfig", "", "Absolute path to the kubeconfig file")
	rootCmd.Flags().StringVar(&argNamespace, "namespace", os.Getenv("NAMESPACE"), "Kubernetes namespace")
	rootCmd.Flags().StringVar(&argLocation, "location", os.Getenv("LOCATION"), "Azure location")
	rootCmd.Flags().StringVar(&argCosmosName, "cosmos-name", os.Getenv("DB_NAME"), "Cosmos database name")
	rootCmd.Flags().StringVar(&argCosmosURL, "cosmos-url", os.Getenv("DB_URL"), "Cosmos database URL")
	rootCmd.Flags().StringVar(&argClustersServiceURL, "clusters-service-url", "https://api.openshift.com", "URL of the OCM API gateway")
	rootCmd.Flags().BoolVar(&argInsecure, "insecure", false, "Skip validating TLS for clusters-service")

	rootCmd.MarkFlagsRequiredTogether("cosmos-name", "cosmos-url")

	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				rootCmd.Version = setting.Value
				break
			}
		}
	}
}

func newKubeconfig(kubeconfig string) (*rest.Config, error) {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	if kubeconfig != "" {
		loader.ExplicitPath = kubeconfig
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loader, nil).ClientConfig()
}

func Run(cmd *cobra.Command, args []string) error {
	handler := slog.NewJSONHandler(os.Stdout, nil)
	logger := slog.New(handler)
	klog.SetLogger(logr.FromSlogHandler(handler))

	// Use pod name as the lock identity.
	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	kubeconfig, err := newKubeconfig(argKubeconfig)
	if err != nil {
		return fmt.Errorf("Failed to create Kubernetes configuration: %w", err)
	}

	leaderElectionLock, err := resourcelock.NewFromKubeconfig(
		resourcelock.LeasesResourceLock,
		argNamespace,
		leaderElectionLockName,
		resourcelock.ResourceLockConfig{
			Identity: hostname,
		},
		kubeconfig,
		leaderElectionRenewDeadline)
	if err != nil {
		return fmt.Errorf("Failed to create leader election lock: %w", err)
	}

	// Create the database client.
	cosmosDatabaseClient, err := database.NewCosmosDatabaseClient(argCosmosURL, argCosmosName)
	if err != nil {
		return fmt.Errorf("failed to create the CosmosDB client: %w", err)
	}

	dbClient, err := database.NewDBClient(context.Background(), cosmosDatabaseClient)
	if err != nil {
		return fmt.Errorf("failed to create the database client: %w", err)
	}

	// Create OCM connection
	ocmConnection, err := ocmsdk.NewUnauthenticatedConnectionBuilder().
		URL(argClustersServiceURL).
		Insecure(argInsecure).
		Build()
	if err != nil {
		return fmt.Errorf("Failed to create OCM connection: %w", err)
	}

	logger.Info(fmt.Sprintf("%s (%s) started", cmd.Short, cmd.Version))

	operationsScanner := NewOperationsScanner(dbClient, ocmConnection)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		logger.Info("Caught interrupt signal")
	}()

	var startedLeading atomic.Bool

	// FIXME Integrate leaderelection.HealthzAdaptor into a /healthz endpoint.
	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          leaderElectionLock,
		LeaseDuration: leaderElectionLeaseDuration,
		RenewDeadline: leaderElectionRenewDeadline,
		RetryPeriod:   leaderElectionRetryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				startedLeading.Store(true)
				go operationsScanner.Run(ctx, logger)
			},
			OnStoppedLeading: func() {
				if startedLeading.Load() {
					operationsScanner.Join()
				}
			},
		},
		ReleaseOnCancel: true,
		Name:            leaderElectionLockName,
	})

	logger.Info(fmt.Sprintf("%s (%s) stopped", cmd.Short, cmd.Version))

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		rootCmd.PrintErrln(rootCmd.ErrPrefix(), err.Error())
		os.Exit(1)
	}
}
