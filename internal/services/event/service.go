package event

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql/event"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
)

var _ Service = (*service)(nil)

type Service interface {
	i()

	List(ctx core.Context, eventData *EventData) (listData []*event.Event, err error)
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
