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
		writer.Header().Add("Content-Type", "application/json")
		if strings.Trim(request.URL.Query().Get("id"), "") == "" {
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(GenerateErrorResponse("package ID is mandatory!"))
			return
		}

		res, errCode := GetPlaystoreData(request)
		if errCode != -1 {
			writer.WriteHeader(errCode)
			if errCode == 404 {
				json.NewEncoder(writer).Encode(GenerateErrorResponse("package ID seems to be invalid"))
			} else {
				json.NewEncoder(writer).Encode(GenerateErrorResponse("an unexpected error occurred"))
			}
			return
		}
		json.NewEncoder(writer).Encode(res)
	})

	port := "5003"
	fmt.Printf("Starting API server on port %s!\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
