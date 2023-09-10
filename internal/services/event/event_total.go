package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

func (s *service) Total(ctx core.Context, eventData *EventData) (total int64, err error) {

	qb := event.NewQueryBuilder()
	if eventData.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, "%"+eventData.Title+"%")
	}
	if eventData.Happend_at != "" {
		qb.WhereHappendAt(mysql.EqualPredicate, eventData.Happend_at)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	total, err = qb.
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
