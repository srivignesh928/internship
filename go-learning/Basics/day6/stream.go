package main

import (
	"fmt"
	"net/http"
	"time"
)

func streamHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	for i := 1; i <= 5; i++ {
		fmt.Fprintf(w, "Message %d\n", i)

		flusher.Flush() // force send

		time.Sleep(1 * time.Second)
	}
}
func main() {
	http.HandleFunc("/stream", streamHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}