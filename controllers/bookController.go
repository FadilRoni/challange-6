package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID string `json: "book_id"`
	Title string `json: "title"`
	Author string `json: "author"`
	Desc string `json: "desc"`
}

var Books = []Book{}

func AddBook(ctx *gin.Context) {
	var addBook Book

	if err := ctx.ShouldBindJSON(&addBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	addBook.ID = fmt.Sprintf("%d", len(Books) + 1)
	Books = append(Books, addBook)

	ctx.JSON(http.StatusCreated, "Created")
}

func GetBook(ctx *gin.Context) {
	ID := ctx.Param("ID")
	condition := false
	var bookData Book

	for i, book := range Books {
		if ID == book.ID {
			condition = true
			bookData = Books[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": http.StatusNotFound,
			"error_message": fmt.Sprintf("Buku dengan id %v tidak ditemukan", ID),
		})

		return
	}

	ctx.JSON(http.StatusOK, bookData)
}

func GetBooks(ctx *gin.Context) {

	if len(Books) <= 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": http.StatusNotFound,
			"error_message": "Belum ada Data Buku",
		})

		return
	}

	ctx.JSON(http.StatusOK, Books)

}

func UpdateBook(ctx *gin.Context) {
	ID := ctx.Param("ID")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range Books {
		if ID == book.ID {
			condition = true
			Books[i] = updateBook
			Books[i].ID = ID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": http.StatusNotFound,
			"error_message": fmt.Sprintf("Buku dengan id %v tidak ditemukan", ID),
		})

		return
	}

	ctx.JSON(http.StatusOK, "Updated")

}

func DeleteBook(ctx *gin.Context) {
	ID := ctx.Param("ID")
	condition := false
	var bookIndex int

	for i, book := range Books {
		if ID == book.ID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": http.StatusNotFound,
			"error_message": fmt.Sprintf("Buku dengan id %v tidak ditemukan", ID),
		})

		return
	}

	copy(Books[bookIndex:], Books[bookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, "Deleted")

}