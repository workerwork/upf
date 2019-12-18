package main

import (
	"N4test/conf"
	"N4test/logger"
	"N4test/redis"
	"N4test/udp"
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
