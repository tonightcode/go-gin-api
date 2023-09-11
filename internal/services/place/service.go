package place

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	entity "github.com/xinliangnote/go-gin-api/internal/repository/mysql/place"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context, where *Where) (list []*entity.Place, err error)
	Detail(ctx core.Context, where *Where) (data *entity.Place, err error)
}

type service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:    db,
		cache: cache,
	}
}

func (s *service) i() {}