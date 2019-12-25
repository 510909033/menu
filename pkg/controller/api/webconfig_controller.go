package api

import (
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/pkg/bo/menu"
)

type WebconfigController struct {
	base.Controller
}

func (ctrl *WebconfigController) GetMenuListAction(ctx *base.BaseContext) {

	ret := menu.GetMenu()
	ctx.Success(nil, ret)

}
