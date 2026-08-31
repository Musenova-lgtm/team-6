package ai

import (
	"github.com/Mira1702-git/tic-tac-toe-arena/internal/board"
)

// FindMove возвращает индекс следующего хода и текстовую причину (для verbose)
func FindMove(b *board.Board) (int, string) {
	// 1. Win
	for i := 0; i < len(b.Cells); i++ {
		if b.IsFree(i) {
			b.Place(i, "O")
			if b.CheckWin("O") {
				b.Undo(i)
				return i, "win"
			}
			b.Undo(i)
		}
	}

	// 2. Block
	for i := 0; i < len(b.Cells); i++ {
		if b.IsFree(i) {
			b.Place(i, "X")
			if b.CheckWin("X") {
				b.Undo(i)
				return i, "block"
			}
			b.Undo(i)
		}
	}

	// 3. Center (индекс 4 для поля 3x3)
	if b.Size == 3 && b.IsFree(4) {
		return 4, "center"
	}

	// 4. Corner (индексы 0, 2, 6, 8)
	corners := []int{0, 2, 6, 8}
	for _, i := range corners {
		if b.IsFree(i) {
			return i, "corner"
		}
	}

	// 5. Side (индексы 1, 3, 5, 7)
	sides := []int{1, 3, 5, 7}
	for _, i := range sides {
		if b.IsFree(i) {
			return i, "side"
		}
	}

	return -1, "none"
}
