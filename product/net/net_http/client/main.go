// ============== //
// CLIENT SERVICE //
// ============== //

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Answer ReadAll error:", err)
		return
	}

	fmt.Println("Server's answer:", string(body))

	jsonBody := []byte(`{"client_message": "hello, server!"}`)
	bodyReader := bytes.NewReader(jsonBody)

	resp, err = http.Post("http://localhost:8080/body", "application/json", bodyReader)
	if err != nil {
		fmt.Println("Post request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Post answer ReadAll error:", err)
		return
	}

	fmt.Println("Post servers answer:", string(body))
}
