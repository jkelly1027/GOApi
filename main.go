package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"errors"
)

// Books stored using a struct (Our data structure)
// Need to be able to be converted to JSON so the API can return the struct directly
// or it can take in the JSON version of the struct and convert it to a struct in go.
type Book struct {
	ID       string `json:"id"` // Field names must be capitalized to be exported
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"` // May be multiple of the same book
}

// Slice of books (Structure to represent the books as we are not using a database)
var books = []Book{ // This structure will need to be converted to json when we return it to the client (done above: `json:"id"`... etc)
	{ID: "1", Title: "The Alchemist", Author: "Paulo Coelho", Quantity: 10},
	{ID: "2", Title: "The Little Prince", Author: "Antoine de Saint-Exupéry", Quantity: 5},
	{ID: "3", Title: "The Da Vinci Code", Author: "Dan Brown", Quantity: 7},
}

// Handles the route of getting all of the different books (Returns the JSON version of the books) (Get Request)
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books) // IndentedJSON returns the JSON version of the books structure, with indentation
}

// Add book to slice
func createBook(c *gin.Context) {
	var newBook Book
	if err := c.BindJSON(&newBook); err != nil { // BindJSON binds the JSON version of the book to &newBook variable.
		return // If there is an error, return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook) // Returns the JSON version of the new book that was added to the slice
}

// Gets the book by ID
func bookById(c *gin.Context) {
	id := c.Param("id") // Param: Path parameter, returns the value of the specified parameter (/books/:id or /books/2)
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"}) // If the book is not found,
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// Returns the book object represented by an ID
func getBookByID(id string) (*Book, error) {
	for i, b := range books { // Iterates through the books slice, looking for the specified book
		if b.ID == id { // Check if the book id is equal to the id we passed to this function
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// Checkout book using the ID
func checkoutBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // What book we want to checkout (Query parameter ?id=1)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID Query Parameter"})
		return
	}

	Book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}

	if Book.Quantity <= 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "book out of stock"})
		return
	}

	Book.Quantity -= 1
	c.IndentedJSON(http.StatusOK, Book)
}

// Returning a Book
func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id") // What book we want to checkout (Query parameter ?id=1)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing ID Query Parameter"})
		return
	}

	Book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}

	Book.Quantity += 1
	c.IndentedJSON(http.StatusOK, Book)
}

func main() {
	router := gin.Default() // Creates a new router/server, to handle different routes/endpoints of our api, allows us to define the
	// routes and the functions that should be called when the routes are hit.

	// Endpoints
	router.GET("/books", getBooks)     // Routed to the getBooks function, which returns the JSON version of the books (localhost:8080/books)
	router.GET("/books/:id", bookById) // Routed to the bookById function, which returns the book with the specified ID (localhost:8080/books/1)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)

	router.Run("localhost:8080") // Runs the server on localhost:8080
}
