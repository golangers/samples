package models

import (
	. "golanger.com/database/activerecord"
	. "golanger.com/middleware"
)

type Class struct {
	Id   int64  `index:"PK" field:"id"`
	Name string `field:"name"`
}

func AddClass(class Class) (int64, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Save(&class)

	return resNum, err
}

func EditClassWithId(id int64, name string) error {
	class := Class{
		Id:   id,
		Name: name,
	}
	var orm = Middleware.Get("orm").(ActiveRecord)
	_, err := orm.Save(&class)
	return err
}

func DeleteClassWithId(id int64) (int64, error) {
	class := Class{
		Id: id,
	}
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Delete(&class)
	return resNum, err
}

func GetClasses() (*[]Class, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	classes := []Class{}

	err := orm.FindAll(&classes)
	return &classes, err
}

func GetClass(id int64) (*Class, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	var class Class
	err := orm.Where("id = ?", id).Find(&class)
	return &class, err
}

func GetClassWithName(name string) (*Class, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	var class Class
	err := orm.Where("name = ?", name).Find(&class)
	return &class, err
}
