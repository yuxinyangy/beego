package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//1.获取name,age,sex
	//2.做数据对比
	//3.根据对比结果进行判断处理
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1042531078@qq.com"
	c.TplName = "index.tpl"
}
//编写一个post方法，用于处理post类型的请求
func (c *MainController) Post(){
	for i:=0;i<10;i++{
		fmt.Printf("第%d次打印\n",i)
	}
}