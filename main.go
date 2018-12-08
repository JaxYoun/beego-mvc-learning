package main

import (
	_ "beego-mvc-learning/models"
	_ "beego-mvc-learning/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
