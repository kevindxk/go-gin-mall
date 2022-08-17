/**
 * @Author: dexukong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/08/15 14:32
 */

package main

import (
	"ginmall/conf"
	"ginmall/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	r.Run(conf.HttpPort)
}
