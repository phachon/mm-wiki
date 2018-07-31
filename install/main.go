package main

import (
	_ "mm-wiki/install/storage"
	"github.com/astaxie/beego"
	"flag"
	"os"
	"log"
)

// install

var (
	port = flag.String("port", "8090", "please input listen port")
)

func main() {
	flag.Parse()

	_, err := os.Stat("../install.lock")
	if err == nil || !os.IsNotExist(err){
		log.Println("MM-Wiki already installed!")
		os.Exit(1)
	}

	//beego.BConfig.RunMode = "prod"
	beego.Run(":"+*port)
}

