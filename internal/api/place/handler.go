package place

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/place"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 事件列表
	// @Tags API.event
	// @Router /api/event [get]
	List() core.HandlerFunc

	// Detail 事件详情
	// @Tags API.event
	// @Router /api/event/{id} [get]
	Detail() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	db           mysql.Repo
	cache        redis.Repo
	hashids      hash.Hash
	placeService place.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:       logger,
		db:           db,
		cache:        cache,
		hashids:      hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		placeService: place.New(db, cache),
	}
}

func (h *handler) i() {}
