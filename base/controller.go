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

func (bc *BaseContext) Success(ret map[string]interface{}) {
	bc.Response = ret
}
