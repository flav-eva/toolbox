package mysqldialector

import (
	"database/sql/driver"
	"reflect"

	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Dialector struct {
	mysql.Dialector
}

func Open(dsn string) gorm.Dialector {
	return &Dialector{Dialector: mysql.Dialector{Config: &mysql.Config{DSN: dsn}}}
}

func (u Dialector) Explain(sql string, vars ...interface{}) string {
	var convertParams func(interface{}, int)
	var cvars = make([]interface{}, len(vars))
	convertParams = func(v interface{}, idx int) {
		switch v := v.(type) {
		case driver.Valuer:
			reflectValue := reflect.ValueOf(v)
			if v != nil && reflectValue.IsValid() && ((reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil()) || reflectValue.Kind() != reflect.Ptr) {
				r, _ := v.Value()
				convertParams(r, idx)
			} else {
				cvars[idx] = v
			}
		case []byte:
			id := &uuid.UUID{}
			if err := id.UnmarshalBinary(v); err == nil {
				cvars[idx] = "UUID:" + id.String()
			} else {
				cvars[idx] = v
			}
		default:
			cvars[idx] = v
		}
	}
	for idx, v := range vars {
		convertParams(v, idx)
	}
	return u.Dialector.Explain(sql, cvars...)
}
