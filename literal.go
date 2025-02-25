package messageformat

import (
	"bytes"
	"fmt"
)

// LiteralExpr represents a string literal
type LiteralExpr struct {
	Values []string
}

func (f *formatter) formatLiteral(expr Expression, ptr_output *bytes.Buffer, pound string) error {
	literal, ok := expr.(LiteralExpr)
	if !ok {
		return fmt.Errorf("InvalidExprType: want LiteralExpr, got: %T", expr)
	}

	for _, val := range literal.Values {
		if val != "" {
			ptr_output.WriteString(val)
		} else if pound != "" {
			ptr_output.WriteString(pound)
		} else {
			ptr_output.WriteRune(PoundChar)
		}
	}

	return nil
}

func (p *parser) parseLiteral(start, end int, ptr_input *[]rune) LiteralExpr {
	var items []int

	input := *ptr_input
	escaped := false

	s, e := start, start
	gap := 0
	for i := start; i < end; i++ {
		c := input[i]

		if EscapeChar == c {
			gap++
			e++
			escaped = true
		} else {
			switch c {
			default:
				e++

			case OpenChar, CloseChar, PoundChar:
				if escaped {
					if i-s > gap {
						if gap > 1 {
							items = append(items, s, i)
						} else {
							items = append(items, s, i-1)
						}
					}
					s = i
				} else {
					if s != e {
						items = append(items, s, e, i, i)
					} else if s != i {
						items = append(items, s, i, i, i)
					} else {
						items = append(items, i, i)
					}
					s = i + 1
				}
				e = s
			}

			escaped = false
			gap = 0
		}
	}

	if s < end {
		items = append(items, s, end)
	}

	n := len(items)
	expr := LiteralExpr{
		Values: make([]string, n/2),
	}

	for i := 0; i < n; i += 2 {
		expr.Values[i/2] = string(input[items[i]:items[i+1]])
	}

	return expr
}
