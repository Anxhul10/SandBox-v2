package main

import (
    "flag"
    "fmt"
	"io/ioutil"
    "encoding/json"
)

type Window struct {
    Width int `json:"width"`
    Height int `json:"height"`
    X int `json:"x"`
    Y int `json:"y"`
}
type Config struct {
    Timeout float32 `json:"timeout"`
    PluginsPath string `json:"pluginsPath"`
    Window Window `json:"window"`
}

func main() {
    items_to_add := flag.String("add", "", "Text to parse. (Required)") // log the id of the item
	update_id := flag.Int("update", -1, "enter id to update")
	delete_id := flag.Int("delete", -1, "enter id to delete")
	progress_id := flag.Int("progress", -1, "enter id to change status to progress")
	complete_id := flag.Int("complete", -1, "enter id to change status to complete")
	list_add := flag.Bool("list", false, "list all the task along with status")
	list_done := flag.Bool("list done", false, "list done ")
	list_todo := flag.Bool("list todo", false, "list todo")
	list_in_progress := flag.Bool("list progress", false, "list progress")

	_ = update_id
	_ = delete_id
	_ = progress_id
	_ = complete_id
	_ = list_add
	_ = list_done
	_ = list_todo
	_ = list_in_progress

    
    flag.Parse()

	if *items_to_add != "" {
		fmt.Printf(*items_to_add)
	}

	fileCount := map[string]int{
        "cpp": 10,
        "js": 8,
        "go": 10,
    }
    bytes, _ := json.Marshal(fileCount)
    fmt.Println(string(bytes))


	// write to file
	//  config := Config {
    //     Timeout: 40.420,
    //     PluginsPath: "~/plugins/etc",
    //     Window: Window {500, 200, 20, 20},
    // }
    // bytes, _ := json.MarshalIndent(config, "", "  ")
    ioutil.WriteFile("config.json", bytes, 0644)

    // fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *items_to_add, *metricPtr, *uniquePtr)
}