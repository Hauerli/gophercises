package urlshort

import (
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Config YAML struct
type Config struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// get path from url
		path := r.URL.Path
		// extract url from map if it exists
		url, ok := pathsToUrls[path]
		// if ok == true
		if ok {
			http.Redirect(w, r, url, http.StatusFound)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	parsedYAML := parseYAML(yml)
	urlMAP := buildMAP(parsedYAML)
	return MapHandler(urlMAP, fallback), nil

}

func parseYAML(yml []byte) []Config {

	var parsedYAML []Config
	err := yaml.Unmarshal(yml, &parsedYAML)
	if err != nil {
		log.Fatal(err)
	}
	return parsedYAML

}
func buildMAP(parsedYAML []Config) map[string]string {

	urlMAP := make(map[string]string)

	for _, line := range parsedYAML {
		urlMAP[line.Path] = line.URL
	}
	return urlMAP

}
