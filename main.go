package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const defaultPort = 8080

func main() {
	fmt.Println("Starting server...")

	// Flags
	port := flag.Int("port", defaultPort, "The port on which to register the HTTP listener to.")
	flag.Parse()

	// Handlers
	http.HandleFunc("/", ok)
	http.HandleFunc("/json", jsonFunc)
	http.HandleFunc("/envs", env)
	http.HandleFunc("/cpu", cpu)

	// Server
	addr := fmt.Sprintf(":%d", *port)
	http.ListenAndServe(addr, nil)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

type Order struct {
	Id           string    `json:"id"`
	Description  string    `json:"description"`
	Price        float32   `json:"price"`
	Status       bool      `json:"status"`
	Date         time.Time `json:"date"`
	DeliveryDate time.Time `json:"deliveryDate"`
}

func jsonFunc(w http.ResponseWriter, r *http.Request) {
	order := &Order{
		Id:           "001",
		Description:  "A very special order",
		Price:        500,
		Status:       true,
		Date:         time.Now(),
		DeliveryDate: time.Now(),
	}
	data, _ := json.Marshal(order)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
}

func env(w http.ResponseWriter, r *http.Request) {
	env := r.URL.Query().Get("env")
	value := os.Getenv(env)
	log.Println(value)
	fmt.Fprint(w, value)
}

func cpu(w http.ResponseWriter, r *http.Request) {
	// Size
	size := r.URL.Query().Get("x")
	i, err := strconv.Atoi(size)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You need to inform the size. Value received: " + size))
		return
	}
	// Print
	printStr := r.URL.Query().Get("print")
	if len(printStr) == 0 {
		printStr = "FALSE"
	}

	b, errPrint := strconv.ParseBool(printStr)
	if errPrint != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("You need to inform the print boolean. Value received: " + printStr))
		return
	}

	// Calculate
	ui := uint(i)
	f := fib(ui)
	if b {
		fmt.Println(strconv.FormatUint(uint64(f), 10))
	}
	fmt.Fprint(w, f)
}

func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
