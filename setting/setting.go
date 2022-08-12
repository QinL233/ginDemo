package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var cfg *ini.File

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
}

var DatabaseSetting = &Database{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func Init() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("错误加载app.ini %v", err)
	}

	cfg.Section("server").MapTo(ServerSetting)
	cfg.Section("database").MapTo(DatabaseSetting)
	cfg.Section("redis").MapTo(RedisSetting)
}
