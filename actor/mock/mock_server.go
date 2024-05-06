// Code generated by MockGen. DO NOT EDIT.
// Source: ./actor/actor.go
//
// Generated by this command:
//
//	mockgen -source ./actor/actor.go -destination ./actor/mock/mock_server.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	actor "github.com/dapr/go-sdk/actor"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockClient) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockClientMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockClient)(nil).ID))
}

// Type mocks base method.
func (m *MockClient) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockClientMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockClient)(nil).Type))
}

// MockServer is a mock of Server interface.
type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}

// MockServerMockRecorder is the mock recorder for MockServer.
type MockServerMockRecorder struct {
	mock *MockServer
}

// NewMockServer creates a new mock instance.
func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockServer) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockServerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockServer)(nil).ID))
}

// SaveState mocks base method.
func (m *MockServer) SaveState() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveState")
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveState indicates an expected call of SaveState.
func (mr *MockServerMockRecorder) SaveState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveState", reflect.TypeOf((*MockServer)(nil).SaveState))
}

// SetID mocks base method.
func (m *MockServer) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockServerMockRecorder) SetID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockServer)(nil).SetID), arg0)
}

// SetStateManager mocks base method.
func (m *MockServer) SetStateManager(arg0 actor.StateManager) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStateManager", arg0)
}

// SetStateManager indicates an expected call of SetStateManager.
func (mr *MockServerMockRecorder) SetStateManager(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateManager", reflect.TypeOf((*MockServer)(nil).SetStateManager), arg0)
}

// Type mocks base method.
func (m *MockServer) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockServerMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockServer)(nil).Type))
}

// WithContext mocks base method.
func (m *MockServer) WithContext() actor.ServerContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext")
	ret0, _ := ret[0].(actor.ServerContext)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockServerMockRecorder) WithContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockServer)(nil).WithContext))
}

// MockServerContext is a mock of ServerContext interface.
type MockServerContext struct {
	ctrl     *gomock.Controller
	recorder *MockServerContextMockRecorder
}

// MockServerContextMockRecorder is the mock recorder for MockServerContext.
type MockServerContextMockRecorder struct {
	mock *MockServerContext
}

// NewMockServerContext creates a new mock instance.
func NewMockServerContext(ctrl *gomock.Controller) *MockServerContext {
	mock := &MockServerContext{ctrl: ctrl}
	mock.recorder = &MockServerContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServerContext) EXPECT() *MockServerContextMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockServerContext) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockServerContextMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockServerContext)(nil).ID))
}

// SaveState mocks base method.
func (m *MockServerContext) SaveState(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveState indicates an expected call of SaveState.
func (mr *MockServerContextMockRecorder) SaveState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveState", reflect.TypeOf((*MockServerContext)(nil).SaveState), arg0)
}

// SetID mocks base method.
func (m *MockServerContext) SetID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetID", arg0)
}

// SetID indicates an expected call of SetID.
func (mr *MockServerContextMockRecorder) SetID(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetID", reflect.TypeOf((*MockServerContext)(nil).SetID), arg0)
}

// SetStateManager mocks base method.
func (m *MockServerContext) SetStateManager(arg0 actor.StateManagerContext) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetStateManager", arg0)
}

// SetStateManager indicates an expected call of SetStateManager.
func (mr *MockServerContextMockRecorder) SetStateManager(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStateManager", reflect.TypeOf((*MockServerContext)(nil).SetStateManager), arg0)
}

// Type mocks base method.
func (m *MockServerContext) Type() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(string)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockServerContextMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockServerContext)(nil).Type))
}

func (mr *MockServerContextMockRecorder) Invoke(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Invoke", reflect.TypeOf((*MockServerContext)(nil).Invoke), arg0, arg1)
}

func (m *MockServerContext) Invoke(ctx context.Context, input string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Invoke", ctx, input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MockReminderCallee is a mock of ReminderCallee interface.
type MockReminderCallee struct {
	ctrl     *gomock.Controller
	recorder *MockReminderCalleeMockRecorder
}

// MockReminderCalleeMockRecorder is the mock recorder for MockReminderCallee.
type MockReminderCalleeMockRecorder struct {
	mock *MockReminderCallee
}

// NewMockReminderCallee creates a new mock instance.
func NewMockReminderCallee(ctrl *gomock.Controller) *MockReminderCallee {
	mock := &MockReminderCallee{ctrl: ctrl}
	mock.recorder = &MockReminderCalleeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReminderCallee) EXPECT() *MockReminderCalleeMockRecorder {
	return m.recorder
}

// ReminderCall mocks base method.
func (m *MockReminderCallee) ReminderCall(arg0 string, arg1 []byte, arg2, arg3 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReminderCall", arg0, arg1, arg2, arg3)
}

// ReminderCall indicates an expected call of ReminderCall.
func (mr *MockReminderCalleeMockRecorder) ReminderCall(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReminderCall", reflect.TypeOf((*MockReminderCallee)(nil).ReminderCall), arg0, arg1, arg2, arg3)
}

// MockStateManager is a mock of StateManager interface.
type MockStateManager struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerMockRecorder
}

