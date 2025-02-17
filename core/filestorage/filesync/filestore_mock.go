// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore (interfaces: FileStore)
//
// Generated by this command:
//
//	mockgen -package filesync -destination filestore_mock.go github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore FileStore
//
// Package filesync is a generated GoMock package.
package filesync

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	localstore "github.com/anyproto/anytype-heart/pkg/lib/localstore"
	filestore "github.com/anyproto/anytype-heart/pkg/lib/localstore/filestore"
	model "github.com/anyproto/anytype-heart/pkg/lib/pb/model"
	storage "github.com/anyproto/anytype-heart/pkg/lib/pb/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockFileStore is a mock of FileStore interface.
type MockFileStore struct {
	ctrl     *gomock.Controller
	recorder *MockFileStoreMockRecorder
}

// MockFileStoreMockRecorder is the mock recorder for MockFileStore.
type MockFileStoreMockRecorder struct {
	mock *MockFileStore
}

// NewMockFileStore creates a new mock instance.
func NewMockFileStore(ctrl *gomock.Controller) *MockFileStore {
	mock := &MockFileStore{ctrl: ctrl}
	mock.recorder = &MockFileStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileStore) EXPECT() *MockFileStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockFileStore) Add(arg0 *storage.FileInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockFileStoreMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFileStore)(nil).Add), arg0)
}

// AddFileKeys mocks base method.
func (m *MockFileStore) AddFileKeys(arg0 ...filestore.FileKeys) error {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddFileKeys", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFileKeys indicates an expected call of AddFileKeys.
func (mr *MockFileStoreMockRecorder) AddFileKeys(arg0 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFileKeys", reflect.TypeOf((*MockFileStore)(nil).AddFileKeys), arg0...)
}

// AddMulti mocks base method.
func (m *MockFileStore) AddMulti(arg0 bool, arg1 ...*storage.FileInfo) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddMulti", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMulti indicates an expected call of AddMulti.
func (mr *MockFileStoreMockRecorder) AddMulti(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulti", reflect.TypeOf((*MockFileStore)(nil).AddMulti), varargs...)
}

// AddTarget mocks base method.
func (m *MockFileStore) AddTarget(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTarget", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTarget indicates an expected call of AddTarget.
func (mr *MockFileStoreMockRecorder) AddTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTarget", reflect.TypeOf((*MockFileStore)(nil).AddTarget), arg0, arg1)
}

// Close mocks base method.
func (m *MockFileStore) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileStoreMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFileStore)(nil).Close), arg0)
}

// DeleteFile mocks base method.
func (m *MockFileStore) DeleteFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockFileStoreMockRecorder) DeleteFile(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockFileStore)(nil).DeleteFile), arg0)
}

// GetByChecksum mocks base method.
func (m *MockFileStore) GetByChecksum(arg0, arg1 string) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByChecksum", arg0, arg1)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByChecksum indicates an expected call of GetByChecksum.
func (mr *MockFileStoreMockRecorder) GetByChecksum(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByChecksum", reflect.TypeOf((*MockFileStore)(nil).GetByChecksum), arg0, arg1)
}

// GetByHash mocks base method.
func (m *MockFileStore) GetByHash(arg0 string) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHash", arg0)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHash indicates an expected call of GetByHash.
func (mr *MockFileStoreMockRecorder) GetByHash(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHash", reflect.TypeOf((*MockFileStore)(nil).GetByHash), arg0)
}

