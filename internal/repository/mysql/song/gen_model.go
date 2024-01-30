package song

import "time"

// Song 声音
//
//go:generate gormgen -structs Song -input .
type Song struct {
	Songid      int32     // 主键
	Name        string    // 名称
	Url         string    // 地址
	Type        int32     // 1 song 2 sound
	Singerid    int32     // 歌手
	Lyric       string    // 歌词
	IsDeleted   int32     // 是否删除 1:是 0 隐藏 -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
