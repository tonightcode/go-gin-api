package person

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
)

type listData struct {
	Id         int32     `json:"id"`         // ID
	Username   string    `json:"username"`   // username
	Intro      string    `json:"intro"`      // intro
	Icon       string    `json:"icon"`       // icon
	Created_at time.Time `json:"created_at"` // create time
	Updated_at time.Time `json:"updated_at"` // update time
}

type listRequest struct{}

type listResponse struct {
	List []listData `json:"list"`
}

// List 人物列表
// @Summary 人物列表
// @Description 人物列表
// @Tags API.person
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/person [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		res := new(listResponse)
		resListData, err := h.personService.List(c)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res.List = make([]listData, len(resListData))
		for k, v := range resListData {
			data := listData{
				Id:         v.Id,
				Username:   v.Username,
				Intro:      v.Intro,
				Icon:       v.Icon,
				Created_at: v.CreatedAt,
				Updated_at: v.UpdatedAt,
			}

			res.List[k] = data
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))
	}
}
