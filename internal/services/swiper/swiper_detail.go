package swiper

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/swiper"
)

func (s *service) Detail(ctx core.Context, where *Where) (data *entity.Swiper, err error) {

	qb := entity.NewQueryBuilder()
	if where.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, where.Id)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	data, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
