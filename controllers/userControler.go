package controllers

import (
	"beego-mvc-learning/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func (c *MainController) Get() {
	c.Data["name"] = "yang"
	c.Data["id"] = "75"
	c.TplName = "test0.html"
}

func (c *MainController) Post() {
	c.Data["name"] = "jax"
	c.Data["id"] = "75"
	c.TplName = "test0.html"
}

func (c *MainController) Put() {
	//InsertUser()
	//SelectUser()
	//UpdateUser()
	DeleteUser()
}

func InsertUser() {
	//1.有ORM对象
	orm := orm.NewOrm()

	//2.有要插入数据的结构体对象
	user := models.User{}

	//3.对结构体对象赋值
	user.Name = "lola"
	user.Age = 27
	user.Password = "kkk"

	//4.插入
	i, e := orm.Insert(&user)
	if e != nil {
		beego.Error("插入失败！")
	} else {
		if i == 1 {
			beego.Info("插入成功！")
		}
	}
}

func SelectUser() {
	//1.有ORM对象
	orm := orm.NewOrm()

	//2.有要插入数据的结构体对象
	user := models.User{}

	//3.指定要查询的字段
	user.Id = 1

	//4.根据主键id查询
	err := orm.Read(&user)
	if err != nil {
		beego.Error(err)
	} else {
		fmt.Println(user)
	}

	//5.根据非主键查询
	user.Name = "yang"
	err0 := orm.Read(&user, "Name")
	if err0 != nil {
		beego.Error(err0)
	} else {
		fmt.Println(user)
	}
}

func UpdateUser() {
	//1.有ORM对象
	orm := orm.NewOrm()

	//2.有要更新数据的结构体对象
	user := models.User{}

	//3.查到要更新的数据
	//user.Id = 5
	//err := orm.Read(&user)

	user.Name = "lola"
	err := orm.Read(&user, "Name")
	if err == nil {
		//4.给数据重新赋值
		user.Name = "jjjjjjjjjjjj"
		user.Password = "pppppppppppp"

		//5.更新
		_, err0 := orm.Update(&user)
		if err0 == nil {
			fmt.Println("更新成功")
		} else {
			beego.Error(err0)
		}
	} else {
		beego.Error(err)
	}
}

func DeleteUser() {
	//1.有ORM对象
	orm := orm.NewOrm()

	//2.有要删除数据的结构体对象
	user := models.User{}

	//3.指定要删除哪一条数据
	user.Id = 6

	//4.删除
	i, err := orm.Delete(&user)
	if err == nil {
		if i > 0 {
			beego.Info("删除成功！")
		}
	} else {
		beego.Error(err)
	}
}
