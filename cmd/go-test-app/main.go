package main

import (
	"fmt"
	"github.com/ianschenck/envflag"
	"net/http"
	"os"
)

var (
	listen = envflag.String("GO_LISTEN", ":8080", "listen")
	usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		envflag.PrintDefaults()
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {
	envflag.Parse()
	http.HandleFunc("/", handler)
	http.ListenAndServe(*listen, nil)
}
