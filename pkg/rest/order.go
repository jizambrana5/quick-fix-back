package rest

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"github.com/jizambrana5/quickfix-back/pkg/lib/errors"
)

// GetOrder maneja la solicitud GET /orders/:id
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("order_id")
	order, err := h.orderService.GetOrder(context.Background(), id)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, order)
}

// GetOrdersByUser maneja la solicitud GET /orders/user/:user_id
func (h *Handler) GetOrdersByUser(c *gin.Context) {
	userID, err := ParseUserIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	orders, err := h.orderService.GetOrdersByUser(context.Background(), userID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

// GetOrderByProfessional maneja la solicitud GET /orders/user/:user_id
func (h *Handler) GetOrderByProfessional(c *gin.Context) {
	professionalID, err := ParseProfessionalIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	orders, err := h.orderService.GetOrdersByProfessional(context.Background(), professionalID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CreateOrder maneja la solicitud POST /orders
func (h *Handler) CreateOrder(c *gin.Context) {
	var orderReq entities.CreateOrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		handleError(c, errors.ErrInvalidCreateOrder)
		return
	}

	if err := orderReq.Validate(); err != nil {
		handleError(c, err)
		return
	}

	createdOrder, err := h.orderService.CreateOrder(context.Background(), orderReq)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, createdOrder)
}

// AcceptOrder maneja la solicitud PUT /orders/:id/accept
func (h *Handler) AcceptOrder(c *gin.Context) {
	orderID, err := ParseOrderIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	updatedOrder, err := h.orderService.AcceptOrder(context.Background(), orderID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}

// CompleteOrder maneja la solicitud PUT /orders/:id/complete
func (h *Handler) CompleteOrder(c *gin.Context) {
	orderID, err := ParseOrderIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	updatedOrder, err := h.orderService.CompleteOrder(context.Background(), orderID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}

// CancelOrder maneja la solicitud PUT /orders/:id/complete
func (h *Handler) CancelOrder(c *gin.Context) {
	orderID, err := ParseOrderIDFromRequest(c)
	if err != nil {
		handleError(c, err)
		return
	}

	updatedOrder, err := h.orderService.CancelOrder(context.Background(), orderID)
	if err != nil {
		handleError(c, err)
		return
	}
	c.JSON(http.StatusOK, updatedOrder)
}
