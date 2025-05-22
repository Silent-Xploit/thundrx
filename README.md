# THUNDRX

**THUNDRX** is an ethical phishing simulation tool for red teamers and educational use only. It allows you to simulate real-world phishing attacks and monitor results in real time, with a focus on social engineering awareness and security training.

---

```
████████╗██╗  ██╗██╗   ██╗███╗   ██╗██████╗ ██████╗ ██╗  ██╗
╚══██╔══╝██║  ██║██║   ██║████╗  ██║██╔══██╗██╔══██╗╚██╗██╔╝
   ██║   ███████║██║   ██║██╔██╗ ██║██║  ██║██████╔╝ ╚███╔╝ 
   ██║   ██╔══██║██║   ██║██║╚██╗██║██║  ██║██╔══██╗ ██╔██╗ 
   ██║   ██║  ██║╚██████╔╝██║ ╚████║██████╔╝██║  ██║██╔╝ ██╗
   ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═══╝╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
```

---

## Features

- **Custom ASCII Banner** on launch
- **Authorization Prompt**: Confirms user is an ethical hacker or red teamer
- **CLI Menu** with options:
  - Camera + Microphone Access (Instagram, Facebook, Twitter, Google clones)
  - Live Location Capture
  - Keylogger Injection
  - Device Info Retrieval
  - Persistent Camera Video Access (Instagram, Facebook, Twitter, Google clones)
- **Ngrok Integration**: Automatically starts ngrok and provides public phishing/result monitoring links
- **Phishing Template Hosting**: Serves responsive HTML/JS templates for each attack
- **Live Logging**: All victim interactions are logged to disk and viewable via a web dashboard
- **Web-based Result Monitoring**: View logs and captured media in real time from anywhere
- **Cross-platform**: Works on Windows, Linux, and macOS
- **Automatic Free Port Selection**: Avoids port conflicts

---

## Quick Start

### 1. Download & Run (Go 1.21+ required)

```bash
# Install directly from GitHub 
go install github.com/Silent-Xploit/thundrx@latest
thundrx
```

**Or clone and run directly:**

```bash
git clone https://github.com/Silent-Xploit/ThunderX.git
cd thundrx
go run main.go
```

### 2. How it Works

```
1. You see the THUNDRX banner and are asked:
   Are you an ethical hacker or authorized red teamer? [y/N]:
2. Select an attack option from the menu (e.g. Camera + Microphone Access)
3. Choose a social media brand if prompted
4. The tool prints a phishing link and a result monitor link (ngrok or localhost)
5. Share the phishing link with your test target
6. Open the result monitor link in your browser to view logs and captured media in real time
```

---

## Directory Structure

```
internal/        # Go source code (handlers, helpers, server, etc.)
templates/       # Phishing HTML/JS templates for each attack
results/         # Captured logs and media (auto-created)
logs/            # Event logs
main.go          # Entry point
```

---

## Legal & Ethical Notice

- **For educational and authorized red team use only.**
- Do not use this tool for illegal activity.
- Always obtain written permission before simulating attacks on any system or user.

---

## Credits
- Built with Go, Gorilla Mux, and ngrok.
- Inspired by real-world red team and security awareness needs.

---

## License
This project is for ethical and educational use only. See LICENSE for details.
