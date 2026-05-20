# 🎨 Go ASCII Art Generator

A sleek, terminal-inspired web application that transforms plain text into stunning ASCII art. Built with Go for the backend and a modern, responsive frontend.

## ✨ Features

- **Customizable ASCII Art**: Generate art using multiple banner styles:
  - `Standard`: The classic ASCII look.
  - `Shadow`: Adds a depth effect.
  - `Thinkertoy`: A distinctive, stylized font.
- **Cyber-Console Aesthetic**: A professional terminal-like interface with a dark-mode default.
- **Dynamic Theme Switching**: Toggle between a high-contrast **Neon Yellow / Obsidian** dark theme and a clean **Orange / White** light theme.
- **User-Friendly Tools**:
    - One-click copy to clipboard.
    - Responsive design (works on desktop and mobile).
    - Persistent theme preference using `localStorage`.
- **Robust Backend**: Powered by Go, ensuring fast processing and efficient request handling.

## 🛠️ Tech Stack

- **Backend**: [Go (Golang)](https://go.dev/)
- **Frontend**: HTML5, CSS3 (Custom Properties / Design Tokens), Vanilla JavaScript
- **Styling**: Custom CSS with a focus on fluid typography and GPU-accelerated animations.

## 🚀 Getting Started

### Prerequisites
- Go 1.21+ installed on your machine.

### Installation & Running
1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/ascii-art-web-stylize.git
   cd ascii-art-web-stylize
   ```

2. Run the server:
   ```bash
   go mod init ascii-art-web-stylize
   go mod tidy
   go run cmd/main.go
   ```

3. Open your browser and navigate to:
   `http://localhost:8080` (or the port specified in the console output).

## 📂 Project Structure

```text
.
├── assets/              # Static files (CSS, Images)
│   └── css/
│       └── styles.css   # Global design tokens and styles
├── banners/             # ASCII font definition files
├── cmd/                 # Application entry point and internal logic
│   ├── main.go          # Server configuration and routing
│   └── internal/
│       ├── asciiart/    # Logic for ASCII generation
│       └── server/      # HTTP handlers and template rendering
├── index.html               # Landing page
├── templates/           # HTML templates
│   ├── asciiart.html    # Result page
│   └── error.html       # Error handling page
└── go.mod               # Go module definitions
```

## 📝 License
This project is developed for educational purposes as part of the 01Talent curriculum.
