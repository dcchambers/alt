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
    // Open our jsonFile
    jsonFile, err := os.Open("programs.json")
    // if we os.Open returns an error then handle it
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened jsonFile as a byte array.
    byteValue, _ := ioutil.ReadAll(jsonFile)

    //initialize programs array
    var programs Programs

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
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
