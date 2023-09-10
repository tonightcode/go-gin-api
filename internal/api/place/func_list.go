package place

import (
	"net/http"
	"strconv"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/place"
)

type listRequest struct {
	name    string
	content string
}

type listData struct {
	Id         int32  `json:"id"`         // ID
	Name       string `json:"name"`       // title
	Content    string `json:"content"`    // content
	Province   string `json:"province"`   // cover
	City       string `json:"city"`       // cover
	County     string `json:"county"`     // cover
	Img        string `json:"img"`        // cover
	Address    string `json:"address"`    // cover
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
		name := c.Query("name")
		content := c.Query("content")
		page_str := c.Query("page")
		limit_str := c.Query("limit")
		placeData := place.PlaceData{
			Name:    name,
			Content: content,
		}
		if page_str == "" {
			placeData.Page = 1
		} else {
			placeData.Page, _ = strconv.Atoi(page_str)
		}
		if limit_str == "" {
			placeData.Limit = 10
		} else {
			placeData.Limit, _ = strconv.Atoi(limit_str)
		}
		list, err := h.placeService.List(c, &placeData)
		if err != nil {
			c.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.MenuListError,
				code.Text(code.MenuListError)).WithError(err),
			)
			return
		}
		res := new(listResponse)
		res.Page = placeData.Page
		res.Total = 0
		res.List = make([]listData, len(list))
		for k, v := range list {
			data := listData{
				Id:         v.Id,
				Name:       v.Name,
				Content:    v.Content,
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
