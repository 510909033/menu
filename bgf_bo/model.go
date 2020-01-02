package bgf_bo

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/config"
	"baotian0506.com/app/menu/pkg/common"
	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}

func init() {
	applog.LogInfo.Printf("init")
	var err error
	var sqlResult sql.Result

	mysqlConfig := config.GetMysqConfig()
	password := mysqlConfig.Password

	db, err = sql.Open("mysql", "root:"+password+"@/menu")
	if err != nil {
		applog.LogError.Printf("open db fail, %v", err)
		panic(err)
	}
	sqlResult, err = db.Exec("set names utf8mb4")
	_ = sqlResult
	if err != nil {
		applog.LogError.Printf("set names utf8mb4 fail, %v", err)
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
	IsNewRow  bool
}

//func (menu *ModelMethod) GetDBName() string {
//	return menu.M.DBName
//}
//func (menu *ModelMethod) GetTableName() string {
//	return menu.M.TableName
//}

func (modelMethod *ModelMethod) Load() (err error) {
	modelMethod.IsNewRow = true
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
		modelMethod.IsNewRow = false
	}

	//	applog.LogInfo.Printf("%v", dataInterface)

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

func (modelMethod *ModelMethod) Query(where string, whereValue []interface{}, limit PageLimit) (ret []interface{}, err error) {

	menu := (*modelMethod).M.(Model)

	fl := getFieldList(menu)

	fieldStrList := make([]string, 0)
	prepareList := make([]string, 0)

	dataInterface := make([]interface{}, len(fl))

	for _, f := range fl {
		fieldStrList = append(fieldStrList, f.ColumnName)
		prepareList = append(prepareList, "?")

	}

	if where != "" {
		where = "where " + where
	}

	if limit.Page == 0 {
		limit.Page = 1
	}
	if limit.Limit < 1 || limit.Limit > 1000 {
		limit.Limit = 10
	}

	whereValue = append(whereValue, (limit.Page-1)*limit.Limit)
	whereValue = append(whereValue, limit.Limit)

	query := fmt.Sprintf("select %s from %s %s order by id desc limit ? , ?",
		strings.Join(fieldStrList, ", "),
		modelMethod.TableName,
		where)

	applog.LogInfo.Printf("query=%s", query)

	var sqlRow *sql.Rows

	sqlRow, err = db.Query(query, whereValue...)

	//debug sql
	debugQueryList := strings.Split(query, "?")

	for k, v := range debugQueryList {
		if k > (len(whereValue) - 1) {
			break
		}
		tmpValue := whereValue[k]
		rt := reflect.TypeOf(tmpValue)
		if rt.Kind().String() == "string" {
			debugQueryList[k] = v + fmt.Sprintf("%s", whereValue[k])
		} else {
			debugQueryList[k] = v + fmt.Sprintf("%d", whereValue[k])
		}

	}
	applog.LogInfo.Printf("debug sql=%s", strings.Join(debugQueryList, " "))

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

		createTs := modelMethod.V.Elem().FieldByName("CreateTs").Interface().(int64)
		createTsFormat := time.Unix(createTs, 0).Format(common.TIME_FORMAT_YMDHIS)

		rv := reflect.ValueOf(createTsFormat)
		modelMethod.V.Elem().FieldByName("CreateTsFormat").Set(rv)
		a := modelMethod.V.Elem().Interface()

		ret = append(ret, a)
	}

	//	applog.LogInfo.Printf("%v", dataInterface)

	return
}

func (modelMethod *ModelMethod) Update() (err error) {
	//create_ts 不更新
	menu := modelMethod.M.(Model)

	update_ts := reflect.ValueOf(time.Now().Unix())
	modelMethod.V.Elem().FieldByName("UpdateTs").Set(update_ts)

	fullName := GetFullName(menu)

	fieldList, okFieldList := ModelMap[fullName]
	if !okFieldList {
		panic("fullName get data fail, fullName:" + fullName)
	}

	v1 := make([]string, 0)      //字段
	v3 := make([]interface{}, 0) //值

	v := reflect.ValueOf(menu)

	//	v = v.Elem()
	//	t := v.Type().Elem()

	for _, field := range fieldList {
		if field.IsPk {
			continue
		}
		if field.ColumnName == "create_ts" {
			continue
		}
		v1 = append(v1, field.ColumnName+" = ?")

		v3 = append(v3, v.Elem().FieldByName(field.StructFieldName).Interface())
	}

	query := fmt.Sprintf("update %s set %s where id = %d limit 1",
		menu.GetTableName(),
		strings.Join(v1, ", "),
		v.Elem().FieldByName("Id").Interface())

	//debug sql
	debugQueryList := strings.Split(query, "?")

	for k, v := range debugQueryList {
		if k > (len(v3) - 1) {
			break
		}
		debugQueryList[k] = v + fmt.Sprintf("%s", v3[k])
	}
	applog.LogInfo.Printf("debug sql=%s", strings.Join(debugQueryList, " "))

	sqlResult, err := db.Exec(query, v3...)

	if err != nil {
		applog.LogError.Printf("err=%v", err)
		return err
	}

	affected, err := sqlResult.RowsAffected()
	if err != nil {
		applog.LogError.Printf("err=%v", err)
		return err
	}

	if affected != 1 {
		err = errors.New("affected !=1")
		applog.LogError.Printf("err=%v", err)
		return err
	}

	return nil
}

func (modelMethod *ModelMethod) Save() (err error) {
	var errRet error
	if modelMethod.IsNewRow {
		if err := modelMethod.Insert(); err != nil {
			errRet = fmt.Errorf("save fail! %w", err)
		}
	} else {
		//update
		if err := modelMethod.Update(); err != nil {
			errRet = fmt.Errorf("Update fail! %w", err)
		}
	}
	return errRet
}

func (modelMethod *ModelMethod) Count(where string, whereValue []interface{}) (cnt int64, err error) {

	if where != "" {
		where = "where " + where
	}

	query := fmt.Sprintf("select count(*) cnt from %s %s  limit 1",
		modelMethod.TableName,
		where)

	applog.LogInfo.Printf("query=%s", query)

	var sqlRow *sql.Rows

	sqlRow, err = db.Query(query, whereValue...)

	//debug sql
	debugQueryList := strings.Split(query, "?")

	for k, v := range debugQueryList {
		if k > (len(whereValue) - 1) {
			break
		}

		tmpValue := whereValue[k]
		rt := reflect.TypeOf(tmpValue)
		if rt.Kind().String() == "string" {
			debugQueryList[k] = v + fmt.Sprintf("%s", whereValue[k])
		} else {
			debugQueryList[k] = v + fmt.Sprintf("%d", whereValue[k])
		}
	}
	applog.LogInfo.Printf("debug sql=%s", strings.Join(debugQueryList, " "))

	if err != nil {
		applog.LogError.Printf("Query fail, err=%v", err)
		return
	}
	defer sqlRow.Close()

	for sqlRow.Next() {

		err = sqlRow.Scan(&cnt)
		if err != nil {
			applog.LogError.Printf("Scan fail, err=%v", err)
			return
		}

	}

	//	applog.LogInfo.Printf("%v", dataInterface)

	return
}
