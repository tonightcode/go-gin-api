package swiper

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/swiper"
)

func (s *service) Total(ctx core.Context, where *Where) (total int64, err error) {

	qb := entity.NewQueryBuilder()
	if where.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, where.Id)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	total, err = qb.
		Count(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return 0, err
	}
	return
}
