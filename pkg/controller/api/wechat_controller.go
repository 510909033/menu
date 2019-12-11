package api

import (
	"fmt"

	"baotian0506.com/app/menu/base"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
)

type WechatController struct {
	base.Controller
}

var memCache = cache.NewMemory()
var config = &wechat.Config{
	AppID:          "wx8156b3306d48031a",
	AppSecret:      "4a9df11334c00d838641ea65b4255dbf",
	Token:          "o0n2gjvGcQ5kSXg9rNzOJfYwR0MM",
	EncodingAESKey: "yourencodingaeskey",
	Cache:          memCache,
}
var wc = wechat.NewWechat(config)

func (ctrl *WechatController) IndexAction(ctx *base.BaseContext) {
	rw := ctx.Writer
	req := ctx.Request

	fmt.Println(req.URL.RawPath)

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {

		//回复消息：演示回复用户发送的消息
		text := message.NewText(msg.Content)
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	//发送回复的消息
	server.Send()
}

func (ctrl *WechatController) SetMenuAction(ctx *base.BaseContext) {
	mu := wc.GetMenu()

	buttons := make([]*menu.Button, 1)
	btn := new(menu.Button)

	//创建click类型菜单
	btn.SetClickButton("name", "key123")
	buttons[0] = btn

	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetSubButton("subButton", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	//发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}
}
