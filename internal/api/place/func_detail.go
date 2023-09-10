package place

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type detailRequest struct{}

type detailResponse struct{}

// Detail 事件详情
// @Summary 事件详情
// @Description 事件详情
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
