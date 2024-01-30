package top_song

import "time"

// TopSong 排行榜-歌曲
//
//go:generate gormgen -structs TopSong -input .
type TopSong struct {
	Topid       int32     `gorm:"primaryKey"` // 排行榜ID
	Songid      int32     `gorm:"primaryKey"` // 声音ID
	IsDeleted   int32     // 是否删除 1:是 0 隐藏 -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
