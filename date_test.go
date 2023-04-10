package messageformat

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	doTest(t, Test{
		"{varname, date, short}",
		[]Expectation{
			{data: map[string]any{"varname": time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}, output: "4/10/2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, medium}",
		[]Expectation{
			{data: map[string]any{"varname": time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}, output: "April 10, 2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, long}",
		[]Expectation{
			{data: map[string]any{"varname": time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}, output: "Monday April 10, 2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, full}",
		[]Expectation{
			{data: map[string]any{"varname": time.Date(2023, 4, 10, 0, 0, 0, 0, time.UTC)}, output: "Monday 10. April 2023"},
		},
	})
}
