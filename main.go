package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/memcache"
	_ "github.com/astaxie/beego/session/redis"
	_ "github.com/astaxie/beego/session/redis_cluster"
	_ "github.com/phachon/mm-wiki/app"
)

func main() {
	beego.Run()
}
