package top

import (
	"time"

	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/top"
)

func (s *service) Total(ctx core.Context, where *Where) (total int64, err error) {

	qb := entity.NewQueryBuilder()
	if where.Name != "" {
		qb.WhereName(mysql.LikePredicate, "%"+where.Name+"%")
	}
	if where.Updated_at != (time.Time{}) {
		qb.WhereUpdatedAt(mysql.EqualPredicate, where.Updated_at)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	total, err = qb.
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
