package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Ok(w http.ResponseWriter, r *http.Request) {
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

func JsonFunc(w http.ResponseWriter, r *http.Request) {
	sizeStr := r.URL.Query().Get("size")
	size := 1
	if len(sizeStr) > 0 {
		sizeTmp, err := strconv.Atoi(sizeStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Size argument is an invalid number: "+sizeStr)
			return
		} else {
			size = sizeTmp
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
	for i := 0; i < size; i++ {
		orders = append(orders, order)
	}
	data, _ := json.Marshal(orders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(data))
}

var mb []byte
var s [][]byte

func Mem(w http.ResponseWriter, r *http.Request) {
	addStr := r.URL.Query().Get("add")
	add, err := strconv.Atoi(addStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid integer received: " + addStr))
		return
	}
	if len(mb) == 0 {
		tmp := make([]byte, 1048576)
		for i := 0; i < len(tmp); i++ {
			tmp[i] = 10
		}
		mb = tmp
	}
	for i := 0; i < add; i++ {
		dst := make([]byte, len(mb))
		copy(dst, mb)
		s = append(s, dst)
	}
}

func Env(w http.ResponseWriter, r *http.Request) {
	env := r.URL.Query().Get("env")
	value := os.Getenv(env)
	log.Println(value)
	fmt.Fprint(w, value)
}

func Tcp(w http.ResponseWriter, r *http.Request) {
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

func Exit(w http.ResponseWriter, r *http.Request) {
	log.Println("Exiting the application")
	os.Exit(3)
}

func LogFunc(w http.ResponseWriter, r *http.Request) {
	m := r.URL.Query().Get("m")
	log.Println(m)
}

func Cpu(w http.ResponseWriter, r *http.Request) {
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
