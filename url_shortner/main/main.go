package main

import (
	"fmt"
	"net/http"
    "github.com/KhNikh/Go-Projects/url_shortner"
)
func main() {
    mux := defaultMux()
    pathsToUrls := map[string]string{
        "/github-nipun": "https://github.com/KhNikh",
        "/linkedin-nipun": "https://linkedin.com/in/nipun-khandelwal", 
    }
    mapHandler := urlshort.MapHandler(pathsToUrls, mux)
    yaml := `
    - path: /github-nipun
      url: https://github.com/KhNikh

    - path: /github-linkedin
      url: https://linkedin.com/in/nipun-khandelwal
    `
    yamlHandler, err := urlshort.YamlHandler([]byte(yaml), mapHandler)
    if err != nil {
        panic(err)
    }
    http.ListenAndServe(":8080", mapHandler)
    http.ListenAndServe(":8081", yamlHandler)
}
func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "Hello, world") 

}
