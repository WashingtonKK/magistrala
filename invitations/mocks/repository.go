// Code generated by mockery v2.43.2. DO NOT EDIT.

// Copyright (c) Abstract Machines

package mocks

import (
	context "context"

	invitations "github.com/absmach/magistrala/invitations"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, invitation
func (_m *Repository) Create(ctx context.Context, invitation invitations.Invitation) error {
	ret := _m.Called(ctx, invitation)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, invitations.Invitation) error); ok {
		r0 = rf(ctx, invitation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, userID, domainID
func (_m *Repository) Delete(ctx context.Context, userID string, domainID string) error {
	ret := _m.Called(ctx, userID, domainID)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, userID, domainID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Retrieve provides a mock function with given fields: ctx, userID, domainID
func (_m *Repository) Retrieve(ctx context.Context, userID string, domainID string) (invitations.Invitation, error) {
	ret := _m.Called(ctx, userID, domainID)

	if len(ret) == 0 {
		panic("no return value specified for Retrieve")
	}

	var r0 invitations.Invitation
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (invitations.Invitation, error)); ok {
		return rf(ctx, userID, domainID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) invitations.Invitation); ok {
		r0 = rf(ctx, userID, domainID)
	} else {
		r0 = ret.Get(0).(invitations.Invitation)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, userID, domainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RetrieveAll provides a mock function with given fields: ctx, page
func (_m *Repository) RetrieveAll(ctx context.Context, page invitations.Page) (invitations.InvitationPage, error) {
	ret := _m.Called(ctx, page)

	if len(ret) == 0 {
		panic("no return value specified for RetrieveAll")
	}

	var r0 invitations.InvitationPage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, invitations.Page) (invitations.InvitationPage, error)); ok {
		return rf(ctx, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, invitations.Page) invitations.InvitationPage); ok {
		r0 = rf(ctx, page)
	} else {
		r0 = ret.Get(0).(invitations.InvitationPage)
	}

	if rf, ok := ret.Get(1).(func(context.Context, invitations.Page) error); ok {
		r1 = rf(ctx, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConfirmation provides a mock function with given fields: ctx, invitation
func (_m *Repository) UpdateConfirmation(ctx context.Context, invitation invitations.Invitation) error {
	ret := _m.Called(ctx, invitation)

	if len(ret) == 0 {
		panic("no return value specified for UpdateConfirmation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, invitations.Invitation) error); ok {
		r0 = rf(ctx, invitation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateToken provides a mock function with given fields: ctx, invitation
func (_m *Repository) UpdateToken(ctx context.Context, invitation invitations.Invitation) error {
	ret := _m.Called(ctx, invitation)

	if len(ret) == 0 {
		panic("no return value specified for UpdateToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, invitations.Invitation) error); ok {
		r0 = rf(ctx, invitation)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
