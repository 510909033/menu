package main

import (
	"fmt"
	"reflect"
)

type MenuBO struct {
	Id       int64  `pk:"" column_name:"id" json:"id"`
	UserId   int64  `column_name:"user_id" json:"user_id"`
	Title    string `column_name:"title" json:"title"`
	Extra    string `column_name:"extra" json:"extra"`
	CreateTs int64  `column_name:"create_ts" json:"create_ts"`
	UpdateTs int64  `column_name:"update_ts" json:"update_ts"`
}

func init() {
	fmt.Println("debug.go init")
	model()
}

func model() {

	bo := &MenuBO{}

	t := reflect.TypeOf(bo)

	fmt.Println(t.Kind(), t.Name())

	t = t.Elem()
	fmt.Println(t.Kind(), t.Name())

	for i := 0; i < t.NumField(); i++ {
		tStructField := t.Field(i)
		fmt.Println("name:" + tStructField.Name)
		fmt.Println("PkgPath:" + tStructField.PkgPath)
		fmt.Println("Kind:", tStructField.Type.Kind())

		tag := tStructField.Tag
		fmt.Println("model:" + tag.Get("model"))
		fmt.Println("json:" + tag.Get("json"))
		if a, b := tag.Lookup("model"); true {
			fmt.Println("lookup:", a, b)
		}

		fmt.Println("\n")
	}

}
