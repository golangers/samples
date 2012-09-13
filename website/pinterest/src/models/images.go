package models

import (
	. "golanger.com/framework/database/activerecord"
	. "golanger.com/framework/middleware"
)

type Images struct {
	Id         int64  `index:"PK" field:"id"`
	Class      int64  `field:"class"`
	Name       string `field:"name"`
	Ext        string `field:"ext"`
	Path       string `field:"path"`
	Status     int64  `field:"status"`
	CreateTime int64  `field:"create_time"`
}

func GetImagesLists(page ...int) (*[]Images, error) {
	var pg int
	if len(page) > 0 {
		pg = page[0] - 1
	}

	if pg < 1 {
		pg = 0
	}

	num := 20
	start := pg * num
	//OnDebug = true
	var orm = Middleware.Get("orm").(ActiveRecord)
	images := []Images{}
	err := orm.Where("status=?", 1).Limit(num, start).OrderBy("id DESC").FindAll(&images)

	return &images, err
}

func GetImage(id int) (*Images, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	var image Images
	err := orm.Where("id = ?", id).Find(&image)
	return &image, err
}

func SaveImages(images Images) (int64, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Save(&images)

	return resNum, err
}

func GetInvalidImages(page ...int) (*[]Images, error) {
	var pg int
	if len(page) > 0 {
		pg = page[0] - 1
	}

	if pg < 1 {
		pg = 0
	}

	num := 20
	start := pg * num
	var orm = Middleware.Get("orm").(ActiveRecord)
	images := []Images{}

	err := orm.Where("status=?", 0).Limit(num, start).OrderBy("id DESC").FindAll(&images)

	return &images, err
}

func DeleteImageWithId(id int64) (int64, error) {
	image := Images{
		Id: id,
	}
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Delete(&image)
	return resNum, err
}

func GetImagesListsWithClassId(classId int64, page ...int) (*[]Images, error) {
	var pg int
	if len(page) > 0 {
		pg = page[0] - 1
	}

	if pg < 1 {
		pg = 0
	}

	num := 20
	start := pg * num
	var orm = Middleware.Get("orm").(ActiveRecord)
	images := []Images{}

	err := orm.Where("status=? and class=?", 1, classId).Limit(num, start).OrderBy("id DESC").FindAll(&images)
	return &images, err
}

func GetImagesListsWithClassName(className string, page ...int) (*[]Images, error) {
	var pg int
	if len(page) > 0 {
		pg = page[0] - 1
	}

	if pg < 1 {
		pg = 0
	}

	num := 20
	start := pg * num
	var orm = Middleware.Get("orm").(ActiveRecord)
	images := []Images{}

	class, err := GetClassWithName(className)
	if err != nil {
		return nil, err
	}

	err = orm.Where("status=? and class=?", 1, class.Id).Limit(num, start).OrderBy("id DESC").FindAll(&images)
	return &images, err
}
