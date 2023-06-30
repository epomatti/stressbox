package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Starting server...")

	port := flag.Int("port", 8080, "The port on which to register the HTTP listener to.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	http.HandleFunc("/envs", func(w http.ResponseWriter, r *http.Request) {
		env := r.URL.Query().Get("env")
		value := os.Getenv(env)
		log.Println(value)
		fmt.Fprint(w, value)
	})

	http.HandleFunc("/cpu", func(w http.ResponseWriter, r *http.Request) {
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
	})

	addr := fmt.Sprintf(":%d", *port)
	http.ListenAndServe(addr, nil)
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
