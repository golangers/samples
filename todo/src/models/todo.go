package models

import (
	"database/sql"
	. "golanger.com/middleware"
	"strconv"
)

type Todo struct {
	Id       int
	Title    string
	Finished bool
	PostDate string
}

func GetTodoLists() (*[]Todo, error) {
	var db = Middleware.Get("db").(*sql.DB)
	todos := []Todo{}
	sql := "SELECT * FROM `todo` ORDER BY `finished` asc, `id` desc"
	rows, err := db.Query(sql)
	for rows.Next() {
		todo := Todo{}
		rows.Scan(&todo.Id, &todo.Title, &todo.Finished, &todo.PostDate)
		todos = append(todos, todo)
	}

	return &todos, err
}

func GetTodo(id int) Todo {
	var db = Middleware.Get("db").(*sql.DB)
	todo := Todo{}
	sql := "SELECT * FROM `todo` WHERE `id` = " + strconv.Itoa(id)
	row := db.QueryRow(sql)
	row.Scan(&todo.Id, &todo.Title, &todo.Finished, &todo.PostDate)

	return todo
}

func SaveTodo(todo Todo) (sql.Result, error) {
	var db = Middleware.Get("db").(*sql.DB)
	sql := "INSERT INTO `todo` (`title`,`post_date`) VALUES(\"" + todo.Title + "\",\"" + todo.PostDate + "\")"
	r, err := db.Exec(sql)

	return r, err
}

func UpdateTodo(todo Todo) (sql.Result, error) {
	var db = Middleware.Get("db").(*sql.DB)
	sql := "UPDATE `todo` SET `title` = \"" + todo.Title + "\", `post_date` = \"" + todo.PostDate + "\" WHERE `id` = " + strconv.Itoa(todo.Id)
	r, err := db.Exec(sql)

	return r, err
}

func DeleteTodo(id int) (sql.Result, error) {
	var db = Middleware.Get("db").(*sql.DB)
	sql := "DELETE FROM `todo` WHERE `id` = " + strconv.Itoa(id)
	r, err := db.Exec(sql)

	return r, err
}

func InitTodo() {
	var db = Middleware.Get("db").(*sql.DB)
	sql := "DELETE FROM `todo`"
	db.Exec(sql)
	sql = "UPDATE sqlite_sequence SET seq = 0 WHERE name = 'todo'"
	db.Exec(sql)
}