// GetBySource mocks base method.
func (m *MockFileStore) GetBySource(arg0, arg1, arg2 string) (*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySource", arg0, arg1, arg2)
	ret0, _ := ret[0].(*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySource indicates an expected call of GetBySource.
func (mr *MockFileStoreMockRecorder) GetBySource(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySource", reflect.TypeOf((*MockFileStore)(nil).GetBySource), arg0, arg1, arg2)
}

// GetChunksCount mocks base method.
func (m *MockFileStore) GetChunksCount(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChunksCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChunksCount indicates an expected call of GetChunksCount.
func (mr *MockFileStoreMockRecorder) GetChunksCount(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChunksCount", reflect.TypeOf((*MockFileStore)(nil).GetChunksCount), arg0)
}

// GetFileKeys mocks base method.
func (m *MockFileStore) GetFileKeys(arg0 string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileKeys", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileKeys indicates an expected call of GetFileKeys.
func (mr *MockFileStoreMockRecorder) GetFileKeys(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileKeys", reflect.TypeOf((*MockFileStore)(nil).GetFileKeys), arg0)
}

// GetFileOrigin mocks base method.
func (m *MockFileStore) GetFileOrigin(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileOrigin", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileOrigin indicates an expected call of GetFileOrigin.
func (mr *MockFileStoreMockRecorder) GetFileOrigin(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileOrigin", reflect.TypeOf((*MockFileStore)(nil).GetFileOrigin), arg0)
}

// GetFileSize mocks base method.
func (m *MockFileStore) GetFileSize(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileSize", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileSize indicates an expected call of GetFileSize.
func (mr *MockFileStoreMockRecorder) GetFileSize(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileSize", reflect.TypeOf((*MockFileStore)(nil).GetFileSize), arg0)
}

// GetSyncStatus mocks base method.
func (m *MockFileStore) GetSyncStatus(arg0 string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncStatus", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncStatus indicates an expected call of GetSyncStatus.
func (mr *MockFileStoreMockRecorder) GetSyncStatus(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncStatus", reflect.TypeOf((*MockFileStore)(nil).GetSyncStatus), arg0)
}

// Indexes mocks base method.
func (m *MockFileStore) Indexes() []localstore.Index {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Indexes")
	ret0, _ := ret[0].([]localstore.Index)
	return ret0
}

// Indexes indicates an expected call of Indexes.
func (mr *MockFileStoreMockRecorder) Indexes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Indexes", reflect.TypeOf((*MockFileStore)(nil).Indexes))
}

// Init mocks base method.
func (m *MockFileStore) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockFileStoreMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockFileStore)(nil).Init), arg0)
}

// IsFileImported mocks base method.
func (m *MockFileStore) IsFileImported(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsFileImported", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsFileImported indicates an expected call of IsFileImported.
func (mr *MockFileStoreMockRecorder) IsFileImported(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsFileImported", reflect.TypeOf((*MockFileStore)(nil).IsFileImported), arg0)
}

// List mocks base method.
func (m *MockFileStore) List() ([]*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List")
	ret0, _ := ret[0].([]*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockFileStoreMockRecorder) List() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockFileStore)(nil).List))
}

// ListByTarget mocks base method.
func (m *MockFileStore) ListByTarget(arg0 string) ([]*storage.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByTarget", arg0)
	ret0, _ := ret[0].([]*storage.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByTarget indicates an expected call of ListByTarget.
func (mr *MockFileStoreMockRecorder) ListByTarget(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByTarget", reflect.TypeOf((*MockFileStore)(nil).ListByTarget), arg0)
}

// ListTargets mocks base method.
func (m *MockFileStore) ListTargets() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTargets")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTargets indicates an expected call of ListTargets.
func (mr *MockFileStoreMockRecorder) ListTargets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTargets", reflect.TypeOf((*MockFileStore)(nil).ListTargets))
}

// Name mocks base method.
func (m *MockFileStore) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFileStoreMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockFileStore)(nil).Name))
}

// RemoveEmptyFileKeys mocks base method.
func (m *MockFileStore) RemoveEmptyFileKeys() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEmptyFileKeys")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEmptyFileKeys indicates an expected call of RemoveEmptyFileKeys.
func (mr *MockFileStoreMockRecorder) RemoveEmptyFileKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEmptyFileKeys", reflect.TypeOf((*MockFileStore)(nil).RemoveEmptyFileKeys))
}

// Run mocks base method.
func (m *MockFileStore) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockFileStoreMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockFileStore)(nil).Run), arg0)
}

// SetChunksCount mocks base method.
func (m *MockFileStore) SetChunksCount(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChunksCount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChunksCount indicates an expected call of SetChunksCount.
func (mr *MockFileStoreMockRecorder) SetChunksCount(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChunksCount", reflect.TypeOf((*MockFileStore)(nil).SetChunksCount), arg0, arg1)
}

// SetFileOrigin mocks base method.
func (m *MockFileStore) SetFileOrigin(arg0 string, arg1 model.ObjectOrigin) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFileOrigin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFileOrigin indicates an expected call of SetFileOrigin.
func (mr *MockFileStoreMockRecorder) SetFileOrigin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileOrigin", reflect.TypeOf((*MockFileStore)(nil).SetFileOrigin), arg0, arg1)
}

// SetFileSize mocks base method.
func (m *MockFileStore) SetFileSize(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFileSize", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFileSize indicates an expected call of SetFileSize.
func (mr *MockFileStoreMockRecorder) SetFileSize(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFileSize", reflect.TypeOf((*MockFileStore)(nil).SetFileSize), arg0, arg1)
}

// SetIsFileImported mocks base method.
func (m *MockFileStore) SetIsFileImported(arg0 string, arg1 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetIsFileImported", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetIsFileImported indicates an expected call of SetIsFileImported.
func (mr *MockFileStoreMockRecorder) SetIsFileImported(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetIsFileImported", reflect.TypeOf((*MockFileStore)(nil).SetIsFileImported), arg0, arg1)
}

// SetSyncStatus mocks base method.
func (m *MockFileStore) SetSyncStatus(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSyncStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSyncStatus indicates an expected call of SetSyncStatus.
func (mr *MockFileStoreMockRecorder) SetSyncStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSyncStatus", reflect.TypeOf((*MockFileStore)(nil).SetSyncStatus), arg0, arg1)
}
