package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	socketPath := "uds.sock"
	socket, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func(p string) {
		<-c
		if err := os.Remove(p); err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}(socketPath)

	m := http.NewServeMux()
	m.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("GET..."))
	})

	m.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("POST..."))
	})

	m.HandleFunc("PUT /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("PUT..."))
	})

	m.HandleFunc("DELETE /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("DELETE..."))
	})

	if err := http.Serve(socket, m); err != nil {
		log.Fatal(err)
	}
}

