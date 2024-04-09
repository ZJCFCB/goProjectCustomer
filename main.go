package main

import (
	_ "github.com/gin-gonic/gin"

	"goProjectCustomer/service"
	"goProjectCustomer/view"
)

func main() {
	var cus view.CustomerView             //控制类
	cus.Cs = service.NewcustomerService() //实例化一个service类
	cus.Cs.AddDemo()                      //添加一个demo
	cus.Run()                             //开始干活了。。
}
