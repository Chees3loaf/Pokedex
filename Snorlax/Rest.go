package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	//t := now.Local()
	t := time.Now()

	fmt.Fprintln(w, "Greetings Fellow Human!")
	fmt.Fprintln(w, "     __      _")
	fmt.Fprintln(w, "   O'')}____//")
	fmt.Fprintln(w, "    ~_/      )")
	fmt.Fprintln(w, "    (_(_/-(_/ ")
	fmt.Fprintln(w, "It's", t.Format(time.Stamp))
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
