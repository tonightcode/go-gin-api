package event

import "time"

// Event 事件表
//
//go:generate gormgen -structs Event -input .
type Event struct {
	Id          int32     // 主键
	Title       string    // 标题
	Content     string    // 详情
	Cover       string    //封面
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
