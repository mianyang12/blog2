package config

import (
	"github.com/BurntSushi/toml"
	"os"
)

type tomlConfig struct {
	Viewer Viewer
	System SystemConfig //因为要提供给外部访问，所以要大写
}
type Viewer struct {
	Title       string
	Description string
	Logo        string
	Navigation  []string
	Bilibili    string
	Avatar      string
	UserName    string
	UserDesc    string
}
type SystemConfig struct {
	AppName         string
	Version         float32
	CurrentDir      string
	CdnURL          string
	QiniuAccessKey  string
	QiniuSecretKey  string
	Valine          bool
	ValineAppid     string
	ValineAppkey    string
	ValineServerURL string
}

var Cfg *tomlConfig //在函数外面定义：因为要定义成全局变量而非全局变量

func init() {
	//程序启动的时候 就会执行init方法
	Cfg = new(tomlConfig) //使用指针：允许程序外部修改
	Cfg.System.AppName = "go-blog"
	Cfg.System.Version = 1.0
	currentDir, _ := os.Getwd()
	Cfg.System.CurrentDir = currentDir
	_, err := toml.DecodeFile("config/config.toml", &Cfg) //进行configg.toml文件对Cfg变量赋值

	if err != nil {
		panic(err)
	}
}
