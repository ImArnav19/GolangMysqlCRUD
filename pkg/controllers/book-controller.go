package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ImArnav/go-bookstore/pkg/models"
	"github.com/ImArnav/go-bookstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()                      //get books from models
	res, _ := json.Marshal(newBooks)                   //json.Marshal for converting it to json
	w.Header().Set("Content-Type", "pkglication/json") //Header
	w.WriteHeader(http.StatusOK)                       //Writeheader
	w.Write(res)                                       //result
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //mux.Vars(r)  access params

	id := vars["bookid"] //extract id

	Id, err := strconv.ParseInt(id, 0, 0) //conv tp string
	if err != nil {
		fmt.Println("Error while Parsing !")
	}
	bookDetails, _ := models.GetBookById(Id) //models.getBook, here _ refers we dont want db

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {

	CreateBook := &models.Book{}   //need a interface to be created
	utils.ParseBody(r, CreateBook) //parse function to unmarshal json
	b := CreateBook.CreateBook()   //get the book struct after storing in DB

	res, _ := json.Marshal(b) //marshal back to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["bookid"]

	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing!")
	}

	book := models.DeleteBook(Id)

	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{} //combination of GetBookById + saving it by changing files here
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	id := vars["bookid"]

	Id, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Error in Parsing!!")
	}

	bookDetails, db := models.GetBookById(Id)

	if UpdateBook.Name != "" {
		UpdateBook.Name = bookDetails.Name
	}
	if UpdateBook.Author != "" {
		UpdateBook.Author = bookDetails.Author
	}
	if UpdateBook.Publication != "" {
		UpdateBook.Publication = bookDetails.Publication
	}

	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
