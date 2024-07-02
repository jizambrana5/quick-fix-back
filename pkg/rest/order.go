package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

// GetOrder maneja la solicitud GET /orders/:id
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.orderService.GetOrder(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetOrdersByUser maneja la solicitud GET /orders/user/:user_id
func (h *Handler) GetOrdersByUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	orders, err := h.orderService.GetOrdersByUser(context.Background(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CreateOrder maneja la solicitud POST /orders
func (h *Handler) CreateOrder(c *gin.Context) {
	var order domain.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdOrder, err := h.orderService.CreateOrder(context.Background(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdOrder)
}

// AdvanceOrder maneja la solicitud PUT /orders/:id/status
func (h *Handler) AdvanceOrder(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedOrder, err := h.orderService.AdvanceOrder(context.Background(), id, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}
