package config

import (
"database/sql"
"fmt"
_ "github.com/go-sql-driver/mysql"
"log"
)

var USER = "root"
var PASS = "password"
var HOST = "localhost"
var PORT = "3309"
var DBNAME = "db_stock"

func CreateDatabase(){

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/", USER, PASS, HOST, PORT)
	db, err := sql.Open("mysql", URL)

	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("create database if not exists db_stock")
	if err != nil{
		panic(err.Error())
	}else {
		log.Println("Criou Banco")
	}

	_, err = db.Exec("use `db_stock`")
	if err != nil{
		panic(err.Error())
	}else {
		log.Println("Usando Banco")
	}

	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS stock(
		id INTEGER AUTO_INCREMENT NOT NULL,
		nome VARCHAR(250),
		ibsn VARCHAR(250),
		autor VARCHAR(250),
		quantidade VARCHAR(250),
		preco VARCHAR(250),
		loja VARCHAR(250),
		PRIMARY KEY (id))`)

	if err != nil{
		panic(err.Error())
	}else {
		log.Println("Preparou tabela")
	}
	_, err = stmt.Exec()

	if err != nil{
		panic(err.Error())
	}else {
		log.Println("Criou tabela")
	}
	defer db.Close()
}

func Connection() *sql.DB{

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := sql.Open("mysql", URL)
	if err != nil{
		panic(err.Error())
	}
	return db
}

