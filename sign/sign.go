package sign

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

type SignUtil struct{}

func (piru *SignUtil) CheckSign(r *http.Request) bool {

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
	//return config.GetSecret()
	return "iambabytreekey!@#$%^&*()"
}

func (piru *SignUtil) GetLoginString(userUniqid string) string {
	return userUniqid + "_" + fmt.Sprintf("%x", md5.Sum([]byte(userUniqid+piru.GetSecretKey())))
}

func (piru *SignUtil) GetUserUniqid(loginString string) (userUniqid string, err error) {
	l := strings.Split(loginString, "_")
	if len(l) != 2 {
		return "", errors.New("解析loginString失败,len!=2")
	}
	if l[1] != fmt.Sprintf("%x", md5.Sum([]byte(l[0]+piru.GetSecretKey()))) {
		return "", errors.New("loginString验签失败")
	}
	return l[0], nil
}
