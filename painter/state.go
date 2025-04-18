package painter

import (
	"image"
	"image/color"

	"golang.org/x/exp/shiny/screen"
)

type State struct {
	background struct {
		color color.Color
		rect  *image.Rectangle
	}

	figurePoints []*image.Point
}

func (s *State) Draw(t screen.Texture) {
	if s.background.color != nil {
		t.Fill(t.Bounds(), s.background.color, screen.Src)
	}

	if s.background.rect != nil {
		t.Fill(*s.background.rect, color.Black, screen.Src)
	}

	for _, f := range s.figurePoints {
		drawFigure(*f, t)
	}
}

func drawFigure(center image.Point, t screen.Texture) {
	yellow := color.RGBA{255, 255, 0, 255}
	blockSide := 100

	stem := image.Rect(
		center.X-blockSide/2,
		center.Y-blockSide,
		center.X+blockSide/2,
		center.Y,
	)

	head := image.Rect(
		center.X-blockSide*3/2,
		center.Y,
		center.X+blockSide*3/2,
		center.Y+blockSide,
	)

	t.Fill(stem, yellow, screen.Src)
	t.Fill(head, yellow, screen.Src)
}
