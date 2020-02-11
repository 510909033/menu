package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Unknwon/goconfig"
)

var c *goconfig.ConfigFile

type Mysql struct {
	Driver string
}

func init() {
	var err error
	var dir string
	//configFilename := "config_dev.ini"
	configFilename1 := flag.String("config", "", "usage configFilename")
	flag.Parse()
	fmt.Println(*configFilename1)
	configFilename := *configFilename1

	if configFilename == "" {
		log.Fatalf("configFilename为空")
		os.Exit(-1)
	}

	dir = "/root/.menu_config/"
	filename := dir + configFilename
	c, err = goconfig.LoadConfigFile(filename) //加载配置文件
	if err != nil {
		dir, err := os.Executable()
		dir = filepath.Dir(dir)

		dir = strings.TrimRight(dir, `/\`) + "/"
		filename := dir + configFilename
		c, err = goconfig.LoadConfigFile(filename) //加载配置文件
		if err != nil {
			log.Fatalf("get config file error, filename=%s", filename)
			os.Exit(-1)
		}
	}
	if GetDomain() == "" {
		log.Fatalf("GetDomain return empty")
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
func GetMysql() *Mysql {
	m := &Mysql{}
	driver, _ := c.GetValue("mysql", "driver")
	m.Driver = driver
	return m
}

func GetLogsDir() (string, error) {

	return c.GetValue("log", "dir")

}
