package swiper

import (
	"net/http"
	"strconv"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	service "github.com/xinliangnote/go-gin-api/internal/services/swiper"
)

type listData struct {
	Id         int32  `json:"id"`         // ID
	Title      string `json:"title"`      // title
	Url        string `json:"url"`        // url
	Ref_type   int32  `json:"ref_type"`   // ref_type
	Ref_id     int32  `json:"ref_id"`     // ref_id
	Created_at string `json:"created_at"` // create time
	Updated_at string `json:"updated_at"` // update time
}

type listResponse struct {
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	List  []listData `json:"list"`
}

// List 事件列表
// @Summary 事件列表
// @Description 事件列表
// @Tags API.event
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body listRequest true "请求信息"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/event [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		page_str := c.Query("page")
		limit_str := c.Query("limit")
		where := service.Where{}
		if page_str == "" {
			where.Page = 1
		} else {
			where.Page, _ = strconv.Atoi(page_str)
		}
		if limit_str == "" {
			where.Limit = 5
		} else {
			where.Limit, _ = strconv.Atoi(limit_str)
		}
		list, err := h.service.List(c, &where)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res := new(listResponse)
		res.Page = where.Page
		res.Total, _ = h.service.Total(c, &where)
		res.List = make([]listData, len(list))
		for k, v := range list {
			data := listData{
				Id:         v.Id,
				Title:      v.Title,
				Url:        v.Url,
				Ref_type:   v.RefType,
				Ref_id:     v.RefId,
				Created_at: v.CreatedAt.Format("2006-01-02 15:04:05"),
				Updated_at: v.UpdatedAt.Format("2006-01-02 15:04:05"),
			}

			res.List[k] = data
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))
	}
}
