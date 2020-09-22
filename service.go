package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

type Book struct {
    BookId    string `json:"bookId,omitempty"`
    Title     string `json:"title,omitempty"`
    Edition   string `json:"edition,omitempty"`
    Copyright string `json:"copyright,omitempty"`
    Language  string `json:"language,omitempty"`
    Pages     string `json:"pages,omitempty"`
    Author    string `json:"author,omitempty"`
    Publisher string `json:"publisher,omitempty"`
}

type Publisher struct {
    PublisherId    string `json:"publisherId,omitempty"`
    Name     string `json:"name,omitempty"`
    Country   string `json:"country,omitempty"`
    Founded string `json:"founded,omitempty"`
    Genere  string `json:"genere,omitempty"`
}

type Author struct {
    AuthorId    string `json:"authorId,omitempty"`
    Name     string `json:"name,omitempty"`
    Nationality   string `json:"nacionality,omitempty"`
    Birth string `json:"birth,omitempty"`
    Genere  string `json:"genere,omitempty"`
}

type bookservice struct {
    logger log.Logger
}

// Service describes the Book service.
type BookService interface {
    CreateBook(ctx context.Context, book Book) (string, error)
    GetBookById(ctx context.Context, id string) (interface{}, error)
    UpdateBook(ctx context.Context, book Book) (string, error)
    DeleteBook(ctx context.Context, id string) (string, error)
}

var books = []Book{
    Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
        Copyright: "2012", Language: "ENGLISH", Pages: "976",
        Author: "Author2", Publisher: "Publisher2"},
    Book{BookId: "Book2", Title: "Computer Networks", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        Author: "Author1", Publisher: "Publisher1"},
}

var authors = []Author{
    Author{AuthorId: "Author1", Name: "Andrew S. Tanenbaum", Nationality: "Canada",
        Birth: "1977", Genere: "Programming"},
    Author{AuthorId: "Author2", Name: "Abraham Silberschatz", Nationality: "US",
        Birth: "1949", Genere: "Programming"},
}

var publishers = []Publisher{
    Publisher{PublisherId: "Publisher1", Name: "John Wiley & Sons", Country: "UK",
        Founded: "1994", Genere: "Programming"},
    Publisher{PublisherId: "Publisher2", Name: "Mc Graw Hill", Country: "US",
        Founded: "1919", Genere: "Teaching"},
}

//Finders

func findBooks(x string) int {
    for i, book := range books {
        if x == book.BookId {
            return i
        }
    }
    return -1
}

func findPublishers(x string) int {
    for i, publisher := range publishers {
        if x == publisher.PublisherId {
            return i
        }
    }
    return -1
}

func findAuthors(x string) int {
    for i, author := range authors {
        if x == author.AuthorId {
            return i
        }
    }
    return -1
}

func findBooksbyAuthors(x string) []int {
	y:=[]int{}
    for i, _ := range books {
		for _,z := range books[i].Author{
			if z == x {
				y = append(y, i)
			}
		} 
    }
    return y
}

func findBooksbyPublishers (x string) []int {
	y:=[]int{}
    for i, book := range books {
		if x == book.Publisher {
			y = append(y, i)
		}
    }
    return y
}


func NewService(logger log.Logger) BookService {
    return &bookservice{
        logger: logger,
    }
}

func (s bookservice) CreateBook(ctx context.Context, book Book) (string, error) {
    var msg = "success"
    books = append(books, book)
    return msg, nil
}

func (s bookservice) GetBookById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var book interface{}
    var empty interface{}
    i := findBooks(id)
    if i == -1 {
        return empty, err
    }
    book = books[i]
    return book, nil
}
func (s bookservice) DeleteBook(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findBooks(id)
    if i == -1 {
        return "", err
    }
    copy(books[i:], books[i+1:])
    books[len(books)-1] = Book{}
    books = books[:len(books)-1]
    return msg, nil
}
func (s bookservice) UpdateBook(ctx context.Context, book Book) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findBooks(book.BookId)
    if i == -1 {
        return empty, err
    }
    books[i] = book
    return msg, nil
}