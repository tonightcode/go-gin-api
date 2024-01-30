///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package place

import (
	"fmt"
	"time"

	"go-gin-api/internal/repository/mysql"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Place {
	return new(Place)
}

func NewQueryBuilder() *placeQueryBuilder {
	return new(placeQueryBuilder)
}

func (t *Place) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type placeQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *placeQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
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

func (qb *placeQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Place{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *placeQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Place{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *placeQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Place{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *placeQueryBuilder) First(db *gorm.DB) (*Place, error) {
	ret := &Place{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *placeQueryBuilder) QueryOne(db *gorm.DB) (*Place, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *placeQueryBuilder) QueryAll(db *gorm.DB) ([]*Place, error) {
	var ret []*Place
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *placeQueryBuilder) Limit(limit int) *placeQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *placeQueryBuilder) Offset(offset int) *placeQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *placeQueryBuilder) WhereId(p mysql.Predicate, value int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereIdIn(value []int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereIdNotIn(value []int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderById(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereName(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereNameIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereNameNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "name", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByName(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "name "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereContent(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereContentIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereContentNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "content", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByContent(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "content "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereProvince(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereProvinceIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereProvinceNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "province", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByProvince(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "province "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereCity(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCityIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCityNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "city", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByCity(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "city "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereCounty(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "county", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCountyIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "county", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCountyNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "county", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByCounty(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "county "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereImg(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "img", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereImgIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "img", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereImgNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "img", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByImg(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "img "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereAddress(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereAddressIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereAddressNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "address", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByAddress(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "address "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereIsDeleted(p mysql.Predicate, value int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereIsDeletedIn(value []int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereIsDeletedNotIn(value []int32) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByIsDeleted(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedAt(p mysql.Predicate, value time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedAtIn(value []time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByCreatedAt(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedUser(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedUserIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereCreatedUserNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByCreatedUser(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedAt(p mysql.Predicate, value time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedAtIn(value []time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByUpdatedAt(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedUser(p mysql.Predicate, value string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedUserIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) WhereUpdatedUserNotIn(value []string) *placeQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *placeQueryBuilder) OrderByUpdatedUser(asc bool) *placeQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}
