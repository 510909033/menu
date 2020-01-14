package menu

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/510909033/menu/pkg/common"
)

type HistoryMenuService struct {
}

func (service *HistoryMenuService) FormatHistoryMenuBO(bo *HistoryMenuBO) {

	bo.WhatTimeFormat = time.Unix(int64(bo.WhatTime), 0).Add(time.Hour * -8).Format(common.TIME_FORMAT_YMDHIS)

	bo.MenuIdListFormat = ""
	menu_id_list := make([]string, 0)

	for _, v := range strings.Split(bo.MenuIdList, ",") {
		fmt.Println(v)
		if v == "" {
			continue
		}
		menu_id, err := strconv.Atoi(v)
		if err != nil {
			menu_id_list = append(menu_id_list, " - ")
			continue
		}
		menu_bo := NewMenuBO(menu_id)
		menu_id_list = append(menu_id_list, menu_bo.Title)
	}
	bo.MenuIdListFormat = strings.Join(menu_id_list, ",")

}
