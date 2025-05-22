package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ShowMenu() string {
	fmt.Println("\nMain Menu:")
	fmt.Println("a. Camera + Microphone Access")
	fmt.Println("b. Live Location Capture")
	fmt.Println("c. Keylogger Injection")
	fmt.Println("d. Device Info Retrieval")
	fmt.Println("f. Persistent Camera Video Access")
	fmt.Println("q. Quit")
	fmt.Print("Select an option: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		return strings.TrimSpace(strings.ToLower(scanner.Text()))
	}
	return ""
}
