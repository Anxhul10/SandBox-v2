package misc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"bufio"
	"io"
)

type Human struct {
    Name string
    Age int
    Address string
}

func readFileLines(logfile string) {
    f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
    if err != nil {
        log.Fatalf("open file error: %v", err)
        return
    }
    defer f.Close()

    rd := bufio.NewReader(f)
    for {
        line, err := rd.ReadString('\n')
        if err != nil {
            if err == io.EOF {
                break
            }

            log.Fatalf("read file line error: %v", err)
            return
        }
        _ = line  // GET the line string
		fmt.Print(line)
    }
}


func ReadWrite() {
	human1 := Human{Name:"Ankit", Age: 23, Address:"New Delhi"} 
       
    // encoding human1 struct 
    // into json format 
    human_enc, err := json.Marshal(human1) 
       
    if err != nil { 
           
        // if error is not nil 
        // print error 
        fmt.Println(err) 
    } 
       
    // as human_enc is in a byte array 
    // format, it needs to be  
    // converted into a string  
    fmt.Println(string(human_enc))

	f, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
		_, writeError := f.Write([]byte("[]"))
        fmt.Print(writeError)
		fmt.Print(err)
    }

	readFileLines("tasks.json")
    // if _, err := f.Write([]byte("data")); err != nil {
    //     log.Fatal(err)
    // }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}