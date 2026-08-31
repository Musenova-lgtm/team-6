package board

import (
	"fmt"
	"strconv"
	"strings"
)

type Board struct {
	Size  int
	Cells []string // "" означает пусто
}

func New(size int) *Board {
	return &Board{
		Size:  size,
		Cells: make([]string, size*size),
	}
}

func (b *Board) Place(index int, mark string) {
	if index >= 0 && index < len(b.Cells) {
		b.Cells[index] = mark
	}
}

func (b *Board) Undo(index int) {
	if index >= 0 && index < len(b.Cells) {
		b.Cells[index] = ""
	}
}

func (b *Board) IsFree(index int) bool {
	if index < 0 || index >= len(b.Cells) {
		return false
	}
	return b.Cells[index] == ""
}

func (b *Board) Full() bool {
	for _, cell := range b.Cells {
		if cell == "" {
			return false
		}
	}
	return true
}

func (b *Board) CheckWin(mark string) bool {
	line := b.WinningLine(mark)
	return len(line) > 0
}

func (b *Board) WinningLine(mark string) []int {
	s := b.Size

	// Проверка строк
	for r := 0; r < s; r++ {
		win := true
		line := make([]int, s)
		for c := 0; c < s; c++ {
			idx := r*s + c
			line[c] = idx
			if b.Cells[idx] != mark {
				win = false
				break
			}
		}
		if win {
			return line
		}
	}

	// Проверка столбцов
	for c := 0; c < s; c++ {
		win := true
		line := make([]int, s)
		for r := 0; r < s; r++ {
			idx := r*s + c
			line[r] = idx
			if b.Cells[idx] != mark {
				win = false
				break
			}
		}
		if win {
			return line
		}
	}

	// Проверка главной диагонали
	win1 := true
	line1 := make([]int, s)
	for i := 0; i < s; i++ {
		idx := i*s + i
		line1[i] = idx
		if b.Cells[idx] != mark {
			win1 = false
			break
		}
	}
	if win1 {
		return line1
	}

	// Проверка побочной диагонали
	win2 := true
	line2 := make([]int, s)
	for i := 0; i < s; i++ {
		idx := i*s + (s - 1 - i)
		line2[i] = idx
		if b.Cells[idx] != mark {
			win2 = false
			break
		}
	}
	if win2 {
		return line2
	}

	return nil
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func (b *Board) colorize(text string, isMark bool, isWin bool, color bool) string {
	if !color {
		return text
	}
	if isWin {
		return "\033[1;32m" + text + "\033[0m" // Зеленый и жирный
	}
	if !isMark {
		return "\033[2m" + text + "\033[0m" // Тусклый для цифр
	}
	if text == "X" {
		return "\033[91m" + text + "\033[0m" // Ярко-красный
	}
	if text == "O" {
		return "\033[94m" + text + "\033[0m" // Ярко-синий
	}
	return text
}

func (b *Board) Render(color bool, winningLine []int) string {
	var out strings.Builder
	s := b.Size

	maxNumLen := len(strconv.Itoa(s * s))

	for r := 0; r < s; r++ {
		for c := 0; c < s; c++ {
			idx := r*s + c
			val := b.Cells[idx]
			isWin := contains(winningLine, idx)

			var display string
			if val != "" {
				display = b.colorize(val, true, isWin, color)
				padding := strings.Repeat(" ", maxNumLen-1)
				display = padding + display
			} else {
				numStr := strconv.Itoa(idx + 1)
				padding := strings.Repeat(" ", maxNumLen-len(numStr))
				display = b.colorize(padding+numStr, false, isWin, color)
			}

			out.WriteString(fmt.Sprintf(" %s ", display))

			if c < s-1 {
				out.WriteString("|")
			}
		}
		out.WriteString("\n")

		if r < s-1 {
			// Разделитель строк
			dashes := strings.Repeat("-", maxNumLen+2)
			for c := 0; c < s; c++ {
				out.WriteString(dashes)
				if c < s-1 {
					out.WriteString("+")
				}
			}
			out.WriteString("\n")
		}
	}
	return out.String()
}

func (b *Board) RenderBig(color bool, winningLine []int) string {
	var out strings.Builder
	s := b.Size

	glyphX := []string{"X   X", "  X  ", "X   X"}
	glyphO := []string{" OOO ", "O   O", " OOO "}

	for r := 0; r < s; r++ {
		// Каждая клетка это 3 строки в высоту
		for line := 0; line < 3; line++ {
			for c := 0; c < s; c++ {
				idx := r*s + c
				val := b.Cells[idx]
				isWin := contains(winningLine, idx)

				var text string
				isMark := false

				switch val {
				case "X":
					text = glyphX[line]
					isMark = true
				case "O":
					text = glyphO[line]
					isMark = true
				default:
					if line == 1 {
						numStr := strconv.Itoa(idx + 1)
						if len(numStr) == 1 {
							text = fmt.Sprintf("  %s  ", numStr)
						} else {
							text = fmt.Sprintf(" %2s  ", numStr)
						}
					} else {
						text = "     "
					}
				}

				out.WriteString(b.colorize(text, isMark, isWin, color))

				if c < s-1 {
					out.WriteString("|")
				}
			}
			out.WriteString("\n")
		}

		if r < s-1 {
			for c := 0; c < s; c++ {
				out.WriteString("-----")
				if c < s-1 {
					out.WriteString("+")
				}
			}
			out.WriteString("\n")
		}
	}
	return out.String()
}
