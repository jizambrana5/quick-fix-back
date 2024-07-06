package rest

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

func TestParseUserIDFromRequest(t *testing.T) {
	tests := []struct {
		param         string
		expectedID    uint64
		expectedError error
	}{
		{"123", 123, nil},
		{"0", 0, errors.ErrInvalidUserID},
		{"abc", 0, errors.ErrInvalidUserID},
		{"", 0, errors.ErrInvalidUserID},
	}

	for _, tt := range tests {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = gin.Params{gin.Param{Key: "user_id", Value: tt.param}}

		userID, err := ParseUserIDFromRequest(c)
		assert.Equal(t, tt.expectedID, userID)
		assert.Equal(t, tt.expectedError, err)
	}
}

func TestParseProfessionalIDFromRequest(t *testing.T) {
	tests := []struct {
		param         string
		expectedID    uint64
		expectedError error
	}{
		{"123", 123, nil},
		{"0", 0, errors.ErrInvalidUserID},
		{"abc", 0, errors.ErrInvalidUserID},
		{"", 0, errors.ErrInvalidUserID},
	}

	for _, tt := range tests {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = gin.Params{gin.Param{Key: "professional_id", Value: tt.param}}

		profID, err := ParseProfessionalIDFromRequest(c)
		assert.Equal(t, tt.expectedID, profID)
		assert.Equal(t, tt.expectedError, err)
	}
}

func TestParseOrderIDFromRequest(t *testing.T) {
	tests := []struct {
		param         string
		expectedID    string
		expectedError error
	}{
		{"order123", "order123", nil},
		{"", "", errors.ErrInvalidOrderID},
	}

	for _, tt := range tests {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		c.Params = gin.Params{gin.Param{Key: "order_id", Value: tt.param}}

		orderID, err := ParseOrderIDFromRequest(c)
		assert.Equal(t, tt.expectedID, orderID)
		assert.Equal(t, tt.expectedError, err)
	}
}
