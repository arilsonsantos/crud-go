// Code generated by MockGen. DO NOT EDIT.
// Source: src/model/service/UserDomainServiceInterface.go
//
// Generated by this command:
//
//	mockgen -source=src/model/service/UserDomainServiceInterface.go -destination=src/test/mocks/user_domain_service_interface_mock.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	errors "github.com/arilsonsantos/crud-go.git/src/errors"
	domain "github.com/arilsonsantos/crud-go.git/src/model/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockUserDomainServiceInterface is a mock of UserDomainServiceInterface interface.
type MockUserDomainServiceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserDomainServiceInterfaceMockRecorder
}

// MockUserDomainServiceInterfaceMockRecorder is the mock recorder for MockUserDomainServiceInterface.
type MockUserDomainServiceInterfaceMockRecorder struct {
	mock *MockUserDomainServiceInterface
}

// NewMockUserDomainServiceInterface creates a new mock instance.
func NewMockUserDomainServiceInterface(ctrl *gomock.Controller) *MockUserDomainServiceInterface {
	mock := &MockUserDomainServiceInterface{ctrl: ctrl}
	mock.recorder = &MockUserDomainServiceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserDomainServiceInterface) EXPECT() *MockUserDomainServiceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockUserDomainServiceInterface) Create(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, *errors.ErrorDto) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", domainInterface)
	ret0, _ := ret[0].(domain.UserDomainInterface)
	ret1, _ := ret[1].(*errors.ErrorDto)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockUserDomainServiceInterfaceMockRecorder) Create(domainInterface any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).Create), domainInterface)
}

// Delete mocks base method.
func (m *MockUserDomainServiceInterface) Delete(arg0 string) *errors.ErrorDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(*errors.ErrorDto)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockUserDomainServiceInterfaceMockRecorder) Delete(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).Delete), arg0)
}

// FindByEmail mocks base method.
func (m *MockUserDomainServiceInterface) FindByEmail(email string) (domain.UserDomainInterface, *errors.ErrorDto) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEmail", email)
	ret0, _ := ret[0].(domain.UserDomainInterface)
	ret1, _ := ret[1].(*errors.ErrorDto)
	return ret0, ret1
}

// FindByEmail indicates an expected call of FindByEmail.
func (mr *MockUserDomainServiceInterfaceMockRecorder) FindByEmail(email any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEmail", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).FindByEmail), email)
}

// FindById mocks base method.
func (m *MockUserDomainServiceInterface) FindById(id string) (domain.UserDomainInterface, *errors.ErrorDto) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(domain.UserDomainInterface)
	ret1, _ := ret[1].(*errors.ErrorDto)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockUserDomainServiceInterfaceMockRecorder) FindById(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).FindById), id)
}

// LoginUserService mocks base method.
func (m *MockUserDomainServiceInterface) LoginUserService(domainInterface domain.UserDomainInterface) (domain.UserDomainInterface, string, *errors.ErrorDto) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginUserService", domainInterface)
	ret0, _ := ret[0].(domain.UserDomainInterface)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(*errors.ErrorDto)
	return ret0, ret1, ret2
}

// LoginUserService indicates an expected call of LoginUserService.
func (mr *MockUserDomainServiceInterfaceMockRecorder) LoginUserService(domainInterface any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginUserService", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).LoginUserService), domainInterface)
}

// Update mocks base method.
func (m *MockUserDomainServiceInterface) Update(arg0 string, arg1 domain.UserDomainInterface) *errors.ErrorDto {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*errors.ErrorDto)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockUserDomainServiceInterfaceMockRecorder) Update(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUserDomainServiceInterface)(nil).Update), arg0, arg1)
}
