package menu

import (
	"fmt"
	"reflect"

	"baotian0506.com/app/menu/bgf_bo"
)

/*

CREATE TABLE `w_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户user_id',
  `title` VARCHAR(1000)  NOT NULL DEFAULT '' COMMENT '名称',
  `create_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `extra` varchar(1000) NOT NULL DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='菜单表';

*/

type MenuBO struct {
	bgfBO *bgf_bo.ModelMethod

	Id       int    `pk:"" column_name:"id" json:"id"`
	UserId   int    `column_name:"user_id" json:"user_id"`
	Title    string `column_name:"title" json:"title"`
	Extra    string `column_name:"extra" json:"extra"`
	CreateTs int64  `column_name:"create_ts" json:"create_ts"`
	UpdateTs int64  `column_name:"update_ts" json:"update_ts"`

	CreateTsFormat string `json:"create_ts_format"`
}

func NewMenuBO(id int) *MenuBO {
	menu := &MenuBO{}
	menu.Id = id

	menu.bgfBO = &bgf_bo.ModelMethod{
		M:         menu,
		V:         reflect.ValueOf(menu),
		TableName: menu.GetTableName(),
		DBName:    menu.GetDBName(),
		IsNewRow:  true,
	}
	if menu.Id > 1 {
		menu.bgfBO.Load()
	}
	return menu
}

func (menu *MenuBO) GetTableName() string {
	return "w_menu"
}

func (menu *MenuBO) GetDBName() string {
	return "menu"
}
func (menu *MenuBO) IsNewRow() bool {
	return menu.bgfBO.IsNewRow
}

func init() {
	fmt.Println("menu_bo init")
	bgf_bo.Register(NewMenuBO(0))
}

func (bo *MenuBO) Insert() error {
	return bo.bgfBO.Insert()
}

func (bo *MenuBO) Save() error {
	return bo.bgfBO.Save()
}

func (bo *MenuBO) Query(where string, whereValue []interface{}, pageLimit bgf_bo.PageLimit) (retList []MenuBO, err error) {
	var ret []interface{}

	ret, err = bo.bgfBO.Query(where, whereValue, pageLimit)

	for _, v := range ret {
		retList = append(retList, v.(MenuBO))
	}
	return
}
