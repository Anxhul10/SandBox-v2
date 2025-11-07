package misc

import (
	_"encoding/json"
	"fmt"
	"os"
	"bufio"
	"strings"
)

func readLine(path string) {
  inFile, err := os.Open(path)
  if err != nil {
     fmt.Println(err.Error() + `: ` + path)
     return
  }
  defer inFile.Close()

  scanner := bufio.NewScanner(inFile)
  for scanner.Scan() {

    fmt.Println(scanner.Text())
  }
}

func RetriveTask(id int) {
	filePath := "tasks.json"
	
	file := filePath

    input, err := os.ReadFile(file)
    if err != nil {
        fmt.Print(err)
    }

    lines := strings.Split(string(input), "\n")

	fmt.Print(lines)
    for i, line := range lines {
        if strings.Contains(line, "\"title\":") {
            // Replace the line
            lines[i] = "\"edited\": "
        }
    }

    // Join lines and write to file
    output := strings.Join(lines, "\n")
    err = os.WriteFile(file, []byte(output), 0644)
    if err != nil {
        fmt.Print(err)
    }
}