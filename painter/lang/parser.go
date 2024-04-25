package lang

import (
	"bufio"
	"errors"
	"image"
	"io"
	"strconv"
	"strings"

	"github.com/roman-mazur/architecture-lab-3/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	uiState UiState
}

func (parser *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	var res []painter.Operation
	for scanner.Scan() {
		commandLine := scanner.Text()
		op, err := parser.parse(commandLine) // parse the line to get Operation
		if err != nil {
			return nil, err
		}
		res = append(res, op...)
	}

	return res, nil
}

func (parser *Parser) parse(cmdLine string) ([]painter.Operation, error) {
	cmdLineParts := strings.Split(cmdLine, " ")
	instruction := cmdLineParts[0]
	switch instruction {
	case "white":
		parser.uiState.SetWhiteBackground()
	case "green":
		parser.uiState.SetGreenBackground()
	case "update":
		parser.uiState.SetUpdatedOperation()
	case "bgrect":
		arguments := cmdLineParts[1:5]
		num, err := strconv.Atoi(arguments[0])
		num2, err := strconv.Atoi(arguments[0])
		num3, err := strconv.Atoi(arguments[0])
		num4, err := strconv.Atoi(arguments[0])
		if err == nil {
			return nil, errors.New("s")
		}
		parser.uiState.SetBackgroundRectangle(image.Point{X: num, Y: num2}, image.Point{X: num3, Y: num4})
	case "figure":
		arguments := cmdLineParts[1:3]
		num, err := strconv.Atoi(arguments[0])
		num2, err := strconv.Atoi(arguments[0])
		if err == nil {
			return nil, errors.New("s")
		}
		parser.uiState.AddFigure(image.Point{X: num, Y: num2})

	case "move":
		arguments := cmdLineParts[1:3]
		num, err := strconv.Atoi(arguments[0])
		num2, err := strconv.Atoi(arguments[0])
		if err == nil {
			return nil, errors.New("s")
		}
		parser.uiState.MoveFigures(image.Point{X: num, Y: num2})
	case "reset":
		parser.uiState.Reset()
	default:
		return nil, errors.New("s")
	}

	return parser.uiState.GetParsedCommands(), nil
}
