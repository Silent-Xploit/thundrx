package internal

import (
	"os/exec"
	"regexp"
	"strings"
)

// StartNgrok starts ngrok and returns the public URL and the process pointer
func StartNgrok(port string) (string, *exec.Cmd, error) {
	cmd := exec.Command("ngrok", "http", port, "--log=stdout")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", nil, err
	}
	if err := cmd.Start(); err != nil {
		return "", nil, err
	}

	// Wait for ngrok to print the public URL
	buf := make([]byte, 4096)
	for {
		n, _ := stdout.Read(buf)
		output := string(buf[:n])
		if strings.Contains(output, "https://") {
			re := regexp.MustCompile(`https://[a-zA-Z0-9\-]+\.ngrok-free\.app`)
			match := re.FindString(output)
			if match != "" {
				return match, cmd, nil
			}
		}
	}
}
