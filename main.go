package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svcB BookService
    svcB = NewServiceB(logger)
	
	var svcP PublisherService
    svcP = NewServiceP(logger)
	
	var svcA AuthorService
    svcA = NewServiceA(logger)
	

	//BOOKS
    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(svcB),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(svcB),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
	GetAllBooks := httptransport.NewServer(
        makeGetAllBooksEndpoint(svcB),
        decodeGetAllBooksRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(svcB),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(svcB),
        decodeUpdateBookRequest,
        encodeResponse,
    )
    http.Handle("/", r)
    http.Handle("/book", CreateBookHandler)
	http.Handle("/book/", GetAllBooks)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
    r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")

	//PUBLISHER
    CreatePublisherHandler := httptransport.NewServer(
        makeCreatePublisherEndpoint(svcP),
        decodeCreatePublisherRequest,
        encodeResponse,
    )
    GetByPublisherIdHandler := httptransport.NewServer(
        makeGetPublisherByIdEndpoint(svcP),
        decodeGetPublisherByIdRequest,
        encodeResponse,
    )
	GetAllPublishersHandler := httptransport.NewServer(
        makeGetAllPublishersEndpoint(svcP),
        decodeGetAllPublishersRequest,
        encodeResponse,
    )
    DeletePublisherHandler := httptransport.NewServer(
        makeDeletePublisherEndpoint(svcP),
        decodeDeletePublisherRequest,
        encodeResponse,
    )
    UpdatePublisherHandler := httptransport.NewServer(
        makeUpdatePublisherendpoint(svcP),
        decodeUpdatePublisherRequest,
        encodeResponse,
    )
    http.Handle("/publisher", CreatePublisherHandler)
	http.Handle("/publisher/", GetAllPublishersHandler)
    http.Handle("/publisher/update", UpdatePublisherHandler)
    r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
    r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")
	
	//AUTHOR
    CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(svcA),
        decodeCreateAuthorRequest,
        encodeResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(svcA),
        decodeGetAuthorByIdRequest,
        encodeResponse,
    )
	GetAllAuthorsHandler := httptransport.NewServer(
        makeGetAllAuthorsEndpoint(svcA),
        decodeGetAllAuthorsRequest,
        encodeResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(svcA),
        decodeDeleteAuthorRequest,
        encodeResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(svcA),
        decodeUpdateAuthorRequest,
        encodeResponse,
    )
    http.Handle("/author", CreateAuthorHandler)
	http.Handle("/author", GetAllAuthorsHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
    r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")
	
	logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}