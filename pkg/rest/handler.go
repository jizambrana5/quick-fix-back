package rest

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
	"github.com/jizambrana5/quickfix-back/pkg/utils"
)

type Handler struct {
	orderService OrderService
	userService  UserService
}

func NewHandler(orderService OrderService, userService UserService) *Handler {
	return &Handler{
		orderService: orderService,
		userService:  userService,
	}
}

type (
	OrderService interface {
		GetOrder(ctx context.Context, ID string) (domain.Order, error)
		GetOrdersByUser(ctx context.Context, userID uint64) ([]domain.Order, error)
		GetOrdersByProfessional(ctx context.Context, professionalID uint64) ([]domain.Order, error)
		CreateOrder(ctx context.Context, order CreateOrderRequest) (domain.Order, error)
		AcceptOrder(ctx context.Context, orderID string) (domain.Order, error)
		CompleteOrder(ctx context.Context, orderID string) (domain.Order, error)
		CancelOrder(ctx context.Context, orderID string) (domain.Order, error)
	}
	UserService interface {
		RegisterUser(ctx context.Context, userReq RegisterUserRequest) (domain.User, error)
		RegisterProfessional(ctx context.Context, professionalReq RegisterProfessionalRequest) (domain.Professional, error)
		GetUser(ctx context.Context, ID uint64) (domain.User, error)
		GetProfessional(ctx context.Context, ID uint64) (domain.Professional, error)
		FindProfessionalsByLocation(ctx context.Context, department string, district string) ([]domain.Professional, error)
	}
)

// PingHandler maneja la solicitud GET /ping
func (h *Handler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}

func handleError(c *gin.Context, err error) {
	if customErr, ok := err.(errors.CustomError); ok {
		c.JSON(customErr.HTTPCode(), gin.H{"code": customErr.InternalCode(), "message": customErr.Error()})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}

// GetLocations handler para obtener las ubicaciones.
func (h *Handler) GetLocations(c *gin.Context) {
	locations, err := utils.LoadLocations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": "", "message": err.Error()})
		return
	}

	fmt.Println("locations:", locations)
	c.JSON(http.StatusOK, locations)
}