// MockStateManagerMockRecorder is the mock recorder for MockStateManager.
type MockStateManagerMockRecorder struct {
	mock *MockStateManager
}

// NewMockStateManager creates a new mock instance.
func NewMockStateManager(ctrl *gomock.Controller) *MockStateManager {
	mock := &MockStateManager{ctrl: ctrl}
	mock.recorder = &MockStateManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateManager) EXPECT() *MockStateManagerMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStateManager) Add(stateName string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", stateName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStateManagerMockRecorder) Add(stateName, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStateManager)(nil).Add), stateName, value)
}

// Contains mocks base method.
func (m *MockStateManager) Contains(stateName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", stateName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains.
func (mr *MockStateManagerMockRecorder) Contains(stateName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockStateManager)(nil).Contains), stateName)
}

// Flush mocks base method.
func (m *MockStateManager) Flush() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush")
}

// Flush indicates an expected call of Flush.
func (mr *MockStateManagerMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockStateManager)(nil).Flush))
}

// Get mocks base method.
func (m *MockStateManager) Get(stateName string, reply any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", stateName, reply)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockStateManagerMockRecorder) Get(stateName, reply any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStateManager)(nil).Get), stateName, reply)
}

// Remove mocks base method.
func (m *MockStateManager) Remove(stateName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", stateName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStateManagerMockRecorder) Remove(stateName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStateManager)(nil).Remove), stateName)
}

// Save mocks base method.
func (m *MockStateManager) Save() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save")
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStateManagerMockRecorder) Save() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStateManager)(nil).Save))
}

// Set mocks base method.
func (m *MockStateManager) Set(stateName string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", stateName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStateManagerMockRecorder) Set(stateName, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStateManager)(nil).Set), stateName, value)
}

// WithContext mocks base method.
func (m *MockStateManager) WithContext() actor.StateManagerContext {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContext")
	ret0, _ := ret[0].(actor.StateManagerContext)
	return ret0
}

// WithContext indicates an expected call of WithContext.
func (mr *MockStateManagerMockRecorder) WithContext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContext", reflect.TypeOf((*MockStateManager)(nil).WithContext))
}

// MockStateManagerContext is a mock of StateManagerContext interface.
type MockStateManagerContext struct {
	ctrl     *gomock.Controller
	recorder *MockStateManagerContextMockRecorder
}

// MockStateManagerContextMockRecorder is the mock recorder for MockStateManagerContext.
type MockStateManagerContextMockRecorder struct {
	mock *MockStateManagerContext
}

// NewMockStateManagerContext creates a new mock instance.
func NewMockStateManagerContext(ctrl *gomock.Controller) *MockStateManagerContext {
	mock := &MockStateManagerContext{ctrl: ctrl}
	mock.recorder = &MockStateManagerContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateManagerContext) EXPECT() *MockStateManagerContextMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStateManagerContext) Add(ctx context.Context, stateName string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, stateName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStateManagerContextMockRecorder) Add(ctx, stateName, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStateManagerContext)(nil).Add), ctx, stateName, value)
}

// Contains mocks base method.
func (m *MockStateManagerContext) Contains(ctx context.Context, stateName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", ctx, stateName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains.
func (mr *MockStateManagerContextMockRecorder) Contains(ctx, stateName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockStateManagerContext)(nil).Contains), ctx, stateName)
}

// Flush mocks base method.
func (m *MockStateManagerContext) Flush(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Flush", ctx)
}

// Flush indicates an expected call of Flush.
func (mr *MockStateManagerContextMockRecorder) Flush(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockStateManagerContext)(nil).Flush), ctx)
}

// Get mocks base method.
func (m *MockStateManagerContext) Get(ctx context.Context, stateName string, reply any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, stateName, reply)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockStateManagerContextMockRecorder) Get(ctx, stateName, reply any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStateManagerContext)(nil).Get), ctx, stateName, reply)
}

// Remove mocks base method.
func (m *MockStateManagerContext) Remove(ctx context.Context, stateName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, stateName)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockStateManagerContextMockRecorder) Remove(ctx, stateName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockStateManagerContext)(nil).Remove), ctx, stateName)
}

// Save mocks base method.
func (m *MockStateManagerContext) Save(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockStateManagerContextMockRecorder) Save(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockStateManagerContext)(nil).Save), ctx)
}

// Set mocks base method.
func (m *MockStateManagerContext) Set(ctx context.Context, stateName string, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", ctx, stateName, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockStateManagerContextMockRecorder) Set(ctx, stateName, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockStateManagerContext)(nil).Set), ctx, stateName, value)
}

// SetWithTTL mocks base method.
func (m *MockStateManagerContext) SetWithTTL(ctx context.Context, stateName string, value any, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWithTTL", ctx, stateName, value, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWithTTL indicates an expected call of SetWithTTL.
func (mr *MockStateManagerContextMockRecorder) SetWithTTL(ctx, stateName, value, ttl any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWithTTL", reflect.TypeOf((*MockStateManagerContext)(nil).SetWithTTL), ctx, stateName, value, ttl)
}
