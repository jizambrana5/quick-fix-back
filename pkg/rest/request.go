package rest

import (
	"strconv"
	"time"

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
		return 0, errors.ErrInvalidProfessionalID
	}
	if profID == 0 {
		return 0, errors.ErrInvalidProfessionalID
	}
	return profID, nil
}

func ParseScheduleToFromRequest(c *gin.Context) (time.Time, error) {
	scheduleToStr := c.Query("schedule_to")
	if scheduleToStr == "" {
		return time.Time{}, nil
	}
	scheduleTo, err := time.Parse(time.RFC3339, scheduleToStr)
	if err != nil {
		return time.Time{}, errors.ErrInvalidScheduleTo
	}
	return scheduleTo, nil
}

func ParseDayScheduleToFromRequest(c *gin.Context) (time.Time, error) {
	day := c.Param("day") // Esto debería ser el día en formato YYYY-MM-DD

	// Parsear la fecha
	parsedDay, err := time.Parse("2006-01-02", day)
	if err != nil {
		return time.Time{}, errors.ErrInvalidDay
	}
	return parsedDay, nil
}

func ParseOrderIDFromRequest(c *gin.Context) (string, error) {
	orderID := c.Param("order_id")
	if orderID == "" {
		return "", errors.ErrInvalidOrderID
	}
	return orderID, nil
}
