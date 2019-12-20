package base

import (
	"context"
	"net/http"

	"baotian0506.com/app/menu/applog"
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

func (bc *BaseContext) Success(msg interface{}, ret map[string]interface{}) {
	bc.Response = make(map[string]interface{}, 0)
	bc.Response["status"] = "success"
	bc.Response["data"] = ret
	bc.Response["msg"] = msg
}
func (bc *BaseContext) Fail(msg interface{}, ret map[string]interface{}) {
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
