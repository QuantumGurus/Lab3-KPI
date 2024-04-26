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
		arguments, err := TakeArguments(cmdLineParts, 5)
		if err != nil {
			return nil, err
		}
		parser.uiState.SetBackgroundRectangle(image.Point{X: arguments[0], Y: arguments[1]},
			image.Point{X: arguments[2], Y: arguments[3]})

	case "figure":
		arguments, err := TakeArguments(cmdLineParts, 3)
		if err != nil {
			return nil, err
		}
		parser.uiState.AddFigure(image.Point{X: arguments[0], Y: arguments[1]})

	case "move":
		arguments, err := TakeArguments(cmdLineParts, 3)
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

func TakeArguments(arguments []string, expectedArguments int) (correctArguments []int, err error) {
	if len(arguments) != expectedArguments {
		return nil, errors.New("wrong number of arguments")
	}
	var args []int
	for _, arg := range arguments[1:] {
		if isInt(arg) {
			intNum, _ := strconv.Atoi(arg)
			args = append(args, intNum)
			continue
		} else if isFloat(arg) {
			floatNum, _ := strconv.ParseFloat(arg, 64)
			intNum := int(floatNum)
			args = append(args, intNum)
		} else {
			return nil, errors.New("wrong type of arguments")
		}
	}
	return args, nil
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
