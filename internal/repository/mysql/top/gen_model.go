package top

import (
	"time"

	"go-gin-api/internal/repository/mysql/song"
)

// Top 排行榜
//
//go:generate gormgen -structs Top -input .
type Top struct {
	Topid       int32       // 主键
	Name        string      // 名称
	Cover       string      // 封面
	IsDeleted   int32       // 是否删除 1:是 0 隐藏 -1:否
	CreatedAt   time.Time   `gorm:"time"` // 创建时间
	CreatedUser string      // 创建人
	UpdatedAt   time.Time   `gorm:"time"` // 更新时间
	UpdatedUser string      // 更新人
	Songs       []song.Song `gorm:"many2many:top_song;foreignKey:Topid;joinForeignKey:Topid;joinReferences:Songid"` // 歌曲
}
