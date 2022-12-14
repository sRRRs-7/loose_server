// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/sRRRs-7/loose_style.git/db/sqlc (interfaces: Store)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/sRRRs-7/loose_style.git/db/sqlc"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAdminUser mocks base method.
func (m *MockStore) CreateAdminUser(arg0 context.Context, arg1 db.CreateAdminUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAdminUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateAdminUser indicates an expected call of CreateAdminUser.
func (mr *MockStoreMockRecorder) CreateAdminUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdminUser", reflect.TypeOf((*MockStore)(nil).CreateAdminUser), arg0, arg1)
}

// CreateCode mocks base method.
func (m *MockStore) CreateCode(arg0 context.Context, arg1 db.CreateCodeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCode indicates an expected call of CreateCode.
func (mr *MockStoreMockRecorder) CreateCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCode", reflect.TypeOf((*MockStore)(nil).CreateCode), arg0, arg1)
}

// CreateCollection mocks base method.
func (m *MockStore) CreateCollection(arg0 context.Context, arg1 db.CreateCollectionParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateCollection indicates an expected call of CreateCollection.
func (mr *MockStoreMockRecorder) CreateCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCollection", reflect.TypeOf((*MockStore)(nil).CreateCollection), arg0, arg1)
}

// CreateMedia mocks base method.
func (m *MockStore) CreateMedia(arg0 context.Context, arg1 db.CreateMediaParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMedia", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMedia indicates an expected call of CreateMedia.
func (mr *MockStoreMockRecorder) CreateMedia(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMedia", reflect.TypeOf((*MockStore)(nil).CreateMedia), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 db.CreateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteCode mocks base method.
func (m *MockStore) DeleteCode(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCode indicates an expected call of DeleteCode.
func (mr *MockStoreMockRecorder) DeleteCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCode", reflect.TypeOf((*MockStore)(nil).DeleteCode), arg0, arg1)
}

// DeleteCollection mocks base method.
func (m *MockStore) DeleteCollection(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockStoreMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockStore)(nil).DeleteCollection), arg0, arg1)
}

// DeleteMedia mocks base method.
func (m *MockStore) DeleteMedia(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMedia", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMedia indicates an expected call of DeleteMedia.
func (mr *MockStoreMockRecorder) DeleteMedia(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMedia", reflect.TypeOf((*MockStore)(nil).DeleteMedia), arg0, arg1)
}

// DeleteUser mocks base method.
func (m *MockStore) DeleteUser(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockStoreMockRecorder) DeleteUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockStore)(nil).DeleteUser), arg0, arg1)
}

