package main
import (
	"fmt"
	"motd/message"
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var name string
	var greeting string
	var prompt bool
	var preview bool

	//Define flags
	flag.StringVar(&name, "name", "", "name to use within the message")
	flag.StringVar(&greeting, "greeting", "", "phrase to use within the greeting")
	flag.BoolVar(&prompt, "prompt", false, "use prompt to input name and message")
	flag.BoolVar(&preview, "preview", false, "use preview to output message without writing to /etc/motd")
	//Parse flags
	flag.Parse()

	//If no arguments passed, show usage
	if prompt == false && (name == "" || greeting == ""){
		flag.Usage()
		os.Exit(1)
	}

	//Optionally print flags and exit if DEBUG is set
	if os.Getenv("DEBUG") != "" {
		fmt.Println("Name:", name)
		fmt.Println("Greeting:", greeting)
		fmt.Println("Prompt:", prompt)
		fmt.Println("Preview:", preview)

		os.Exit(0)
	}
	
	//Conditionally read from stdin
	if prompt {
		name, greeting = renderPrompt()
	}
	
	//Generate message
	m := message.Greeting(name, greeting)

	//Either preview the message or write to file
	if preview {
		fmt.Println(m)
	} else {
        // Open file to write
        f, err := os.OpenFile("/etc/motd", os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println("Error: Unable to open to /etc/motd")
            os.Exit(1)
        }
        defer f.Close()
        
        // Write to the file
        err = ioutil.WriteFile("/etc/motd", []byte(m), 0644)
        if err != nil {
            fmt.Println("Unable to write /etc/mod")
            os.Exit(1)
        }
	}
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Your Greeting: ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Your Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)

	return
}

