package main

import (
	"context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    TM29 "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/TMpocket"

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
    // Setup a channel to listen for interrupt signals
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

    // Define HTTP server
    server := &http.Server{Addr: ":8090"}

    // Setup request handlers
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/eatme", eatMeHandler)
    http.HandleFunc("/drinkme", drinkMeHandler)

    // Start the server in a separate goroutine
    go func() {
        fmt.Println("Starting server on port 8090")
        if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            fmt.Printf("listen: %s\n", err)
        }
    }()

    // Wait for interrupt signal
    <-stop
    fmt.Println("Shutting down server...")

    // Initiate graceful shutdown from TM29
    TM29.PrepareForShutdown()

    // Create a deadline for the shutdown process
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Shutdown the server
    if err := server.Shutdown(ctx); err != nil {
        fmt.Printf("Server shutdown error: %v\n", err)
    }

    fmt.Println("Server shut down gracefully")
}
