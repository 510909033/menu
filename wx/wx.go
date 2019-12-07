package main

import (
	"fmt"
	"net/http"

	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/cache"
	"github.com/silenceper/wechat/message"
)

func hello(rw http.ResponseWriter, req *http.Request) {

	memCache := cache.NewMemory()

	//配置微信参数
	config := &wechat.Config{
		AppID:          "wx8156b3306d48031a",
		AppSecret:      "4a9df11334c00d838641ea65b4255dbf",
		Token:          "o0n2gjvGcQ5kSXg9rNzOJfYwR0MM",
		EncodingAESKey: "yourencodingaeskey",
		Cache:          memCache,
	}
	wc := wechat.NewWechat(config)

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
	server.SetDebug(true)
	fmt.Printf("openid:%s\n", server.GetOpenID())
	if ret, err := server.GetAccessToken(); true {
		fmt.Printf("GetAccessToken ret:%s\n", ret)
		if err != nil {
			fmt.Printf("GetAccessToken err:%v\n", err)
		}
	}
	if ret, ret2, err := server.GetAuthrInfo(server.AppID); true {
		if err != nil {
			fmt.Printf("GetAuthrInfo err:%v\n", err)
		} else {
			fmt.Println(ret2.AccessToken, ret2.Appid)
			fmt.Printf("GetAuthrInfo ret:%s\n", ret)
		}
	}

	//发送回复的消息
	server.Send()
}
