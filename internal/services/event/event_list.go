package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

type EventData struct {
	Pid int32 // 父类ID
}

func (s *service) List(ctx core.Context, eventData *EventData) (listData []*event.Event, err error) {

	qb := event.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
