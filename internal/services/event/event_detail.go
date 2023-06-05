package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

type EventParams struct {
	Id int32
}

func (s *service) Detail(ctx core.Context, eventParams *EventParams) (personData *event.Event, err error) {

	qb := event.NewQueryBuilder()
	if eventParams.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, eventParams.Id)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	personData, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
