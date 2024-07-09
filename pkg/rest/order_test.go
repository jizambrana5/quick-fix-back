package rest

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/jizambrana5/quickfix-back/pkg/domain"
)

func (t *handlerSuite) Test_GetOrdersByProfessional() {
	t.orderService.GetOrdersByProfessionalAndScheduleToFunc = func(background context.Context, professionalID uint64, scheduleTo time.Time) ([]domain.Order, error) {
		return []domain.Order{{ID: "as"}}, nil
	}
	params := []gin.Param{
		{
			Key:   "professional_id",
			Value: "1",
		},
	}

	queryParams := url.Values{}
	queryParams.Add("schedule_to", "2021-01-01T00:00:00Z")
	MockRequest(t.ctx, params, queryParams, http.MethodGet)
	t.handler.GetOrderByProfessional(t.ctx)
	t.Equal(http.StatusOK, t.w.Code)
	t.Equal("[{\"id\":\"as\",\"user_id\":0,\"professional_id\":0,\"status\":\"\",\"dates\":{\"created_at\":\"0001-01-01T00:00:00Z\",\"updated_at\":\"0001-01-01T00:00:00Z\",\"schedule_to\":\"0001-01-01T00:00:00Z\"}}]", t.w.Body.String())
}
