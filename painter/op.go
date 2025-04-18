package painter

import (
	"image"
	"image/color"
	"image/draw"

	"golang.org/x/exp/shiny/screen"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

type BgRect struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

func (op *BgRect) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.X1, op.Y1, op.X2, op.Y2), color.Black, screen.Src)
	return false
}

type TShape struct {
	X int
	Y int
}

func (op *TShape) Do(t screen.Texture) bool {
	yellow := color.RGBA{R: 255, G: 255, B: 0, A: 255}
	sideLen := 100

	stem := image.Rect(
		op.X-sideLen/2,
		op.Y-sideLen,
		op.X+sideLen/2,
		op.Y,
	)

	head := image.Rect(
		op.X-sideLen*3/2,
		op.Y,
		op.X+sideLen*3/2,
		op.Y+sideLen,
	)

	t.Fill(stem, yellow, draw.Src)
	t.Fill(head, yellow, draw.Src)
	return false
}

type Move struct {
	Dx     int
	Dy     int
	Shapes []*TShape
}

func (op *Move) Do(t screen.Texture) bool {
	for _, fig := range op.Shapes {
		fig.X += op.Dx
		fig.Y += op.Dy
	}
	return false
}

type Reset struct{}

func (op *Reset) Do(t screen.Texture) bool {
	t.Fill(t.Bounds(), color.Black, screen.Src)
	return false
}
