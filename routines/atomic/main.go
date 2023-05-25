package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64

func main() {
	// go run -race main.go
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		atomic.AddUint64(&number, 1)
		number++
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("%d", number)))
		http.ListenAndServe(":3000", nil)
	})
}
