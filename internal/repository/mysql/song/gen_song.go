///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package song

import (
	"fmt"
	"time"

	"go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Song {
	return new(Song)
}

func NewQueryBuilder() *songQueryBuilder {
	return new(songQueryBuilder)
}

func (t *Song) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Songid, nil
}

type songQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *songQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *songQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Song{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *songQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Song{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *songQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Song{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *songQueryBuilder) First(db *gorm.DB) (*Song, error) {
	ret := &Song{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *songQueryBuilder) QueryOne(db *gorm.DB) (*Song, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *songQueryBuilder) QueryAll(db *gorm.DB) ([]*Song, error) {
	var ret []*Song
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *songQueryBuilder) Limit(limit int) *songQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *songQueryBuilder) Offset(offset int) *songQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *songQueryBuilder) WhereSongid(p mysql.Predicate, value int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "songid", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereSongidIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "songid", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereSongidNotIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "songid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderBySongid(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "songid "+order)
	return qb
}

func (qb *songQueryBuilder) WhereName(p mysql.Predicate, value string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereNameIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereNameNotIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByName(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *songQueryBuilder) WhereUrl(p mysql.Predicate, value string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUrlIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUrlNotIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "url", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByUrl(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "url "+order)
	return qb
}

func (qb *songQueryBuilder) WhereType(p mysql.Predicate, value int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereTypeIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereTypeNotIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "type", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByType(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "type "+order)
	return qb
}

func (qb *songQueryBuilder) WhereSingerid(p mysql.Predicate, value int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereSingeridIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereSingeridNotIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "singerid", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderBySingerid(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "singerid "+order)
	return qb
}

func (qb *songQueryBuilder) WhereLyric(p mysql.Predicate, value string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lyric", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereLyricIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lyric", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereLyricNotIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "lyric", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByLyric(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "lyric "+order)
	return qb
}

func (qb *songQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereIsDeletedIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereIsDeletedNotIn(value []int32) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByIsDeleted(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *songQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereCreatedAtIn(value []time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByCreatedAt(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *songQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereCreatedUserIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereCreatedUserNotIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByCreatedUser(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedAtIn(value []time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByUpdatedAt(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedUserIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) WhereUpdatedUserNotIn(value []string) *songQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *songQueryBuilder) OrderByUpdatedUser(asc bool) *songQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
