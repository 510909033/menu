package api

import (
	"fmt"
	"strings"

	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/pkg/bo/menu"
)

type MenuController struct {
	base.Controller
}

func (ctrl *MenuController) SaveAction(ctx *base.BaseContext) {
	var err error
	ret := make(map[string]interface{}, 0)

	title := ctx.Request.FormValue("title")

	title = strings.TrimSpace(title)

	if title == "" {
		ctx.Fail(nil, nil)
		return
	}

	menuBO := menu.NewMenuBO(0)
	menuBO.Title = title
	err = menuBO.Insert()
	if err != nil {
		ctx.Fail("保存失败", nil)
		return
	}
	ret["id"] = menuBO.Id
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
