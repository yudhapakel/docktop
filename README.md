# docktop

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8.svg?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Framework](https://img.shields.io/badge/TUI-Bubble%20Tea-7D56F4.svg?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Styling](https://img.shields.io/badge/Style-Lip%20Gloss-EE6FF8.svg?style=flat-square)](https://github.com/charmbracelet/lipgloss)

A terminal user interface (TUI) for monitoring and inspecting Docker containers, images, and system resources, written in Go.

---

## Overview

`docktop` provides a fast, keyboard-driven interface to observe local and remote Docker engines without leaving the terminal. Built on top of the Charm ecosystem (`bubbletea` and `lipgloss`), it offers a clean visual layout following the Model-Update-View (Elm) architecture.

---

## Features

- **Interactive TUI:** Smooth keyboard navigation with Vim-style bindings (`j`/`k` or arrow keys).
- **Modern Terminal Styling:** Styled layouts with dynamic color palettes via `lipgloss`.
- **Modular Views:** Structured multi-panel architecture covering containers, images, system metrics, and real-time logs.
- **Resource Efficient:** Single binary execution with zero runtime dependencies.

---

## Keybindings

| Key | Action |
| :--- | :--- |
| `up` / `k` | Move selection up |
| `down` / `j` | Move selection down |
| `enter` | Select item / Open detailed view |
| `q` / `ctrl+c` | Quit application |

---

## Getting Started

### Prerequisites

- Go 1.22 or newer
- Docker Engine / Docker Desktop (running locally or accessible via socket)

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yudhapakel/docktop.git
   cd docktop
   ```

2. Download dependencies:
   ```bash
   go mod tidy
   ```

3. Run directly:
   ```bash
   go run .
   ```

4. Build a standalone binary:
   ```bash
   go build -o docktop .
   ./docktop
   ```

---

## Project Structure

```text
docktop/
├── main.go        # Application entrypoint and Bubble Tea lifecycle
├── go.mod         # Module dependencies definition
├── go.sum         # Checksums for module dependencies
└── .gitignore     # Git ignore rules for compiled binaries
```

---

## Roadmap

- [ ] Docker Engine SDK integration (`moby/moby/client`)
- [ ] Container lifecycle management (start, stop, restart, kill)
- [ ] Live resource usage streaming (CPU, RAM, Network I/O)
- [ ] Container log inspection with search and auto-scroll
- [ ] Multi-host Docker daemon connection support

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
