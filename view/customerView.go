package view

import (
	"fmt"
	"goProjectCustomer/model"
	"goProjectCustomer/service"
)

type CustomerView struct {
	Cs *service.CustomerService
}

//菜单显示

func (C *CustomerView) Run() {

	for {
		var chance string
		var confire bool
		fmt.Println("\n\n----------------客户信息管理系统----------------")
		fmt.Println("                 1.添加客户                 ")
		fmt.Println("                 2.修改客户                 ")
		fmt.Println("                 3.删除客户                 ")
		fmt.Println("                 4.客户列表                 ")
		fmt.Println("                 5.退出                 ")
		fmt.Println("                 请选择(1-5)                 ")
		fmt.Printf("你的选择(1-5): ")
		fmt.Scanln(&chance)

		switch chance {
		case "1":
			C.add()
		case "2":
			C.update()
		case "3":
			C.delete()
		case "4":
			C.list()
		case "5":
			conform := C.exit()
			if conform {
				confire = true
			}

		default:
			fmt.Println("------------------请输入正确的选项-----------------")
		}

		if confire {
			break
		}
	}
}

func (C *CustomerView) list() { //打印一下列表
	inform := C.Cs.List()
	fmt.Println("------------------客户列表begin-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for _, value := range inform {
		fmt.Println(value.Getinfo())
	}
	fmt.Println("------------------客户列表ending-----------------")
}

func (C *CustomerView) add() { //添加客户
	var name, gender, phone, email string
	var id, age int
	fmt.Println("------------------添加客户begin-----------------")
	fmt.Printf("Name : ")
	fmt.Scanln(&name)
	fmt.Printf("gender : ")
	fmt.Scanln(&gender)
	fmt.Printf("age : ")
	fmt.Scanln(&age)
	fmt.Printf("phone : ")
	fmt.Scanln(&phone)
	fmt.Printf("email : ")
	fmt.Scanln(&email)
	id = C.Cs.GetId() + 1
	tempCustomer := model.NewCustomer(id, name, gender, age, phone, email)
	if C.Cs.AddCustomer(tempCustomer) {
		fmt.Println("添加完成")
	}
	fmt.Println("------------------添加客户end-----------------")
}

func (C *CustomerView) delete() { //删除客户
	fmt.Println("------------------删除客户begin-----------------")
	var id int = -1
	fmt.Printf("请输入待删除客户编号(-1退出)")
	fmt.Scanln(&id)
	if id != -1 {
		var choice string
		fmt.Printf("请确认是否删除用户信息(y/n)? ")
		fmt.Scanln(&choice)
		if choice == "y" || choice == "Y" {
			isSuccess := C.Cs.Delete(id)
			if isSuccess {
				fmt.Println("------------------成功删除客户end--------------")
			} else {
				fmt.Println("------------------id 不存在--------------")
			}
		}
	}
}

func (C *CustomerView) update() { // 更新客户信息，不修改的项目直接敲回车
	fmt.Println("------------------更新客户begin-----------------")
	var id int = -1
	fmt.Printf("请输入待更新客户编号(-1退出)")
	fmt.Scanln(&id)
	if id != -1 {
		index := C.Cs.FindById(id)
		if index != -1 { //存在
			context := C.Cs.GetInform(index)
			var name, gender, phone, email string
			var age int
			fmt.Println("------------------修改客户begin-----------------")
			fmt.Printf("Name(%s) : ", context.Name)
			fmt.Scanln(&name)
			if name == "" {
				name = context.Name
			}
			fmt.Printf("gender(%s) : ", context.Gender)
			fmt.Scanln(&gender)
			if gender == "" {
				gender = context.Gender
			}
			fmt.Printf("age(%d) : ", context.Age)
			fmt.Scanln(&age)
			if age == 0 {
				age = context.Age
			}
			fmt.Printf("phone(%s) : ", context.Phone)
			fmt.Scanln(&phone)
			if phone == "" {
				phone = context.Phone
			}
			fmt.Printf("email(%s): ", context.Email)
			fmt.Scanln(&email)
			if email == "" {
				email = context.Email
			}

			C.Cs.Update(index, name, gender, age, phone, email)

			fmt.Println("------------------修改客户成功-----------------")

		} else { //不存在用户输入的id
			fmt.Println("------------------客户id 不存在-----------------")
		}
	}
}

func (C *CustomerView) exit() bool { //确认一下是否真的要退出
	var conform string
	fmt.Printf("确定要退出吗(y/n) ")
	fmt.Scanln(&conform)
	if conform == "y" || conform == "Y" || conform == "n" || conform == "N" {
		if conform == "y" || conform == "Y" {
			return true
		} else {
			return false
		}
	} else {
		fmt.Println("------------------请重新输入-----------------")
	}
	return false
}
