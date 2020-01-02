package api

import (
	"fmt"
	"strings"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/bgf_bo"
	"baotian0506.com/app/menu/pkg/bo/menu"
	"baotian0506.com/app/menu/pkg/bo/menu/menu_input"
)

type FoodController struct {
	base.Controller
}

func (ctrl *FoodController) SaveAction(ctx *base.BaseContext) {
	var err error
	var input = &menu_input.FoodEdit{}

	err = ctx.Bind(input)

	userId := ctx.GetUserId()
	if userId < 1 {
		ctx.Fail("请先登录", nil)
		return
	}
	input.UserId = userId

	if err != nil {
		err = fmt.Errorf("bind fail, err=%w", err)
		applog.LogError.Println(err)
		ctx.Fail("获取参数失败", nil)
		return
	}

	ret := make(map[string]interface{}, 0)

	input.Title = strings.TrimSpace(input.Title)

	if input.Title == "" {
		ctx.Fail("食材名称不能为空", nil)
		return
	}

	if input.UserId < 1 {
		ctx.Fail("用户不能为空", nil)
		return
	}

	input.CategoryId = menu.CATEGORY_FOOD

	menuBO := menu.NewMenuBO(0)
	menuBO.Title = input.Title
	menuBO.UserId = input.UserId
	menuBO.CategoryId = input.CategoryId

	where := "category_id = ? and title = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, menu.CATEGORY_FOOD, input.Title)
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
				fmt.Sprintf("title:%s已存在", input.Title),
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
	ret["modal-first-btn-text"] = "返回列表"
	ret["modal-first-btn-url"] = "/default/menu/?layout=food_list"
	ret["modal-second-btn-text"] = "继续添加"
	ret["modal-second-btn-url"] = "refresh"

	ctx.Success("编辑成功", ret)

}

func (ctrl *FoodController) ListAction(ctx *base.BaseContext) {
	var input = &menu_input.FoodList{}
	var retList []menu.MenuBO
	var err error

	err = ctx.Bind(input)

	if err != nil {
		err = fmt.Errorf("bind fail, err=%w", err)
		applog.LogError.Println(err)
		ctx.Fail("获取参数失败", nil)
		return
	}

	menuBO := menu.NewMenuBO(0)
	ret := make(map[string]interface{}, 0)

	if input.Page < 1 {
		input.Page = 1
	}
	input.Pagesize = 20
	input.CategoryId = menu.CATEGORY_FOOD

	where := "category_id = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, input.CategoryId)
	pageLimit := bgf_bo.PageLimit{
		Page:  input.Page,
		Limit: input.Pagesize,
	}
	retList, err = menuBO.Query(where, whereValue, pageLimit)
	if err != nil {
		ctx.Fail("获取列表失败1", nil)
		return
	}

	if count, err := menuBO.Count(where, whereValue); err == nil {
		ret["count"] = count
	} else {
		ctx.Fail("获取列表失败2", nil)
		return
	}

	ret["pagesize"] = input.Pagesize
	ret["list"] = retList
	ctx.Success(nil, ret)

}

// ListAllAction
func (ctrl *FoodController) ListAllAction(ctx *base.BaseContext) {

	r := ctx.Request

	search_key := r.FormValue("search_key")

	menuBO := menu.NewMenuBO(0)
	ret := make(map[string]interface{}, 0)
	var retList []menu.MenuBO
	var err error

	where := "category_id = ?"
	whereValue := make([]interface{}, 0)
	whereValue = append(whereValue, menu.CATEGORY_FOOD)
	pageLimit := bgf_bo.PageLimit{
		Page:  1,
		Limit: 2000,
	}
	retList, err = menuBO.Query(where, whereValue, pageLimit)
	if err != nil {
		ctx.Fail("获取列表失败", nil)
		return
	}

	if search_key != "" {
		searchKeyUtil := &menu.SearchKeyUtil{}
		retList = searchKeyUtil.Filter(search_key, retList)

	}

	ret["list"] = retList
	ctx.Success(nil, ret)

}
