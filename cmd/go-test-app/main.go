package main

import (
	"fmt"
	"github.com/ianschenck/envflag"
	"net/http"
)

var (
	listen = envflag.String("GO_LISTEN", ":8080", "listen")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	envflag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(*listen, nil)
}
