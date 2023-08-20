package main

import (
	"flag"
	"fmt"
	"net/http"

	"main/handlers"
)

func main() {
	fmt.Println("Starting server...")

	// Flags
	port := flag.Int("port", 8080, "The port on which to register the HTTP listener to.")
	flag.Parse()

	// Handlers
	http.HandleFunc("/", handlers.Ok)
	http.HandleFunc("/json", handlers.JsonFunc)
	http.HandleFunc("/envs", handlers.Env)
	http.HandleFunc("/tcp", handlers.Tcp)
	http.HandleFunc("/cpu", handlers.Cpu)
	http.HandleFunc("/mem", handlers.Mem)
	http.HandleFunc("/exit", handlers.Exit)
	http.HandleFunc("/log", handlers.LogFunc)

	// Server
	addr := fmt.Sprintf(":%d", *port)
	http.ListenAndServe(addr, nil)
}
