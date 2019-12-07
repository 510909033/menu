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
	applog.LogInfo.Printf("init")
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
	M         interface{}
	V         reflect.Value
	TableName string
	DBName    string
	Def       map[string]string // [user_id]UserId
}

//func (menu *ModelMethod) GetDBName() string {
//	return menu.M.DBName
//}
//func (menu *ModelMethod) GetTableName() string {
//	return menu.M.TableName
//}

func (modelMethod *ModelMethod) Load() (err error) {

	menu := (*modelMethod).M.(Model)
	//	if menu.Id < 1 {
	//		return nil
	//	}

	fl := getFieldList(menu)

	fieldStrList := make([]string, 0)
	prepareList := make([]string, 0)
	dataInterface := make([]interface{}, len(fl))

	for _, f := range fl {
		fieldStrList = append(fieldStrList, f.ColumnName)
		prepareList = append(prepareList, "?")

	}

	query := fmt.Sprintf("select %s from %s where id=%d limit 1",
		strings.Join(fieldStrList, ", "),
		modelMethod.TableName,
		modelMethod.V.Elem().FieldByName("Id").Interface())

	applog.LogInfo.Printf("query=%s", query)

	var sqlRow *sql.Rows
	sqlRow, err = db.Query(query)
	if err != nil {
		applog.LogError.Printf("Query fail, err=%v", err)
		return
	}
	defer sqlRow.Close()

	var columnsList []string

	columnsList, err = sqlRow.Columns()
	if err != nil {
		return
	}

	for k, columnName := range columnsList {
		kk := GetFullName(menu) + "_" + columnName
		structColumnName := ModelDbColumnMapStructField[kk]

		ptr := modelMethod.V.Elem().FieldByName(structColumnName).Addr().Interface()
		dataInterface[k] = ptr
	}

	for sqlRow.Next() {
		err = sqlRow.Scan(dataInterface...)
		if err != nil {
			applog.LogError.Printf("Scan fail, err=%v", err)
			return
		}
	}

	applog.LogInfo.Printf("%v", dataInterface)

	return nil
}

func (modelMethod *ModelMethod) Insert() (err error) {

	menu := modelMethod.M.(Model)
	if modelMethod.V.Elem().FieldByName("CreateTs").Interface().(int64) == 0 {
		create_ts := reflect.ValueOf(time.Now().Unix())
		modelMethod.V.Elem().FieldByName("CreateTs").Set(create_ts)
	}

	if modelMethod.V.Elem().FieldByName("UpdateTs").Interface().(int64) == 0 {
		update_ts := reflect.ValueOf(time.Now().Unix())
		modelMethod.V.Elem().FieldByName("UpdateTs").Set(update_ts)
	}

	fullName := GetFullName(menu)

	fieldList, okFieldList := ModelMap[fullName]
	if !okFieldList {
		panic("fullName get data fail, fullName:" + fullName)
	}

	v1 := make([]string, 0)
	v2 := make([]string, 0)
	v3 := make([]interface{}, 0)

	v := reflect.ValueOf(menu)

	//	v = v.Elem()
	//	t := v.Type().Elem()

	for _, field := range fieldList {
		if field.IsPk {
			continue
		}
		v1 = append(v1, field.ColumnName)
		v2 = append(v2, "?")
		v3 = append(v3, v.Elem().FieldByName(field.StructFieldName).Interface())
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

	modelMethod.V.Elem().FieldByName("Id").Set(reflect.ValueOf(int(id)))

	return nil
}

func (modelMethod *ModelMethod) Query() (ret []interface{}, err error) {

	menu := (*modelMethod).M.(Model)
	//	if menu.Id < 1 {
	//		return nil
	//	}

	fl := getFieldList(menu)

	fieldStrList := make([]string, 0)
	prepareList := make([]string, 0)

	dataInterface := make([]interface{}, len(fl))

	for _, f := range fl {
		fieldStrList = append(fieldStrList, f.ColumnName)
		prepareList = append(prepareList, "?")

	}

	query := fmt.Sprintf("select %s from %s order by id desc",
		strings.Join(fieldStrList, ", "),
		modelMethod.TableName)

	applog.LogInfo.Printf("query=%s", query)

	var sqlRow *sql.Rows

	sqlRow, err = db.Query(query)
	if err != nil {
		applog.LogError.Printf("Query fail, err=%v", err)
		return
	}
	defer sqlRow.Close()

	var columnsList []string

	columnsList, err = sqlRow.Columns()
	if err != nil {
		return
	}

	for k, columnName := range columnsList {
		kk := GetFullName(menu) + "_" + columnName
		structColumnName := ModelDbColumnMapStructField[kk]

		ptr := modelMethod.V.Elem().FieldByName(structColumnName).Addr().Interface()
		dataInterface[k] = ptr
	}

	for sqlRow.Next() {
		//		copyDataInterface :=dataInterface
		err = sqlRow.Scan(dataInterface...)
		if err != nil {
			applog.LogError.Printf("Scan fail, err=%v", err)
			return
		}

		a := modelMethod.V.Elem().Interface()

		ret = append(ret, a)
	}

	applog.LogInfo.Printf("%v", dataInterface)

	return
}
