///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package singer

import (
	"fmt"
	"time"

	"go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Singer {
	return new(Singer)
}

func NewQueryBuilder() *singerQueryBuilder {
	return new(singerQueryBuilder)
}

func (t *Singer) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Singerid, nil
}

type singerQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *singerQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *singerQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Singer{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *singerQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Singer{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *singerQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Singer{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *singerQueryBuilder) First(db *gorm.DB) (*Singer, error) {
	ret := &Singer{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *singerQueryBuilder) QueryOne(db *gorm.DB) (*Singer, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *singerQueryBuilder) QueryAll(db *gorm.DB) ([]*Singer, error) {
	var ret []*Singer
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *singerQueryBuilder) Limit(limit int) *singerQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *singerQueryBuilder) Offset(offset int) *singerQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *singerQueryBuilder) WhereSingerid(p mysql.Predicate, value int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereSingeridIn(value []int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereSingeridNotIn(value []int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderBySingerid(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "singerid "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereName(p mysql.Predicate, value string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereNameIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereNameNotIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByName(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereHead(p mysql.Predicate, value string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "head", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereHeadIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "head", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereHeadNotIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "head", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByHead(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "head "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereIsDeletedIn(value []int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereIsDeletedNotIn(value []int32) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByIsDeleted(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedAtIn(value []time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByCreatedAt(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedUserIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereCreatedUserNotIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByCreatedUser(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedAtIn(value []time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByUpdatedAt(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedUserIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) WhereUpdatedUserNotIn(value []string) *singerQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *singerQueryBuilder) OrderByUpdatedUser(asc bool) *singerQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
