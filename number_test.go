package messageformat

import (
	"testing"
)

func TestNumber(t *testing.T) {
	doTest(t, Test{
		"{varname, number}",
		[]Expectation{
			{data: map[string]any{"varname": 1234}, output: "1,234"},
		},
	})
	doTest(t, Test{
		"{varname, number}",
		[]Expectation{
			{data: map[string]any{"varname": 1234.5}, output: "1,234.5"},
		},
	})
	doTest(t, Test{
		"{varname, number}",
		[]Expectation{
			{data: map[string]any{"varname": float64(12345.678)}, output: "12,345.678"},
		},
	})
	doTest(t, Test{
		"{varname, number}",
		[]Expectation{
			{data: map[string]any{"varname": uint64(12345678)}, output: "12,345,678"},
		},
	})
}
