	var l net.Listener
	if os.Getenv("SERVER_STARTER_PORT") != "" {
		listeners, err := listener.ListenAll()
		if len(listeners) > 0 {
			l = listeners[0]
		}
	}
	if l == nil {
		var err error
		l, err = net.Listen("tcp", ":8080")
	}
	fmt.Println(http.Serve(l, newHandler()))
