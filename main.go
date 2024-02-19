package main

import (
	"errors"
	"fmt"
	"gdrcdn/handlers"
	"net/http"
	"os"
)

func CheckPaths() {
	_, err := os.Stat("files")
	if os.IsNotExist(err) {
		os.Mkdir("files", 0755)
	}
}


func main() {
	CheckPaths()
	
	var port string = "9000"
	fmt.Println("Server started on port " + port)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	http.HandleFunc("/files/", handlers.HandlerFiles)

	err := http.ListenAndServe(":"+port, nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
