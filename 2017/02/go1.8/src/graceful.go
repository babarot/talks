package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lestrrat/go-server-starter/listener"
)

func main() {
	listeners, err := listener.ListenAll()
	if err != nil || len(listeners) == 0 {
		log.Fatal(err)
	}

	server := &http.Server{Handler: newHandler()}
	go func() {
		if err := server.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM)

	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("[INFO] Shutting down server")
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s", r.Method, r.URL.Path)
		time.Sleep(100 * time.Millisecond)
	})
	return mux
}
