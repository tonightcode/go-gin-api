package person

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/person"
)

func (s *service) List(ctx core.Context) (listData []*person.Person, err error) {

	qb := person.NewQueryBuilder()
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
