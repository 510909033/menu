package base

import (
	"context"
	"net/http"

	"github.com/510909033/menu/applog"
	"github.com/510909033/menu/pkg/bo/user"
	"github.com/510909033/menu/sign"
	"github.com/go-playground/form/v4"
)

type Controller struct {
}

type BaseContext struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	Context  context.Context
	Response map[string]interface{}
}

func (bc *BaseContext) Success(msg interface{}, ret interface{}) {
	bc.Response = make(map[string]interface{}, 0)
	bc.Response["status"] = "success"
	bc.Response["data"] = ret
	bc.Response["msg"] = msg
}
func (bc *BaseContext) Fail(msg interface{}, ret interface{}) {
	bc.Response = make(map[string]interface{}, 0)
	bc.Response["status"] = "fail"
	bc.Response["data"] = ret
	bc.Response["msg"] = msg
}

func (bc *BaseContext) Bind(input interface{}) (err error) {

	decoder := form.NewDecoder()
	err = bc.Request.ParseForm()
	if err != nil {
		applog.LogError.Printf("ParseForm fail, err=%w", err)
		panic(err)
	}
	err = decoder.Decode(input, bc.Request.Form)
	if err != nil {
		applog.LogError.Printf("Decode fail, err=%w", err)
		panic(err)
	}

	return nil
}

func (bc *BaseContext) GetUserId() int {
	u := &user.UserUtil{}
	signUtil := &sign.SignUtil{}
	if encUserId, err := signUtil.GetUserUniqid(bc.Request.FormValue("login_string")); err == nil {
		return u.GetUserId(encUserId)
	}
	return 0
}
