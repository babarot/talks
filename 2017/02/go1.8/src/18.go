package main

func main() {
	listeners, err := listener.ListenAll()
	if err != nil || len(listeners) == 0 {
		log.Fatal(err)
	}

	// subscribe to SIGINT signals
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM)

	go18(listeners[0], sigCh)
	log.Println("[INFO] Gracefully shutdown")
}

func go18(l net.Listener, sigCh chan os.Signal) {
	server := &http.Server{Handler: newHandler()}
	go func() {
		if err := server.Serve(l); err != nil {
			log.Println(err)
		}
	}()

	<-sigCh
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

}

func newHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[INFO] %s %s\n", r.Method, r.URL.Path)
		time.Sleep(100 * time.Millisecond)
	})
	return mux
}
