package main

import (
	_"mm-wiki/app"
	_"github.com/astaxie/beego/session/memcache"
	_"github.com/astaxie/beego/session/redis"
	_"github.com/astaxie/beego/session/redis_cluster"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
