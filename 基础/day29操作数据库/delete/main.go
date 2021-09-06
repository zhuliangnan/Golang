package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {

	database, err := sqlx.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/go-test")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database

}

func main() {

	/*
	   _, err := Db.Exec("delete from person where user_id=?", 1)
	   if err != nil {
	       fmt.Println("exec failed, ", err)
	       return
	   }
	*/
	defer Db.Close() // 注意这行代码要写在上面err判断的下面
	res, err := Db.Exec("delete from person where user_id=?", 2)
	if err != nil {
		fmt.Println("exec failed, ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("rows failed, ", err)
	}

	fmt.Println("delete succ: ", row)
}
