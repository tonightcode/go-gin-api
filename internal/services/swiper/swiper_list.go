package swiper

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/swiper"
)

type Where struct {
	Id    int32
	Page  int
	Limit int
}

func (s *service) List(ctx core.Context, where *Where) (list []*entity.Swiper, err error) {

	qb := entity.NewQueryBuilder()
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
