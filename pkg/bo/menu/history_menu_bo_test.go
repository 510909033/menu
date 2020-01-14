package menu

import (
	"fmt"
	"testing"
	"time"

	"github.com/510909033/menu/bgf_bo"
)

func TestInsertSuccessHistoryMenuBO(t *testing.T) {
	var err error
	msg := ""
	history_menu_var := NewHistoryMenuBO(0)

	history_menu_var.UserId = int(time.Now().Unix())
	history_menu_var.MenuIdList = "123,456,789"

	err = history_menu_var.Insert()
	if err == nil {
		msg = fmt.Sprintf("insert success, id=%d", history_menu_var.Id)
		//		applog.LogInfo.Println(msg)
		t.Log(msg)
	} else {
		msg = fmt.Sprintf("insert fail, err=%v", err)
		//		applog.LogInfo.Println(msg)
		t.Fatal(msg)
	}

}

func TestSelectByExistsIdHistoryMenuBO(t *testing.T) {

	history_menu_var := NewHistoryMenuBO(1)
	if history_menu_var.IsNewRow() {
		t.Fatal("TestSelectByExistsId , result IsNewRow")
		return
	}
	t.Log("TestSelectByExistsId , success")
}

func TestQueryWhereFieldHistoryMenuBO(t *testing.T) {

	history_menu_var := NewHistoryMenuBO(3)
	where := "menu_id_list = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, "123,456,789")
	pageLimit := bgf_bo.PageLimit{
		Page:  1,
		Limit: 22,
	}
	if _, err := history_menu_var.Query(where, whereValue, pageLimit); err != nil {
		t.Fatalf("TestQueryWhereField fail, err=%v", err)
	} else {
		t.Log("TestQueryWhereField , success")
	}
}

func TestUpdateExistsIdHistoryMenuBO(t *testing.T) {
	var err error
	history_menu_var := NewHistoryMenuBO(1)
	if history_menu_var.IsNewRow() {
		t.Fatal("TestSelectByExistsId , result IsNewRow")
		return
	}
	//	applog.LogInfo.Println(history_menu_var.UserId)
	history_menu_var.UserId = 99999
	err = history_menu_var.Save()
	if err != nil {
		t.Fatal("TestUpdateExistsId , fail")
		return
	}

	t.Log("TestSelectByExistsId , success")
}
