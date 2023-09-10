package main

import (
	"fmt"
	"net/http"
	"time"
	"os"
	
	
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

func eatMeHandler(w http.ResponseWriter, req *http.Request) {
	files, err := os.ReadDir("./EatMe")
	if err != nil {
		http.Error(w, "Unable to read EatMe directory", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "The Finest of Tavern Foods")
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}

func drinkMeHandler(w http.ResponseWriter, req *http.Request) {
	files, err := os.ReadDir("./DrinkMe")
	if err != nil {
		http.Error(w, "Unable to read DrinkMe directory", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "A Very Marry Un-Birthday to You!")
	fmt.Fprintln(w, "     The Tea Party Awaits")
	for _, file := range files {
		fmt.Fprintln(w, file.Name())
	}
}


func main() {
	

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/eatme", eatMeHandler)
	http.HandleFunc("/drinkme", drinkMeHandler)

	http.ListenAndServe(":8090", nil)


}
