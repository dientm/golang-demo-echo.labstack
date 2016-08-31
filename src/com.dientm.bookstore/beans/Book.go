/* Book */

package beans

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	Desc string `json:"desc"`
	PublishDate string `json:"publishDate"`
}

func NewBook1(book Book) (b Book) {
	b = book
	return
}

func NewBook(title string, author string, desc string, publishDate string) (b Book) {
	b.Title = title
	b.Author = author
	b.Desc = desc
	b.PublishDate = publishDate
	return
}

func (b Book) GetTitle() string {
	return b.Title
}