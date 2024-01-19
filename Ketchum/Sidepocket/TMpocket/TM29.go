package TM29

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	Potion "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket"
	Cyberball "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket"
)

var (
	openConnections []*net.Conn
	openFiles		[]*os.File
	mutex 			sync.Mutex
)

func RegisterConnection(conn *net.Conn) {
    mutex.Lock()
    defer mutex.Unlock()
    openConnections = append(openConnections, conn)
}

func RegisterFile(file *os.File) {
    mutex.Lock()
    defer mutex.Unlock()
    openFiles = append(openFiles, file)
}

func CloseResources() {
    mutex.Lock()
    defer mutex.Unlock()

    for _, conn := range openConnections {
        if conn != nil {
            if err := (*conn).Close(); err != nil {
                fmt.Printf("Error closing connection: %v\n", err)
            }
        }
    }
    openConnections = []*net.Conn{}

    for _, file := range openFiles {
        if file != nil {
            if err := file.Close(); err != nil {
                fmt.Printf("Error closing file: %v\n", err)
            }
        }
    }
    openFiles = []*os.File{}

    fmt.Println("All resources have been closed.")
}

func PrepareForShutdown() {
    fmt.Println("Preparing for shutdown...")
    CloseResources()
    // Additional shutdown preparation tasks can be added here.
}

func Wait(seconds int) {
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		close(done)
	}()

	<-done
	fmt.Printf("Waited for %d seconds./n,", seconds)

}

// Clears screen in cmd
func Cls() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func StringPrompt(label string) (string, error) {
	var s string
	r := bufio.NewReader(os.Stdin)

	fmt.Fprint(os.Stderr, label+" ")
	s, err := r.ReadString('\n')
	if err != nil {
		return "", err // Return an empty string and the error if one occurs
	}
	
	return strings.TrimSpace(s), nil // Return the trimmed string and no error
}


func ListAndSelectFiles(directory string) (string, error) {
    files, err := Potion.Map(directory)
    if err != nil {
        fmt.Println("Error listing files:", err)
        return "", err
    }

    fmt.Println("Files available:")
    for i, file := range files {
        fmt.Printf("%d. %s\n", i+1, file)
    }

    choice, err := StringPrompt("Select a file by number to read its content:")
    if err != nil {
        fmt.Println("Error reading input:", err)
        return "", err
    }

    index, err := strconv.Atoi(choice)
    if err != nil || index < 1 || index > len(files) {
        fmt.Println("Invalid selection.")
        return "", fmt.Errorf("invalid selection")
    }

    selectedFileName := files[index-1] // Assuming this is the name of the file

    content, err := Potion.Load(directory + "/" + selectedFileName)
    if err != nil {
        fmt.Println("Error reading file:", err)
        return "", err
    } else {
        fmt.Println(content)
    }

    printOption, err := StringPrompt("Save for the coming network apocalypse? (Yes/No)")
    if printOption == "Yes" {
        fmt.Println("Good choice")
        err = Cyberball.GeneratePDF(selectedFileName, content)
        if err != nil {
            fmt.Println("Error generating PDF:", err)
            return "", err
        } else {
            fmt.Println("Keep it secret, keep it safe.", selectedFileName+".pdf")
        }
    } else if printOption == "No" {
        fmt.Println("Suit yourself")
    }

    return selectedFileName, nil
}


// FetchFromServer makes an HTTP GET request to a specified endpoint on a predefined server.
// It returns the response body as a string and any error encountered during the process.
func FetchFromServer(endpoint string) (string, error) {
    // Define the server address.
    serverAddress := "10.0.0.55:55555"

    // Perform an HTTP GET request to the server using the provided endpoint.
    resp, err := http.Get(serverAddress + endpoint)
    if err != nil {
        // Return a formatted error if the request fails.
        return "", fmt.Errorf("error making request to server: %v", err)
    }
    // Ensure the response body is closed after the function exits.
    defer resp.Body.Close()

    // Check if the server's response status is not 'OK'.
    if resp.StatusCode != http.StatusOK {
        // Return a formatted error if the status is not 'OK'.
        return "", fmt.Errorf("server returned non-OK status: %s", resp.Status)
    }

    // Use a buffered reader to read the response body.
    reader := bufio.NewReader(resp.Body)
    var response strings.Builder
    // Read the response body line by line until EOF is reached.
    for {
        line, err := reader.ReadString('\n')
        if err != nil && err != io.EOF {
            // Return a formatted error if there's an issue reading the response body.
            return "", fmt.Errorf("error reading response body: %v", err)
        }
        if err == io.EOF {
            // Break the loop if EOF is reached.
            break
        }
        // Append each line to the response string.
        response.WriteString(line)
    }

    // Return the complete response body as a string.
    return response.String(), nil
}
