package lang

import (
	"image/color"
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
			want:  []painter.Operation{&painter.Fill{Color: color.White}},
		},
		{
			name:  "Green fill operation",
			input: "green",
			want:  []painter.Operation{&painter.Fill{Color: color.RGBA{R: 0, G: 255, B: 0, A: 255}}},
		},

		{
			name:  "Update operation",
			input: "update",
			want:  []painter.Operation{&painter.UpdateOp{}},
		},

		{
			name:  "BgRect operation with valid coordinates",
			input: "bgrect 0 0 0.1 0.2",
			want:  []painter.Operation{&painter.BgRect{X1: 0, Y1: 0, X2: 80, Y2: 160}},
		},
		{
			name:    "BgRect with invalid coordinates count",
			input:   "bgrect 0 0 0.1",
			wantErr: true,
		},
		{
			name:    "BgRect with non-numeric coordinates",
			input:   "bgrect a b c d",
			wantErr: true,
		},

		{
			name:  "Figure (TShape) operation with valid coordinates",
			input: "figure 0.25 0.5",
			want:  []painter.Operation{&painter.Figure{X: 200, Y: 400}},
		},
		{
			name:    "Figure with invalid coordinates count",
			input:   "figure 0.25",
			wantErr: true,
		},
		{
			name:    "Figure with non-numeric coordinates",
			input:   "figure x y",
			wantErr: true,
		},

		{
			name:  "Move operation with valid coordinates",
			input: "move 0.1 0.5",
			want:  []painter.Operation{&painter.Move{X: 80, Y: 400}},
		},
		{
			name:    "Move with invalid coordinates count",
			input:   "move 0.1",
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
