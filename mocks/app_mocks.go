// Code generated by MockGen. DO NOT EDIT.
// Source: app.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	aquareo "github.com/pedrobfernandes/aquareo/internal/aquareo"
)

// MockWebServer is a mock of WebServer interface.
type MockWebServer struct {
	ctrl     *gomock.Controller
	recorder *MockWebServerMockRecorder
}

// MockWebServerMockRecorder is the mock recorder for MockWebServer.
type MockWebServerMockRecorder struct {
	mock *MockWebServer
}

// NewMockWebServer creates a new mock instance.
func NewMockWebServer(ctrl *gomock.Controller) *MockWebServer {
	mock := &MockWebServer{ctrl: ctrl}
	mock.recorder = &MockWebServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebServer) EXPECT() *MockWebServerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockWebServer) Start(addr string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start", addr)
}

// Start indicates an expected call of Start.
func (mr *MockWebServerMockRecorder) Start(addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockWebServer)(nil).Start), addr)
}

// Stop mocks base method.
func (m *MockWebServer) Stop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", ctx)
}

// Stop indicates an expected call of Stop.
func (mr *MockWebServerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWebServer)(nil).Stop), ctx)
}

// MockController is a mock of Controller interface.
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// Config mocks base method.
func (m *MockController) Config() aquareo.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(aquareo.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockControllerMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockController)(nil).Config))
}

// Install mocks base method.
func (m *MockController) Install(s aquareo.Subsystem) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", s)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockControllerMockRecorder) Install(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockController)(nil).Install), s)
}

// Start mocks base method.
func (m *MockController) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockControllerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockController)(nil).Start))
}

// Stop mocks base method.
func (m *MockController) Stop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", ctx)
}

// Stop indicates an expected call of Stop.
func (mr *MockControllerMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockController)(nil).Stop), ctx)
}

// Storage mocks base method.
func (m *MockController) Storage() aquareo.Storage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Storage")
	ret0, _ := ret[0].(aquareo.Storage)
	return ret0
}

// Storage indicates an expected call of Storage.
func (mr *MockControllerMockRecorder) Storage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Storage", reflect.TypeOf((*MockController)(nil).Storage))
}

// MockGPIODriver is a mock of GPIODriver interface.
type MockGPIODriver struct {
	ctrl     *gomock.Controller
	recorder *MockGPIODriverMockRecorder
}

// MockGPIODriverMockRecorder is the mock recorder for MockGPIODriver.
type MockGPIODriverMockRecorder struct {
	mock *MockGPIODriver
}

// NewMockGPIODriver creates a new mock instance.
func NewMockGPIODriver(ctrl *gomock.Controller) *MockGPIODriver {
	mock := &MockGPIODriver{ctrl: ctrl}
	mock.recorder = &MockGPIODriverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGPIODriver) EXPECT() *MockGPIODriverMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockGPIODriver) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockGPIODriverMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockGPIODriver)(nil).Close))
}

// Open mocks base method.
func (m *MockGPIODriver) Open() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockGPIODriverMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockGPIODriver)(nil).Open))
}

// MockSubsystem is a mock of Subsystem interface.
type MockSubsystem struct {
	ctrl     *gomock.Controller
	recorder *MockSubsystemMockRecorder
}

// MockSubsystemMockRecorder is the mock recorder for MockSubsystem.
type MockSubsystemMockRecorder struct {
	mock *MockSubsystem
}

// NewMockSubsystem creates a new mock instance.
func NewMockSubsystem(ctrl *gomock.Controller) *MockSubsystem {
	mock := &MockSubsystem{ctrl: ctrl}
	mock.recorder = &MockSubsystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubsystem) EXPECT() *MockSubsystemMockRecorder {
	return m.recorder
}

// Install mocks base method.
func (m *MockSubsystem) Install(ctrl aquareo.Controller) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Install", ctrl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Install indicates an expected call of Install.
func (mr *MockSubsystemMockRecorder) Install(ctrl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Install", reflect.TypeOf((*MockSubsystem)(nil).Install), ctrl)
}

// Start mocks base method.
func (m *MockSubsystem) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockSubsystemMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockSubsystem)(nil).Start))
}

// Stop mocks base method.
func (m *MockSubsystem) Stop(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop", ctx)
}

// Stop indicates an expected call of Stop.
func (mr *MockSubsystemMockRecorder) Stop(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockSubsystem)(nil).Stop), ctx)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// CreateBucket mocks base method.
func (m *MockStorage) CreateBucket(bucket string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBucket", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBucket indicates an expected call of CreateBucket.
func (mr *MockStorageMockRecorder) CreateBucket(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBucket", reflect.TypeOf((*MockStorage)(nil).CreateBucket), bucket)
}

// MetricStore mocks base method.
func (m *MockStorage) MetricStore(bucket string) aquareo.MetricStore {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MetricStore", bucket)
	ret0, _ := ret[0].(aquareo.MetricStore)
	return ret0
}

// MetricStore indicates an expected call of MetricStore.
func (mr *MockStorageMockRecorder) MetricStore(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MetricStore", reflect.TypeOf((*MockStorage)(nil).MetricStore), bucket)
}

// MockMetricStore is a mock of MetricStore interface.
type MockMetricStore struct {
	ctrl     *gomock.Controller
	recorder *MockMetricStoreMockRecorder
}

// MockMetricStoreMockRecorder is the mock recorder for MockMetricStore.
type MockMetricStoreMockRecorder struct {
	mock *MockMetricStore
}

// NewMockMetricStore creates a new mock instance.
func NewMockMetricStore(ctrl *gomock.Controller) *MockMetricStore {
	mock := &MockMetricStore{ctrl: ctrl}
	mock.recorder = &MockMetricStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetricStore) EXPECT() *MockMetricStoreMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockMetricStore) List(size int) ([]aquareo.MetricEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", size)
	ret0, _ := ret[0].([]aquareo.MetricEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMetricStoreMockRecorder) List(size interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMetricStore)(nil).List), size)
}

// Put mocks base method.
func (m *MockMetricStore) Put(timespan int64, value float32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", timespan, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *MockMetricStoreMockRecorder) Put(timespan, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockMetricStore)(nil).Put), timespan, value)
}
