package bo_menu

import (
	"testing"
	"time"

	"baotian0506.com/app/menu/pkg/common"
)

func TestMyUnit1(t *testing.T) {

	menu := &MenuBO{}

	menu.UserId = time.Now().Unix()
	menu.Title = time.Now().Format(common.TIME_FORMAT_YMDHIS)

	if menu.Insert() == nil {
		t.Log("the result is ok")
	} else {
		t.Fatal("the result is wrong")
	}

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
