package event

import (
	"net/http"
	"strconv"
	"time"

	"go-gin-api/configs"
	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	service "go-gin-api/internal/services/event"
)

type listData struct {
	Id         int32  `json:"id"`         // ID
	Title      string `json:"title"`      // title
	Content    string `json:"content"`    // content
	Cover      string `json:"cover"`      // cover
	Happend_at string `json:"happend_at"` // cover
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
		title := c.Query("title")
		happend_at := c.Query("happend_at")
		page_str := c.Query("page")
		limit_str := c.Query("limit")
		where := service.Where{
			Title:      title,
			Happend_at: happend_at,
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
		total, _ := h.service.Total(c, &where)
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
		res.Total = total
		res.List = make([]listData, len(list))
		for k, v := range list {
			t, _ := time.Parse("2006-01-02", v.HappendAt)
			data := listData{
				Id:         v.Id,
				Title:      v.Title,
				Content:    v.Content,
				Cover:      v.Cover,
				Happend_at: t.Format("2006-01-02"),
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
