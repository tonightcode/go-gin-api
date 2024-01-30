package singer

import "time"

// Singer 歌手
//
//go:generate gormgen -structs Singer -input .
type Singer struct {
	Singerid    int32     // 主键
	Name        string    // 名称
	Head        string    // 头像
	IsDeleted   int32     // 是否删除 1:是 0 隐藏 -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
