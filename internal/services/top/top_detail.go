package top

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/top"
)

func (s *service) Detail(ctx core.Context, where *Where) (data *entity.Top, err error) {

	qb := entity.NewQueryBuilder()
	if where.Topid != 0 {
		qb.WhereTopid(mysql.EqualPredicate, where.Topid)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	data, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
