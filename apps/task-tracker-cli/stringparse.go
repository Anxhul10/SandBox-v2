package main

import (
    "flag"
    "fmt"
    // "os"
)

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

    metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines};.")
    uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric.")
    flag.Parse()

    // if *items_to_add == "" {
    //     flag.PrintDefaults()
    //     os.Exit(1)
    // }

    fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *items_to_add, *metricPtr, *uniquePtr)
}