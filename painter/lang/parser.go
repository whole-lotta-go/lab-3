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
type Parser struct {
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	var res []painter.Operation

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd := parts[0]
		args := parts[1:]

		switch cmd {
		case "reset":
			res = append(res, &painter.Reset{})
		case "green":
			res = append(res, painter.OperationFunc(painter.GreenFill))
		case "white":
			res = append(res, painter.OperationFunc(painter.WhiteFill))
		case "T":
			if len(args) != 2 {
				return nil, fmt.Errorf("T command requires exactly 2 arguments")
			}
			x, err1 := strconv.Atoi(args[0])
			y, err2 := strconv.Atoi(args[1])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid arguments for T command")
			}
			res = append(res, &painter.TShape{X: x, Y: y})
		case "bgrect":
			if len(args) != 4 {
				return nil, fmt.Errorf("bgrect command requires exactly 4 arguments")
			}
			x1, err1 := strconv.Atoi(args[0])
			y1, err2 := strconv.Atoi(args[1])
			x2, err3 := strconv.Atoi(args[2])
			y2, err4 := strconv.Atoi(args[3])
			if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
				return nil, fmt.Errorf("invalid arguments for bgrect command")
			}
			res = append(res, &painter.BgRect{X1: x1, Y1: y1, X2: x2, Y2: y2})
		default:
			return nil, fmt.Errorf("unknown command: %s", cmd)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
