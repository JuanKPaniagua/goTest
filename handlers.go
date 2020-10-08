package main

import (
    "encoding/json"
    "net/http"
    "path"
)

func findTeam(x string) int {
    for i, team := range teams {
        if x == team.Id {
            return i
        }
    }
    return -1
}

func deleteTeam(x string) int {
    for i, team := range teams {
        if x == team.Id {
			copy(teams[i:], teams[i+1:]) // Shift a[i+1:] left one index.
			teams = teams[:len(teams)-1]     // Truncate slice.
            return i
        }
    }
    return -1
}

func handleTeamGet(w http.ResponseWriter, r *http.Request) (err error) {
    id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := findTeam(id)
	dataJson,err := json.Marshal(teams)
	if i == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)	
		return
	}
	dataJson,err = json.Marshal(teams[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)	
    return
}

func handleTeamPost(w http.ResponseWriter, r *http.Request) (err error) {
    len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    team := Team{}
    json.Unmarshal(body, &team)
    teams = append(teams, team)
    w.WriteHeader(200)
    return
}

func handleTeamPut(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	i := findTeam(id)
	if i == -1 {
		return
	}	
	//r.ParseForm()
	len := r.ContentLength
    body := make([]byte, len)
    r.Body.Read(body)
    team := Team{}
    json.Unmarshal(body, &team)
	if team.Name != ""{
		teams[i].Name=team.Name
	}
	dataJson,err := json.Marshal(teams)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
	return
}
func handleTeamDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := deleteTeam(id)
	if i == -1 {
		//w.Header().Set("Content-Type", "application/json")
		//w.Write("Id typed cannot be founded")	
		return
	}
	dataJson,err := json.Marshal(teams)
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson)
    //w.WriteHeader(200)
    return
}