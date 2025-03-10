# 🗨️ Blabber

A modern terminal-based public chat application built with Go, featuring real-time communication through WebSocket technology.

## ✨ Features

- 💬 Real-time global public chatrooms
- 🔐 Server-side session authentication
- 🎨 Beautiful TUI client powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea)
- 💾 Persistent message storage with PostgreSQL
- ⌨️ Intuitive keyboard-driven interface

## 🚀 Getting Started

### Prerequisites

- Go 1.24.1 or higher
- PostgreSQL database

### Installation

```bash
# Clone the repository
git clone https://github.com/theHimanshuShekhar/blabber.git
cd blabber

# Install dependencies
go mod download

# Run the client
go run cmd/client/main.go
```

## 🎮 Usage

- Type your message and press `Enter` to send
- Press `Ctrl+C` or `Esc` to exit
- Navigate through chat history using arrow keys

## 🛣️ Roadmap

- [ ] Multiple chat room support
- [ ] Server side session management
- [ ] Message history
- [ ] Emoji support
- [ ] Anonymous chatting

## 🛠️ Built With

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions
- [PostgreSQL](https://www.postgresql.org/) - Database
- WebSocket - Real-time communication

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check the [issues page](https://github.com/theHimanshuShekhar/blabber/issues).
