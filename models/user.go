package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//表的设计
type User struct {
	Id       int
	Name     string
	Age      int
	Password string
}

func init() {
	//1.注册数据库连接
	//aliasName：声明数据库别名，至少要有一个名字为"default"的数据库别名
	orm.RegisterDataBase("default", "mysql", "root:mysql@tcp(127.0.0.1:3306)/yang_dev?charset=utf8")

	//2.注册数据模型（Mapping）
	orm.RegisterModel(new(User))

	//3.同步数据模型与数据库表
	//name：注册步骤声明的数据库别名
	//force：是否强制更新模型和表的结构【删了重建】，true：每次执行到这里都会删了重建数据表结构，false：表结构不会联动更新
	//verbose：是否可见创建过程
	orm.RunSyncdb("default", false, true)
}
