package place

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/place"
)

type PlaceParams struct {
	Id int32
}

func (s *service) Detail(ctx core.Context, placeParams *PlaceParams) (placeData *place.Place, err error) {

	qb := place.NewQueryBuilder()
	if placeParams.Id != 0 {
		qb.WhereId(mysql.EqualPredicate, placeParams.Id)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	placeData, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
