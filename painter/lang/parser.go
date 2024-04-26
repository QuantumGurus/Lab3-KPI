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

const screenSize = 800

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	uiState UiState
}

func (parser *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	parser.uiState.ResetOperations()
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
		if len(cmdLineParts) != 1 {
			return nil, errors.New("wrong number of arguments")
		}
		parser.uiState.SetWhiteBackground()
	case "green":
		if len(cmdLineParts) != 1 {
			return nil, errors.New("wrong number of arguments")
		}
		parser.uiState.SetGreenBackground()
	case "update":
		if len(cmdLineParts) != 1 {
			return nil, errors.New("wrong number of arguments")
		}
		parser.uiState.SetUpdatedOperation()

	case "bgrect":
		arguments, err := TakeArguments(cmdLineParts, 5, screenSize)
		if err != nil {
			return nil, err
		}
		parser.uiState.SetBackgroundRectangle(image.Point{X: arguments[0], Y: arguments[1]},
			image.Point{X: arguments[2], Y: arguments[3]})

	case "figure":
		arguments, err := TakeArguments(cmdLineParts, 3, screenSize)
		if err != nil {
			return nil, err
		}
		parser.uiState.AddFigure(image.Point{X: arguments[0], Y: arguments[1]})

	case "move":
		arguments, err := TakeArguments(cmdLineParts, 3, screenSize)
		if err != nil {
			return nil, err
		}
		parser.uiState.MoveFigures(image.Point{X: arguments[0], Y: arguments[1]})

	case "reset":
		if len(cmdLineParts) != 1 {
			return nil, errors.New("wrong number of arguments")
		}
		parser.uiState.ResetStateAndBackground()
	default:
		return nil, errors.New("some error")
	}
	return parser.uiState.GetParsedCommands(), nil
}

func TakeArguments(arguments []string, expectedArguments int, screenSize int) (correctArguments []int, err error) {
	if len(arguments) != expectedArguments {
		return nil, errors.New("wrong number of arguments")
	}
	var args []int
	args, err = ArgsToInt(arguments, screenSize)
	return args, nil
}

func ArgsToInt(tokens []string, screenSize int) ([]int, error) {
	args := make([]int, 0, len(tokens)-1)
	for _, arg := range tokens[1:] {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			return nil, err
		}
		args = append(args, int(val*float64(screenSize)))
	}
	return args, nil
}
