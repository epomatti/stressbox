package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
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
	http.HandleFunc("/tcp", tcp)
	http.HandleFunc("/cpu", cpu)
	http.HandleFunc("/exit", exit)

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
	countStr := r.URL.Query().Get("count")
	count := 1
	if len(countStr) > 0 {
		countTmp, err := strconv.Atoi(countStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Count argument is an invalid number: "+countStr)
			return
		} else {
			count = countTmp
		}
	}
	order := &Order{
		Id:           "001",
		Description:  "A very special order",
		Price:        500,
		Status:       true,
		Date:         time.Now(),
		DeliveryDate: time.Now(),
	}
	orders := []*Order{}
	for i := 0; i < count; i++ {
		orders = append(orders, order)
	}
	data, _ := json.Marshal(orders)
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

func tcp(w http.ResponseWriter, r *http.Request) {
	addr := r.URL.Query().Get("addr")
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		errMsg := "Resolution failed:" + err.Error()
		log.Println(errMsg)
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, errMsg)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		errMsg := "Dial failed:" + err.Error()
		log.Println(errMsg)
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, errMsg)
		return
	}
	defer conn.Close()
	fmt.Fprint(w, "TCP connection: OK\n")
}

func exit(w http.ResponseWriter, r *http.Request) {
	log.Println("Exiting the application")
	os.Exit(3)
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
