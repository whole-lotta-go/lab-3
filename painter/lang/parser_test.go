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
			input: "white",
			want:  []painter.Operation{&painter.WhiteFill{}},
		},
		{
			name:  "Green fill operation",
			input: "green",
			want:  []painter.Operation{&painter.GreenFill{}},
		},

		{
			name:  "Update operation",
			input: "update",
			want:  []painter.Operation{painter.UpdateOp},
		},

		{
			name:  "BgRect operation with valid coordinates",
			input: "bgrect 0 0 100 100",
			want:  []painter.Operation{&painter.BgRect{X1: 0, Y1: 0, X2: 100, Y2: 100}},
		},
		{
			name:    "BgRect with invalid coordinates count",
			input:   "bgrect 0 0 100",
			wantErr: true,
		},
		{
			name:    "BgRect with non-numeric coordinates",
			input:   "bgrect a b c d",
			wantErr: true,
		},

		{
			name:  "Figure (TShape) operation with valid coordinates",
			input: "figure 150 200",
			want:  []painter.Operation{&painter.TShape{X: 150, Y: 200}},
		},
		{
			name:    "Figure with invalid coordinates count",
			input:   "figure 150",
			wantErr: true,
		},
		{
			name:    "Figure with non-numeric coordinates",
			input:   "figure x y",
			wantErr: true,
		},

		{
			name:  "Move operation with valid coordinates",
			input: "move 10 -20",
			want:  []painter.Operation{&painter.Move{Dx: 10, Dy: -20}},
		},
		{
			name:    "Move with invalid coordinates count",
			input:   "move 10",
			wantErr: true,
		},
		{
			name:    "Move with non-numeric coordinates",
			input:   "move x y",
			wantErr: true,
		},

		{
			name:  "Reset operation",
			input: "reset",
			want:  []painter.Operation{&painter.Reset{}},
		},

		{
			name:    "Unknown command",
			input:   "unknown",
			wantErr: true,
		},
		{
			name:    "Empty command",
			input:   "",
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
