/** BookService **/

package beans

import "fmt"

const BOOKSTORE_CAPACITY = 5;

type BookStore struct {
	Name string `json:"bookstore"`
	Size  int `json:"amount"`
	Books []Book `json:"bookList"`
}

func New() (bs *BookStore) {
	bs.Name = "Book Store"
	bs.Size = 0;
	bs.Books = make([]Book, 0, BOOKSTORE_CAPACITY)
	return
}

func NewBookStore(name string) (bs BookStore) {
	bs.Name = name
	bs.Size = 0
	bs.Books = make([]Book, BOOKSTORE_CAPACITY)
	return
}

func (bs *BookStore) ExtendSize() {
	
	l := len(bs.Books)
	newarray := make([]Book, l * 2)
	copy(newarray, bs.Books[:])
	bs.Books = newarray
	return
}

func (bs BookStore) GetBookAt(index int) (b Book) {
	return bs.Books[index]
}

func (bs BookStore) List() {

}

func (bs *BookStore) AddBook(b Book) {
	size := bs.Size
	len := len(bs.Books)

	if (size < len) {
		bs.Books[size] = b
	} else {
		
		bs.ExtendSize()
		bs.Books[size] = b
	}
	bs.Size += 1
	
	fmt.Println("Size: ", bs.Size)
	fmt.Println("Len: ", len)
	fmt.Println("Titile: ", b.GetTitle())
}

func (bs BookStore) Delete(id int) {

}

func (bs BookStore) Update(b Book) {

}

func (bs BookStore) GetBook(id int) {

}

func BookStoreInfo() {
	
}

func (bs *BookStore) InitBookStoreData() {

	fmt.Println("Initilize book store.......")
	book_0 := NewBook("Nhat Ban Den Va Yeu", "Duong Linh", "", "");
	book_1 := NewBook("Con Meo Trieu Kiep", "Sano Yoko", "", "");
	book_2 := NewBook("The Maze Runner Series", "James Dashner", "", "");
	book_3 := NewBook("1Q84", "Haruki Murakami", "", "");
	book_4 := NewBook("Cau Khong Lo Ti Hon", "Janet Foxley", "", "");
	book_5 := NewBook("Hoang Tu Hanh Phuc", "Oscar Wilde", "", "");
	book_6 := NewBook("Chiec Hop Giang Sinh", "Richard Paul Evans", "", "");
	book_7 := NewBook("Bon Mua, Troi va Dat", "Marai Sandor", "", "");
	
	bs.AddBook(book_0);
	bs.AddBook(book_1);
	bs.AddBook(book_2);
	bs.AddBook(book_3);
	bs.AddBook(book_4);
	bs.AddBook(book_5);
	bs.AddBook(book_6);
	bs.AddBook(book_7);

	fmt.Println("Done InitBookStoreData() <-------")

}