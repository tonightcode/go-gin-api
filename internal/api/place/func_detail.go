package place

import (
	"net/http"
	"strconv"
	"time"

	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	service "go-gin-api/internal/services/place"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Id         int32     `json:"id"`         // ID
	Name       string    `json:"name"`       // name
	Content    string    `json:"content"`    // content
	Province   string    `json:"province"`   // province
	City       string    `json:"city"`       // city
	County     string    `json:"county"`     // county
	Img        string    `json:"img"`        // img
	Address    string    `json:"address"`    // address
	Created_at time.Time `json:"created_at"` // create time
	Updated_at time.Time `json:"updated_at"` // update time
}

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
		id, _ := strconv.ParseInt(req.Id, 10, 32)

		where := new(service.Where)
		where.Id = int32(id)

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
			Id:         data.Id,
			Name:       data.Name,
			Content:    data.Content,
			Province:   data.Province,
			City:       data.City,
			County:     data.County,
			Img:        data.Img,
			Address:    data.Address,
			Created_at: data.CreatedAt,
			Updated_at: data.UpdatedAt,
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))
	}
}
