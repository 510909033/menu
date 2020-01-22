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

	"github.com/510909033/menu/applog"
	"github.com/510909033/menu/base"
	"github.com/510909033/menu/config"
	"github.com/510909033/menu/pkg/controller/api"
	"github.com/510909033/menu/pkg/controller/api/demo"
)

type myHandler struct {
	action map[string]reflect.Value
}

func init() {
	applog.LogInfo.Printf("main init")

}
func main() {
	config.GetDomain() //初始化配置

	//	(&api.WechatController{}).QrcodeAction(nil)
	//	return

	http.Handle("/default/", http.FileServer(http.Dir("template")))
	http.Handle("/public/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	h := &myHandler{
		action: make(map[string]reflect.Value),
	}
	h.registerController(&api.MenuController{})
	h.registerController(&api.WechatController{})
	h.registerController(&api.HistoryMenuController{})
	h.registerController(&api.FoodController{})
	h.registerController(&api.WebconfigController{})
	h.registerController(&api.UserController{})
	h.registerController(&demo.DemoEnvController{})

	http.Handle("/", h)

	err := http.ListenAndServe("0.0.0.0:9679", nil)
	applog.LogError.Printf("%v", err)
}

func (this *myHandler) registerController(controller interface{}) {

	//	getKey("history_menu_345_ok/get_user_1_list")

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

	if len(l) == 2 {
	} else if len(l) == 3 {
		l = l[1:]
	} else {
		err := errors.New("path 格式错误：" + path)
		applog.LogError.Printf("%v", err)
		return "", err
	}

	controllerName := l[0]
	calcControllerName := ""
	for _, v := range strings.Split(controllerName, "_") {
		for _, v1 := range v {
			calcControllerName += string(unicode.ToUpper(v1)) + v[1:]
			break
		}
	}

	actionName := l[1]
	calcActionName := ""
	for _, v := range strings.Split(actionName, "_") {
		for _, v1 := range v {
			calcActionName += string(unicode.ToUpper(v1)) + v[1:]
			break
		}
	}

	applog.Log(calcControllerName+"Controller-"+calcActionName+"Action", applog.DEBUG)

	return calcControllerName + "Controller-" + calcActionName + "Action", nil
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
	applog.LogInfo.Printf("r.Url.Path=%s", r.URL.Path)

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
