# Boodo

A simple terminal-based to-do list application built with Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

![Demo](https://github.com/user-attachments/assets/5706de88-950e-4fea-97df-563501af6848)


## Features

- Navigate tasks with ↑ and ↓
- Add new tasks (press `n`)
- Toggle completion (press `d`)
- Quit application (press `q` or `Ctrl+C`)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/blaze-d83/boodo.git
   cd boodo
   ```
2. Fetch dependencies:
   ```bash
   go mod tidy
   ```
3. Run the app:
   ```bash
   go run main.go
   ```

## Usage

Use the following keys to interact:

| Key         | Action              |
| ----------- | ------------------- |
| `n`         | Add a new task      |
| `d`         | Toggle task status  |
| `↑` / `↓`   | Move cursor         |
| `q` / `Ctrl+C` | Quit the app     |

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT © Your Name

