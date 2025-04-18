package painter

import (
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(s *State) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(s *State) (ready bool) {
	for _, o := range ol {
		ready = o.Do(s) || ready
	}
	return
}

type UpdateOp struct{}

func (op *UpdateOp) Do(s *State) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(s *State)

func (f OperationFunc) Do(s *State) bool {
	f(s)
	return false
}

type Fill struct {
	Color color.Color
}

func (op *Fill) Do(s *State) bool {
	s.background.color = op.Color
	return false
}

type BgRect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (op *BgRect) Do(s *State) bool {
	rect := image.Rect(op.X1, op.Y1, op.X2, op.Y2)
	s.background.rect = &rect
	return false
}

type Figure struct {
	X int
	Y int
}

func (op *Figure) Do(s *State) bool {
	p := image.Point{op.X, op.Y}
	s.figurePoints = append(s.figurePoints, &p)
	return false
}

type Move struct {
	X int
	Y int
}

func (op *Move) Do(s *State) bool {
	for _, fp := range s.figurePoints {
		fp.X = op.X
		fp.Y = op.Y
	}
	return false
}

type Reset struct{}

func (op *Reset) Do(s *State) bool {
	s.background.color = color.Black
	s.background.rect = nil
	s.figurePoints = s.figurePoints[:0]
	return false
}
