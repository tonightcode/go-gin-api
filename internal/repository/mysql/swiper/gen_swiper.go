///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package swiper

import (
	"fmt"
	"time"

	"go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Swiper {
	return new(Swiper)
}

func NewQueryBuilder() *swiperQueryBuilder {
	return new(swiperQueryBuilder)
}

func (t *Swiper) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type swiperQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *swiperQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *swiperQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Swiper{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *swiperQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Swiper{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *swiperQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Swiper{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *swiperQueryBuilder) First(db *gorm.DB) (*Swiper, error) {
	ret := &Swiper{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *swiperQueryBuilder) QueryOne(db *gorm.DB) (*Swiper, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *swiperQueryBuilder) QueryAll(db *gorm.DB) ([]*Swiper, error) {
	var ret []*Swiper
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *swiperQueryBuilder) Limit(limit int) *swiperQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *swiperQueryBuilder) Offset(offset int) *swiperQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *swiperQueryBuilder) WhereId(p mysql.Predicate, value int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereIdIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereIdNotIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderById(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereTitle(p mysql.Predicate, value string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereTitleIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereTitleNotIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "title", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByTitle(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "title "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereUrl(p mysql.Predicate, value string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUrlIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUrlNotIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByUrl(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "url "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereRefType(p mysql.Predicate, value int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_type", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereRefTypeIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_type", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereRefTypeNotIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByRefType(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ref_type "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereRefId(p mysql.Predicate, value int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_id", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereRefIdIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_id", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereRefIdNotIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "ref_id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByRefId(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "ref_id "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereIsDeletedIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereIsDeletedNotIn(value []int32) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByIsDeleted(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedAtIn(value []time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByCreatedAt(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedUserIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereCreatedUserNotIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByCreatedUser(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedAtIn(value []time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByUpdatedAt(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedUserIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) WhereUpdatedUserNotIn(value []string) *swiperQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *swiperQueryBuilder) OrderByUpdatedUser(asc bool) *swiperQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
