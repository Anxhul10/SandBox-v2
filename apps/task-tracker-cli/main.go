package main

import (
    "flag"
    _"fmt"
	_"os"
    _"encoding/json"
	"github.com/Anxhul10/Sandbox-v2/apps/task-tracker-cli/misc"
)
// boilerplate for the task
// {
// 	"id":{
// 		"title":"my title",
// 		"description": "my description",
// 		"status": "todo/in-progress/done",
// 		"createdAt":"time of creation",
// 		"updatedAt": "time of last update"
// 	}
// }
// type Window struct {
//     Width int `json:"width"`
//     Height int `json:"height"`
//     X int `json:"x"`
//     Y int `json:"y"`
// }

// type Config struct {
//     Timeout float32 `json:"timeout"`
//     PluginsPath string `json:"pluginsPath"`
//     Window Window `json:"window"`
// }

// func createJSON(key int, value string) {
// 	myData := map[int] string {
// 		key:value,
// 	}
// 	marshalData, _ := json.Marshal(myData)
//     os.WriteFile("data.json", marshalData, 0644)
// }

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
	get_json := flag.Bool("json", false, "get json")

	_ = items_to_add
	_ = update_id
	_ = delete_id
	_ = progress_id
	_ = complete_id
	_ = list_add
	_ = list_done
	_ = list_todo
	_ = list_in_progress
	_ = get_json

    
    flag.Parse()

	// misc.CreateTask("title1", "description1", "pending")
	misc.RetriveTask(12)
	// if *items_to_add != "" {
	// 	// fmt.Printf(*items_to_add)
	// 	va := map[string]map[string]string {"first": {"name":"anshul"}}
	// 	fmt.Println(va)
	// }

	// createJSON(2, "nig")

	// bytes, err := os.ReadFile("data.json")

	// _ = bytes

	// if err != nil {
	// 	// create json
	// 	// return 
	// 	createJSON(0,"temp")
	// 	return
	// }

	// config := Config {
    //     Timeout: 40.420,
    //     PluginsPath: "~/plugins/etc",
    //     Window: Window {500, 200, 20, 20},
    // }
    // tyes, _ := json.MarshalIndent(config, "", "  ")
    // os.WriteFile("config.json", tyes, 0644)

	// fmt.Print(string(tyes))
    // fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *items_to_add, *metricPtr, *uniquePtr)
}