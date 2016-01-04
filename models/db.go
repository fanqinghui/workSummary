package models

import (
	_ "database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"workSummary/utils"
)

type User struct {
	Id         int
	LoginName  string
	Password   string
	UserName   string
	Email      string
	createData string
}

func Login(loginName, password string) (User, bool) {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT id,loginName,password,userName,email FROM t_user where loginName=? and password=?", loginName, password)
	var user = new(User)

	err := row.Scan(&user.Id, &user.LoginName, &user.Password, &user.UserName, &user.Email)

	if err != nil {
		fmt.Println(err)
		return *user, false
	}

	fmt.Println(user.Id)

	if user.Id > 0 {
		return *user, true
	}

	return *user, false
}

func InsertUser(user User) (User, bool) {
	db := OpenDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT t_user (loginName,password,userName,email) values (?,?,?,?)")
	defer stmt.Close()
	utils.CheckErr(err)
	res, err := stmt.Exec(user.LoginName, user.Password, user.UserName, user.Email)
	utils.CheckErr(err)
	id, err := res.LastInsertId()
	utils.CheckErr(err)
	fmt.Println(id)
	return user, true
}
