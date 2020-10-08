package main

import (
    "encoding/csv"
    "log"
    "os"
)

type Team struct {
    Id        string `json:"id"`
    Name      string `json:"name"`
}

var teams []Team

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func readTeamData(filePath string) {
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()
    teams = []Team{}
    for _, record := range records {
        team := Team{
            Id:  record[0],
            Name:    record[1]}
        teams = append(teams, team)
    }
    file.Close()
}

func writeTeamData(filePath string) {
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()

    file.Seek(0, 0)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, team := range teams {
        record := []string{team.Id,team.Name}
        err := writer.Write(record)
        checkError("Cannot write to file", err)
    }
    writer.Flush()
    file.Close()
}