package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lestrrat/go-server-starter/listener"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
	w.(http.Flusher).Flush()
	time.Sleep(time.Millisecond * 100)
}

func main() {
	listeners, err := listener.ListenAll()
	if err != nil && err != listener.ErrNoListeningTarget {
		log.Print(err)
	}

	server := &http.Server{Handler: http.HandlerFunc(handler)}
	// サーバとは別の goroutine で待ち受ける
	go func() {
		if err := server.Serve(listeners[0]); err != nil {
			log.Print(err)
		}
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	<-sigCh
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
}
