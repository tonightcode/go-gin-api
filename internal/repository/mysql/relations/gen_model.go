package relations

import "time"

// Relations 事件-人物-地点关系表
//
//go:generate gormgen -structs Relations -input .
type Relations struct {
	Id         int32     //
	EventId    int32     //
	PlaceId    int32     //
	PersonId   int32     //
	Createtime time.Time `gorm:"time"` //
	Updatetime time.Time `gorm:"time"` //
}