// GetAdminUser mocks base method.
func (m *MockStore) GetAdminUser(arg0 context.Context, arg1 db.GetAdminUserParams) (*db.Adminuser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdminUser", arg0, arg1)
	ret0, _ := ret[0].(*db.Adminuser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAdminUser indicates an expected call of GetAdminUser.
func (mr *MockStoreMockRecorder) GetAdminUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdminUser", reflect.TypeOf((*MockStore)(nil).GetAdminUser), arg0, arg1)
}

// GetAllCodes mocks base method.
func (m *MockStore) GetAllCodes(arg0 context.Context, arg1 db.GetAllCodesParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCodes", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCodes indicates an expected call of GetAllCodes.
func (mr *MockStoreMockRecorder) GetAllCodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCodes", reflect.TypeOf((*MockStore)(nil).GetAllCodes), arg0, arg1)
}

// GetAllCodesByKeyword mocks base method.
func (m *MockStore) GetAllCodesByKeyword(arg0 context.Context, arg1 db.GetAllCodesByKeywordParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCodesByKeyword", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCodesByKeyword indicates an expected call of GetAllCodesByKeyword.
func (mr *MockStoreMockRecorder) GetAllCodesByKeyword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCodesByKeyword", reflect.TypeOf((*MockStore)(nil).GetAllCodesByKeyword), arg0, arg1)
}

// GetAllCodesByTag mocks base method.
func (m *MockStore) GetAllCodesByTag(arg0 context.Context, arg1 db.GetAllCodesByTagParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCodesByTag", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCodesByTag indicates an expected call of GetAllCodesByTag.
func (mr *MockStoreMockRecorder) GetAllCodesByTag(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCodesByTag", reflect.TypeOf((*MockStore)(nil).GetAllCodesByTag), arg0, arg1)
}

// GetAllCodesSortedAccess mocks base method.
func (m *MockStore) GetAllCodesSortedAccess(arg0 context.Context, arg1 db.GetAllCodesSortedAccessParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCodesSortedAccess", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCodesSortedAccess indicates an expected call of GetAllCodesSortedAccess.
func (mr *MockStoreMockRecorder) GetAllCodesSortedAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCodesSortedAccess", reflect.TypeOf((*MockStore)(nil).GetAllCodesSortedAccess), arg0, arg1)
}

// GetAllCodesSortedStar mocks base method.
func (m *MockStore) GetAllCodesSortedStar(arg0 context.Context, arg1 db.GetAllCodesSortedStarParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCodesSortedStar", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCodesSortedStar indicates an expected call of GetAllCodesSortedStar.
func (mr *MockStoreMockRecorder) GetAllCodesSortedStar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCodesSortedStar", reflect.TypeOf((*MockStore)(nil).GetAllCodesSortedStar), arg0, arg1)
}

// GetAllCollections mocks base method.
func (m *MockStore) GetAllCollections(arg0 context.Context, arg1 db.GetAllCollectionsParams) ([]*db.GetAllCollectionsRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCollections", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetAllCollectionsRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCollections indicates an expected call of GetAllCollections.
func (mr *MockStoreMockRecorder) GetAllCollections(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCollections", reflect.TypeOf((*MockStore)(nil).GetAllCollections), arg0, arg1)
}

// GetAllCollectionsBySearch mocks base method.
func (m *MockStore) GetAllCollectionsBySearch(arg0 context.Context, arg1 db.GetAllCollectionsBySearchParams) ([]*db.GetAllCollectionsBySearchRow, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCollectionsBySearch", arg0, arg1)
	ret0, _ := ret[0].([]*db.GetAllCollectionsBySearchRow)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCollectionsBySearch indicates an expected call of GetAllCollectionsBySearch.
func (mr *MockStoreMockRecorder) GetAllCollectionsBySearch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCollectionsBySearch", reflect.TypeOf((*MockStore)(nil).GetAllCollectionsBySearch), arg0, arg1)
}

// GetAllOwnCodes mocks base method.
func (m *MockStore) GetAllOwnCodes(arg0 context.Context, arg1 db.GetAllOwnCodesParams) ([]*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOwnCodes", arg0, arg1)
	ret0, _ := ret[0].([]*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOwnCodes indicates an expected call of GetAllOwnCodes.
func (mr *MockStoreMockRecorder) GetAllOwnCodes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOwnCodes", reflect.TypeOf((*MockStore)(nil).GetAllOwnCodes), arg0, arg1)
}

// GetCode mocks base method.
func (m *MockStore) GetCode(arg0 context.Context, arg1 int64) (*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCode", arg0, arg1)
	ret0, _ := ret[0].(*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCode indicates an expected call of GetCode.
func (mr *MockStoreMockRecorder) GetCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCode", reflect.TypeOf((*MockStore)(nil).GetCode), arg0, arg1)
}

// GetCollection mocks base method.
func (m *MockStore) GetCollection(arg0 context.Context, arg1 int64) (*db.Codes, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollection", arg0, arg1)
	ret0, _ := ret[0].(*db.Codes)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollection indicates an expected call of GetCollection.
func (mr *MockStoreMockRecorder) GetCollection(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollection", reflect.TypeOf((*MockStore)(nil).GetCollection), arg0, arg1)
}

// GetMedia mocks base method.
func (m *MockStore) GetMedia(arg0 context.Context, arg1 int64) (*db.Media, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMedia", arg0, arg1)
	ret0, _ := ret[0].(*db.Media)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedia indicates an expected call of GetMedia.
func (mr *MockStoreMockRecorder) GetMedia(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedia", reflect.TypeOf((*MockStore)(nil).GetMedia), arg0, arg1)
}

// GetUserByID mocks base method.
func (m *MockStore) GetUserByID(arg0 context.Context, arg1 int64) (*db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", arg0, arg1)
	ret0, _ := ret[0].(*db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID.
func (mr *MockStoreMockRecorder) GetUserByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockStore)(nil).GetUserByID), arg0, arg1)
}

// GetUserByUsername mocks base method.
func (m *MockStore) GetUserByUsername(arg0 context.Context, arg1 string) (*db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByUsername", arg0, arg1)
	ret0, _ := ret[0].(*db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockStoreMockRecorder) GetUserByUsername(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockStore)(nil).GetUserByUsername), arg0, arg1)
}

// ListMedia mocks base method.
func (m *MockStore) ListMedia(arg0 context.Context, arg1 db.ListMediaParams) ([]*db.Media, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMedia", arg0, arg1)
	ret0, _ := ret[0].([]*db.Media)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMedia indicates an expected call of ListMedia.
func (mr *MockStoreMockRecorder) ListMedia(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMedia", reflect.TypeOf((*MockStore)(nil).ListMedia), arg0, arg1)
}

// LoginUser mocks base method.
func (m *MockStore) LoginUser(arg0 context.Context, arg1 db.LoginUserParams) (*db.Users, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUser", arg0, arg1)
	ret0, _ := ret[0].(*db.Users)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoginUser indicates an expected call of LoginUser.
func (mr *MockStoreMockRecorder) LoginUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUser", reflect.TypeOf((*MockStore)(nil).LoginUser), arg0, arg1)
}

// UpdateAccess mocks base method.
func (m *MockStore) UpdateAccess(arg0 context.Context, arg1 db.UpdateAccessParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccess", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccess indicates an expected call of UpdateAccess.
func (mr *MockStoreMockRecorder) UpdateAccess(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccess", reflect.TypeOf((*MockStore)(nil).UpdateAccess), arg0, arg1)
}

// UpdateCode mocks base method.
func (m *MockStore) UpdateCode(arg0 context.Context, arg1 db.UpdateCodeParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCode", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCode indicates an expected call of UpdateCode.
func (mr *MockStoreMockRecorder) UpdateCode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCode", reflect.TypeOf((*MockStore)(nil).UpdateCode), arg0, arg1)
}

// UpdateMedia mocks base method.
func (m *MockStore) UpdateMedia(arg0 context.Context, arg1 db.UpdateMediaParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMedia", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMedia indicates an expected call of UpdateMedia.
func (mr *MockStoreMockRecorder) UpdateMedia(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMedia", reflect.TypeOf((*MockStore)(nil).UpdateMedia), arg0, arg1)
}

// UpdateStar mocks base method.
func (m *MockStore) UpdateStar(arg0 context.Context, arg1 db.UpdateStarParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStar", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateStar indicates an expected call of UpdateStar.
func (mr *MockStoreMockRecorder) UpdateStar(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStar", reflect.TypeOf((*MockStore)(nil).UpdateStar), arg0, arg1)
}

// UpdateUser mocks base method.
func (m *MockStore) UpdateUser(arg0 context.Context, arg1 db.UpdateUserParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockStoreMockRecorder) UpdateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockStore)(nil).UpdateUser), arg0, arg1)
}
