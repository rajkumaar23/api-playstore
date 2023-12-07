package main

import (
	"fmt"
	"net/http"
	"os"
)

var HttpClient = &http.Client{}

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		readme, err := os.ReadFile("README.md")
		if err != nil {
			panic(err)
		}

		fmt.Fprint(writer, string(mdToHTML(readme)))
	})

	port := "5003"
	fmt.Printf("Starting API server on port %s!", port)
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		panic(err)
	}
}
