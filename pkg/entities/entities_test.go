package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

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
	validTime := time.Now().Add(24 * time.Hour).Format(Layout)

	tests := []struct {
		req CreateOrderRequest
		err error
	}{
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: validTime}, errors.EmptyAddress},
		{CreateOrderRequest{UserID: 0, ProfessionalID: 1, ScheduleTo: validTime}, errors.ErrInvalidUserID},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 0, ScheduleTo: validTime}, errors.ErrInvalidProfessionalID},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: ""}, errors.ErrInvalidScheduleTo},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: "invalid_time"}, errors.ErrInvalidScheduleTo},
		{CreateOrderRequest{UserID: 1, ProfessionalID: 1, ScheduleTo: time.Now().Add(-24 * time.Hour).Format(Layout)}, errors.EmptyAddress},
	}

	for _, tt := range tests {
		err := tt.req.Validate()
		assert.Equal(t, tt.err, err)
	}
}

func TestValidateRegisterUserRequest(t *testing.T) {
	tests := []struct {
		name     string
		request  RegisterUserRequest
		expected error
	}{
		{
			name: "Valid Request",
			request: RegisterUserRequest{
				Username: "user123",
				Name:     "John",
				LastName: "Doe",
				Phone:    "123456789",
				Address:  "123 Street, City",
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			expected: nil,
		},
		{
			name: "Empty Username",
			request: RegisterUserRequest{
				Username: "",
				Name:     "John",
				LastName: "Doe",
				Phone:    "123456789",
				Address:  "123 Street, City",
				Email:    "john.doe@example.com",
				Password: "password123",
			},
			expected: errors.EmptyUserName,
		},
		{
			name: "Empty Email",
			request: RegisterUserRequest{
				Username: "user123",
				Name:     "John",
				LastName: "Doe",
				Phone:    "123456789",
				Address:  "123 Street, City",
				Email:    "",
				Password: "password123",
			},
			expected: errors.EmptyEmail,
		},
		// Agrega más casos de prueba según sea necesario para cubrir todos los escenarios posibles
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			assert.Equal(t, tt.expected, err)
		})
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
