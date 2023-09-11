package person

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	entity "github.com/xinliangnote/go-gin-api/internal/repository/mysql/person"
)

type Where struct {
	Id       int32
	Username string
	Page     int
	Limit    int
}

func (s *service) List(ctx core.Context, where *Where) (listData []*entity.Person, err error) {

	qb := entity.NewQueryBuilder()
	if where.Username != "" {
		qb.WhereUsername(mysql.LikePredicate, "%"+where.Username+"%")
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.Offset((where.Page - 1) * where.Limit)
	qb.Limit(where.Limit)

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
