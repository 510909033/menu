package menu

import (
	"fmt"
	"reflect"

	"baotian0506.com/app/menu/bgf_bo"
)

/*

CREATE TABLE `history_menu`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',

  `user_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '用户user_id',
  `menu_id_list` varchar(2000) NOT NULL DEFAULT '',
  `what_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '',
  `status` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '',
  `extra` varchar(3000) NOT NULL DEFAULT '',
  `create_ts` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_ts` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COMMENT = '';

*/

type HistoryMenuBO struct {
	bgfBO *bgf_bo.ModelMethod

	Id         int    `pk:"" column_name:"id" json:"id"`
	UserId     int    `column_name:"user_id" json:"user_id"`
	MenuIdList string `column_name:"menu_id_list" json:"menu_id_list"`
	WhatTime   int    `column_name:"what_time" json:"what_time"`
	Status     int    `column_name:"status" json:"status"`
	Extra      string `column_name:"extra" json:"extra"`
	CreateTs   int64  `column_name:"create_ts" json:"create_ts"`
	UpdateTs   int64  `column_name:"update_ts" json:"update_ts"`

	CreateTsFormat string `json:"create_ts_format"`
}

func NewHistoryMenuBO(id int) *HistoryMenuBO {
	history_menu := &HistoryMenuBO{}
	history_menu.Id = id

	history_menu.bgfBO = &bgf_bo.ModelMethod{
		M:         history_menu,
		V:         reflect.ValueOf(history_menu),
		TableName: history_menu.GetTableName(),
		DBName:    history_menu.GetDBName(),
		IsNewRow:  true,
	}
	if history_menu.Id > 0 {
		history_menu.bgfBO.Load()
	}
	return history_menu
}

func (history_menu *HistoryMenuBO) GetTableName() string {
	return "history_menu"
}

func (history_menu *HistoryMenuBO) GetDBName() string {
	return "menu"
}
func (history_menu *HistoryMenuBO) IsNewRow() bool {
	return history_menu.bgfBO.IsNewRow
}

func init() {
	fmt.Println("history_menu_bo init")
	bgf_bo.Register(NewHistoryMenuBO(0))
}

func (bo *HistoryMenuBO) Insert() error {
	return bo.bgfBO.Insert()
}

func (bo *HistoryMenuBO) Save() error {
	return bo.bgfBO.Save()
}

func (bo *HistoryMenuBO) Query(where string, whereValue []interface{}, pageLimit bgf_bo.PageLimit) (retList []HistoryMenuBO, err error) {
	var ret []interface{}

	ret, err = bo.bgfBO.Query(where, whereValue, pageLimit)

	for _, v := range ret {
		retList = append(retList, v.(HistoryMenuBO))
	}
	return
}
