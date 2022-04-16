package main

import (
	"homework-3-week-5-boyrazhamdi/author"
	"homework-3-week-5-boyrazhamdi/book"
	"homework-3-week-5-boyrazhamdi/database"
	"net/http"

	"log"

	_ "homework-3-week-5-boyrazhamdi/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// @title Book API
// @version 1.0
// @description This is a sample an exercise written in Go.

// @contact.name Hamdi

// @host localhost:6363

// @accept json
// @produce json

// @BasePath /api
// @schemes http https

func main() {
	// Set env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	repo, err := database.NewRepositories()
	if err != nil {
		log.Fatal("Postgres cannot init: ", err)
	}
	/*
		repo.Migration()

		book := utils.BookCsvToStruct("./dummy.csv")

		for _, i := range book {
			id, err := repo.Book.Create(i)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id)
		}

		author := utils.AuthorCsvToStruct("./dummy1.csv")

		for _, i := range author {
			id, err := repo.Author.Create(i)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id)
		}
	*/

	bookService := book.NewBookService(repo.Book)
	authorService := author.NewAuthorService(repo.Author)

	bookHandler := book.NewBookHandler(bookService)
	authorHandler := author.NewAuthorHandler(authorService)

	router := mux.NewRouter()
	router.HandleFunc("/api/books", bookHandler.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/search/{word}", bookHandler.SearchBooks).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}", bookHandler.GetBookByID).Methods(http.MethodGet)
	router.HandleFunc("/api/authors/{id}", authorHandler.GetAuthorByID).Methods(http.MethodGet)
	router.HandleFunc("/api/books/{id}/sell/{quantity}", bookHandler.SellBook).Methods(http.MethodPut)
	router.HandleFunc("/api/books/{id}", bookHandler.DeleteBook).Methods(http.MethodDelete)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(":6363", router)

}
