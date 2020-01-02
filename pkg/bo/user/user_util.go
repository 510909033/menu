package user

import "baotian0506.com/app/menu/applog"

type UserUtil struct{}

var userIDMap = map[string]int{
	"o1jnu1XtP4g-1mNUtO-NQrCyfuBQ": 12345,
	"o1jnu1WEBN177aDc9d5BVLtpDEkw": 54321,
}

func (u *UserUtil) GetUserId(encUserId string) (userId int) {
	if v, ok := userIDMap[encUserId]; ok {
		return v
	}
	applog.LogError.Printf("encUserId未配置user_id, encUserId=%s\n", encUserId)
	return 0
}
