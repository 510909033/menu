package api

import (
	"github.com/510909033/menu/base"
	"github.com/510909033/menu/pkg/bo/menu"
)

type WebconfigController struct {
	base.Controller
}

func (ctrl *WebconfigController) GetMenuListAction(ctx *base.BaseContext) {
	ret := menu.GetMenu()
	ctx.Success(nil, ret)

}
