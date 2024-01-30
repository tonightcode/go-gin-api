package singer

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/singer"
)

func (s *service) Detail(ctx core.Context, where *Where) (data *entity.Singer, err error) {

	qb := entity.NewQueryBuilder()
	if where.Singerid != 0 {
		qb.WhereSingerid(mysql.EqualPredicate, where.Singerid)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	data, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
