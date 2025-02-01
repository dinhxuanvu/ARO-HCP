// Code generated by MockGen. DO NOT EDIT.
// Source: ../database/database.go
//
// Generated by this command:
//
//	mockgen-v0.5.0 -source=../database/database.go -destination=dbclient.go -package mocks github.com/Azure/ARO-HCP/internal/database DBClient
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	database "github.com/Azure/ARO-HCP/internal/database"
	arm "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	gomock "go.uber.org/mock/gomock"
)

// MockDBClientIterator is a mock of DBClientIterator interface.
type MockDBClientIterator[T any] struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientIteratorMockRecorder[T]
	isgomock struct{}
}

// MockDBClientIteratorMockRecorder is the mock recorder for MockDBClientIterator.
type MockDBClientIteratorMockRecorder[T any] struct {
	mock *MockDBClientIterator[T]
}

// NewMockDBClientIterator creates a new mock instance.
func NewMockDBClientIterator[T any](ctrl *gomock.Controller) *MockDBClientIterator[T] {
	mock := &MockDBClientIterator[T]{ctrl: ctrl}
	mock.recorder = &MockDBClientIteratorMockRecorder[T]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClientIterator[T]) EXPECT() *MockDBClientIteratorMockRecorder[T] {
	return m.recorder
}

// GetContinuationToken mocks base method.
func (m *MockDBClientIterator[T]) GetContinuationToken() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContinuationToken")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetContinuationToken indicates an expected call of GetContinuationToken.
func (mr *MockDBClientIteratorMockRecorder[T]) GetContinuationToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContinuationToken", reflect.TypeOf((*MockDBClientIterator[T])(nil).GetContinuationToken))
}

// GetError mocks base method.
func (m *MockDBClientIterator[T]) GetError() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetError")
	ret0, _ := ret[0].(error)
	return ret0
}

// GetError indicates an expected call of GetError.
func (mr *MockDBClientIteratorMockRecorder[T]) GetError() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetError", reflect.TypeOf((*MockDBClientIterator[T])(nil).GetError))
}

// Items mocks base method.
func (m *MockDBClientIterator[T]) Items(ctx context.Context) database.DBClientIteratorItem[T] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Items", ctx)
	ret0, _ := ret[0].(database.DBClientIteratorItem[T])
	return ret0
}

// Items indicates an expected call of Items.
func (mr *MockDBClientIteratorMockRecorder[T]) Items(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Items", reflect.TypeOf((*MockDBClientIterator[T])(nil).Items), ctx)
}

// MockDBClient is a mock of DBClient interface.
type MockDBClient struct {
	ctrl     *gomock.Controller
	recorder *MockDBClientMockRecorder
	isgomock struct{}
}

// MockDBClientMockRecorder is the mock recorder for MockDBClient.
type MockDBClientMockRecorder struct {
	mock *MockDBClient
}

// NewMockDBClient creates a new mock instance.
func NewMockDBClient(ctrl *gomock.Controller) *MockDBClient {
	mock := &MockDBClient{ctrl: ctrl}
	mock.recorder = &MockDBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBClient) EXPECT() *MockDBClientMockRecorder {
	return m.recorder
}

