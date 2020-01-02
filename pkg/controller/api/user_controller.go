package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/config"
	"baotian0506.com/app/menu/sign"
)

type UserController struct {
	base.Controller
}

func (ctrl *UserController) TestAction(ctx *base.BaseContext) {
	w := ctx.Writer
	signUtil := sign.SignUtil{}
	params := make(map[string]string)
	params["a"] = "a"
	params["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	params["login_string"] = signUtil.GetLoginString(ctx.Request.FormValue("u"))

	secret := signUtil.CalcSign(params)

	params["sign"] = secret

	u := make(url.Values)
	for k, v := range params {
		u.Set(k, v)
	}
	url := fmt.Sprintf(`<a href="%s/user/wechatLogin?%s">ttt</a>`, config.GetDomain(), u.Encode())

	w.Write([]byte(url))
}

func (ctrl *UserController) WechatLoginAction(ctx *base.BaseContext) {
	signUtil := sign.SignUtil{}

	applog.LogInfo.Println("login_string=" + ctx.Request.FormValue("login_string"))

	if signUtil.CheckSign(ctx.Request) {
		url := fmt.Sprintf("%s%s?%s", config.GetDomain(), "/default/menu", "layout=default&login_string="+ctx.Request.FormValue("login_string"))
		http.Redirect(ctx.Writer, ctx.Request, url, 301)
		return
	} else {
		//redirect
		ctx.Fail("no", nil)
	}
}

//验证login_string是否正确
func (ctrl *UserController) CheckLoginAction(ctx *base.BaseContext) {
	signUtil := sign.SignUtil{}

	if _, err := signUtil.GetUserUniqid(ctx.Request.FormValue("login_string")); err == nil {
		ctx.Success("success", nil)
	} else {

		ctx.Fail(err.Error(), nil)
	}
}
