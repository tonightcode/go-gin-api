package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	entity "github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

type Where struct {
	Id         int32
	Title      string
	Happend_at string
	Page       int
	Limit      int
}

func (s *service) List(ctx core.Context, where *Where) (list []*entity.Event, err error) {

	qb := entity.NewQueryBuilder()
	if where.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, "%"+where.Title+"%")
	}
	if where.Happend_at != "" {
		qb.WhereHappendAt(mysql.EqualPredicate, where.Happend_at)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.Offset((where.Page - 1) * where.Limit)
	qb.Limit(where.Limit)

	list, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
