package main

import (
    "net/http"
    "os"
)

func handlerTeam(writer http.ResponseWriter, request *http.Request) {
    var err error
    readTeamData("teams.csv")
    switch request.Method {
    case "GET":
        err = handleTeamGet(writer, request)
    case "POST":
        err = handleTeamPost(writer, request)
    case "PUT":
        err = handleTeamPut(writer, request)
    case "DELETE":
        err = handleTeamDelete(writer, request)
    }
    writeTeamData("teams.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    http.HandleFunc("/teams/", handlerTeam)
    if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
        panic(err)
    }
}