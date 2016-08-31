package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"fmt"
	"time"
	"com.dientm.bookstore/beans"
	"com.dientm.bookstore/service"
)


const SERVER_PORT string = ":9999";

func main() {
	e := echo.New()
	
	fmt.Println("This is ", time.Now())

	tikiBook := beans.NewBookStore("Tiki Book")
	tikiBook.InitBookStoreData();
	var bookservice service.BookStoreService

	e.GET("/book/:id", func(c echo.Context) error {
		id := c.Param("id")
		switch {
		case id == "tiki" :
			bookservice = service.NewBookStoreService(tikiBook)
		default:
			bookservice = service.NewBookStoreService(tikiBook)
		}
		return bookservice.ListBook(c)
	})


	e.POST("/save", func(c echo.Context) error {
		storeId := c.FormValue("storeId")
		title := c.FormValue("title")
		author := c.FormValue("author")
		desc := c.FormValue("desc")
		pubDate := c.FormValue("pubDate")
		switch {
		case storeId == "tiki" :
			bookservice = service.NewBookStoreService(tikiBook)
		default:
			bookservice = service.NewBookStoreService(tikiBook)
		}
		return bookservice.AddBookToStore(c, title, author, desc, pubDate)
/* test
{"storeId":"tiki","title":"Life of DienTM", "desc":"This is the book of my life", "pubDate":"25/03/2040"}
*/

	})


	fmt.Println("......Starting server at ", SERVER_PORT )
	e.Run(standard.New(SERVER_PORT))

}