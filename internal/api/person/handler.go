package person

import (
	"go-gin-api/configs"
	"go-gin-api/internal/pkg/core"
	"go-gin-api/internal/repository/mysql"
	"go-gin-api/internal/repository/redis"
	service "go-gin-api/internal/services/person"
	"go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// List 事件列表
	// @Tags API.person
	// @Router /api/person [get]
	List() core.HandlerFunc

	// Detail 事件详情
	// @Tags API.person
	// @Router /api/person/{id} [get]
	Detail() core.HandlerFunc
}

type handler struct {
	logger  *zap.Logger
	db      mysql.Repo
	cache   redis.Repo
	hashids hash.Hash
	service service.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:  logger,
		db:      db,
		cache:   cache,
		hashids: hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		service: service.New(db, cache),
	}
}

func (h *handler) i() {}
