package Oak

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





func main() {
	//Ends user interactions do server can shutdown
	TM29.Psychic()

	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine so that it doesn't block
	server := &http.Server{Addr: ":8090"}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	// Block until we receive our signal.
	<-stop

	// Create a deadline for the shutdown process.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline or all connections have closed.
	server.Shutdown(ctx)

	fmt.Println("Shutting down gracefully...")

}
