package painter

import (
	"golang.org/x/exp/shiny/screen"
	"image"
	"image/color"
	"image/draw"
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

type Rectangle struct {
	LeftPoint  image.Point
	RightPoint image.Point
}

func (op *Rectangle) Do(t screen.Texture) bool {
	t.Fill(image.Rect(op.LeftPoint.X, op.LeftPoint.Y, op.RightPoint.X, op.RightPoint.Y), color.Black, screen.Src)
	return false
}

type Figure struct {
	FigureCentralPos image.Point
	Color            color.RGBA
}

func (op *Figure) Do(t screen.Texture) bool {
	red := color.RGBA{255, 0, 0, 255}
	hRect := image.Rect(op.FigureCentralPos.X-60, op.FigureCentralPos.Y-50, op.FigureCentralPos.X+60, op.FigureCentralPos.Y-30)
	vRect := image.Rect(op.FigureCentralPos.X-10, op.FigureCentralPos.Y-50, op.FigureCentralPos.X+10, op.FigureCentralPos.Y+50)
	t.Fill(hRect, red, draw.Src)
	t.Fill(vRect, red, draw.Src)
	return false
}

type Move struct {
	NewPointCenter image.Point
	Figures        []*Figure
}

func (op *Move) Do(t screen.Texture) bool {
	for i := range op.Figures {
		op.Figures[i].FigureCentralPos.X += op.NewPointCenter.X
		op.Figures[i].FigureCentralPos.Y += op.NewPointCenter.Y
	}
	return false
}

func Reset(t screen.Texture) {
	t.Fill(t.Bounds(), color.Black, screen.Src)
}
