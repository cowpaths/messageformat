package messageformat

import (
	"bytes"
	"fmt"

	"github.com/gotnospirit/makeplural/plural"
	"golang.org/x/text/language"
)

var (
	DefaultLocale = language.English
)

// pluralFunc describes a function used to produce a named key when processing a plural or selectordinal expression.
type pluralFunc func(interface{}, bool) string

type Formatter interface {
	Format(*ParseTree) (string, error)
	FormatMap(*ParseTree, map[string]any) (string, error)
}

type FormatterOption func(f *formatter)

func WithLocale(locale language.Tag) FormatterOption {
	return func(f *formatter) {
		f.SetLocale(locale)
	}
}

// NewFormatter creates a new formatter with the given options
func NewFormatter(opts ...FormatterOption) (Formatter, error) {
	f := formatter{}

	for _, opt := range opts {
		opt(&f)
	}

	if f.locale == language.Und {
		err := f.SetLocale(DefaultLocale)
		if err != nil {
			return nil, err
		}
	}

	return &f, nil
}

type formatter struct {
	locale language.Tag
	plural pluralFunc
	date   DateFormatter
}

func (x *formatter) SetLocale(tag language.Tag) error {
	x.locale = tag

	switch tag {
	case language.AmericanEnglish:
		x.date = createAmericanDateFormatter()
	case language.German:
		x.date = createGermanDateFormatter()
	default:
		x.date = createAmericanDateFormatter()
	}

	if x.plural == nil {
		culture := tag
		for culture != language.Und {
			fn, err := plural.GetFunc(culture.String())
			if err != nil {
				x.plural = fn
				break
			}
			culture = culture.Parent()
		}
	}

	return nil
}

func (x *formatter) SetPluralFunction(fn pluralFunc) error {
	if nil == fn {
		return fmt.Errorf("PluralFunctionRequired")
	}
	x.plural = fn

	return nil
}

func (f *formatter) Format(n *ParseTree) (string, error) {
	return f.FormatMap(n, nil)
}

func (f *formatter) FormatMap(n *ParseTree, data map[string]any) (string, error) {
	var buf bytes.Buffer

	err := f.format(n, &buf, data, "")
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (f *formatter) format(n *ParseTree, buf *bytes.Buffer, data map[string]any, value string) error {
	err := n.ForEach(func(n *Node) error {
		switch n.Type {
		case "date":
			return f.formatDate(n.Expr, buf, data)
		case "literal":
			return f.formatLiteral(n.Expr, buf, value)
		case "plural":
			return f.formatPlural(n.Expr, buf, data)
		case "select":
			return f.formatSelect(n.Expr, buf, data)
		case "selectordinal":
			return f.formatOrdinal(n.Expr, buf, data)
		case "time":
			return fmt.Errorf("formatter not implemented for time")
		case "var":
			return f.formatVar(n.Expr, buf, data)
		default:
			return fmt.Errorf("formatter not implemented for expression of type %s", n.Type)
		}
	})
	if nil != err {
		return err
	}

	return nil
}

func (f *formatter) getNamedKey(value interface{}, ordinal bool) (string, error) {
	if nil == f.plural {
		return "", fmt.Errorf("UndefinedPluralFunc")
	}

	return f.plural(value, ordinal), nil
}
