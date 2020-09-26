package db_mysql

import (
	"beego1/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

//在初始化函数当中连接数据库
func init() {
	fmt.Println("连接数据库")
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIP := config.String("db_ip")
	dbName := config.String("db_name")
	connUrl := dbUser + ":" + dbPassword + "@tcp(" + dbIP + ")/" + dbName + "?charset=utf8"
	db, err := sql.Open(dbDriver, connUrl)
	if err != nil {
		panic("数据库连接错误，请检查错误")
	}
	//为全局赋值
	Db = db
}
func InserUser(user models.Personto)(int64,error){
	//1.将用户密码进行hash脱敏，使用md5计算密码hash值，并储存hash值
	hashMd5:=md5.New()
	hashMd5.Write([]byte(user.Password))
	bytes :=hashMd5.Sum(nil)
	user.Password=hex.EncodeToString(bytes)
	fmt.Println("将要保存的用户名：",user.Nick,user.Birthday,user.Address,user.Password)
	result,err:=Db.Exec("insert into user(nick,birthday,address,password) value(?,?,?,?)",user.Nick,user.Birthday,user.Address,user.Password)
	if err != nil {
		return  -1,err
	}
	id,err :=result.RowsAffected()
	if err != nil {
		return  -1,err
	}
	return id,err
}

/*
查询用户
 */
func QueryUser(){
	Db.QueryRow("select * from ")
}
