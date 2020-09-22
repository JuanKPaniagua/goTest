package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)

//Books

func makeCreateBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateBookRequest)
        msg, err := s.CreateBook(ctx, req.book)
        return CreateBookResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetBookByIdEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetBookByIdRequest)
        bookDetails, err := s.GetBookById(ctx, req.Id)
        if err != nil {
            return GetBookByIdResponse{Book: bookDetails, Err: "Id not found"}, nil
        }
        return GetBookByIdResponse{Book: bookDetails, Err: ""}, nil
    }
}
func makeDeleteBookEndpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteBookRequest)
        msg, err := s.DeleteBook(ctx, req.Bookid)
        if err != nil {
            return DeleteBookResponse{Msg: msg, Err: err}, nil
        }
        return DeleteBookResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateBookendpoint(s BookService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateBookRequest)
        msg, err := s.UpdateBook(ctx, req.book)
        return msg, err
    }
}

func decodeCreateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateBookRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetBookByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetBookByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetBookByIdRequest{
        Id: vars["bookid"],
    }
    return req, nil
}
func decodeDeleteBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteBookRequest
    vars := mux.Vars(r)
    req = DeleteBookRequest{
        Bookid: vars["bookid"],
    }
    return req, nil
}
func decodeUpdateBookRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateBookRequest
    if err := json.NewDecoder(r.Body).Decode(&req.book); err != nil {
        return nil, err
    }
    return req, nil
}

//Publishers

func makeCreatePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreatePublisherRequest)
        msg, err := s.CreatePublisher(ctx, req.publisher)
        return CreatePublisherResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetPublisherByIdEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetPublisherByIdRequest)
        publisherDetails, err := s.GetPublisherById(ctx, req.Id)
        if err != nil {
            return GetPublisherByIdResponse{Publisher: publisherDetails, Err: "Id not found"}, nil
        }
        return GetPublisherByIdResponse{Publisher: publisherDetails, Err: ""}, nil
    }
}
func makeDeletePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeletePublisherRequest)
        msg, err := s.DeletePublisher(ctx, req.Publisherid)
        if err != nil {
            return DeletePublisherResponse{Msg: msg, Err: err}, nil
        }
        return DeletePublisherResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdatePublisherendpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdatePublisherRequest)
        msg, err := s.UpdatePublisher(ctx, req.publisher)
        return msg, err
    }
}

func decodeCreatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreatePublisherRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetPublisherByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetPublisherByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetPublisherByIdRequest{
        Id: vars["publisherid"],
    }
    return req, nil
}
func decodeDeletePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeletePublisherRequest
    vars := mux.Vars(r)
    req = DeletePublisherRequest{
        Publisherid: vars["publisherid"],
    }
    return req, nil
}
func decodeUpdatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdatePublisherRequest
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}

//ENCODE

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    //Book
	CreateBookRequest struct {
        book Book
    }
    CreateBookResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetBookByIdRequest struct {
        Id string `json:"bookid"`
    }
    GetBookByIdResponse struct {
        Book interface{} `json:"book,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteBookRequest struct {
        Bookid string `json:"bookid"`
    }

    DeleteBookResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateBookRequest struct {
        book Book
    }
    UpdateBookResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
	
	//Publishers

	CreatePublisherRequest struct {
        publisher Publisher
    }
    CreatePublisherResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetPublisherByIdRequest struct {
        Id string `json:"publisherid"`
    }
    GetPublisherByIdResponse struct {
        Publisher interface{} `json:"publisher,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeletePublisherRequest struct {
        Publisherid string `json:"publisherid"`
    }

    DeletePublisherResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdatePublisherRequest struct {
        publisher Publisher
    }
    UpdatePublisherResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)