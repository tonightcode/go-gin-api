package singer

import (
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	entity "go-gin-api/internal/repository/mysql/singer"
	"go-gin-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context, where *Where) (list []*entity.Singer, err error)
	Total(ctx core.Context, where *Where) (total int64, err error)
	Detail(ctx core.Context, where *Where) (data *entity.Singer, err error)
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
