package mySql

import (
	"database/sql"
	"log"

	_ "github.com/go-SQL-Driver/MySQL"
)

const (
	DB_HOST = "tcp(127.0.0.1:3306)"
	DB_NAME = "pruebago"
	DB_USER = "root"
	DB_PASS = "root"
)

func ExecuteQuerys(query string) {
	//fmt.Println(query)
	dsn := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Exec(query)
}
