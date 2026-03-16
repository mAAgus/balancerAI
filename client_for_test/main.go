package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func main() {
	jsonData := []byte(`{
        "type": "chat",
        "payload": {"message": "Привет, мир!"}
    }`)

	resp, err := http.Post("http://127.0.0.1:9080/api/v1/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
