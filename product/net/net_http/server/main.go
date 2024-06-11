// ============== //
// SERVER SERVICE //
// ============== //

package main

import (
	"fmt"
	"io"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloHandler called")
	io.WriteString(w, "Hello, World!\n")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("rootHandler called")
	io.WriteString(w, "You are on main page!\n")
}

func main() {
	simplest()

	// withMux()

	// coolerServer()

}

func simplest() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func withMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/hello", helloHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", mux)
}
