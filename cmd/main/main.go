package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"unicode"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/pkg/controller/api"
)

type myHandler struct {
	action map[string]reflect.Value
}

func init() {
	applog.LogInfo.Printf("main init")
}
func main() {

	http.Handle("/default/", http.FileServer(http.Dir("../../template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	h := &myHandler{
		action: make(map[string]reflect.Value),
	}
	h.registerController(&api.MenuController{})
	h.registerController(&api.WechatController{})

	http.Handle("/", h)

	err := http.ListenAndServe("0.0.0.0:9678", nil)
	applog.LogError.Printf("%v", err)
}

func (this *myHandler) registerController(controller interface{}) {

	t := reflect.TypeOf(controller)
	v := reflect.ValueOf(controller)

	m := make(map[string]interface{})

	m["t.Name()"] = t.Name()
	m["t.Elem().Name()"] = t.Elem().Name()
	//	m["t.Elem().Elem().Name()"] = t.Elem().Elem().Name()
	m["v.NumMethod()"] = v.NumMethod()

	for i := 0; i < v.NumMethod(); i++ {
		method := t.Method(i)
		s := make([]reflect.Value, 0)
		s = append(s, reflect.ValueOf(context.Background()))

		key := t.Elem().Name() + "-" + method.Name
		this.action[key] = v.Method(i)
	}

	//	fmt.Printf("%v\n", m)

}

func getKey(path string) (string, error) {
	//	uri := r.RequestURI

	path = strings.Trim(path, "/")
	if !strings.Contains(path, "/") {
		err := errors.New("path not contains / , path=" + path)
		applog.LogError.Printf("%v", err)
		return "", err
	}

	l := strings.Split(path, "/")
	if len(l) != 2 {
		err := errors.New("path 格式错误：" + path)
		applog.LogError.Printf("%v", err)
		return "", err
	}

	controllerName := l[0]
	for _, v := range controllerName {

		controllerName = string(unicode.ToUpper(v)) + controllerName[1:]
		break
	}

	actionName := l[1]
	for _, v := range actionName {
		actionName = string(unicode.ToUpper(v)) + actionName[1:]
		break
	}

	return controllerName + "Controller-" + actionName + "Action", nil
}

func (this *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		r.URL.Path = "/wechat/index"
	}
	key, err := getKey(r.URL.Path)
	if err != nil {
		w.Write([]byte("fail getKey"))
		return
	}

	if v, ok := this.action[key]; ok {

		baseContext := &base.BaseContext{}
		baseContext.Request = r
		baseContext.Writer = w
		baseContext.Context = context.Background()

		s := make([]reflect.Value, 0)
		s = append(s, reflect.ValueOf(baseContext))
		v.Call(s)

		if len(baseContext.Response) == 0 {
			return
		}
		b, err := json.Marshal(baseContext.Response)
		if err != nil {
			w.Write([]byte("Marshal fail"))
			return
		}

		w.Write(b)
		return
	}

	msg := fmt.Sprintf("key=%s, action[key] not exits", key)
	w.Write([]byte(msg))
	//	fmt.Println("ServeHTTP" + r.RequestURI)

	//	//	bo_menu.TestQuery()

	//	m := &bo_menu.Menu{}
	//	m.Title = time.Now().Format(common.TIME_FORMAT_YMDHIS)
	//	m.UserId = time.Now().Unix()
	//	m.Extra = ""
	//	m.Insert()

	//	w.Write([]byte(strconv.FormatInt(m.Id, 10)))
}
