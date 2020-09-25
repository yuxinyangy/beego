package controllers

import (
	"beego1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

//该结构用于处理/register请求

type MaintoController struct {
	beego.Controller
}
//该结构用于处理POST请求
func (c *MaintoController) Post(){
	var person models.Personto
	dataBytes,err:=ioutil.ReadAll(c.Ctx.Request.Body)
	if err!=nil{
		c.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err =json.Unmarshal(dataBytes,&person)
	if err !=nil{
		c.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	fmt.Println("姓名",person.Name)
	fmt.Println("生日",person.Birthday)
	fmt.Println("地址",person.Address)
	fmt.Println("昵称",person.Nick)
	c.Ctx.WriteString("数据解析成功")
}