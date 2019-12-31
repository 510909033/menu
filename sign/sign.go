package sign

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"baotian0506.com/app/menu/base"
)

type SignUtil struct{}

func (piru *SignUtil) CheckSign(ctx *base.BaseContext) bool {
	r := ctx.Request
	r.ParseForm()

	signRet := ""
	params := make(map[string]string)
	for k, _ := range r.Form {
		if k == "sign" {
			signRet = r.Form.Get(k)
		} else {
			params[k] = r.Form.Get(k)
		}
	}
	fmt.Printf("signRet=%s\n", signRet)
	fmt.Println(piru.CalcSign(params))

	return piru.CalcSign(params) == signRet
}

//计算签名
func (piru *SignUtil) CalcSign(params map[string]string) string {

	signParams := make(map[string]string)
	for k, v := range params {
		k = strings.ToLower(k)
		signParams[k] = v
	}

	keys := make([]string, 0)
	for k, _ := range signParams {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	keyValues := make([]string, 0)
	for _, v := range keys {
		keyValues = append(keyValues, url.QueryEscape(v)+"="+url.QueryEscape(signParams[v]))
	}

	b := md5.Sum([]byte(strings.Join(keyValues, "") + piru.GetSecretKey()))
	return fmt.Sprintf("%x", b)
}

func (piru *SignUtil) GetSecretKey() string {
	return "iambabytreekey!@#$%^&*()"
}