// CreateOperationDoc mocks base method.
func (m *MockDBClient) CreateOperationDoc(ctx context.Context, doc *database.OperationDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOperationDoc", ctx, doc)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOperationDoc indicates an expected call of CreateOperationDoc.
func (mr *MockDBClientMockRecorder) CreateOperationDoc(ctx, doc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOperationDoc", reflect.TypeOf((*MockDBClient)(nil).CreateOperationDoc), ctx, doc)
}

// CreateResourceDoc mocks base method.
func (m *MockDBClient) CreateResourceDoc(ctx context.Context, doc *database.ResourceDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResourceDoc", ctx, doc)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResourceDoc indicates an expected call of CreateResourceDoc.
func (mr *MockDBClientMockRecorder) CreateResourceDoc(ctx, doc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResourceDoc", reflect.TypeOf((*MockDBClient)(nil).CreateResourceDoc), ctx, doc)
}

// CreateSubscriptionDoc mocks base method.
func (m *MockDBClient) CreateSubscriptionDoc(ctx context.Context, doc *database.SubscriptionDocument) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSubscriptionDoc", ctx, doc)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateSubscriptionDoc indicates an expected call of CreateSubscriptionDoc.
func (mr *MockDBClientMockRecorder) CreateSubscriptionDoc(ctx, doc any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSubscriptionDoc", reflect.TypeOf((*MockDBClient)(nil).CreateSubscriptionDoc), ctx, doc)
}

// DBConnectionTest mocks base method.
func (m *MockDBClient) DBConnectionTest(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DBConnectionTest", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DBConnectionTest indicates an expected call of DBConnectionTest.
func (mr *MockDBClientMockRecorder) DBConnectionTest(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DBConnectionTest", reflect.TypeOf((*MockDBClient)(nil).DBConnectionTest), ctx)
}

// DeleteResourceDoc mocks base method.
func (m *MockDBClient) DeleteResourceDoc(ctx context.Context, resourceID *arm.ResourceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResourceDoc", ctx, resourceID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResourceDoc indicates an expected call of DeleteResourceDoc.
func (mr *MockDBClientMockRecorder) DeleteResourceDoc(ctx, resourceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResourceDoc", reflect.TypeOf((*MockDBClient)(nil).DeleteResourceDoc), ctx, resourceID)
}

// GetLockClient mocks base method.
func (m *MockDBClient) GetLockClient() *database.LockClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLockClient")
	ret0, _ := ret[0].(*database.LockClient)
	return ret0
}

// GetLockClient indicates an expected call of GetLockClient.
func (mr *MockDBClientMockRecorder) GetLockClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLockClient", reflect.TypeOf((*MockDBClient)(nil).GetLockClient))
}

// GetOperationDoc mocks base method.
func (m *MockDBClient) GetOperationDoc(ctx context.Context, operationID string) (*database.OperationDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperationDoc", ctx, operationID)
	ret0, _ := ret[0].(*database.OperationDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOperationDoc indicates an expected call of GetOperationDoc.
func (mr *MockDBClientMockRecorder) GetOperationDoc(ctx, operationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperationDoc", reflect.TypeOf((*MockDBClient)(nil).GetOperationDoc), ctx, operationID)
}

// GetResourceDoc mocks base method.
func (m *MockDBClient) GetResourceDoc(ctx context.Context, resourceID *arm.ResourceID) (*database.ResourceDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResourceDoc", ctx, resourceID)
	ret0, _ := ret[0].(*database.ResourceDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResourceDoc indicates an expected call of GetResourceDoc.
func (mr *MockDBClientMockRecorder) GetResourceDoc(ctx, resourceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResourceDoc", reflect.TypeOf((*MockDBClient)(nil).GetResourceDoc), ctx, resourceID)
}

// GetSubscriptionDoc mocks base method.
func (m *MockDBClient) GetSubscriptionDoc(ctx context.Context, subscriptionID string) (*database.SubscriptionDocument, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptionDoc", ctx, subscriptionID)
	ret0, _ := ret[0].(*database.SubscriptionDocument)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptionDoc indicates an expected call of GetSubscriptionDoc.
func (mr *MockDBClientMockRecorder) GetSubscriptionDoc(ctx, subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptionDoc", reflect.TypeOf((*MockDBClient)(nil).GetSubscriptionDoc), ctx, subscriptionID)
}

// ListAllOperationDocs mocks base method.
func (m *MockDBClient) ListAllOperationDocs() database.DBClientIterator[database.OperationDocument] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllOperationDocs")
	ret0, _ := ret[0].(database.DBClientIterator[database.OperationDocument])
	return ret0
}

// ListAllOperationDocs indicates an expected call of ListAllOperationDocs.
func (mr *MockDBClientMockRecorder) ListAllOperationDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllOperationDocs", reflect.TypeOf((*MockDBClient)(nil).ListAllOperationDocs))
}

// ListAllSubscriptionDocs mocks base method.
func (m *MockDBClient) ListAllSubscriptionDocs() database.DBClientIterator[database.SubscriptionDocument] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllSubscriptionDocs")
	ret0, _ := ret[0].(database.DBClientIterator[database.SubscriptionDocument])
	return ret0
}

// ListAllSubscriptionDocs indicates an expected call of ListAllSubscriptionDocs.
func (mr *MockDBClientMockRecorder) ListAllSubscriptionDocs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllSubscriptionDocs", reflect.TypeOf((*MockDBClient)(nil).ListAllSubscriptionDocs))
}

// ListOperationDocs mocks base method.
func (m *MockDBClient) ListOperationDocs(subscriptionID string) database.DBClientIterator[database.OperationDocument] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOperationDocs", subscriptionID)
	ret0, _ := ret[0].(database.DBClientIterator[database.OperationDocument])
	return ret0
}

// ListOperationDocs indicates an expected call of ListOperationDocs.
func (mr *MockDBClientMockRecorder) ListOperationDocs(subscriptionID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOperationDocs", reflect.TypeOf((*MockDBClient)(nil).ListOperationDocs), subscriptionID)
}

// ListResourceDocs mocks base method.
func (m *MockDBClient) ListResourceDocs(prefix *arm.ResourceID, maxItems int32, continuationToken *string) database.DBClientIterator[database.ResourceDocument] {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourceDocs", prefix, maxItems, continuationToken)
	ret0, _ := ret[0].(database.DBClientIterator[database.ResourceDocument])
	return ret0
}

// ListResourceDocs indicates an expected call of ListResourceDocs.
func (mr *MockDBClientMockRecorder) ListResourceDocs(prefix, maxItems, continuationToken any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourceDocs", reflect.TypeOf((*MockDBClient)(nil).ListResourceDocs), prefix, maxItems, continuationToken)
}

// UpdateOperationDoc mocks base method.
func (m *MockDBClient) UpdateOperationDoc(ctx context.Context, operationID string, callback func(*database.OperationDocument) bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOperationDoc", ctx, operationID, callback)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOperationDoc indicates an expected call of UpdateOperationDoc.
func (mr *MockDBClientMockRecorder) UpdateOperationDoc(ctx, operationID, callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOperationDoc", reflect.TypeOf((*MockDBClient)(nil).UpdateOperationDoc), ctx, operationID, callback)
}

// UpdateResourceDoc mocks base method.
func (m *MockDBClient) UpdateResourceDoc(ctx context.Context, resourceID *arm.ResourceID, callback func(*database.ResourceDocument) bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResourceDoc", ctx, resourceID, callback)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResourceDoc indicates an expected call of UpdateResourceDoc.
func (mr *MockDBClientMockRecorder) UpdateResourceDoc(ctx, resourceID, callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResourceDoc", reflect.TypeOf((*MockDBClient)(nil).UpdateResourceDoc), ctx, resourceID, callback)
}

// UpdateSubscriptionDoc mocks base method.
func (m *MockDBClient) UpdateSubscriptionDoc(ctx context.Context, subscriptionID string, callback func(*database.SubscriptionDocument) bool) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSubscriptionDoc", ctx, subscriptionID, callback)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSubscriptionDoc indicates an expected call of UpdateSubscriptionDoc.
func (mr *MockDBClientMockRecorder) UpdateSubscriptionDoc(ctx, subscriptionID, callback any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSubscriptionDoc", reflect.TypeOf((*MockDBClient)(nil).UpdateSubscriptionDoc), ctx, subscriptionID, callback)
}
