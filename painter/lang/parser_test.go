package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"image"
	"image/color"
	"strings"
	"testing"
)

func Test_parse_struct(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "background rectangle",
			command: "bgrect 0 0 0.25 0.25",
			op:      &painter.Rectangle{LeftPoint: image.Point{X: 0, Y: 0}, RightPoint: image.Point{X: 200, Y: 200}},
		},
		{
			name:    "figure",
			command: "figure 0.2 0.2",
			op:      &painter.Figure{FigureCentralPos: image.Point{X: 160, Y: 160}, Color: color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		},
		{
			name:    "move",
			command: "move 0.1 0.1",
			op:      &painter.Move{NewPointCenter: image.Point{X: 80, Y: 80}},
		},
		{
			name:    "update",
			command: "update",
			op:      painter.UpdateOp,
		},
		{
			name:    "invalid command",
			command: "abrakadabra",
			op:      nil,
		},
		{
			name:    "not enough args",
			command: "figure 1",
			op:      nil,
		},
		{
			name:    "float numbers",
			command: "bgrect 10 10 122.3 122.3",
			op:      &painter.Rectangle{LeftPoint: image.Point{X: 10, Y: 10}, RightPoint: image.Point{X: 122, Y: 122}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parser := &Parser{}
			ops, err := parser.Parse(strings.NewReader(tc.command))
			if tc.op == nil {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.IsType(t, tc.op, ops[1])
				assert.Equal(t, tc.op, ops[1])
			}
		})
	}
}

func Test_parse_func(t *testing.T) {
	tests := []struct {
		name    string
		command string
		op      painter.Operation
	}{
		{
			name:    "white fill",
			command: "white",
			op:      painter.OperationFunc(painter.WhiteFill),
		},
		{
			name:    "green fill",
			command: "green",
			op:      painter.OperationFunc(painter.GreenFill),
		},
		{
			name:    "reset",
			command: "reset",
			op:      painter.OperationFunc(painter.Reset),
		},
	}

	parser := &Parser{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ops, err := parser.Parse(strings.NewReader(tc.command))

			require.NoError(t, err)
			require.Len(t, ops, 1)
			assert.IsType(t, tc.op, ops[0])

		})
	}
}
