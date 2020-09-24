package controllers

import (
	"beego1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type MainController struct {
	beego.Controller
}
func (c *MainController) Get() {
	//1.获取请求数据
	userName :=c.Ctx.Input.Query("user")
	password :=c.Ctx.Input.Query("pad")
	//2.使用固定数据进行数据校验
	if userName !="admin" || password !="123" {
		//代表错误处理
		c.Ctx.ResponseWriter.Write([]byte("对不起，数据校验错误。"))
		return
	}
	//校验正确的情况
	c.Ctx.ResponseWriter.Write([]byte("恭喜，数据校验成功"))
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1042531078@qq.com"
	c.TplName = "index.tpl"
}
//编写一个post方法，用于处理post类型的请求
//该方法用于处理post请求
/*func (c *MainController) Post(){
	//1.接收（解析）post请求的参数
	name :=c.Ctx.Request.FormValue("name")
	age :=c.Ctx.Request.FormValue("age")
	sex :=c.Ctx.Request.FormValue("sex")
	fmt.Println(name,age,sex)
	//2.进行数据校验
	if name != "admin" && age != "18" {
		c.Ctx.WriteString("数据校验失败")
		return
	}
	c.Ctx.WriteString("数据校验成功")
}*/

//该方法用于处理post请求
func (c *MainController) Post(){
	//1.解析前端提交的json格式的数据
	var person models.Person
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
	fmt.Println("年龄",person.Age)
	fmt.Println("性别",person.Sex)
	c.Ctx.WriteString("数据解析成功")
}

//该方法用于处理delete请求
func (c *MainController) Delete(){

}