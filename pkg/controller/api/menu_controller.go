package api

import (
	"fmt"
	"strconv"
	"strings"

	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/pkg/bo/menu"
)

type MenuController struct {
	base.Controller
}

func (ctrl *MenuController) SaveAction(ctx *base.BaseContext) {
	var err error
	var userId int
	ret := make(map[string]interface{}, 0)

	title := ctx.Request.FormValue("title")
	userIdString := ctx.Request.FormValue("user_id")

	userId, err = strconv.Atoi(userIdString)

	title = strings.TrimSpace(title)

	if title == "" {
		ctx.Fail("标题不能为空", nil)
		return
	}

	if userId < 1 {
		ctx.Fail("用户不能为空", nil)
		return
	}

	menuBO := menu.NewMenuBO(0)
	menuBO.Title = title
	menuBO.UserId = userId
	err = menuBO.Insert()
	if err != nil {
		ctx.Fail("保存失败", nil)
		return
	}
	ret["id"] = menuBO.Id
	ctx.Success(nil, ret)

}

func (ctrl *MenuController) ListAction(ctx *base.BaseContext) {

	menuBO := menu.NewMenuBO(0)
	ret := make(map[string]interface{}, 0)
	var retList []menu.MenuBO
	var err error
	retList, err = menuBO.Query()
	if err != nil {
		ctx.Fail("获取列表失败", nil)
		return
	}

	ret["list"] = retList
	ctx.Success(nil, ret)

}

func (ctrl *MenuController) Save1Action(ctx *base.BaseContext) {

	fmt.Println("call SaveAction")
	ret := make(map[string]interface{}, 0)

	data := map[string]interface{}{
		"user_id":  12345,
		"nickname": "hellow",
	}

	d2 := map[string]interface{}{
		"0": data,
		"1": data,
		"2": []string{
			"a",
			"b",
			"c",
		},
	}

	ret["status"] = "success"
	ret["data"] = d2

	ctx.Success(nil, ret)

}
