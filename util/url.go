package util

import (
	"fmt"
	"net/url"

	"github.com/510909033/menu/config"
)

func GetUrl(path string, params url.Values) string {
	return fmt.Sprintf("%s%s?%s", config.GetDomain(), path, params.Encode())
}
