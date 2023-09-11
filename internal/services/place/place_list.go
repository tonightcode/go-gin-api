package place

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	entity "github.com/xinliangnote/go-gin-api/internal/repository/mysql/place"
)

type Where struct {
	Id      int32
	Name    string
	Content string
	Page    int
	Limit   int
}

func (s *service) List(ctx core.Context, where *Where) (list []*entity.Place, err error) {

	qb := entity.NewQueryBuilder()
	if where.Name != "" {
		qb.WhereName(mysql.LikePredicate, "%"+where.Name+"%")
	}
	if where.Content != "" {
		qb.WhereContent(mysql.LikePredicate, "%"+where.Content+"%")
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
