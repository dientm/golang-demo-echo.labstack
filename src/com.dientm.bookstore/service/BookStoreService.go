/* compose service for book store */

// package bookstoreservice
package service

import "github.com/labstack/echo"
import "net/http"
import "com.dientm.bookstore/beans"
import "fmt"
type BookStoreService struct {
	b beans.BookStore
}

func NewBookStoreService(b beans.BookStore) (bs BookStoreService) {
	bs.b = b
	return
}

func (bs BookStoreService) ListBook(c echo.Context) error {
	
	return c.JSON(http.StatusOK, bs.b)
}

func (bs BookStoreService) AddBookToStore(c echo.Context, title string, author string, desc string, pubDate string) error {
	fmt.Println("Title: ", title)

	newbook := beans.NewBook(title, author, desc, pubDate)
	bs.b.AddBook(newbook)
	return c.JSON(http.StatusCreated, beans.Message{"Add successfully"})
}