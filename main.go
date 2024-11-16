package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve static files (HTML, CSS, JS, images) from the "static" folder
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the server
	port := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
