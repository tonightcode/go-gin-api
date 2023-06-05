package person

import (
	"net/http"
	"strconv"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/person"
)

type detailRequest struct {
	Id string `uri:"id"` // HashID
}

type detailResponse struct {
	Id         int32     `json:"id"`         // ID
	Username   string    `json:"username"`   // username
	Intro      string    `json:"intro"`      // intro
	Icon       string    `json:"icon"`       // icon
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
		id, err := strconv.ParseInt(req.Id, 10, 32)

		personParams := new(person.PersonParams)
		personParams.Id = int32(id)

		data, err := h.personService.Detail(c, personParams)

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
			Username:   data.Username,
			Intro:      data.Intro,
			Icon:       data.Icon,
			Created_at: data.CreatedAt,
			Updated_at: data.UpdatedAt,
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))

	}
}
