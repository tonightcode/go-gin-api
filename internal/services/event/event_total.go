package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	entity "github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
)

func (s *service) Total(ctx core.Context, where *Where) (total int64, err error) {

	qb := entity.NewQueryBuilder()
	if where.Title != "" {
		qb.WhereTitle(mysql.LikePredicate, "%"+where.Title+"%")
	}
	if where.Happend_at != "" {
		qb.WhereHappendAt(mysql.EqualPredicate, where.Happend_at)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	total, err = qb.
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
