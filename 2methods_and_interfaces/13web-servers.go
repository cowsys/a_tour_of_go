package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

// NOTE: ServeHTTPを実装するとhandlerインタフェースを満たす
// type Handler interface {
//     ServeHTTP(ResponseWriter, *Request)
// }
func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// NOTE:満たすべき実装は以下
	// ServeHTTP should write reply headers and data to the ResponseWriter and
	// then return. Returning signals that the request is finished and that the
	// HTTP server can move on to the next request on the connection.
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello

	err := http.ListenAndServe("localhost:4444", h)
	if err != nil {
		log.Fatal(err)
	}
}
