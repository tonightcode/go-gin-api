package singer

import (
	"net/http"
	"strconv"

	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	service "go-gin-api/internal/services/singer"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Singerid   int32  `json:"singerid"`   // singerid
	Name       string `json:"name"`       // name
	Head       string `json:"head"`       // url
	Is_deleted string `json:"is_deleted"` // is_deleted
	Created_at string `json:"created_at"` // create time
	Updated_at string `json:"updated_at"` // update time
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
	return func(c core.Context) {
		req := new(detailRequest)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithError(err),
			)
			return
		}
		singerid, _ := strconv.ParseInt(req.Id, 10, 32)

		where := new(service.Where)
		where.Singerid = int32(singerid)
		data, err := h.service.Detail(c, where)

		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res := &detailResponse{
			Singerid:   data.Singerid,
			Name:       data.Name,
			Head:       data.Head,
			Created_at: data.CreatedAt.Format("2006-01-02 15:04:05"),
			Updated_at: data.UpdatedAt.Format("2006-01-02 15:04:05"),
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))
	}
}
