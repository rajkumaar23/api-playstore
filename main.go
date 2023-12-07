package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
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

	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		if strings.Trim(request.URL.Query().Get("id"), "") == "" {
			fmt.Fprint(writer, "'id' is mandatory!")
			return
		}
		res := GetPlaystoreData(request)
		json.NewEncoder(writer).Encode(res)
	})

	port := "5003"
	fmt.Printf("Starting API server on port %s!\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
