package place

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/place"
)

type PlaceData struct {
	Name    string
	Content string
	Page    int
	Limit   int
}

func (s *service) List(ctx core.Context, placeData *PlaceData) (listData []*place.Place, err error) {

	qb := place.NewQueryBuilder()
	if placeData.Name != "" {
		qb.WhereName(mysql.LikePredicate, "%"+placeData.Name+"%")
	}
	if placeData.Content != "" {
		qb.WhereContent(mysql.LikePredicate, "%"+placeData.Content+"%")
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	qb.Offset((placeData.Page - 1) * placeData.Limit)
	qb.Limit(placeData.Limit)

	listData, err = qb.
		QueryAll(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
