package api

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"baotian0506.com/app/menu/applog"
	"baotian0506.com/app/menu/base"
	"baotian0506.com/app/menu/sign"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/menu"
	"github.com/silenceper/wechat/message"
	"github.com/silenceper/wechat/qr"
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

func (ctrl *WechatController) QrcodeAction(ctx *base.BaseContext) {

	qrcode := wc.GetQR()

	// uniqid_key

	request := qr.NewTmpQrRequest(time.Second*300, "q=123&p=我")

	ticket, err := qrcode.GetQRTicket(request)
	if err != nil {
		fmt.Println(err)
	}

	url := qr.ShowQRCode(ticket)

	fmt.Println(url)
}

func (ctrl *WechatController) IndexAction(ctx *base.BaseContext) {
	rw := ctx.Writer
	req := ctx.Request

	fmt.Println(req.URL.RawPath)
	req.ParseForm()
	fmt.Println(req.Form.Encode())

	// 传入request和responseWriter
	server := wc.GetServer(req, rw)
	//设置接收消息的处理方法
	server.SetMessageHandler(func(v message.MixMessage) *message.Reply {

		applog.LogInfo.Printf("MixMessage:%v", v)
		applog.LogInfo.Printf("MsgType:%s", v.MsgType)

		switch v.MsgType {
		//文本消息
		case message.MsgTypeText:
			//do something
			//回复消息：演示回复用户发送的消息
			text := message.NewText(v.Content)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

			//图片消息
		case message.MsgTypeImage:
			//do something

			//语音消息
		case message.MsgTypeVoice:
			//do something

			//视频消息
		case message.MsgTypeVideo:
			//do something

			//小视频消息
		case message.MsgTypeShortVideo:
			//do something

			//地理位置消息
		case message.MsgTypeLocation:
			//do something

			//链接消息
		case message.MsgTypeLink:
			//do something

			//事件推送消息
		case message.MsgTypeEvent:
			applog.LogInfo.Printf("Event:%s", v.Event)
			//事件推送消息
			switch v.Event {
			//EventSubscribe 订阅
			case message.EventSubscribe:
				//do something

				//取消订阅
			case message.EventUnsubscribe:
				//do something

				//用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
			case message.EventScan:
				//do something

				// 上报地理位置事件
			case message.EventLocation:
				//do something

				// 点击菜单拉取消息时的事件推送
			case message.EventClick:
				//do something
				//回复消息：演示回复用户发送的消息
				params := make(map[string]string)
				params["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
				params["openid"] = v.OpenID

				signUtil := sign.SignUtil{}
				params["sign"] = signUtil.CalcSign(params)
				u := make(url.Values)
				for k, v := range params {
					u.Set(k, v)
				}
				url := fmt.Sprintf(`<a href="http://39.106.133.49:9678/user/login?%s">菜单列表</a>`, u.Encode())

				text := message.NewText(url)
				return &message.Reply{MsgType: message.MsgTypeText, MsgData: text}

				// 点击菜单跳转链接时的事件推送
			case message.EventView:
				//do something

				// 扫码推事件的事件推送
			case message.EventScancodePush:
				//do something

				// 扫码推事件且弹出“消息接收中”提示框的事件推送
			case message.EventScancodeWaitmsg:
				//do something

				// 弹出系统拍照发图的事件推送
			case message.EventPicSysphoto:
				//do something

				// 弹出拍照或者相册发图的事件推送
			case message.EventPicPhotoOrAlbum:
				//do something

				// 弹出微信相册发图器的事件推送
			case message.EventPicWeixin:
				//do something

				// 弹出地理位置选择器的事件推送
			case message.EventLocationSelect:
				//do something

			}

		}

		//回复消息：演示回复用户发送的消息
		text := message.NewText(`<a href="http://39.106.133.49:9678/default/menu">菜单列表</a>`)
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
	//btn.SetViewButton("view btn1", "http://39.106.133.49:9678/wechat/S1")
	buttons[0] = btn

	buttons = make([]*menu.Button, 0)
	//设置btn为二级菜单
	btn2 := new(menu.Button)
	btn2.SetClickButton("获取链接", "key123")
	//btn2.SetSubButton("获取链接", buttons)

	buttons2 := make([]*menu.Button, 1)
	buttons2[0] = btn2

	//发送请求
	err := mu.SetMenu(buttons2)
	if err != nil {
		fmt.Printf("err= %v", err)
		return
	}
}

func (ctrl *WechatController) S1Action(ctx *base.BaseContext) {
	oauth := wc.GetOauth()
	//	url := oauth.GetRedirectURL()
	//	err := oauth.Redirect("跳转的绝对地址", "snsapi_userinfo", "123dd123")
	err := oauth.Redirect(ctx.Writer, ctx.Request, "http://39.106.133.49:9678/wechat/s1", "snsapi_userinfo", "123dd123")
	if err != nil {
		fmt.Println(err)
	}
}

func (ctrl *WechatController) S2Action(ctx *base.BaseContext) {
	oauth := wc.GetOauth()
	code := ctx.Request.FormValue("code")
	applog.LogInfo.Printf("code=%s", code)
	resToken, err := oauth.GetUserAccessToken(code)
	applog.LogInfo.Printf("resToken:%v", resToken)
	if err != nil {
		applog.LogError.Printf("GetUserAccessToken err, err=%v", err)
		return
	}

	//getUserInfo
	userInfo, err := oauth.GetUserInfo(resToken.AccessToken, resToken.OpenID)
	if err != nil {
		applog.LogError.Printf("GetUserInfo err, err=%v", err)
		return
	}
	applog.LogInfo.Printf("userInfo:%v", userInfo)
}
