package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var number int64

func main() {
	// go run -race main.go
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("%d", number)))
		http.ListenAndServe(":3000", nil)
	})
}
