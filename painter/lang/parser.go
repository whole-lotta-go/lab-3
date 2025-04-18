package lang

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/whole-lotta-go/lab-3/painter"
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
	args := tokens[1:]

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
		x1, err1 := strconv.Atoi(args[0])
		y1, err2 := strconv.Atoi(args[1])
		x2, err3 := strconv.Atoi(args[2])
		y2, err4 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			return nil, fmt.Errorf("invalid bgrect coordinates")
		}
		return &painter.BgRect{X1: x1, Y1: y1, X2: x2, Y2: y2}, nil
	case "figure":
		if len(args) != 2 {
			return nil, fmt.Errorf("figure requires exactly 2 arguments")
		}
		x, err1 := strconv.Atoi(args[0])
		y, err2 := strconv.Atoi(args[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid figure coordinates")
		}
		return &painter.TShape{X: x, Y: y}, nil
	case "move":
		if len(args) != 2 {
			return nil, fmt.Errorf("move requires exactly 2 arguments")
		}
		dx, err1 := strconv.Atoi(args[0])
		dy, err2 := strconv.Atoi(args[1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("invalid move coordinates")
		}
		return &painter.Move{Dx: dx, Dy: dy}, nil
	case "reset":
		return &painter.Reset{}, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", cmd)
	}
}
