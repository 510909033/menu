package bgf_bo

import (
	"fmt"
	"reflect"
	"sync"

	"baotian0506.com/app/menu/applog"
)

var ModelMap = make(map[string][]*Field)
var muModelLock sync.Mutex

func Register(m Model) {
	registerModel(m)
}

type Field struct {
	IsPk            bool
	ColumnName      string
	StructFieldName string
}

func registerModel(m Model) {

	t := reflect.TypeOf(m)

	t = t.Elem()

	fieldList := make([]*Field, 0)

	for i := 0; i < t.NumField(); i++ {
		field := &Field{}
		field.IsPk = false

		tStructField := t.Field(i)
		tag := tStructField.Tag
		if val, ok := tag.Lookup("column_name"); ok {
			field.ColumnName = val
		}
		if field.ColumnName == "" {
			continue
		}
		field.StructFieldName = tStructField.Name

		//		fmt.Println("name:" + tStructField.Name)
		//		fmt.Println("PkgPath:" + tStructField.PkgPath)
		//		fmt.Println("Kind:", tStructField.Type.Kind())

		//		fmt.Println("model:" + tag.Get("model"))
		//		fmt.Println("json:" + tag.Get("json"))

		if _, ok := tag.Lookup("pk"); ok {
			field.IsPk = true
		}

		fieldList = append(fieldList, field)
	}

	setModelMap(m, fieldList)

}

func setModelMap(m Model, f []*Field) {
	muModelLock.Lock()
	defer muModelLock.Unlock()
	fullName := GetFullName(m)
	ModelMap[fullName] = f
}

func getFieldList(m Model) []*Field {
	muModelLock.Lock()
	defer muModelLock.Unlock()
	fullName := GetFullName(m)
	return ModelMap[fullName]
}

func GetFullName(m Model) string {
	applog.LogError.Printf("%s_%s\n", m.GetDBName(), m.GetTableName())

	return fmt.Sprintf("%s_%s", m.GetDBName(), m.GetTableName())
}
