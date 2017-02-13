package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/lestrrat/go-server-starter/listener"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello DEMO")
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func main() {
	var l net.Listener

	if os.Getenv("SERVER_STARTER_PORT") != "" {
		listeners, err := listener.ListenAll()
		if err != nil {
			log.Println(err)
			return
		}

		if len(listeners) > 0 {
			l = listeners[0]
		}
	}

	if l == nil {
		var err error
		l, err = net.Listen("tcp", ":8080")

		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Println("Start to serve")
	fmt.Println(http.Serve(l, newHandler()))
}
