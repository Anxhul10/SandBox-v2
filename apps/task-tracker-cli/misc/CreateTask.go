package misc

import (
	"encoding/json"
	"fmt"
	"os"
)

type Item struct {
	ID    int64  `json:"id"`
	Title  string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
}

func CreateTask(title string ,description string, status string) {
	filePath := "tasks.json"

	var data []Item

	if _, err := os.Stat(filePath); err == nil {
		fileBytes, err := os.ReadFile(filePath)
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

	newItem := Item{
		ID:    makeTimestamp(),
		Title:  title,
		Description: description,
		Status: status,
	}
	data = append(data, newItem)

	fileBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile(filePath, fileBytes, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("✅ Object added successfully!")
}

func makeTimestamp() int64 {
	return int64(os.Getpid()) + int64(os.Getuid())
}