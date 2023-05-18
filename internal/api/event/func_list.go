package event

import (
	"net/http"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/services/event"
)

type listData struct {
	Id         int32     `json:"id"`         // ID
	Title      string    `json:"title"`      // title
	Content    string    `json:"content"`    // content
	Cover      string    `json:"cover"`      // cover
	Created_at time.Time `json:"created_at"` // create time
	Updated_at time.Time `json:"updated_at"` // update time
}
type listRequest struct{}

type listResponse struct {
	List []listData `json:"list"`
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
		res := new(listResponse)
		resListData, err := h.eventService.List(c, new(event.EventData))
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
				Title:      v.Title,
				Content:    v.Content,
				Cover:      v.Cover,
				Created_at: v.CreatedAt,
				Updated_at: v.UpdatedAt,
			}

			res.List[k] = data
		}
		c.Payload(res)
	}
}
