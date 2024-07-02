package rest

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

func ParseUserIDFromRequest(c *gin.Context) (uint64, error) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return 0, errors.ErrInvalidUserID
	}
	if userID == 0 {
		return 0, errors.ErrInvalidUserID
	}
	return userID, nil
}

func ParseProfessionalIDFromRequest(c *gin.Context) (uint64, error) {
	profIDStr := c.Param("professional_id")
	profID, err := strconv.ParseUint(profIDStr, 10, 64)
	if err != nil {
		return 0, errors.ErrInvalidUserID
	}
	if profID == 0 {
		return 0, errors.ErrInvalidUserID
	}
	return profID, nil
}

func ParseOrderIDFromRequest(c *gin.Context) (string, error) {
	orderID := c.Param("order_id")
	if orderID == "" {
		return "", errors.ErrInvalidOrderID
	}
	return orderID, nil
}
