package menu

import (
	"testing"
	"time"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/pkg/common"
)

func TestMyUnit1(t *testing.T) {

	menu := NewMenuBO(0)

	menu.UserId = int(time.Now().Unix())
	menu.Title = time.Now().Format(common.TIME_FORMAT_YMDHIS)

	if menu.Insert() == nil {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

}

//func TestMyUnit2(t *testing.T) {

//	menu := NewMenuBO(0)

//	menu.UserId = int(time.Now().Unix())
//	menu.Title = time.Now().Format(common.TIME_FORMAT_YMDHIS)

//	if menu.Insert() == nil {
//		t.Log("the result is ok")
//	} else {
//		t.Fatal("the result is wrong")
//	}

//}

func TestMyUnit3(t *testing.T) {

	menu := NewMenuBO(3)
	applog.LogInfo.Printf("%v", menu)
	applog.LogInfo.Printf("%s", menu.Title)

	_ = menu
}

func TestMyUnit4(t *testing.T) {

	menu := NewMenuBO(3)
	menu.Query()

	_ = menu
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
