package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//book represents a book about a book data.
type book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Writer string  `json:"writer"`
	Price  float64 `json:"price"`
}

// books slice to seed book data.
var books = []book{
	{ID: "1", Title: "Dune", Writer: "Frank Herbert", Price: 49.68},
	{ID: "2", Title: "Persepolis", Writer: "Marjane Satrapi", Price: 35.90},
	{ID: "3", Title: "Maus", Writer: "Art Spiegelman", Price: 49.99},
}

//getBooks responds with the list of all books as JSON string

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// Gin Context: It carries request details, validates and serializes JSON, and more.
// Context.IndentedJSON: serialize the struct into JSON and add it to the response.

func postBooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusOK, books)
}

func getBookByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range books {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", postBooks)
	router.GET("/books/:id", getBookByID)

	router.Run("localhost:8080")
}

// Initializa a Gin router using Default
// Use GET function to associates the GET HTTP method and /books path with a handler function
// Use the Run function to attach the router to an htt.Server and start the server
