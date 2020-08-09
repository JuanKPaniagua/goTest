package main

import (
    "encoding/json"
    "net/http"
    "path"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func delete(x string) int {
    for i, book := range books {
        if x == book.Id {
			copy(books[i:], books[i+1:]) // Shift a[i+1:] left one index.
			books = books[:len(books)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	dataJson,err := json.Marshal(books)
	if i == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)	
		return
	}
	dataJson,err = json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)	
    return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    book := Book{}
    json.Unmarshal(body, &book)
    books = append(books, book)
    w.WriteHeader(200)
    return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	i := find(id)
	if i == -1 {
		return
	}	
	r.ParseForm()
	book :=books[i]
	for key,value := range r.Form{
		switch key {
			case "title":
				book.Title=value[0]
		}
	}
	books[i] =book
	books = append(books, book)
	dataJson,err := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := delete(id)
	if i == -1 {
		return
	}
	dataJson,err := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
    w.WriteHeader(200)
    return
}