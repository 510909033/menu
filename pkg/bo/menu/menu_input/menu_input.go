package menu_input

type MenuEdit struct {
	UserId     int    `form:"user_id"`
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
	MenuIdList string `form:"menu_id_list"`
}

type MenuList struct {
	CategoryId int `form:"category_id"`
	Page       int `form:"page"`
	Pagesize   int `form:"pagesize"`
}

type HistoryMenuEdit struct {
	UserId     int    `form:"user_id"`
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
}

type HistoryMenuList struct {
	CategoryId int `form:"category_id"`
	Page       int `form:"page"`
	Pagesize   int `form:"pagesize"`
}

type FoodEdit struct {
	UserId     int    `form:"user_id"`
	Title      string `form:"title"`
	CategoryId int    `form:"category_id"`
}

type FoodList struct {
	CategoryId int `form:"category_id"`
	Page       int `form:"page"`
	Pagesize   int `form:"pagesize"`
}
