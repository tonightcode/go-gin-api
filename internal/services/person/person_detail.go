package person

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/person"
)

type PersonParams struct {
	Id int32
}

func (s *service) Detail(ctx core.Context, personParams *PersonParams) (personData *person.Person, err error) {

	qb := person.NewQueryBuilder()
	if personParams.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, personParams.Id)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	personData, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
