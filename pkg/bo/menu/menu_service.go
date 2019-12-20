package menu

import (
	"fmt"
	"strconv"
	"strings"
)

type MenuService struct {
}

func (service *MenuService) FormatBO(bo *MenuBO) {

	ExtraMenuIdList, exists := bo.getExtra("menu_id_list")
	if !exists {
		return
	}
	ExtraMenuIdListValue, ok := ExtraMenuIdList.(string)
	if !ok {
		ExtraMenuIdListValue = ""
	}

	menu_id_list := make([]string, 0)
	bo.ExtraFormat = make(map[string]interface{}, 0)

	for _, v := range strings.Split(ExtraMenuIdListValue, ",") {
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
	bo.ExtraFormat["menu_id_list"] = strings.Join(menu_id_list, ",")

}
