package bgf_bo

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"baotian0506.com/app/menu/applog"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:root@/menu")
	fmt.Println("init")
	if err != nil {
		applog.LogError.Printf("open db fail, %v", err)
		panic(err)
	}
}

type Model interface {
	GetTableName() string
	GetDBName() string
}

type ModelMethod struct {
	M interface{}
}

//func (menu *ModelMethod) GetDBName() string {
//	return menu.M.DBName
//}
//func (menu *ModelMethod) GetTableName() string {
//	return menu.M.TableName
//}

func (modelMethod *ModelMethod) Load() (err error) {

	menu := modelMethod.M.(Model)
	//	if menu.Id < 1 {
	//		return nil
	//	}

	fl := getFieldList(menu)

	fieldStrList := make([]string, 0)
	prepareList := make([]string, 0)
	dataInterface := make([]interface{}, len(fl))

	fmt.Println(len(fl))

	for _, f := range fl {
		fieldStrList = append(fieldStrList, f.ColumnName)
		prepareList = append(prepareList, "?")

	}

	query := fmt.Sprintf("select %s from %s where id=%d limit 1",
		strings.Join(fieldStrList, ", "),
		menu.TableName,
		menu.Id)
	sqlRow := db.QueryRow(query)

	err = sqlRow.Scan(dataInterface...)
	fmt.Println(dataInterface)
	if err != nil {
		return err
	}

	return nil
}

func (modelMethod *ModelMethod) Insert() (err error) {

	menu := modelMethod.M
	if menu.CreateTs == 0 {
		menu.CreateTs = time.Now().Unix()
	}

	if menu.UpdateTs == 0 {
		menu.UpdateTs = time.Now().Unix()
	}

	fullName := GetFullName(menu)

	fieldList, okFieldList := ModelMap[fullName]
	if !okFieldList {
		panic("fullName get data fail, fullName:" + fullName)
	}

	v1 := make([]string, 0)
	v2 := make([]string, 0)
	v3 := make([]interface{}, 0)

	//	fmt.Println(menu)
	v := reflect.ValueOf(*menu)

	//	v = v.Elem()
	//	t := v.Type().Elem()

	for _, field := range fieldList {
		if field.IsPk {
			continue
		}
		v1 = append(v1, field.ColumnName)
		v2 = append(v2, "?")
		fmt.Println("field.StructFieldName" + field.StructFieldName)
		fmt.Println(v.FieldByName(field.StructFieldName))
		v3 = append(v3, v.FieldByName(field.StructFieldName).Interface())
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) values(%s)",
		menu.GetTableName(), strings.Join(v1, ", "), strings.Join(v2, ", "))

	sqlResult, err := db.Exec(query, v3...)

	if err != nil {
		applog.LogError.Printf("err=%v", err)
		return err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		applog.LogError.Printf("err=%v", err)
	}
	menu.Id = int(id)

	return nil
}
