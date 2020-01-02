package menu

import (
	"encoding/json"
	"fmt"
	"reflect"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/bgf_bo"
)

/*

CREATE TABLE `menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `user_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户user_id',
  `title` VARCHAR(1000)  NOT NULL DEFAULT '' COMMENT '名称',
  `create_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_ts` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  `extra` varchar(1000) NOT NULL DEFAULT '' COMMENT '扩展字段',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COMMENT='菜单表';

*/

const (
	CATEGORY_MENU = 1
	CATEGORY_FOOD = 2
)

type MenuBO struct {
	bgfBO *bgf_bo.ModelMethod

	Id         int    `pk:"" column_name:"id" json:"id"`
	CategoryId int    `column_name:"category_id" json:"category_id"`
	UserId     int    `column_name:"user_id" json:"user_id"`
	Title      string `column_name:"title" json:"title"`
	Extra      string `column_name:"extra" json:"extra"`
	CreateTs   int64  `column_name:"create_ts" json:"create_ts"`
	UpdateTs   int64  `column_name:"update_ts" json:"update_ts"`

	CreateTsFormat string     `json:"create_ts_format"`
	ExtraFormat    *MenuExtra `json:"extra_format"`
}

type MenuExtra struct {
	MenuIdList string `json:"menu_id_list"`
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
	if menu.Id > 0 {
		menu.bgfBO.Load()
	}
	return menu
}

func (menu *MenuBO) GetTableName() string {
	return "menu"
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

func (bo *MenuBO) AddExtra(key string, value interface{}) error {
	extra := bo.Extra
	dbResult := make(map[string]interface{})
	if extra != "" {
		if err := json.Unmarshal([]byte(extra), dbResult); err != nil {
			//反解析失败，则extra置空
			//@TODO 记录日志
		}
	}
	dbResult[key] = value
	if bSlice, err := json.Marshal(dbResult); err == nil {
		bo.Extra = string(bSlice)
	} else {
		//Marshal 失败
		//@TODO 记录日志
	}
	return nil
}
func (bo *MenuBO) getExtra(key string) (interface{}, bool) {
	extra := bo.Extra

	dbResult := make(map[string]interface{})
	if extra == "" {
		return nil, false
	}
	if err := json.Unmarshal([]byte(extra), &dbResult); err != nil {
		//反解析失败，则extra置空
		//@TODO 记录日志
		applog.LogError.Printf("Unmarshal fail, err:%+v, extra=%s", err, extra)
		return nil, false
	}
	applog.LogDebug.Printf("dbResult:%+v", dbResult)
	if v, ok := dbResult[key]; ok {
		return v, true
	}
	return nil, false
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
		retList = append(retList, (v.(MenuBO)))
	}
	return
}

func (bo *MenuBO) Count(where string, whereValue []interface{}) (count int64, err error) {
	count, err = bo.bgfBO.Count(where, whereValue)
	if err != nil {
		applog.Log(err, applog.ERROR)
	}
	return
}
