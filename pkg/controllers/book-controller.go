package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/raghavgh/bookmanagement/pkg/models"
	"github.com/raghavgh/bookmanagement/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

// GetBooks calls internal models function GetAllBooks and generate response
func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	utils.SetResponse(http.StatusOK, newBooks, map[string]interface{}{"content-type": "application/json"}, &w, nil)
}

// GetBookById fetch id from request and call internal models function to getBookById and generate response
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	utils.SetResponse(http.StatusOK, bookDetails, map[string]interface{}{"content-type": "application/json"}, &w, nil)
}

// CreateBook fetch boo data from request and call internal odels function createBook and generate response
func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	book := createBook.CreateBook()
	utils.SetResponse(http.StatusOK, book, map[string]interface{}{"content-type": "application/json"}, &w, nil)
}

// DeleteBook fetch id from request and call internal models delete func and generate response
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	utils.SetResponse(http.StatusOK, book, map[string]interface{}{"content-type": "application/json"}, &w, nil)
}

// UpdateBook fetch newbook data and id from request and do folowing things
// 		1. call internal models func getBookBy id to get book that we need to update
// 		2. updates book fields and save into db
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(ID)
	utils.CopyBookModelIfNotEmpty(updateBook, bookDetails)
	db.Save(bookDetails)
	utils.SetResponse(http.StatusOK, bookDetails, map[string]interface{}{"content-type": "application/json"}, &w, nil)
}
