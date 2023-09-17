package main

import (
	"fmt"
	"net/http"
    "../"
)


func main() {
    mux := defaultMux()
    pathsToUrls := map[string]string{
        "/github-nipun": "https://github.com/KhNikh",
        "/linkedin-nipun": "https://linkedin.com/in/nipun-khandelwal", 
    }
    mapHandler := urlshort.MapHandler(pathsToUrls, mux)
}
func defaultMux() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/", hello)
    return mux
}
func hello(w http.ResponseWriter, r *http.Request) {
   fmt.Println(w, "Hello, world") 

}
