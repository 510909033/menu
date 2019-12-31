package api

import (
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/sign"
)

type UserController struct {
	base.Controller
}

func (ctrl *UserController) LoginAction(ctx *base.BaseContext) {

	signUtil := sign.SignUtil{}

	if signUtil.CheckSign(ctx) {
		ctx.Success("ok", nil)
	} else {
		//redirect
		ctx.Fail("no", nil)
	}

}
