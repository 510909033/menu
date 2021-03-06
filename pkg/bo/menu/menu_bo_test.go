package menu

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/510909033/menu/applog"
	"github.com/510909033/menu/bgf_bo"
	"github.com/510909033/menu/pkg/common"
)

func TestNewMenuBO(t *testing.T) {

	err := errors.New("first error!")
	err1 := fmt.Errorf("1111 \n%w", err)
	err2 := fmt.Errorf("222 \n%w", err1)

	fmt.Println(err2)
	fmt.Println(errors.Unwrap(err2))
	fmt.Println(err2)

	menu := NewMenuBO(0)

	v := reflect.ValueOf(menu)

	for i := 0; i < v.Elem().NumField(); i++ {
		fieldV := v.Elem().Field(i)

		if v.Elem().Type().Field(i).PkgPath == "" {
			applog.LogInfo.Println(v.Elem().Type().Field(i).Name)
			applog.LogInfo.Println(fieldV.Interface())
		}

		//

	}

}

func TestInsertSuccess(t *testing.T) {
	var err error
	msg := ""
	menu := NewMenuBO(0)

	menu.UserId = int(time.Now().Unix())
	menu.Title = time.Now().Format(common.TIME_FORMAT_YMDHIS)

	err = menu.Insert()
	if err == nil {
		msg = fmt.Sprintf("insert success, id=%d", menu.Id)
		applog.LogInfo.Println(msg)
		t.Log(msg)
	} else {
		msg = fmt.Sprintf("insert fail, err=%v", err)
		applog.LogInfo.Println(msg)
		t.Fatal(msg)
	}

}

func TestSelectByExistsId(t *testing.T) {

	menu := NewMenuBO(3)
	if menu.IsNewRow() {
		t.Fatal("TestSelectByExistsId , result IsNewRow")
		return
	}
	t.Log("TestSelectByExistsId , success")
}

func TestQueryWhereField(t *testing.T) {

	menu := NewMenuBO(3)
	where := "title = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, "12345")
	pageLimit := bgf_bo.PageLimit{
		Page:  1,
		Limit: 22,
	}
	if _, err := menu.Query(where, whereValue, pageLimit); err != nil {
		t.Fatalf("TestQueryWhereField fail, err=%v", err)
	} else {
		t.Log("TestQueryWhereField , success")
	}
}

func TestUpdateExistsId(t *testing.T) {
	var err error
	menu := NewMenuBO(3)
	if menu.IsNewRow() {
		t.Fatal("TestSelectByExistsId , result IsNewRow")
		return
	}
	applog.LogInfo.Println(menu.Title)
	menu.Title = "66666111"
	err = menu.Save()
	if err != nil {
		t.Fatal("TestUpdateExistsId , fail")
		return
	}

	t.Log("TestSelectByExistsId , success")
}

//func TestMyUnit1V2(t *testing.T) {
//	sum := my_unit1(1, 2)
//	if sum == 3 {
//		t.Log("the result is ok")
//	} else {
//		t.Fatal("the result is wrong")
//	}

//	sum = my_unit1(3, 4)
//	if sum == 7 {
//		t.Log("the result is ok")
//	} else {
//		t.Fatal("the result is wrong")
//	}
//}

//func Benchmark(b *testing.B) {
//	for i := 0; i < b.N; i++ { // b.N，测试循环次数
//		my_unit1(4, 5)
//	}
//}
