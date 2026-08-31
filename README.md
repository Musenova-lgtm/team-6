# Tic-Tac-Toe Arena

A terminal-based Tic-Tac-Toe game written in Go. The project supports two game modes, a rule-based computer opponent, colored output, large board rendering, custom player names, different board sizes, and session statistics.

## Features

* Two human players mode
* Human vs Computer mode
* Rule-based AI opponent
* Configurable board size
* Colored terminal output
* Large board rendering
* Custom player names
* Player turn selection
* Session statistics
* Verbose statistics mode
* Input validation and error handling

## Project Structure

```text
tic-tac-toe-arena/
├── go.mod
├── main.go
├── README.md
└── internal/
    ├── cli/
    │   └── cli.go
    ├── board/
    │   └── board.go
    ├── game/
    │   └── game.go
    └── ai/
        └── ai.go
```

The project is divided into several internal packages:

* `internal/cli` — command-line flags, validation, and usage information.
* `internal/board` — board state, rendering, win and draw detection.
* `internal/game` — game loop, player input, and session statistics.
* `internal/ai` — rule-based computer opponent.

## Installation

Make sure Go is installed on your computer.

Clone the repository:

```bash
git clone <YOUR_REPOSITORY_URL>
cd tic-tac-toe-arena
```

Download dependencies:

```bash
go mod download
```

## How to Run

Run the game from the project root:

```bash
go run .
```

Exactly one game mode must be specified.

### Two Players

```bash
go run . --players
```

Two human players take turns. X and O play on the same computer.

### Against AI

```bash
go run . --ai
```

The human player is X and the computer is O.

### Available Flags

| Flag           | Description                        |
| -------------- | ---------------------------------- |
| `--players`    | Two human players                  |
| `--ai`         | Play against the computer          |
| `--color`      | Enable colored terminal output     |
| `--big`        | Render the board with large glyphs |
| `--verbose`    | Show extended statistics           |
| `--first X\|O` | Choose who moves first             |
| `--name A,B`   | Set custom names for X and O       |
| `--size N`     | Set board size to N×N, where N ≥ 3 |
| `--help`       | Show usage information             |
| `-h`           | Show usage information             |

Examples:

```bash
go run . --players --color
```

```bash
go run . --players --big
```

```bash
go run . --players --color --big --verbose
```

```bash
go run . --players --first O --name Alice,Bob
```

```bash
go run . --players --size 4
```

## Game Rules

The board cells are numbered from `1` to `N²`, from left to right and top to bottom.

For a standard 3×3 board:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9
```

Players choose a cell by entering its number.

A player wins when they complete a full row, column, or diagonal.

If the board is full and neither player has won, the game ends in a draw.

After each game, players can choose whether to play again. The session statistics are kept in memory.

Example:

```text
=== Stats ===
Games: 3   X: 1   O: 1   Draws: 1
```

## AI

In `--ai` mode, the human player is X and the computer is O.

The AI uses a deterministic rule-based strategy. It does not use recursion or minimax.

The AI checks the following rules in order:

1. **Win** — if O can win immediately, take that cell.
2. **Block** — if X can win on the next move, block X.
3. **Center** — take the center cell if it is free.
4. **Corner** — choose the first free corner in this order: `1, 3, 7, 9`.
5. **Side** — choose the first free side in this order: `2, 4, 6, 8`.

For example:

```text
O O .
X X .
. . .
```

The AI chooses cell `3` because it can complete its row and win immediately.

## Input Validation

Invalid moves are rejected without crashing the program.

Examples:

```text
X move: hello
Error: enter a number 1-9
```

```text
X move: 99
Error: enter a number 1-9
```

```text
X move: 5
Error: cell 5 is taken
```

The program continues asking for a valid move.

## Statistics

The game keeps track of:

* Total games
* X wins
* O wins
* Draws

With `--verbose`, additional information is displayed:

```text
=== Stats ===
Games: 3   X: 1   O: 1   Draws: 1
Moves this game: 7
Win rate — X: 33%  O: 33%
```

## Example Game

Running:

```bash
go run . --players
```

Example:

```text
 1 | 2 | 3
---+---+---
 4 | 5 | 6
---+---+---
 7 | 8 | 9

X move: 5

 1 | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9

O move: 1

 O | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | 9

X move: 9

 O | 2 | 3
---+---+---
 4 | X | 6
---+---+---
 7 | 8 | X

...

X wins!

=== Stats ===
Games: 1   X: 1   O: 0   Draws: 0

Play again? (y/n): n
```

## Team

Musenova Gaide — Musenova-lgtm
Adil Tolegen — Grond3749
Merey Zhaxybayeva — Mira1702-git

## Technologies

* Go
* Go standard library
* Git
* GitHub

## License

This project was created as a group programming project.
