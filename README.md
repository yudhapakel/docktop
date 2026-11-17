# docktop

[![Go](https://img.shields.io/badge/Go-1.22+-00ADD8.svg?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Framework](https://img.shields.io/badge/TUI-Bubble%20Tea-7D56F4.svg?style=flat-square)](https://github.com/charmbracelet/bubbletea)
[![Styling](https://img.shields.io/badge/Style-Lip%20Gloss-EE6FF8.svg?style=flat-square)](https://github.com/charmbracelet/lipgloss)

Terminal User Interface (TUI) for inspecting Docker containers, images, volumes, and system resources, written in Go.

---

## Overview

`docktop` provides a keyboard-driven interface to observe local Docker environments directly from the terminal. Built on Charm's `bubbletea` and `lipgloss`, it follows the Model-Update-View architecture to deliver responsive navigation and minimal resource consumption.

---

## Features

- **Keyboard Navigation:** Vim-style keybindings (`j`/`k`, arrows, enter, esc).
- **Terminal Styling:** Structured panel layouts with dynamic terminal color palettes via `lipgloss`.
- **Modular Views:** Dedicated views for container status tables, image catalogs, volume inspectors, and system resource overviews.
- **On-Demand Refresh:** Immediate data re-polling via `r` keybinding across all active views.
- **Lightweight Binary:** Single static binary with zero external runtime dependencies.

---

## Keybindings

| Key | Action |
| :--- | :--- |
| `up` / `k` | Move selection up |
| `down` / `j` | Move selection down |
| `enter` | Select active view |
| `r` | Refresh active view data |
| `esc` / `b` | Return to main menu |
| `q` / `ctrl+c` | Exit application |

---

## Getting Started

### Prerequisites

- Go 1.22 or newer
- Local Docker Engine or Docker socket access

### Installation & Run

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

4. Build binary:
   ```bash
   go build -o docktop .
   ./docktop
   ```

---

## Project Structure

```text
docktop/
├── main.go        # Application lifecycle, state models, and view rendering
├── go.mod         # Module dependencies
├── go.sum         # Checksums
└── .gitignore     # Build artifacts ignore list
```

---

## Roadmap

- [x] Interactive container table with status and port bindings
- [x] Docker image inspection table with tags and size metrics
- [x] Docker volume inspector view with driver and mount points
- [x] System resource dashboard (CPU, memory, container counts)
- [x] On-demand view state refresh trigger (`r`)
- [ ] Real-time container log viewer with follow mode and search *(in progress)*
- [ ] Docker Engine SDK client integration (`moby/moby/client`)
- [ ] Container lifecycle actions (start, stop, restart, kill)
- [ ] Real-time container resource metrics streaming (CPU, RAM, Network I/O)
- [ ] Remote Docker daemon socket connection support

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
