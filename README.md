![Boodo Banner](https://github.com/user-attachments/assets/9b29ca98-d70d-49a6-ad7d-ba4f68da72a8)


# BooDo — Terminal TODO App

## Overview

BooDo is a simple, elegant, and keyboard-friendly TODO list application running directly in your terminal. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea) and styled via [Lipgloss](https://github.com/charmbracelet/lipgloss), it offers:

* Framed, modern TUI layout
* Quick task addition, toggling, and deletion
* Persistent storage in `$XDG_CONFIG_HOME/boodo/tasks.json`
* Intuitive keyboard controls (j/k navigation, d toggle, x delete, n new, q quit)

## Features

* **Add Tasks**
* **Toggle Done**
* **Delete Tasks**
* **Navigate**
* **Persistent Storage**

## Installation

1. **Clone the repo**:

   ```bash
   git clone https://github.com/yourusername/boodo.git
   cd boodo
   ```

2. **Build the binary**:

   ```bash
   go build -o boodo main.go
   ```

3. **Run**:

   ```bash
   ./boodo
   ```

## Configuration & Data Location

By default, BooDo stores tasks in:

```
$XDG_CONFIG_HOME/boodo/tasks.json
```

If `$XDG_CONFIG_HOME` is not set, it falls back to `$HOME/.config/boodo/tasks.json`.

## Usage

```bash
# Launch the app
./boodo

# Keyboard shortcuts:
# n      Open new task input
# Enter  Save new task
# esc    Cancel input mode
# j/k    Move cursor up/down
# d      Toggle task done
# x      Delete task
# q      Quit (auto-saves)
```

## Screenshot / Demo

![Demo](https://github.com/user-attachments/assets/5706de88-950e-4fea-97df-563501af6848)


## License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.
