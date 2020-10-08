package swagger

import (
    "encoding/json"
    "net/http"
    "path"
)

var books = []Book{
    Book{BookId: "Book1", Title: "Operating System Concepts", Edition: "9th",
        Copyright: "2012", Language: "ENGLISH", Pages: "976",
        Author: "Abraham Silberschatz", Publisher: "John Wiley & Sons"},
    Book{BookId: "Book3", Title: "Computer Networks", Edition: "5th",
        Copyright: "2010", Language: "ENGLISH", Pages: "960",
        Author: "Andrew S. Tanenbaum", Publisher: "Andrew S. Tanenbaum"},
}

func find(x string) int {
    for i, book := range books {
        if x == book.BookId {
            return i
        }
    }
    return -1
}

func delete(x string) int {
    for i, book := range books {
        if x == book.BookId {
			copy(books[i:], books[i+1:]) // Shift a[i+1:] left one index.
			books = books[:len(books)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

func BooksBookIdDelete(w http.ResponseWriter, r *http.Request) {
	id := path.Base(r.URL.Path)
	i := delete(id)
	if i == -1 {
		//w.Header().Set("Content-Type", "application/json")
		//w.Write("Id typed cannot be founded")	
		return
	}
	dataJson,_ := json.Marshal(books)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdGet(w http.ResponseWriter, r *http.Request) {
    id := path.Base(r.URL.Path)
    i := find(id)
	dataJson, _ := json.Marshal(books)
    if i == -1 {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(dataJson)
		w.WriteHeader(http.StatusOK)
		return
    }
    dataJson, _ = json.Marshal(books[i])
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Write(dataJson)
    w.WriteHeader(http.StatusOK)
}

func BooksBookIdPut(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}

func BooksPost(w http.ResponseWriter, r *http.Request) {
    var book Book
    err := json.NewDecoder(r.Body).Decode(&book)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    books = append(books, book)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
}