package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConfirmEthicalHacker() bool {
	fmt.Print("Are you an ethical hacker or authorized red teamer? [y/N]: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		resp := strings.TrimSpace(strings.ToLower(scanner.Text()))
		return resp == "y"
	}
	return false
}
