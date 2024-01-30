package person

import (
	"net/http"
	"strconv"
	"time"

	"go-gin-api/configs"
	"go-gin-api/internal/pkg/core"
	service "go-gin-api/internal/services/person"
)

type listData struct {
	Id         int32     `json:"id"`         // ID
	Username   string    `json:"username"`   // username
	Intro      string    `json:"intro"`      // intro
	Icon       string    `json:"icon"`       // icon
	Created_at time.Time `json:"created_at"` // create time
	Updated_at time.Time `json:"updated_at"` // update time
}

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
		username := c.Query("username")
		page_str := c.Query("page")
		limit_str := c.Query("limit")
		where := service.Where{
			Username: username,
		}
		if page_str == "" {
			where.Page = 1
		} else {
			where.Page, _ = strconv.Atoi(page_str)
		}
		if limit_str == "" {
			where.Limit = configs.ApiLimit
		} else {
			where.Limit, _ = strconv.Atoi(limit_str)
		}
		list, _ := h.service.List(c, &where)
		res := new(listResponse)
		res.List = make([]listData, len(list))
		for k, v := range list {
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
