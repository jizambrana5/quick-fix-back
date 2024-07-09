package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

// CreateUser maneja la creación de usuarios
func (h *Handler) CreateUser(c *gin.Context) {
	var user entities.RegisterUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		handleError(c, errors.ErrInvalidCreateOrder)
		return
	}

	if err := user.Validate(); err != nil {
		handleError(c, err)
		return
	}

	createdUser, err := h.userService.RegisterUser(c.Request.Context(), user)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

// GetUser maneja la obtención de usuarios por ID
func (h *Handler) GetUser(c *gin.Context) {
	userID, err := ParseUserIDFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateProfessional maneja la creación de profesionales
func (h *Handler) CreateProfessional(c *gin.Context) {
	var professional entities.RegisterProfessionalRequest
	if err := c.ShouldBindJSON(&professional); err != nil {
		handleError(c, err)
		return
	}

	if err := professional.Validate(); err != nil {
		handleError(c, err)
		return
	}

	createdProfessional, err := h.userService.RegisterProfessional(c.Request.Context(), professional)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, createdProfessional)
}

// GetProfessional maneja la obtención de profesionales por ID
func (h *Handler) GetProfessional(c *gin.Context) {
	profID, err := ParseProfessionalIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	professional, err := h.userService.GetProfessional(c.Request.Context(), profID)
	if err != nil {
		handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, professional)
}

func (h *Handler) GetProfessionalsByLocation(c *gin.Context) {
	department := c.Param("department")
	district := c.Param("district")

	professionals, err := h.userService.FindProfessionalsByLocation(c.Request.Context(), department, district)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, professionals)
}
