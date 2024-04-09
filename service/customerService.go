package service

import (
	"goProjectCustomer/model"
)

//这里主要完成对customer的操作 包括增删改查

type CustomerService struct {
	customerNum int
	customers   []model.Customer
}

// 工厂模式
func NewcustomerService() *CustomerService {
	return &CustomerService{
		customerNum: 0,
	}
}

func (c *CustomerService) AddDemo() {
	demo := model.NewCustomer(1, "张三", "男", 20, "112", "zs@baidu.com")
	c.customerNum += 1
	c.customers = append(c.customers, demo)
}

func (c *CustomerService) AddCustomer(customer model.Customer) bool {
	c.customerNum += 1
	c.customers = append(c.customers, customer)
	return true
}

func (c *CustomerService) GetId() int {
	return c.customerNum
}

func (c *CustomerService) List() []model.Customer {
	return c.customers
}

func (c *CustomerService) Delete(id int) bool {

	var index int = c.FindById(id)
	if index == -1 {
		return false
	}

	c.customers = append(c.customers[:index], c.customers[index+1:]...)
	//c.customerNum -= 1
	return true
}

func (c *CustomerService) FindById(id int) int {
	var index int = -1
	for key, value := range c.customers {
		if value.Id == id {
			index = key
			break
		}
	}
	return index
}

func (c *CustomerService) GetInform(id int) model.Customer {
	return c.customers[id]
}

func (c *CustomerService) Update(id int, name string, gender string, age int, phone string, email string) bool {

	customer := &c.customers[id]
	customer.Name = name
	customer.Gender = gender
	customer.Phone = phone
	customer.Email = email
	customer.Age = age
	return true
}
