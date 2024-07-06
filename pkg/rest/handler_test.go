package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
	"github.com/jizambrana5/quickfix-back/pkg/entities"
	"github.com/jizambrana5/quickfix-back/pkg/rest/mocks"
)

type handlerSuite struct {
	suite.Suite
	ctx          *gin.Context
	userService  *mocks.UserServiceMock
	orderService *mocks.OrderServiceMock
	handler      *Handler
	w            *httptest.ResponseRecorder
}

func (t *handlerSuite) SetupTest() {
	t.w = httptest.NewRecorder()
	gin.SetMode(gin.TestMode)
	t.ctx, _ = gin.CreateTestContext(t.w)
	t.ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	t.userService = &mocks.UserServiceMock{
		FindProfessionalsByLocationFunc: func(ctx context.Context, department string, district string) ([]domain.Professional, error) {
			return []domain.Professional{{ID: 1, Username: "test"}}, nil
		},
		GetProfessionalFunc: func(ctx context.Context, ID uint64) (domain.Professional, error) {
			return domain.Professional{ID: 1}, nil
		},
		GetUserFunc: func(ctx context.Context, ID uint64) (domain.User, error) {
			return domain.User{ID: 1}, nil
		},
		RegisterProfessionalFunc: func(ctx context.Context, professionalReq entities.RegisterProfessionalRequest) (domain.Professional, error) {
			return domain.Professional{ID: 1}, nil
		},
		RegisterUserFunc: func(ctx context.Context, userReq entities.RegisterUserRequest) (domain.User, error) {
			return domain.User{ID: 1}, nil
		},
	}
	t.orderService = &mocks.OrderServiceMock{
		AcceptOrderFunc: func(ctx context.Context, orderID string) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		CancelOrderFunc: func(ctx context.Context, orderID string) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		CompleteOrderFunc: func(ctx context.Context, orderID string) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		CreateOrderFunc: func(ctx context.Context, order entities.CreateOrderRequest) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		GetOrderFunc: func(ctx context.Context, ID string) (domain.Order, error) {
			return domain.Order{ID: "aaaa"}, nil
		},
		GetOrdersByProfessionalFunc: func(ctx context.Context, professionalID uint64) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaaa"}}, nil
		},
		GetOrdersByUserFunc: func(ctx context.Context, userID uint64) ([]domain.Order, error) {
			return []domain.Order{{ID: "aaa"}}, nil
		},
	}
	t.handler = NewHandler(t.orderService, t.userService)
}

func (t *handlerSuite) Test_NewHandler() {
	t.NotNil(NewHandler(t.orderService, t.userService))
}

func TestHandler(t *testing.T) {
	suite.Run(t, new(handlerSuite))
}

func MockRequest(c *gin.Context, params gin.Params, u url.Values, method string) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")

	// set path params
	c.Params = params

	// set query params
	c.Request.URL.RawQuery = u.Encode()
}
