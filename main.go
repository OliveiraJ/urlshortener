package main

import (
	"fmt"
	"net/http"

	"github.com/OliveiraJ/urlshortener/handler"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := handler.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
  - path: /urlshort
    url: https://github.com/gophercises/urlshort
  - path: /urlshort-final
    url: https://github.com/gophercises/urlshort/tree/solution
`
	// yamlHandler, err := handler.YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	json := `
		[
			{
				"path":"/urlshortjson",
				"url":"https://pkg.go.dev/encoding/json"
			}
		]
	`
	// jsonHandler, err := handler.JSONHandler([]byte(json), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	allHandler, err := handler.MainHandler([]byte(yaml), []byte(json), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", allHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
