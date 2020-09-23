package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1042531078@qq.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post(){
	for i:=0;i<10;i++{
		fmt.Printf("第%d次打印\n",i)
	}
}