package config

import (
	"os"
	"strings"

	"baotian0506.com/app/menu/applog"
	"github.com/Unknwon/goconfig"
)

const CONFIG_FILENAME = "config.ini"

var c *goconfig.ConfigFile

type Mysql struct {
	Password string
}

func init() {
	var err error
	var dir string

	if err != nil {
		applog.LogError.Printf("Getwd fail, err=%w", err)
		os.Exit(-1)
	}

	dir = "/root/.menu_config/"
	filename := dir + CONFIG_FILENAME
	c, err = goconfig.LoadConfigFile(filename) //加载配置文件
	if err != nil {
		dir, err = os.Getwd()
		dir = strings.TrimRight(dir, `/\`) + "/../../"
		filename := dir + CONFIG_FILENAME
		c, err = goconfig.LoadConfigFile(filename) //加载配置文件
		if err != nil {
			applog.LogError.Printf("get config file error, filename=%s", filename)
			os.Exit(-1)
		}
	}
	if GetDomain() == "" {
		applog.LogError.Printf("GetDomain return empty")
		os.Exit(-1)
	}
}

func GetDomain() string {
	val, err := c.GetValue("sys", "domain")
	if err != nil {
		return ""
	}

	return val
}

func GetSecret() string {
	val, err := c.GetValue("sys", "secret")
	if err != nil {
		return ""
	}

	return val
}

func GetVersion() string {
	val, err := c.GetValue("sys", "version")
	if err != nil {
		return ""
	}

	return val
}
func GetMysqConfig() *Mysql {
	m := &Mysql{}
	password, _ := c.GetValue("mysql", "password")

	m.Password = password

	return m
}
