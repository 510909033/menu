package api

import (
	"fmt"
	"strconv"
	"strings"

	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/bgf_bo"
	"baotian0506.com/app/menu/pkg/bo/menu"
)

type MenuController struct {
	base.Controller
}

func (ctrl *MenuController) SaveAction(ctx *base.BaseContext) {
	var err error
	var userId int
	ret := make(map[string]interface{}, 0)

	title := ctx.Request.FormValue("title")
	userIdString := ctx.Request.FormValue("user_id")

	userId, err = strconv.Atoi(userIdString)

	title = strings.TrimSpace(title)

	if title == "" {
		ctx.Fail("标题不能为空", nil)
		return
	}

	if userId < 1 {
		ctx.Fail("用户不能为空", nil)
		return
	}

	menuBO := menu.NewMenuBO(0)
	menuBO.Title = title
	menuBO.UserId = userId

	where := "title = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, title)
	pageLimit := bgf_bo.PageLimit{
		Page:  1,
		Limit: 1,
	}
	if tmpList, err := menuBO.Query(where, whereValue, pageLimit); true {
		if err != nil {
			ctx.Fail("系统错误", map[string]interface{}{})
			return
		}
		if len(tmpList) > 0 {
			ctx.Fail(
				fmt.Sprintf("title:%s已存在", title),
				map[string]interface{}{})
			return
		}
	}

	err = menuBO.Save()
	if err != nil {
		ctx.Fail("保存失败", nil)
		return
	}
	ret["id"] = menuBO.Id
	ctx.Success(nil, ret)

}

func (ctrl *MenuController) ListAction(ctx *base.BaseContext) {

	menuBO := menu.NewMenuBO(0)
	ret := make(map[string]interface{}, 0)
	var retList []menu.MenuBO
	var err error
	var page int

	pageStr := ctx.Request.FormValue("page")
	if page, err = strconv.Atoi(pageStr); err != nil {
		page = 1
	}
	if page < 1 {
		page = 1
	}

	where := ""
	whereValue := make([]interface{}, 0)
	pageLimit := bgf_bo.PageLimit{
		Page:  page,
		Limit: 20,
	}
	retList, err = menuBO.Query(where, whereValue, pageLimit)
	if err != nil {
		ctx.Fail("获取列表失败", nil)
		return
	}

	ret["list"] = retList
	ctx.Success(nil, ret)

}
