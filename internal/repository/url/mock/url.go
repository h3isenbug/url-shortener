// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/url/url.go

// Package mock_url is a generated GoMock package.
package mock_url

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/h3isenbug/url-shortener/internal/types"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateShortUrl mocks base method.
func (m *MockRepository) CreateShortUrl(ctx context.Context, originalUrl, slug string, accountID uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateShortUrl", ctx, originalUrl, slug, accountID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateShortUrl indicates an expected call of CreateShortUrl.
func (mr *MockRepositoryMockRecorder) CreateShortUrl(ctx, originalUrl, slug, accountID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateShortUrl", reflect.TypeOf((*MockRepository)(nil).CreateShortUrl), ctx, originalUrl, slug, accountID)
}

// GetByAccountID mocks base method.
func (m *MockRepository) GetByAccountID(ctx context.Context, accountID uint64, cursor string) ([]types.Url, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByAccountID", ctx, accountID, cursor)
	ret0, _ := ret[0].([]types.Url)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetByAccountID indicates an expected call of GetByAccountID.
func (mr *MockRepositoryMockRecorder) GetByAccountID(ctx, accountID, cursor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByAccountID", reflect.TypeOf((*MockRepository)(nil).GetByAccountID), ctx, accountID, cursor)
}

// GetBySlug mocks base method.
func (m *MockRepository) GetBySlug(ctx context.Context, slug string) (*types.Url, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySlug", ctx, slug)
	ret0, _ := ret[0].(*types.Url)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBySlug indicates an expected call of GetBySlug.
func (mr *MockRepositoryMockRecorder) GetBySlug(ctx, slug interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySlug", reflect.TypeOf((*MockRepository)(nil).GetBySlug), ctx, slug)
}

// IncrementVisits mocks base method.
func (m *MockRepository) IncrementVisits(ctx context.Context, slug string, newVisit bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IncrementVisits", ctx, slug, newVisit)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrementVisits indicates an expected call of IncrementVisits.
func (mr *MockRepositoryMockRecorder) IncrementVisits(ctx, slug, newVisit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrementVisits", reflect.TypeOf((*MockRepository)(nil).IncrementVisits), ctx, slug, newVisit)
}

// SetUrlState mocks base method.
func (m *MockRepository) SetUrlState(ctx context.Context, accountID uint64, slug string, disabled bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUrlState", ctx, accountID, slug, disabled)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetUrlState indicates an expected call of SetUrlState.
func (mr *MockRepositoryMockRecorder) SetUrlState(ctx, accountID, slug, disabled interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUrlState", reflect.TypeOf((*MockRepository)(nil).SetUrlState), ctx, accountID, slug, disabled)
}
