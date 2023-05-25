package place

import "time"

// Place 地点表
//
//go:generate gormgen -structs Place -input .
type Place struct {
	Id          int32     // 主键
	Name        string    // 名称
	Content     string    // 详情
	Province    string    // 省份
	City        string    // 城市
	County      string    // 区县
	Img         string    // 图片
	Address     string    // 地址
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
