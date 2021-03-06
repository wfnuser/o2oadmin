package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type BeTransactionHeader struct {
	Id                  int       `orm:"column(id);auto"`
	MerchantId          int       `orm:"column(merchant_id)"`
	SupplierId          int       `orm:"column(supplier_id)"`
	TransactionType     string    `orm:"column(transaction_type);size(3);null"`
	OrderNumber         string    `orm:"column(order_number);size(45);null"`
	OrderAmount         float32   `orm:"column(order_amount);null"`
	ActualAmount        float32   `orm:"column(actual_amount);null"`
	OrderTime           time.Time `orm:"column(order_time);type(datetime);null"`
	ActualReceiveTime   time.Time `orm:"column(actual_receive_time);type(datetime);null"`
	ExpectedReceiveTime time.Time `orm:"column(expected_receive_time);type(datetime);null"`
	AveragePrice        float32   `orm:"column(average_price);null"`
	CreateUser          int       `orm:"column(create_user);null"`
	UpdateUser          int       `orm:"column(update_user);null"`
	CreateTime          time.Time `orm:"column(create_time);type(timestamp);null"`
	UpdateTime          time.Time `orm:"column(update_time);type(timestamp);null;auto_now"`
	TrasactionStatus    string    `orm:"column(trasaction_status);size(3);null"`
}

func init() {
	orm.RegisterModel(new(BeTransactionHeader))
}

// AddBeTransactionHeader insert a new BeTransactionHeader into database and returns
// last inserted Id on success.
func AddBeTransactionHeader(m *BeTransactionHeader) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBeTransactionHeaderById retrieves BeTransactionHeader by Id. Returns error if
// Id doesn't exist
func GetBeTransactionHeaderById(id int) (v *BeTransactionHeader, err error) {
	o := orm.NewOrm()
	v = &BeTransactionHeader{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllBeTransactionHeader retrieves all BeTransactionHeader matches certain condition. Returns empty list if
// no records exist
func GetAllBeTransactionHeader(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BeTransactionHeader))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []BeTransactionHeader
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateBeTransactionHeader updates BeTransactionHeader by Id and returns error if
// the record to be updated doesn't exist
func UpdateBeTransactionHeaderById(m *BeTransactionHeader) (err error) {
	o := orm.NewOrm()
	v := BeTransactionHeader{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBeTransactionHeader deletes BeTransactionHeader by Id and returns error if
// the record to be deleted doesn't exist
func DeleteBeTransactionHeader(id int) (err error) {
	o := orm.NewOrm()
	v := BeTransactionHeader{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BeTransactionHeader{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
