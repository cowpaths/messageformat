package messageformat

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tz, _ := time.LoadLocation("America/New_York")
	dateTime := time.Date(2023, 4, 10, 5, 19, 42, 543, tz)

	doTest(t, Test{
		"{varname, date}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "4/10/2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, short}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "4/10/2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, medium}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "April 10, 2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, long}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "Monday April 10, 2023"},
		},
	})
	doTest(t, Test{
		"{varname, date, full}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "Monday 10. April 2023"},
		},
	})
}

func TestParseTime(t *testing.T) {
	tz, _ := time.LoadLocation("America/New_York")
	dateTime := time.Date(2023, 4, 10, 5, 19, 42, 543, tz)

	doTest(t, Test{
		"{varname, time}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "5:19 AM"},
		},
	})
	doTest(t, Test{
		"{varname, time, short}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "5:19 AM"},
		},
	})
	doTest(t, Test{
		"{varname, time, medium}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "5:19:42 AM"},
		},
	})
	doTest(t, Test{
		"{varname, time, long}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "5:19:42 AM EDT"},
		},
	})
	doTest(t, Test{
		"{varname, time, full}",
		[]Expectation{
			{data: map[string]any{"varname": dateTime}, output: "5:19:42 AM EDT"},
		},
	})
}
