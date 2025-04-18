package lang

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/whole-lotta-go/lab-3/painter"
	"github.com/whole-lotta-go/lab-3/ui"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct{}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var ops []painter.Operation

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		op, err := p.ParseLine(line)
		if err != nil {
			return nil, err
		}
		ops = append(ops, op)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scanner error: %w", err)
	}

	return ops, nil
}

func (p *Parser) ParseLine(line string) (painter.Operation, error) {
	if line == "" {
		return nil, fmt.Errorf("empty line")
	}

	tokens := strings.Fields(line)
	cmd := tokens[0]
	args, err := argsFloatToInt(tokens[1:])
	if err != nil {
		return nil, err
	}

	switch cmd {
	case "white":
		return &painter.WhiteFill{}, nil
	case "green":
		return &painter.GreenFill{}, nil
	case "update":
		return painter.UpdateOp, nil
	case "bgrect":
		if len(args) != 4 {
			return nil, fmt.Errorf("bgrect requires exactly 4 arguments")
		}
		x1, y1, x2, y2 := args[0], args[1], args[2], args[3]
		return &painter.BgRect{X1: x1, Y1: y1, X2: x2, Y2: y2}, nil
	case "figure":
		if len(args) != 2 {
			return nil, fmt.Errorf("figure requires exactly 2 arguments")
		}
		x, y := args[0], args[1]
		return &painter.TShape{X: x, Y: y}, nil
	case "move":
		if len(args) != 2 {
			return nil, fmt.Errorf("move requires exactly 2 arguments")
		}
		dx, dy := args[0], args[1]
		return &painter.Move{Dx: dx, Dy: dy}, nil
	case "reset":
		return &painter.Reset{}, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", cmd)
	}
}

func argsFloatToInt(args []string) ([]int, error) {
	coords := make([]int, 0)
	for _, arg := range args {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}
		coords = append(coords, coordFloatToInt(val))
	}
	return coords, nil
}

func coordFloatToInt(coord float64) int {
	return int(coord * ui.WindowSide)
}
