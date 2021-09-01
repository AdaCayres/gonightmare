package controller

import (
	"encoding/json"
	"fmt"
	"github.com/aaaaaaaaaaa/entities"
	"github.com/aaaaaaaaaaa/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request){
	var httpError = entities.ErrorResponse{
		Code: http.StatusNotFound, Message: "NO BOOK FOUND IN THE DATABASE",
	}
	jsonResponse := service.GetAllBook()
	if jsonResponse == nil{
		ReturnErrorResponse(w,r, httpError)
	}else{
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(jsonResponse)
	}
}

func GetOne(w http.ResponseWriter, r *http.Request){
	var httpError = entities.ErrorResponse{
		Code: http.StatusNotFound, Message: "NO BOOK FOUND IN THE DATABASE",
	}
	bookId := mux.Vars(r)["id"]
	jsonResponse := service.GetSingleBook(bookId)
	w.Header().Set("Content-Type", "application/json")
	if (jsonResponse == entities.Stock{}){
		ReturnErrorResponse(w,r,httpError)
	}else {
		json.NewEncoder(w).Encode(jsonResponse)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var httpError = entities.ErrorResponse{
		Code: http.StatusInternalServerError,Message: "CAN'T INSERT BOOK IN THE DATABASE",
	}
	var newBook entities.Stock
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newBook)
	if err != nil{
		ReturnErrorResponse(w,r, httpError)
	}else {
		httpError.Code = http.StatusBadRequest
		if newBook.Name == ""{
			httpError.Message = "Name cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else if newBook.IBSN == "" {
			httpError.Message = "IBSN cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else if newBook.Author == ""{
			httpError.Message = "Author cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else if newBook.Quantity == "" {
			httpError.Message = "Quantity cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else if newBook.Price == ""{
			httpError.Message = "Price cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else if newBook.Store == ""{
			httpError.Message = "Store cannot be null"
			ReturnErrorResponse(w, r, httpError)
		}else {
			isInsert := service.InsertNewBook(newBook)
			if isInsert{
				jsonResponse := service.GetAllBook()
				if jsonResponse == nil{
					ReturnErrorResponse(w,r,httpError)
				}else {
					w.Header().Set("Content-Type", "application/json")
					//w.Write(jsonResponse)
					json.NewEncoder(w).Encode(jsonResponse)
				}
			}else {
				ReturnErrorResponse(w, r, httpError)
			}
		}
	}
}

func Delete(w http.ResponseWriter, r *http.Request){
	var httpError = entities.ErrorResponse{
		Code: http.StatusNotFound,Message: "CAN'T DELETE BOOK IN THE DATABASE",
	}

	bookId := mux.Vars(r)["id"]

	if bookId == ""{
		httpError.Message = "book ID cannot be empty"
		ReturnErrorResponse(w,r,httpError)
	}else{
		deletedBook := service.DeleteBook(bookId)
		if (deletedBook == entities.Stock{}){
			ReturnErrorResponse(w,r,httpError)

		}else{
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(deletedBook)
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request){
	var httpError404 = entities.ErrorResponse{
		Code: http.StatusNotFound, Message: "CAN'T FIND BOOK IN THE DATABASE",
	}
	var httpError500 = entities.ErrorResponse{
		Code: http.StatusInternalServerError, Message: "CAN'T UPDATE BOOK IN THE DATABASE DUE TO SERVER ERROR",
	}
	bookId := mux.Vars(r)["id"]

	if bookId == ""{
		httpError500.Message = "book ID cannot be empty"
		ReturnErrorResponse(w,r,httpError500)
	}else{

		var book entities.Stock
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		defer r.Body.Close()
		if err != nil{
			ReturnErrorResponse(w,r,httpError500)
		}else{
			worked := service.UpdateBook(book,bookId)
			if worked{

				jsonResponse := service.GetSingleBook(bookId)
				if (jsonResponse == entities.Stock{}){
					fmt.Print(bookId)
					ReturnErrorResponse(w,r,httpError404)
				}else{
					w.Header().Set("Content-Type","application/json")
					json.NewEncoder(w).Encode(jsonResponse)
				}
			}
		}

	}
}

func ReturnErrorResponse(response http.ResponseWriter, request *http.Request, errorMessage entities.ErrorResponse) {
	httpResponse := &entities.ErrorResponse{Code: errorMessage.Code, Message: errorMessage.Message}
	jsonResponse, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(errorMessage.Code)
	response.Write(jsonResponse)
}