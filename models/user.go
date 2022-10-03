package models

import (
	"fmt"
	"server/db"
	"time"
)

type User struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const UsersTable string = "users"

const UsersSchema string = `CREATE TABLE users (
	id INT(6) UNIQUE INDEX,
	user_name VARCHAR(30) NOT NULL,
	password VARCHAR(100) NOT NULL,
	email 	 VARCHAR(50),
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP NULL,
	deleted_at TIMESTAMP NULL
)`

func (user *User) insert() {
	sql := "INSERT users SET id=? user_name=?, password=?, email=?  "
	res, _ := db.Exec(sql, user.Id, user.UserName, user.Password, user.Email)
	user.Id, _ = res.LastInsertId()
}

func (user *User) update() {
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE users SET id=? user_name=?, password=?, email=?, updated_at=?  WHERE id = ? "
	db.Exec(sql, user.Id, user.UserName, user.Password, user.Email, now, user.Id)
}

func (user *User) Save() {
	if user.Id == 0 {
		user.insert()
	} else {
		user.update()
	}
}

func (user *User) Delete() {
	if user.Id != 0 {
		DeleteUser(int(user.Id))
	}
}

func GetLastId(table string) (int64, error) {
	sql := fmt.Sprintf("SELECT id FROM %s ORDER BY id DESC LIMIT 1", table)
	var id int64
	rows, err := db.Query(sql)
	if err != nil {
		return id, err
	} else {
		for rows.Next() {
			rows.Scan(id)
		}
		return id, nil
	}
}

func NewUser(id int64, name, pass, email string) *User {

	user := &User{Id: id, UserName: name, Password: pass, Email: email}
	return user
}

func CreateUser(name, pass, email string) *User {
	id, _ := GetLastId(UsersTable)
	id = id + 1
	user := NewUser(id, name, pass, email)
	user.insert()
	return user
}

func ListUsers() ([]User, error) {
	sql := "SELECT id, user_name, password, email FROM `users` WHERE deleted_at IS NULL "
	users := []User{}
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	} else {
		for rows.Next() {
			user := User{}
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
			users = append(users, user)
		}
		return users, nil
	}

}

func GetUserId(id int) (User, error) {
	sql := "SELECT id, user-name, password, email FROM `users` WHERE id = ? AND deleted_at IS NULL "
	user := User{}
	rows, err := db.Query(sql, id)
	if err != nil {
		return user, err
	} else {
		for rows.Next() {
			rows.Scan(&user.Id, &user.UserName, &user.Password, &user.Email)
		}
		return user, nil
	}
}

func DeleteUser(id int) {
	now := time.Now().Format("2006-01-02 15:04:05")
	sql := "UPDATE users SET deleted_at=? WHERE id = ? "
	db.Exec(sql, now, id)
}
