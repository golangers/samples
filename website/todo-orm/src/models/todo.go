package models

import (
	"database/sql"
	. "golanger.com/database/activerecord"
	. "golanger.com/middleware"
)

type Todo struct {
	Id       int64  `index:"PK" field:"id"`
	Title    string `field:"title"`
	Finished bool   `field:"finished"`
	PostDate string `field:"post_date"`
}

func GetTodoLists() (*[]Todo, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	todos := []Todo{}
	err := orm.OrderBy("finished ASC, id DESC").FindAll(&todos)

	return &todos, err
}

func GetTodo(id int64) Todo {
	var orm = Middleware.Get("orm").(ActiveRecord)
	todo := Todo{}
	orm.Where("id = ?", id).Find(&todo)

	return todo
}

func SaveTodo(todo Todo) (int64, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Save(&todo)

	return resNum, err
}

func UpdateTodo(todo Todo) (int64, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.Save(&todo)

	return resNum, err
}

func DeleteTodo(id int64) (int64, error) {
	var orm = Middleware.Get("orm").(ActiveRecord)
	resNum, err := orm.SetTable("todo").Where("id = ?", id).DeleteRow()
	return resNum, err
}

func InitTodo() {
	var db = Middleware.Get("db").(*sql.DB)
	sql := "DELETE FROM `todo`"
	db.Exec(sql)
	sql = "UPDATE sqlite_sequence SET seq = 0 WHERE name = 'todo'"
	db.Exec(sql)
}
