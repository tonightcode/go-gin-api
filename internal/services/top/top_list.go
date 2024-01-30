package top

import (
	"time"

	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/top"
)

type Where struct {
	Topid      int32
	Name       string
	Updated_at time.Time
	Page       int
	Limit      int
}

func (s *service) List(ctx core.Context, where *Where) (list []*entity.Top, err error) {

	qb := entity.NewQueryBuilder()
	if where.Name != "" {
		qb.WhereName(mysql.LikePredicate, "%"+where.Name+"%")
	}
	if where.Updated_at != (time.Time{}) {
		qb.WhereUpdatedAt(mysql.EqualPredicate, where.Updated_at)
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
