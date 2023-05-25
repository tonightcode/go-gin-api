package person

import "time"

// Person 人物表
//
//go:generate gormgen -structs Person -input .
type Person struct {
	Id          int32     // 主键
	Username    string    // 名字
	Intro       string    // 详情
	Icon        string    // 头像
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
