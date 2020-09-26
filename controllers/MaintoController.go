package controllers

import (
	"beego1/db_mysql"
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
func (r *MaintoController) Post(){
	fmt.Println(r == nil)
	fmt.Println(r.Ctx == nil)
	fmt.Println(r.Ctx.Request == nil)
	var person models.Personto
	dataBytes,err:=ioutil.ReadAll(r.Ctx.Request.Body)
	if err!=nil{
		r.Ctx.WriteString("数据接收错误，请重试")
		return
	}
	err =json.Unmarshal(dataBytes,&person)
	if err !=nil{
		r.Ctx.WriteString("数据解析失败，请重试")
		return
	}
	//fmt.Println("姓名",person.Name)
	//fmt.Println("生日",person.Birthday)
	//fmt.Println("地址",person.Address)
	//fmt.Println("昵称",person.Nick)
	//fmt.Println("密码",person.Password)
	r.Ctx.WriteString("数据解析成功")

	id,err := db_mysql.InserUser(person)
	if err!=nil {
		r.Ctx.WriteString("用户保存失败")
		return
	}
	fmt.Println(id)
	r.Ctx.WriteString("恭喜，用户保存成功")
	result := models.ResponseResult{
		Code:    0,
		Message: "保存成功",
		Data:    nil,
	}
	r.Data["json"] = &result
	r.ServeJSON()
}