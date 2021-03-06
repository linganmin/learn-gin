package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File

	AppEnv       string
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("config/app.ini")

	if err != nil {
		log.Fatalf("加载配置文件 config/app.ini 失败: %v", err)
	}

	loadBase()
	loadServer()
	loadApp()
}

func loadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func loadServer() {
	sec, err := Cfg.GetSection("server")

	if err != nil {
		log.Fatalf("加载配置文件的 server 配置项失败：%v", err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func loadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("加载配置文件的 app 配置项失败：%v", err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("123456")
}
