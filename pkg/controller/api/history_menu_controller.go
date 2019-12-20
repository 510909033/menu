package api

import (
	"fmt"
	"strconv"
	"time"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/bgf_bo"
	"baotian0506.com/app/menu/pkg/bo/menu"
	"baotian0506.com/app/menu/pkg/bo/menu/menu_input"
)

type HistoryMenuController struct {
	base.Controller
}

func (ctrl *HistoryMenuController) SaveAction(ctx *base.BaseContext) {
	var err error
	var userId int
	ret := make(map[string]interface{}, 0)

	menu_id_list := ctx.Request.FormValue("menu_id_list")
	what_time := ctx.Request.FormValue("what_time")
	userIdString := ctx.Request.FormValue("user_id")

	userId, err = strconv.Atoi(userIdString)

	if menu_id_list == "" {
		ctx.Fail("菜单必选", nil)
		return
	}

	if userId < 1 {
		ctx.Fail("用户不能为空", nil)
		return
	}

	var whatTime int64
	if t, err := time.Parse("2006-01-02 15:04", what_time); err != nil {
		ctx.Fail("时间必选", nil)
		return
	} else {
		whatTime = t.Unix()
	}

	historyMenuBO := menu.NewHistoryMenuBO(0)

	historyMenuBO.MenuIdList = menu_id_list
	historyMenuBO.WhatTime = int(whatTime)
	historyMenuBO.UserId = userId

	err = historyMenuBO.Save()
	if err != nil {
		ctx.Fail("保存失败", nil)
		return
	}
	ret["id"] = historyMenuBO.Id
	ret["redirect_url"] = "/default/menu/index.html?layout=history_menu_list"
	ctx.Success("编辑成功", ret)

}

// ListAction 列表
func (ctrl *HistoryMenuController) ListAction(ctx *base.BaseContext) {
	input := &menu_input.HistoryMenuList{}
	historyMenuBO := menu.NewHistoryMenuBO(0)
	ret := make(map[string]interface{}, 0)
	var retList []menu.HistoryMenuBO
	var err error

	err = ctx.Bind(input)

	if err != nil {
		err = fmt.Errorf("bind fail, err=%w", err)
		applog.LogError.Println(err)
		ctx.Fail("获取参数失败", nil)
		return
	}

	input.Pagesize = 20

	if input.Page < 1 {
		input.Page = 1
	}

	where := ""
	whereValue := make([]interface{}, 0)
	pageLimit := bgf_bo.PageLimit{
		Page:  input.Page,
		Limit: input.Pagesize,
	}
	retList, err = historyMenuBO.Query(where, whereValue, pageLimit)
	if err != nil {
		ctx.Fail("获取列表失败", nil)
		return
	}

	history_menu_service := &menu.HistoryMenuService{}
	for k, v := range retList {
		history_menu_service.FormatHistoryMenuBO(&v)
		retList[k] = v
	}

	if count, err := historyMenuBO.Count(where, whereValue); err == nil {
		ret["count"] = count
	} else {
		ctx.Fail("获取列表失败2", nil)
		return
	}

	ret["pagesize"] = input.Pagesize
	ret["page"] = input.Page
	ret["list"] = retList
	ctx.Success(nil, ret)

}
