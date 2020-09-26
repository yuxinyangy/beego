package main

import (
	_ "beego1/db_mysql"
	_ "beego1/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
