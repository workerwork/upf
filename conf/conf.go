package conf

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"time"
)

/**
const (
	SERVER_IP       = "127.0.0.1"
	SERVER_PORT     = 10006
	SERVER_RECV_LEN = 10
)*/

type server struct {
	IP      string `mapstructure:"ip"`
	Port    int    `mapstructure:"port"`
	RecvLen int    `mapstructure:"recv_len"`
}

var ServerConf = &server{}

// redis 缓存配置结构
type redis struct {
	Host        string        `mapstructure:"host"`
	Port        int           `mapstructure:"port"`
	Password    string        `mapstructure:"password"`
	DBNum       int           `mapstructure:"db"`
	MaxIdle     int           `mapstructure:"maxIdle"`
	MaxActive   int           `mapstructure:"maxActive"`
	IdleTimeout time.Duration `mapstructure:"idleTimeout"`
}

// RedisConf 缓存配置
var RedisConf = &redis{}

// logger 日志配置结构
type logger struct {
	Level  string `mapstructure:"level"`
	Pretty bool   `mapstructure:"pretty"`
	Color  bool   `mapstructure:"color"`
}

// LoggerConf 日志配置
var LoggerConf = &logger{}

// Setup 生成服务配置
func Setup() {
	viper.SetConfigType("YAML")
	// 读取配置文件内容
	data, err := ioutil.ReadFile("../config/config.yaml")
	if err != nil {
		log.Fatalf("Read 'config.yaml' fail: %v\n", err)
	}
	// 配置内容解析
	viper.ReadConfig(bytes.NewBuffer(data))
	// 解析配置赋值
	viper.UnmarshalKey("server", ServerConf)
	viper.UnmarshalKey("redis", RedisConf)
	viper.UnmarshalKey("logger", LoggerConf)
}
