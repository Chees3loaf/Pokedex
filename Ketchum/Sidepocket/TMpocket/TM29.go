package TM29

import(
	"bufio"
	"io"
	"os/exec"
	"strconv"
	"strings"

	Potion "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket"
	Cyberball "github.com/Chees3loaf/Pokedex/Ketchum/Sidepocket/Ballpocket"
)

func Psychic() {
	
	func wait(seconds int) {
		done := make(chan struct{})
	
		go func() {
			time.Sleep(time.Duration(seconds) * time.Second)
			close(done)
		}()
	
		<-done
		//fmt.Printf("Waited for %d seconds./n,", seconds)
	
	}
	
// Clears screen in cmd
func cls() {
	
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	
}
	
func stringPrompt(label string) string {
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
	
func listAndSelectFiles(directory string) {
	files, err := Potion.Map(directory)
	if err != nil {
		fmt.Println("Error listing files:", err)
		return
	}
	
	fmt.Println("Files available:")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}
	
	choice := stringPrompt("Select a file by number to read its content:")
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(files) {
		fmt.Println("Invalid selection.")
		return
	}
	
	content, err := Potion.Load(directory + "/" + files[index-1])
		if err != nil {
			fmt.Println("Error reading file:", err)
		} else {
			fmt.Println(content)
		}
	
		printOption := stringPrompt("Save for the coming network apocalypse? (Yes/No)")
		if printOption == "Yes" {
			fmt.Println("Good choice")
			err = Cyberball.GeneratePDF(files[index-1], content)
			if err != nil {
				fmt.Println("Error generating PDF:", err)
			} else {
				fmt.Println("Keep it secret, keep it safe.", files[index-1]+".pdf")
			}
		} else if printOption == "No" {
			fmt.Println("Suit yourself")
		}
	}
	
func fetchFromServer(endpoint string) string {

	serverAddress := "10.0.0.55:55555"
	
		resp, err := http.Get(serverAddress + endpoint)
		if err != nil {
			fmt.Println("Error making request:", err)
			return ""
		}
		defer resp.Body.Close()
	
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Server returned an error:", resp.Status)
			return ""
		}
	
		reader := bufio.NewReader(resp.Body)
		content, err := reader.ReadString('\x00') // Read until a null character (which won't be found, so it reads the entire content)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading response:", err)
			return ""
		}
	
		return content
	}
	
}