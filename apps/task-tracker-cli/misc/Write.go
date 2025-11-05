package misc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)


type Rankings struct {
    Keyword  string 
    GetCount uint32 
    Engine   string 
    Locale   string 
    Mobile   bool   
}

func ReadWrite() {
	var jsonBlob = []byte(`
        {"keyword":"hipaa compliance form", "get_count":157, "engine":"google", "locale":"en-us", "mobile":false}
    `)
    rankings := Rankings{}
    err := json.Unmarshal(jsonBlob, &rankings)
    if err != nil {
        // nozzle.printError("opening config file", err.Error())
    }

    rankingsJson, _ := json.Marshal(rankings)
    err = os.WriteFile("output.json", rankingsJson, 0644)
    fmt.Printf("%+v", rankings)


	f, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte("some data")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}