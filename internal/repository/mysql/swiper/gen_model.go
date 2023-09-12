package swiper

import "time"

// Swiper 轮播图
//
//go:generate gormgen -structs Swiper -input .
type Swiper struct {
	Id          int32     // 主键
	Title       string    // 标题
	Url         string    // 地址
	RefType     int32     // 类型 1事件 2地点 3人物 4购物 5自定义
	RefId       int32     // 引用的ID
	IsDeleted   int32     // 是否删除 1:是  -1:否
	CreatedAt   time.Time `gorm:"time"` // 创建时间
	CreatedUser string    // 创建人
	UpdatedAt   time.Time `gorm:"time"` // 更新时间
	UpdatedUser string    // 更新人
}
