package main

import (
	"fmt"
	"net/http"
)

func uploadfile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "file uploading")
}

func setupRoute() {
	http.HandleFunc("/upload", uploadfile)
	http.ListenAndServe(":8090", nil)
}

func main() {
	fmt.Println("upload file test")
	setupRoute()
}
