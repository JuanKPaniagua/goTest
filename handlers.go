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

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
	//err=err;
	checkError("Parse error", err)
	if id == "" {
		dataJson,err = json.Marshal(books)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
	}else{
		i := find(id)
		if i == -1 {
			return
		}
		dataJson2,err = json.Marshal(books[i])
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson2)
	}
	
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
    w.WriteHeader(200)
    return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
    w.WriteHeader(200)
    return
}