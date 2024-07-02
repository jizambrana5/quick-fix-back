package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

type Handler struct {
	orderService OrderService
}

func NewHandler(orderService OrderService) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

type OrderService interface {
	GetOrder(ctx context.Context, ID string) (domain.Order, error)
	GetOrdersByUser(ctx context.Context, userID uint64) ([]domain.Order, error)
	GetOrdersByProfessional(ctx context.Context, professionalID uint64) ([]domain.Order, error)
	CreateOrder(ctx context.Context, order CreateOrderRequest) (domain.Order, error)
	AcceptOrder(ctx context.Context, orderID string) (domain.Order, error)
	CompleteOrder(ctx context.Context, orderID string) (domain.Order, error)
	CancelOrder(ctx context.Context, orderID string) (domain.Order, error)
}

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
