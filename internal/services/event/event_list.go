package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

type EventData struct {
	Title      string
	Happend_at string
	Page       int
	Limit      int
}

func (s *service) List(ctx core.Context, eventData *EventData) (listData []*event.Event, err error) {

	qb := event.NewQueryBuilder()
	if eventData.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, "%"+eventData.Title+"%")
	}
	if eventData.Happend_at != "" {
		qb.WhereHappendAt(mysql.EqualPredicate, eventData.Happend_at)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.Offset((eventData.Page - 1) * eventData.Limit)
	qb.Limit(eventData.Limit)

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
