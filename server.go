package main

import (
	"github.com/workerwork/upf/conf"
	"github.com/workerwork/upf/logger"
	"github.com/workerwork/upf/redis"
	"github.com/workerwork/upf/udp"
)

func main() {
	//基本配置初始化
	conf.Setup()
	//日志初始化
	logger.Setup()
	//缓存初始化
	redis.Setup()
	//UDP处理
	udp.Setup()
}
