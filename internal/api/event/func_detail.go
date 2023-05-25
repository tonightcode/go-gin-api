package event

import (
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type detailRequest struct{}

type detailResponse struct {
	Id         int32     `json:"id"`         // ID
	Title      string    `json:"title"`      // title
	Content    string    `json:"content"`    // content
	Cover      string    `json:"cover"`      // cover
	Created_at time.Time `json:"created_at"` // create time
	Updated_at time.Time `json:"updated_at"` // update time
}

// Detail 事件
// @Summary 事件
// @Description 事件
// @Tags API.event
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body detailRequest true "请求信息"
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/event/{id} [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(ctx core.Context) {

	}
}
