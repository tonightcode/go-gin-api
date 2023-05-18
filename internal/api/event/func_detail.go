package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type detailRequest struct{}

type detailResponse struct{}

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
