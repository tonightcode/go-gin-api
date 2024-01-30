package top

import (
	"net/http"
	"strconv"
	"time"

	"go-gin-api/configs"
	"go-gin-api/internal/code"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql/song"
	service "go-gin-api/internal/services/top"
)

type listData struct {
	Topid      int32       `json:"topid"`      // topid
	Name       string      `json:"name"`       // name
	Cover      string      `json:"cover"`      // cover
	Created_at string      `json:"created_at"` // create time
	Updated_at string      `json:"updated_at"` // update time
	Songs      []song.Song `json:"songs"`
}

type listResponse struct {
	Total int64      `json:"total"`
	Page  int        `json:"page"`
	List  []listData `json:"list"`
}

func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		name := c.Query("name")
		updated_at_str := c.Query("Updated_at")
		updated_at, _ := time.Parse(time.RFC3339, updated_at_str)
		page_str := c.Query("page")
		limit_str := c.Query("limit")
		where := service.Where{
			Name:       name,
			Updated_at: updated_at,
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
			data := listData{
				Topid:      v.Topid,
				Name:       v.Name,
				Cover:      v.Cover,
				Created_at: v.CreatedAt.Format("2006-01-02 15:04:05"),
				Updated_at: v.UpdatedAt.Format("2006-01-02 15:04:05"),
				Songs:      v.Songs,
			}

			res.List[k] = data
		}
		c.Payload(core.ApiSuccess(
			http.StatusOK,
			"Success",
			res))
	}
}
