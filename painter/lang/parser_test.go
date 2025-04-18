package lang

import (
	"reflect"
	"strings"
	"testing"

	"github.com/whole-lotta-go/lab-3/painter"
)

func Test_Parse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    []painter.Operation
		wantErr bool
	}{
		{
			name:  "White fill operation",
			input: "white\n",
			want:  []painter.Operation{painter.OperationFunc(painter.WhiteFill)},
		},
		{
			name:  "Green fill operation",
			input: "green\n",
			want:  []painter.Operation{painter.OperationFunc(painter.GreenFill)},
		},

		{
			name:  "Update operation",
			input: "update\n",
			want:  []painter.Operation{painter.UpdateOp},
		},

		{
			name:  "BgRect operation with valid coordinates",
			input: "bgrect 0 0 100 100\n",
			want:  []painter.Operation{&painter.BgRect{X1: 0, Y1: 0, X2: 100, Y2: 100}},
		},
		{
			name:    "BgRect with invalid coordinates count",
			input:   "bgrect 0 0 100\n",
			wantErr: true,
		},
		{
			name:    "BgRect with non-numeric coordinates",
			input:   "bgrect a b c d\n",
			wantErr: true,
		},

		{
			name:  "Figure (TShape) operation with valid coordinates",
			input: "figure 150 200\n",
			want:  []painter.Operation{&painter.TShape{X: 150, Y: 200}},
		},
		{
			name:    "Figure with invalid coordinates count",
			input:   "figure 150\n",
			wantErr: true,
		},
		{
			name:    "Figure with non-numeric coordinates",
			input:   "figure x y\n",
			wantErr: true,
		},

		{
			name:  "Move operation with valid coordinates",
			input: "move 10 -20\n",
			want:  []painter.Operation{&painter.Move{Dx: 10, Dy: -20}},
		},
		{
			name:    "Move with invalid coordinates count",
			input:   "move 10\n",
			wantErr: true,
		},
		{
			name:    "Move with non-numeric coordinates",
			input:   "move x y\n",
			wantErr: true,
		},

		{
			name:  "Reset operation",
			input: "reset\n",
			want:  []painter.Operation{&painter.Reset{}},
		},

		{
			name:    "Unknown command",
			input:   "unknown\n",
			wantErr: true,
		},
		{
			name:    "Empty command",
			input:   "\n",
			wantErr: false,
			want:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{}
			ops, err := p.Parse(strings.NewReader(tt.input))
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(ops, tt.want) {
				t.Errorf("Parse() = %v, want %v", ops, tt.want)
			}
		})
	}
}
