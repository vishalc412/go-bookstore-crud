package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"github.com/vishalc412/go-bookstore-crud/pkg/models"
)

var NewBooks models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBook()
	res, err := json.Marshal(newBooks)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	bookId := param["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err == nil {

		bookDetails, _ := models.GetBookById(int(ID))
		res, err := json.Marshal(bookDetails)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	bookId := param["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err == nil {
		bookDetails := models.DeleteBook(int(ID))
		res, err := json.Marshal(bookDetails)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(res)
		}
	}

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	bookId := param["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err == nil {
		bookDetails, _ := models.GetBookById(int(ID))
		_ = json.NewDecoder(r.Body).Decode(&NewBooks)
		bookDetails.Name = NewBooks.Name
		bookDetails.Author = NewBooks.Author
		bookDetails.Publication = NewBooks.Publication
		res, _ := json.Marshal(bookDetails)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	_ = json.NewDecoder(r.Body).Decode(&NewBooks)
	bookDetails := NewBooks.CreateBook()
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
