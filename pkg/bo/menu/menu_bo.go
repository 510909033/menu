package bo_menu

import (
	"fmt"
	"time"

	"baotian0506.com/app/menu/applog"
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
type menuDef struct {
	TableName string
}
type MenuBO struct {
	Id       int64
	UserId   int64
	Title    string
	Extra    string
	CreateTs int64
	UpdateTs int64
}

func (def *MenuBO) NewMenuDef() *menuDef {
	return &menuDef{
		TableName: "w_menu",
	}
}

func (menu *MenuBO) Insert() (err error) {
	def := menu.NewMenuDef()
	if menu.CreateTs == 0 {
		menu.CreateTs = time.Now().Unix()
	}

	if menu.UpdateTs == 0 {
		menu.UpdateTs = time.Now().Unix()
	}
	query := fmt.Sprintf("INSERT INTO %s (user_id,title,extra,create_ts,update_ts) values(?,?,?,?,?)", def.TableName)
	args := []interface{}{
		menu.UserId,
		menu.Title,
		menu.Extra,
		menu.CreateTs,
		menu.UpdateTs,
	}
	sqlResult, err := db.Exec(query, args...)

	if err != nil {
		applog.LogError.Printf("err=%v", err)
		return err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		applog.LogError.Printf("err=%v", err)
	}
	menu.Id = id

	return nil
}


