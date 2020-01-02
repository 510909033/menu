package demo

import (
	"fmt"
	"log"
	"os"
	"time"

	"baotian0506.com/app/menu/base"
	"github.com/Unknwon/goconfig"
	"github.com/kelseyhightower/envconfig"
)

type DemoEnvController struct {
	base.Controller
}
type Specification struct {
	Debug      bool
	Port       int
	User       string
	Users      []string
	Rate       float32
	Timeout    time.Duration
	ColorCodes map[string]int
	Client     string
}

func (ctrl *DemoEnvController) Test1Action(ctx *base.BaseContext) {
	var cfg *goconfig.ConfigFile
	config, err := goconfig.LoadConfigFile(`C:\Users\Administrator\go\src\baotian0506.com\app\menu\config.ini`) //加载配置文件
	if err != nil {
		fmt.Println("get config file error")
		os.Exit(-1)
	}
	cfg = config

	glob, _ := cfg.GetSection("sys") //读取全部mysql配置
	fmt.Println(glob)

}

func (ctrl *DemoEnvController) TestAction(ctx *base.BaseContext) {
	var s Specification
	//BSPRINT_CLIENT
	err := envconfig.Process("bsprint", &s)

	if err != nil {
		log.Fatal(err.Error())
	}
	format := "Debug: %v\nPort: %d\nUser: %s\nRate: %f\nTimeout: %s,client:%s\n"
	_, err = fmt.Printf(format, s.Debug, s.Port, s.User, s.Rate, s.Timeout, s.Client)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Users:")
	for _, u := range s.Users {
		fmt.Printf("  %s\n", u)
	}

	fmt.Println("Color codes:")
	for k, v := range s.ColorCodes {
		fmt.Printf("  %s: %d\n", k, v)
	}
}
