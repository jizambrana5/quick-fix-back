package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"github.com/jizambrana5/quickfix-back/pkg/utils"
)

// Mock de LoadLocations para los tests
func mockLoadLocations() (utils.Locations, error) {
	return utils.Locations{
		Locations: map[string][]string{
			"Department1": {"District1", "District2"},
			"Department2": {"District3", "District4"},
		},
	}, nil
}

func TestAdvanceOrderRequest_Validate(t *testing.T) {
	tests := []struct {
		status domain.Status
		err    error
	}{
		{domain.OrderStatusPending, nil},
		{domain.OrderStatusAccepted, nil},
		{domain.OrderStatusCompleted, nil},
		{domain.OrderStatusCancelled, nil},
		{"invalid_status", errors.ErrInvalidInput},
	}

	for _, tt := range tests {
		r := AdvanceOrderRequest{Status: tt.status}
		err := r.Validate()
		assert.Equal(t, tt.err, err)
	}
}

func TestCreateOrderRequest_Validate(t *testing.T) {
	validTime := time.Now().Add(24 * time.Hour).Format(layout)

	tests := []struct {
		req CreateOrderRequest
		err error
	}{
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: validTime}, nil},
		{CreateOrderRequest{UserID: 0, ProfessionalID: 1, ScheduleTo: validTime}, errors.ErrInvalidUserID},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 0, ScheduleTo: validTime}, errors.ErrInvalidProfessionalID},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: ""}, errors.ErrInvalidScheduleTo},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: "invalid_time"}, errors.ErrInvalidScheduleTo},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: time.Now().Add(-24 * time.Hour).Format(layout)}, errors.ErrInvalidScheduleTo},
	}

	for _, tt := range tests {
		err := tt.req.Validate()
		assert.Equal(t, tt.err, err)
	}
}

func TestRegisterUserRequest_Validate(t *testing.T) {
	tests := []struct {
		req RegisterUserRequest
		err error
	}{
		{RegisterUserRequest{Username: "user", Email: "user@example.com", Password: "password"}, nil},
		{RegisterUserRequest{Username: "", Email: "user@example.com", Password: "password"}, errors.EmptyUserName},
		{RegisterUserRequest{Username: "user", Email: "", Password: "password"}, errors.EmptyEmail},
		{RegisterUserRequest{Username: "user", Email: "user@example.com", Password: ""}, errors.EmptyPassword},
	}

	for _, tt := range tests {
		err := tt.req.Validate()
		assert.Equal(t, tt.err, err)
	}
}

/*
func TestRegisterProfessionalRequest_Validate(t *testing.T) {
	//originalLoadLocations := utils.LoadLocations
	//utils.LoadLocations = mockLoadLocations
	//defer func() { utils.LoadLocations = originalLoadLocations }()

	validLocation := Location{Department: "Department1", District: "District1"}

	tests := []struct {
		req RegisterProfessionalRequest
		err error
	}{
		{RegisterProfessionalRequest{Username: "pro", Email: "pro@example.com", Password: "password", Profession: "PLOMERO", Location: validLocation}, nil},
		{RegisterProfessionalRequest{Username: "", Email: "pro@example.com", Password: "password", Profession: "PLOMERO", Location: validLocation}, errors.EmptyUserName},
		{RegisterProfessionalRequest{Username: "pro", Email: "", Password: "password", Profession: "PLOMERO", Location: validLocation}, errors.EmptyEmail},
		{RegisterProfessionalRequest{Username: "pro", Email: "pro@example.com", Password: "", Profession: "PLOMERO", Location: validLocation}, errors.EmptyPassword},
		{RegisterProfessionalRequest{Username: "pro", Email: "pro@example.com", Password: "password", Profession: "INVALID_PROFESSION", Location: validLocation}, errors.ErrInvalidProfession},
		{RegisterProfessionalRequest{Username: "pro", Email: "pro@example.com", Password: "password", Profession: "PLOMERO", Location: Location{Department: "InvalidDepartment", District: "District1"}}, errors.ErrInvalidLocation},
	}

	for _, tt := range tests {
		err := tt.req.Validate()
		assert.Equal(t, tt.err, err)
	}
}*/

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email string
		err   error
	}{
		{"test@example.com", nil},
		{"invalid-email", errors.ErrInvalidEmail},
		{"another.test@domain.co", nil},
		{"missing@at.sign", nil},
		{"@missingusername.com", errors.ErrInvalidEmail},
	}

	for _, tt := range tests {
		err := IsValidEmail(tt.email)
		if tt.err == nil {
			assert.NoError(t, err)
		} else {
			assert.ErrorIs(t, err, tt.err)
		}
	}
}
