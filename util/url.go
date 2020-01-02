package util

import (
	"fmt"
	"net/url"

	"baotian0506.com/app/menu/config"
)

func GetUrl(path string, params url.Values) string {
	return fmt.Sprintf("%s%s?%s", config.GetDomain(), path, params.Encode())
}
