package handler

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	pathUrls, err := parseYaml(yamlBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pathsToUrls := buildMap(pathUrls)
// 	return MapHandler(pathsToUrls, fallback), nil
// }

// func JSONHandler(jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
// 	pathUrls, err := parseJson(jsonBytes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	pathsToUrls := buildMap(pathUrls)
// 	return MapHandler(pathsToUrls, fallback), nil
// }

// Aggregates both YAMLHandle and JSONHandle, making it easier to use both formats for the exercise, both yamlBytes and jsonBytes are
// slices of bytes that have data in both formats and needs to be parsed to an slice of the pathUrl type, then this slices are joined
// by the build map function, maintaining the original of the exercise resolution.
func MainHandler(yamlBytes []byte, jsonBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	yamlPathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}
	jsonPathUrls, err := parseJson(jsonBytes)
	if err != nil {
		return nil, err
	}
	pathsToUrls := buildMap(yamlPathUrls, jsonPathUrls)

	return MapHandler(pathsToUrls, fallback), nil
}

func parseJson(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(data, &pathUrls)
	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

// Joins both  yamlPathUrls and jsonPathUrls in a map[string]string and returns this map, originally this function used to received only
// an slice of pathUrl type, but to fit the bonus exercise and be able to read from yaml and json formats, a new slice with the jsonPaths
// is now received to and so, both are joined in a nap to be returned.
func buildMap(yamlPathUrls, jsonPathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, ypu := range yamlPathUrls {
		pathsToUrls[ypu.Path] = ypu.URL
	}
	for _, jpu := range jsonPathUrls {
		pathsToUrls[jpu.Path] = jpu.URL
	}

	return pathsToUrls
}

type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}
