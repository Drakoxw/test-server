package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const url = "root:@tcp(localhost:3306)/go_db"

var db *sql.DB

func Connect() {
	conx, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	// fmt.Println("conexion a base de datos exitosa")
	db = conx
}

func Close() {
	db.Close()
}

// VERIFICA LA CONEXION A BASE DE DATOS
func VerifyPing() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func ExistsTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := Query(sql)
	if err != nil {
		println("Error:", err)
		return false
	}
	exist := rows.Next()
	if exist {
		println("Ya existe la tabla:", tableName)
	}
	return exist

}

// CREA UNA TABLA
func CreateTable(schema string, name string) {
	if !ExistsTable(name) {
		ok, err := Exec(schema)
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}
}

func Exec(query string, arg ...interface{}) (sql.Result, error) {
	Connect()
	res, err := db.Exec(query, arg...)
	Close()

	if err != nil {
		fmt.Println(err)
	}
	return res, err
}

func Query(query string, arg ...interface{}) (*sql.Rows, error) {
	Connect()
	res, err := db.Query(query, arg...)
	Close()
	if err != nil {

		fmt.Println(err)
	}
	return res, err
}
