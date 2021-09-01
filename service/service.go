package service

import (
	"encoding/json"
	"github.com/aaaaaaaaaaa/config"
	"github.com/aaaaaaaaaaa/entities"
	"log"
)

func GetAllBook()[]entities.Stock{
	db := config.Connection()
	var books []entities.Stock

	rows, err := db.Query("SELECT * FROM stock")
	if err != nil{
		panic(err.Error())
	}

	for rows.Next(){
		var book entities.Stock
		err := rows.Scan(&book.Id, &book.Name,&book.IBSN,&book.Author,&book.Quantity,&book.Price,&book.Store)
		if err != nil{
			log.Fatal(err)
		}

		books = append(books, book)
	}

	defer rows.Close()

	_ , jsonError := json.Marshal(books)
	if jsonError != nil{
		panic(jsonError.Error())
	}
	return books
}

func GetSingleBook(bookId string) entities.Stock{
	db := config.Connection()
	book, queryError := db.Query("SELECT * FROM stock WHERE id =?",bookId)
	if queryError != nil{
		panic(queryError.Error())
	}
	var response entities.Stock
	for book.Next(){

		err := book.Scan(&response.Id, &response.Name,&response.IBSN,&response.Author,&response.Quantity,&response.Price,&response.Store);
		if err != nil{
			log.Fatal(err)
		}
	}
	return response
}

func InsertNewBook(newBook entities.Stock) bool{
	db := config.Connection()
	stmt, err := db.Prepare("INSERT INTO stock(nome,ibsn,autor,quantidade,preco,loja) VALUES(?,?,?,?,?,?);")
	if err != nil{
		panic(err.Error())
	}

	_, queryError := stmt.Exec(newBook.Name, newBook.IBSN, newBook.Author, newBook.Quantity, newBook.Price, newBook.Store)
	if queryError != nil{
		panic(queryError.Error())

	}

	defer db.Close()
	return true
}

func UpdateBook(book entities.Stock, bookId string) bool{
	db := config.Connection()
	stmt, err := db.Prepare("UPDATE stock SET nome=?,ibsn=?, autor=?, quantidade=?, preco=?, loja=? WHERE id=?")
	if err != nil{
		panic(err.Error())
	}
	_,err2 := stmt.Exec(book.Name,book.IBSN,book.Author,book.Quantity, book.Price, book.Store,bookId)
	if err2 != nil{
		panic(err2.Error())
	}
	defer db.Close()
	return true
}

func DeleteBook(bookId string) entities.Stock{
	var book entities.Stock
	book = GetSingleBook(bookId)
	db := config.Connection()
	stmt, err := db.Prepare("DELETE FROM stock WHERE id =?")

	if err != nil{
		panic(err.Error())
	}
	_,queryError :=stmt.Exec(bookId)
	if queryError != nil{
		panic(queryError.Error())
	}
	return book
}