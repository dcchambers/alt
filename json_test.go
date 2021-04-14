package main

import (
    "fmt"
    "os"
    "encoding/json"
    "io/ioutil"
)

type Programs struct {
    Programs []Program `json:"programs"`
}

type Program struct {
    Name         string   `json:"name"`
    FullName     string   `json:"fullName"`
    Description  string   `json:"description"`
    URL          string   `json:"url"`
    Alternatives []string `json:"alternatives"`
}

func main() {
    // Open jsonFile
    jsonFile, err := os.Open("programs.json")
    // handle error os.Open may return
    if err != nil {
      fmt.Println(err)
    }

    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    //initialize programs array
    var programs Programs

    // unmarshal the byteArray from jsonFile
    json.Unmarshal(byteValue, &programs)

    // we iterate through every program within our programs array and
    // print out some attributes
    for i := 0; i < len(programs.Programs); i++ {
        fmt.Println("Name: " + programs.Programs[i].Name)
        fmt.Println("Full Name: " + programs.Programs[i].FullName)
        fmt.Println("Description: " + programs.Programs[i].Description)
        fmt.Println("URL: " + programs.Programs[i].URL)
        for j :=0; j<len(programs.Programs[i].Alternatives); j++ {
            fmt.Println("Alternatives: " + programs.Programs[i].Alternatives[j])
        }
        fmt.Println("---")
    }
}
