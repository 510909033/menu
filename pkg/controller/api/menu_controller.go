package api

import (
	"fmt"

	"baotian0506.com/app/menu/base"
)

type MenuController struct {
	base.Controller
}

func (ctrl *MenuController) SaveAction(ctx *base.BaseContext) {

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

	ctx.Success(ret)

}
