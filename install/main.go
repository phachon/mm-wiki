package main

import (
	_ "mm-wiki/install/storage"
	"github.com/astaxie/beego"
	"flag"
)

// install

var (
	port = flag.String("port", "8090", "please input listen port")
)

func main() {
	flag.Parse()
	beego.Run(":"+*port)
}

