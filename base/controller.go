package base

import (
	"context"
	"net/http"
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
