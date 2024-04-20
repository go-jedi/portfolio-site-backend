// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	project "github.com/go-jedi/portfolio/internal/model/project"
	review "github.com/go-jedi/portfolio/internal/model/review"
	gomock "github.com/golang/mock/gomock"
)

// MockProjectRepository is a mock of ProjectRepository interface.
type MockProjectRepository struct {
	ctrl     *gomock.Controller
	recorder *MockProjectRepositoryMockRecorder
}

// MockProjectRepositoryMockRecorder is the mock recorder for MockProjectRepository.
type MockProjectRepositoryMockRecorder struct {
	mock *MockProjectRepository
}

// NewMockProjectRepository creates a new mock instance.
func NewMockProjectRepository(ctrl *gomock.Controller) *MockProjectRepository {
	mock := &MockProjectRepository{ctrl: ctrl}
	mock.recorder = &MockProjectRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProjectRepository) EXPECT() *MockProjectRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProjectRepository) Create(ctx context.Context, dto project.Create) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dto)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockProjectRepositoryMockRecorder) Create(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProjectRepository)(nil).Create), ctx, dto)
}

// Get mocks base method.
func (m *MockProjectRepository) Get(ctx context.Context, page, limit int) ([]project.Get, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, page, limit)
	ret0, _ := ret[0].([]project.Get)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockProjectRepositoryMockRecorder) Get(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProjectRepository)(nil).Get), ctx, page, limit)
}

// GetByID mocks base method.
func (m *MockProjectRepository) GetByID(ctx context.Context, id int) (project.Get, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(project.Get)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockProjectRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockProjectRepository)(nil).GetByID), ctx, id)
}

// MockImageRepository is a mock of ImageRepository interface.
type MockImageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockImageRepositoryMockRecorder
}

// MockImageRepositoryMockRecorder is the mock recorder for MockImageRepository.
type MockImageRepositoryMockRecorder struct {
	mock *MockImageRepository
}

// NewMockImageRepository creates a new mock instance.
func NewMockImageRepository(ctrl *gomock.Controller) *MockImageRepository {
	mock := &MockImageRepository{ctrl: ctrl}
	mock.recorder = &MockImageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockImageRepository) EXPECT() *MockImageRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockImageRepository) Create(ctx context.Context, id int, paths []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, id, paths)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockImageRepositoryMockRecorder) Create(ctx, id, paths interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImageRepository)(nil).Create), ctx, id, paths)
}

// MockReviewRepository is a mock of ReviewRepository interface.
type MockReviewRepository struct {
	ctrl     *gomock.Controller
	recorder *MockReviewRepositoryMockRecorder
}

// MockReviewRepositoryMockRecorder is the mock recorder for MockReviewRepository.
type MockReviewRepositoryMockRecorder struct {
	mock *MockReviewRepository
}

// NewMockReviewRepository creates a new mock instance.
func NewMockReviewRepository(ctrl *gomock.Controller) *MockReviewRepository {
	mock := &MockReviewRepository{ctrl: ctrl}
	mock.recorder = &MockReviewRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReviewRepository) EXPECT() *MockReviewRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockReviewRepository) Create(ctx context.Context, dto review.Create) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, dto)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockReviewRepositoryMockRecorder) Create(ctx, dto interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockReviewRepository)(nil).Create), ctx, dto)
}

// Delete mocks base method.
func (m *MockReviewRepository) Delete(ctx context.Context, id int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockReviewRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockReviewRepository)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockReviewRepository) Get(ctx context.Context, page, limit int) ([]review.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, page, limit)
	ret0, _ := ret[0].([]review.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockReviewRepositoryMockRecorder) Get(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockReviewRepository)(nil).Get), ctx, page, limit)
}

// GetByID mocks base method.
func (m *MockReviewRepository) GetByID(ctx context.Context, id int) (review.Review, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(review.Review)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockReviewRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockReviewRepository)(nil).GetByID), ctx, id)
}

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockUserRepository) Get(ctx context.Context, id int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUserRepositoryMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepository)(nil).Get), ctx, id)
}
