package misc

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"bufio"
	"io"
	"io/ioutil"
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


// func ReadWrite() {
// 	human1 := Human{Name:"Ankit", Age: 23, Address:"New Delhi"} 
       
//     // encoding human1 struct 
//     // into json format 
//     human_enc, err := json.Marshal(human1) 
       
//     if err != nil { 
           
//         // if error is not nil 
//         // print error 
//         fmt.Println(err) 
//     } 
       
//     // as human_enc is in a byte array 
//     // format, it needs to be  
//     // converted into a string  
//     fmt.Println(string(human_enc))

// 	f, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     if err != nil {
// 		_, writeError := f.Write([]byte("[]"))
//         fmt.Print(writeError)
// 		fmt.Print(err)
//     }

// 	readFileLines("tasks.json")
//     // if _, err := f.Write([]byte("data")); err != nil {
//     //     log.Fatal(err)
//     // }
//     if err := f.Close(); err != nil {
//         log.Fatal(err)
//     }
// }

type Item struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func ReadWrite() {
	filePath := "data.json"

	// Step 1: Read file (create if not exists)
	var data []Item

	if _, err := os.Stat(filePath); err == nil {
		fileBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		if len(fileBytes) > 0 {
			err = json.Unmarshal(fileBytes, &data)
			if err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}
		}
	}

	// Step 2: Push new object
	newItem := Item{
		ID:    makeTimestamp(),
		Name:  "Anshul",
		Value: 42,
	}
	data = append(data, newItem)

	// Step 3: Write updated JSON back to file
	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = ioutil.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("✅ Object added successfully!")
}

func makeTimestamp() int64 {
	return int64(os.Getpid()) + int64(os.Getuid())
}
