package main

import (
    "context"
    "github.com/go-kit/kit/log"
)

//structs
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
    AuthorId    	string `json:"authorId,omitempty"`
    Name     		string `json:"name,omitempty"`
    Nationality   	string `json:"nacionality,omitempty"`
    Birth 			string `json:"birth,omitempty"`
    Genere  		string `json:"genere,omitempty"`
}

type bookservice struct {
    logger log.Logger
}

type publisherservice struct {
    logger log.Logger
}

type authorservice struct {
    logger log.Logger
}

// Services
type BookService interface {
    CreateBook(ctx context.Context, book Book) (string, error)
    GetBookById(ctx context.Context, id string) (interface{}, error)
    UpdateBook(ctx context.Context, book Book) (string, error)
    DeleteBook(ctx context.Context, id string) (string, error)
}

type PublisherService interface {
    CreatePublisher(ctx context.Context, publisher Publisher) (string, error)
    GetPublisherById(ctx context.Context, id string) (interface{}, error)
    UpdatePublisher(ctx context.Context, publisher Publisher) (string, error)
    DeletePublisher(ctx context.Context, id string) (string, error)
}

type AuthorService interface {
    CreateAuthor(ctx context.Context, author Author) (string, error)
    GetAuthorById(ctx context.Context, id string) (interface{}, error)
    UpdateAuthor(ctx context.Context, author Author) (string, error)
    DeleteAuthor(ctx context.Context, id string) (string, error)
}

//Data
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

//Books

func NewServiceB(logger log.Logger) BookService {
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

//PublisherService

func NewServiceP(logger log.Logger) PublisherService {
    return &publisherservice{
        logger: logger,
    }
}

func (s publisherservice) CreatePublisher(ctx context.Context, publisher Publisher) (string, error) {
    var msg = "success"
    publishers = append(publishers, publisher)
    return msg, nil
}

func (s publisherservice) GetPublisherById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var publisher interface{}
    var empty interface{}
    i := findPublishers(id)
    if i == -1 {
        return empty, err
    }
    publisher = publishers[i]
    return publisher, nil
}

func (s publisherservice) DeletePublisher(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findPublishers(id)
    if i == -1 {
        return "", err
    }
    copy(publishers[i:], publishers[i+1:])
    publishers[len(publishers)-1] = Publisher{}
    publishers = publishers[:len(publishers)-1]
    return msg, nil
}

func (s publisherservice) UpdatePublisher(ctx context.Context, publisher Publisher) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findPublishers(publisher.PublisherId)
    if i == -1 {
        return empty, err
    }
    publishers[i] = publisher
    return msg, nil
}

//AuthorService

func NewServiceA(logger log.Logger) AuthorService {
    return &authorservice{
        logger: logger,
    }
}

func (s authorservice) CreateAuthor(ctx context.Context, author Author) (string, error) {
    var msg = "success"
    authors = append(authors, author)
    return msg, nil
}

func (s authorservice) GetAuthorById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var author interface{}
    var empty interface{}
    i := findAuthors(id)
    if i == -1 {
        return empty, err
    }
    author = authors[i]
    return author, nil
}
func (s authorservice) DeleteAuthor(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findAuthors(id)
    if i == -1 {
        return "", err
    }
    copy(authors[i:], authors[i+1:])
    authors[len(authors)-1] = Author{}
    authors = authors[:len(authors)-1]
    return msg, nil
}
func (s authorservice) UpdateAuthor(ctx context.Context, author Author) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findAuthors(author.AuthorId)
    if i == -1 {
        return empty, err
    }
    authors[i] = author
    return msg, nil
}