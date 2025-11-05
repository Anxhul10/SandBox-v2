package misc

import (
	_"encoding/json"
	_"fmt"
	"log"
	"os"
)

func ReadWrite() {
	f, err := os.OpenFile("tasks.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    if _, err := f.Write([]byte("data")); err != nil {
        log.Fatal(err)
    }
    if err := f.Close(); err != nil {
        log.Fatal(err)
    }
}