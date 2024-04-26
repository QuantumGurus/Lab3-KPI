package lang

import (
	"github.com/roman-mazur/architecture-lab-3/painter"
	"image"
	"image/color"
)

type UiState struct {
	BackgroundColor     painter.Operation
	BackgroundRectangle *painter.Rectangle
	Figures             []*painter.Figure
	UpdateOperation     painter.Operation
	MoveOperations      []painter.Operation
}

func (uiState *UiState) SetWhiteBackground() {
	uiState.BackgroundColor = painter.OperationFunc(painter.WhiteFill)
}

func (uiState *UiState) SetGreenBackground() {
	uiState.BackgroundColor = painter.OperationFunc(painter.GreenFill)
}

func (uiState *UiState) SetBackgroundRectangle(leftPoint image.Point, rightPoint image.Point) {
	uiState.BackgroundRectangle = &painter.Rectangle{
		LeftPoint:  leftPoint,
		RightPoint: rightPoint,
	}
}

func (uiState *UiState) AddFigure(centralPoint image.Point) {
	figure := painter.Figure{
		FigureCentralPos: centralPoint,
		Color:            color.RGBA{R: 255, G: 255, B: 255, A: 255}}
	uiState.Figures = append(uiState.Figures, &figure)
}

func (uiState *UiState) Reset() {
	uiState.BackgroundColor = nil
	uiState.Figures = nil
	uiState.UpdateOperation = nil
	uiState.BackgroundRectangle = nil
}

func (uiState *UiState) SetUpdatedOperation() {
	uiState.UpdateOperation = painter.UpdateOp
}

func (uiState *UiState) MoveFigures(movePoint image.Point) {
	moveOper := painter.Move{NewPointCenter: movePoint, Figures: uiState.Figures}
	uiState.MoveOperations = append(uiState.MoveOperations, &moveOper)
}

func (uiState *UiState) GetParsedCommands() []painter.Operation {
	var res []painter.Operation
	if uiState.BackgroundColor != nil {
		res = append(res, uiState.BackgroundColor)
	}
	if uiState.BackgroundRectangle != nil {
		res = append(res, uiState.BackgroundRectangle)
	}
	if len(uiState.MoveOperations) != 0 {
		res = append(res, uiState.MoveOperations...)
	}
	uiState.MoveOperations = nil
	if len(uiState.Figures) != 0 {
		for _, figure := range uiState.Figures {
			res = append(res, figure)
		}
	}
	if uiState.UpdateOperation != nil {
		res = append(res, uiState.UpdateOperation)
	}
	return res
}

func (uiState *UiState) ResetOperations() {
	if uiState.BackgroundColor == nil {
		uiState.BackgroundColor = painter.OperationFunc(painter.Reset)
	}
	if uiState.UpdateOperation != nil {
		uiState.UpdateOperation = nil
	}
}

func (uiState *UiState) ResetStateAndBackground() {
	uiState.Reset()
	uiState.BackgroundColor = painter.OperationFunc(painter.Reset)
}
