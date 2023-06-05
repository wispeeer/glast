package info

import (
	"fmt"
)

var Services string     // 微服务项目名称
var MicroService string // 微服务模块名称

var ServicesVersion string //微服务项目版本

var MicroServiceName string = fmt.Sprintf("%s(%s)", Services, MicroService)
