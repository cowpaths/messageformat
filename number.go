package messageformat

import (
	"bytes"
	"errors"
	"fmt"

	"golang.org/x/text/message"
)

type NumberExpr struct {
	Name string
}

func (p *parser) parseNumber(varName string, char rune, start, end int, ptr_input *[]rune) (Expression, int, error) {
	var result = NumberExpr{
		Name: varName,
	}

	pos := start

	if char != CloseChar {
		format, _, cursor, err := readVar(pos+1, end, ptr_input)
		if err != nil {
			return nil, pos, errors.New("failed to parse number format")
		}
		pos = cursor

		return nil, pos, fmt.Errorf("number format not implemented: %s", format)
	}

	return &result, pos, nil
}

func (f *formatter) formatNumber(expr Expression, ptrOutput *bytes.Buffer, data map[string]any) error {
	e, ok := expr.(*NumberExpr)
	if !ok {
		return fmt.Errorf("InvalidExprType: want NumberExpr, got %T", e)
	}

	n, ok := data[e.Name]
	if !ok {
		return fmt.Errorf("InvalidArgType: want number, got %T", n)
	}

	p := message.NewPrinter(f.locale)
	ptrOutput.WriteString(p.Sprint(n))

	return nil
}
