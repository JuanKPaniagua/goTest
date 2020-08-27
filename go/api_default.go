package swagger

import (
	"encoding/json"
    "net/http"
    "path"
	"fmt"
	"strings"
)

//Variables

var books = []Book{
    Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
        Copyright: "2012", Language: "ENGLISH", Pages: "976",
        Author: []string{"Author2"}, Publisher: "Publisher2"},
    Book{BookId: "Book3", Title: "Computer Networks", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        Author: []string{"Author1"}, Publisher: "Publisher1"},
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

//Delete

func deleteBook(x string) int {
    for i, book := range books {
        if x == book.BookId {
			copy(books[i:], books[i+1:]) // Shift a[i+1:] left one index.
			books = books[:len(books)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

func deleteAuthor(x string) int {
    for i, author := range authors {
        if x == author.AuthorId {
			copy(authors[i:], authors[i+1:]) // Shift a[i+1:] left one index.
			authors = authors[:len(authors)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

func deletePublisher(x string) int {
    for i, publisher := range publishers {
        if x == publisher.PublisherId {
			copy(publishers[i:], publishers[i+1:]) // Shift a[i+1:] left one index.
			publishers = publishers[:len(publishers)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

//Books

func BooksPut(w http.ResponseWriter, r *http.Request) {
	var book Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	lAuthors := []string{}
	for _,authorB := range book.Author{
		k := findAuthors(authorB)
		if k != -1 {
			lAuthors = append(lAuthors, authorB)
		}
	}
	if(len(lAuthors) == 0){
		return;
	}
	book.Author = lAuthors
	if(findPublishers(book.Publisher) == -1){
		return;
	}
    books = append(books, book)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func BooksGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
    i := findBooks(id)
    if i == -1 {
        return
    }
    dataJson, _ := json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdPost(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findBooks(id)
	if i == -1 {
		return
	}	
	leng := r.ContentLength
    body := make([]byte, leng)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
	if book.Title != ""{
		books[i].Title=book.Title
	}
	if book.Edition != ""{
		books[i].Edition=book.Edition
	}
	if book.Copyright != ""{
		books[i].Copyright=book.Copyright
	}
	if book.Language != ""{
		books[i].Language=book.Language
	}
	if book.Pages != ""{
		books[i].Pages=book.Pages
	}
	if len(book.Author) != 0{
		lAuthors := []string{}
		for _,authorB := range book.Author{
			k := findAuthors(authorB)
			if k != -1 {
				lAuthors = append(lAuthors, authorB)
			}
		}
		if(len(lAuthors) != 0){
			fmt.Printf(lAuthors[0])
			fmt.Printf(lAuthors[1])
			books[i].Author=book.Author
		}
	}
	if (book.Publisher != "") && (findPublishers(book.Publisher) == 0){
		books[i].Publisher=book.Publisher
	}
	dataJson,_ := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func BooksBookIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := deleteBook(id)
	if i == -1 {
		return
	}
	dataJson,_ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdAuthorsGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = strings.Replace(id, "/authors", "", -1)
	id = strings.Replace(id, "/books/", "", -1)
    i := findBooks(id)
    if i == -1 {
		fmt.Printf(id)
        return
    }
	dataJson, _ := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	for _,authorB := range books[i].Author{
		k := findAuthors(authorB)
		if k != -1 {
			dataJson, _ = json.Marshal(authors[k])
			w.Write(dataJson)
		}
	}
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdPublishersGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = strings.Replace(id, "/publishers", "", -1)
	id = strings.Replace(id, "/books/", "", -1)
    i := findBooks(id)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	 dataJson, _ := json.Marshal(books)
    if i == -1 {
        return
    }
	j := findPublishers(books[i].Publisher)
    if j == -1 {
        return
    }
    dataJson, _ = json.Marshal(publishers[j])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

//Authors

func AuthorsPut(w http.ResponseWriter, r *http.Request) {
	var author Author
    err := json.NewDecoder(r.Body).Decode(&author)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    authors = append(authors, author)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func AuthorsGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(authors)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
    i := findAuthors(id)
    if i == -1 {
        return
    }
    dataJson, _ := json.Marshal(authors[i])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdPost(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findAuthors(id)
	if i == -1 {
		return
	}	
	len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    author := Author{}
    json.Unmarshal(body, &author)
	if author.Name != ""{
		authors[i].Name=author.Name
	}
	if author.Nationality != ""{
		authors[i].Nationality=author.Nationality
	}
	if author.Birth != ""{
		authors[i].Birth=author.Birth
	}
	if author.Genere != ""{
		authors[i].Genere=author.Genere
	}
	dataJson,_ := json.Marshal(authors[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := deleteAuthor(id)
	if i == -1 {
		return
	}
	dataJson,_ := json.Marshal(authors)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func AuthorsAuthorIdBooksGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = strings.Replace(id, "/books", "", -1)
	id = strings.Replace(id, "/authors/", "", -1)
    i := findBooksbyAuthors(id)
	dataJson, _ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    for _, x := range i {
		dataJson, _ = json.Marshal(books[x])
		w.Write(dataJson)
	}
    w.WriteHeader(http.StatusOK)
}

//Publisher

func PublishersPut(w http.ResponseWriter, r *http.Request) {
	var publisher Publisher
    err := json.NewDecoder(r.Body).Decode(&publisher)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    publishers = append(publishers, publisher)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func PublishersGet(w http.ResponseWriter, r *http.Request) {
	dataJson, _ := json.Marshal(publishers)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdGet(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
    i := findPublishers(id)
    if i == -1 {
        return
    }
    dataJson, _ := json.Marshal(publishers[i])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdPost(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := findPublishers(id)
	if i == -1 {
		return
	}	
	len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    publisher := Publisher{}
    json.Unmarshal(body, &publisher)
	if publisher.Name != ""{
		publishers[i].Name=publisher.Name
	}
	if publisher.Country != ""{
		publishers[i].Country=publisher.Country
	}
	if publisher.Founded != ""{
		publishers[i].Founded=publisher.Founded
	}
	if publisher.Genere != ""{
		publishers[i].Genere=publisher.Genere
	}
	dataJson,_ := json.Marshal(publishers[i])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
	w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := deletePublisher(id)
	if i == -1 {
		return
	}
	dataJson,_ := json.Marshal(publishers)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func PublishersPublisherIdBooksGet(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path
	id = strings.Replace(id, "/books", "", -1)
	id = strings.Replace(id, "/publishers/", "", -1)
    i := findBooksbyPublishers(id)
	dataJson, _ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    for _, x := range i {
		dataJson, _ = json.Marshal(books[x])
		w.Write(dataJson)
	}
    w.WriteHeader(http.StatusOK)
}