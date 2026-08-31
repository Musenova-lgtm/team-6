package cli

import (
	"errors"
	"strconv"
	"strings"
)

type Config struct {
	Mode    string // "players" или "ai"
	Color   bool
	Big     bool
	Verbose bool
	First   string
	NameX   string
	NameO   string
	Size    int
	Help    bool
}

func Usage() string {
	return `Usage: go run main.go (--players | --ai) [options]

Modes (exactly one required):
  --players        two human players take turns
  --ai             play against the computer (you are X)

Options:
  --color          enable colored output (default: plain)
  --big            render the board with large glyphs
  --verbose        show extended statistics
  --first X|O      who moves first (default: X)
  --name A,B       custom names: X=A, O=B (e.g. --name Alice,Bob)
  --size N         board is N×N, win = N in a row (default: 3)
  --help, -h       print this help and exit 0
`
}

func Parse(args []string) (Config, error) {
	cfg := Config{
		First: "X",
		NameX: "X",
		NameO: "O",
		Size:  3,
	}

	hasPlayers := false
	hasAI := false
	hasSize := false

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "--help", "-h":
			cfg.Help = true
			return cfg, nil
		case "--players":
			hasPlayers = true
			cfg.Mode = "players"
		case "--ai":
			hasAI = true
			cfg.Mode = "ai"
		case "--color":
			cfg.Color = true
		case "--big":
			cfg.Big = true
		case "--verbose":
			cfg.Verbose = true
		case "--first":
			if i+1 >= len(args) {
				return cfg, errors.New("--first requires a value (X or O)")
			}
			cfg.First = strings.ToUpper(args[i+1])
			if cfg.First != "X" && cfg.First != "O" {
				return cfg, errors.New("--first must be X or O")
			}
			i++
		case "--name":
			if i+1 >= len(args) {
				return cfg, errors.New("--name requires a value (e.g. Alice,Bob)")
			}
			names := strings.Split(args[i+1], ",")
			if len(names) != 2 || names[0] == "" || names[1] == "" {
				return cfg, errors.New("--name value must contain a comma and two non-empty names")
			}
			cfg.NameX = names[0]
			cfg.NameO = names[1]
			i++
		case "--size":
			if i+1 >= len(args) {
				return cfg, errors.New("--size requires an integer value")
			}
			size, err := strconv.Atoi(args[i+1])
			if err != nil || size < 3 {
				return cfg, errors.New("--size must be an integer >= 3")
			}
			cfg.Size = size
			hasSize = true
			i++
		default:
			return cfg, errors.New("unknown flag: " + arg)
		}
	}

	if hasPlayers && hasAI {
		return cfg, errors.New("choose exactly one of --players or --ai")
	}
	if !hasPlayers && !hasAI {
		return cfg, errors.New("choose exactly one of --players or --ai")
	}
	if hasAI && hasSize {
		return cfg, errors.New("--ai and --size cannot be combined (AI is 3×3 only)")
	}

	return cfg, nil
}
