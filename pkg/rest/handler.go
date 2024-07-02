package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
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
	CreateOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	AdvanceOrder(ctx context.Context, ID string, status string) (domain.Order, error)
}

// PingHandler maneja la solicitud GET /ping
func (h *Handler) PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}
