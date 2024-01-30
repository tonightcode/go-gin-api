package song

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/song"
)

func (s *service) Detail(ctx core.Context, where *Where) (data *entity.Song, err error) {

	qb := entity.NewQueryBuilder()
	if where.Songid != 0 {
		qb.WhereSongid(mysql.EqualPredicate, where.Songid)
	}
	qb.WhereIsDeleted(mysql.EqualPredicate, -1)
	data, err = qb.
		QueryOne(s.db.GetDbR().WithContext(ctx.RequestContext()))
	if err != nil {
		return nil, err
	}
	return
}
