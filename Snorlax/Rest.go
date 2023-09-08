package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"bufio"
	"os"
	"os/exec"
	"Pokedex/Snorlax/moves/sleeptalk.go"
	
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

func wait(seconds int) {
	done := make(chan struct{})

	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		close(done)
	}()

	<-done
	fmt.Printf("Waited for %d seconds./n,", seconds)

}
//Clears screen in cmd
func Cls() {

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != " " {
			break
		}
	}
	return strings.TrimSpace(s)

}

func main() {
	mode := 0 // 0: unset, 1: Eat Me, 2: Drink Me, 3: Wake Up
	var err error
	var resp []byte
	var interfaces, interfaceList []string
	var adapterNum uint32 = 0

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)

	wait(3)


MainMenu:
	mode = 0
	fmt.Println("Choose One")
	fmt.Println("1. Eat Me")
	fmt.Println("2. Drink Me")
	fmt.Println("3. Wake Up")
	res := StringPrompt("One or Two, What shall I do?")
	if res == "1" {
		mode = 1
		goto EatMe
	} else if res == "2" {
		mode = 2
		goto DrinkMe
	} else if res == "3" {
		Cls()
		fmt.Println("Goodbye Friend!")
		time.Sleep(3 * time.seconds)
		goto Exit
	} else {
		fmt.Println("1, 2 or 3, it's all the same to me...")
		time.Sleep(3 * time.seconds)
		goto MainMenu
	}

EatMe:
	Cls()
	mode = 1
	files, err := SleepTalk.Map("Pokedex/Snorlax/Backpack/EatMe")
	if err != nil {
    	fmt.Println("Error:", err)
	} else {
    	fmt.Println("The Finest of Tavern Foods")
    	for _, filename := range files {
        	fmt.Println(filename)
    	}
	}



DrinkMe:
	Cls()
	mode = 2
	files, err := SleepTalk.Map("Pokedex/Snorlax/Backpack/DrinkMe")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("A Very Marry Un-Birthday to You!")
		fmt.Println("    The Tea Party Awaits")
		for _, filename := range files {
			fmt.Println(filename)
		}
	}


	

Exit:
	os.Exit(0)
}
