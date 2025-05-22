package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var logDir = "logs"
var eventLog = filepath.Join(logDir, "event.log")

func LogEvent(event string) {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}
	f, err := os.OpenFile(eventLog, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[!] Failed to log event:", err)
		return
	}
	defer f.Close()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	f.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, event))
}
