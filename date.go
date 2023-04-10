package messageformat

import (
	"bytes"
	"errors"
	"fmt"
	"time"
)

// DateTimeFormatter is an interface for a type that formats
// a date/time in a variety of formats.
type DateTimeFormatter interface {
	ShortDate(time.Time) string
	MediumDate(time.Time) string
	LongDate(time.Time) string
	FullDate(time.Time) string
	ShortTime(time.Time) string
	MediumTime(time.Time) string
	LongTime(time.Time) string
	FullTime(time.Time) string
}

type AmericanDateTimeFormatter struct {
	Month   map[time.Month]string
	Weekday map[time.Weekday]string
}

func createAmericanDateTimeFormatter() DateTimeFormatter {
	return &AmericanDateTimeFormatter{
		Month: map[time.Month]string{
			time.January:   "January",
			time.February:  "February",
			time.March:     "March",
			time.April:     "April",
			time.May:       "May",
			time.June:      "June",
			time.July:      "July",
			time.August:    "August",
			time.September: "September",
			time.October:   "October",
			time.November:  "November",
			time.December:  "December",
		},
		Weekday: map[time.Weekday]string{
			time.Monday:    "Monday",
			time.Tuesday:   "Tuesday",
			time.Wednesday: "Wednesday",
			time.Thursday:  "Thursday",
			time.Friday:    "Friday",
			time.Saturday:  "Saturday",
			time.Sunday:    "Sunday",
		},
	}
}

// 10/16/1996
func (en *AmericanDateTimeFormatter) ShortDate(t time.Time) string {
	return fmt.Sprintf("%d/%d/%d", t.Month(), t.Day(), t.Year())
}

// October 16, 1996
func (en *AmericanDateTimeFormatter) MediumDate(t time.Time) string {
	return fmt.Sprintf("%s %d, %d", en.Month[t.Month()], t.Day(), t.Year())
}

// Tuesday October 16, 1996
func (en *AmericanDateTimeFormatter) LongDate(t time.Time) string {
	return fmt.Sprintf("%s %s %d, %d", en.Weekday[t.Weekday()], en.Month[t.Month()], t.Day(), t.Year())
}

func (en *AmericanDateTimeFormatter) FullDate(t time.Time) string {
	// TODO: implement format
	return fmt.Sprintf("%s %d. %s %d", en.Weekday[t.Weekday()], t.Day(), en.Month[t.Month()], t.Year())
}

// 3:04 PM
func (en *AmericanDateTimeFormatter) ShortTime(t time.Time) string {
	return t.Format("3:04 PM")
}

// 3:04:05 PM
func (en *AmericanDateTimeFormatter) MediumTime(t time.Time) string {
	return t.Format("3:04:05 PM")
}

// 3:04:05 PM MST
func (en *AmericanDateTimeFormatter) LongTime(t time.Time) string {
	return t.Format("3:04:05 PM MST")
}

// 3:04:05 PM MST - same as long format
func (en *AmericanDateTimeFormatter) FullTime(t time.Time) string {
	return t.Format("3:04:05 PM MST")
}

type GermanDateTimeFormatter struct {
	Month   map[time.Month]string
	Weekday map[time.Weekday]string
}

func createGermanDateTimeFormatter() DateTimeFormatter {
	return &GermanDateTimeFormatter{
		Month: map[time.Month]string{
			time.January:   "Januar",
			time.February:  "Februar",
			time.March:     "März",
			time.April:     "April",
			time.May:       "Mai",
			time.June:      "Juni",
			time.July:      "Juli",
			time.August:    "August",
			time.September: "September",
			time.October:   "Oktober",
			time.November:  "November",
			time.December:  "Dezember",
		},
		Weekday: map[time.Weekday]string{
			time.Monday:    "Montag",
			time.Tuesday:   "Dienstag",
			time.Wednesday: "Mittwoch",
			time.Thursday:  "Donnerstag",
			time.Friday:    "Freitag",
			time.Saturday:  "Samstag",
			time.Sunday:    "Sonntag",
		},
	}
}

func (de *GermanDateTimeFormatter) ShortDate(t time.Time) string {
	return fmt.Sprintf("%d.%d.%d", t.Day(), t.Month(), t.Year())
}

func (de *GermanDateTimeFormatter) MediumDate(t time.Time) string {
	return fmt.Sprintf("%d. %s %d", t.Day(), de.Month[t.Month()], t.Year())
}

func (de *GermanDateTimeFormatter) LongDate(t time.Time) string {
	return fmt.Sprintf("%s %d. %s %d", de.Weekday[t.Weekday()], t.Day(), de.Month[t.Month()], t.Year())
}

func (de *GermanDateTimeFormatter) FullDate(t time.Time) string {
	// TODO: implement format
	return fmt.Sprintf("%s %d. %s %d", de.Weekday[t.Weekday()], t.Day(), de.Month[t.Month()], t.Year())
}

// 3:04 PM
func (en *GermanDateTimeFormatter) ShortTime(t time.Time) string {
	return t.Format("15:04")
}

// 3:04:05 PM
func (en *GermanDateTimeFormatter) MediumTime(t time.Time) string {
	return t.Format("15:04:05")
}

// 3:04:05 PM MST
func (en *GermanDateTimeFormatter) LongTime(t time.Time) string {
	return t.Format("15:04:05 MST")
}

// 3:04:05 PM MST - same as long format
func (en *GermanDateTimeFormatter) FullTime(t time.Time) string {
	return t.Format("15:04:05 MST")
}

type DateTimeExpr struct {
	// Key is the key for the data map when this
	// expression is being formatted.
	Key string `json:"key"`
	// The type: date or time
	Type string `json:"type"`
	// Format represents the DateFormat enum value
	Format DateTimeFormat `json:"format"`
}

type DateTimeFormat = string

const (
	Short  DateTimeFormat = "short"
	Medium DateTimeFormat = "medium"
	Long   DateTimeFormat = "long"
	Full   DateTimeFormat = "full"
	// skeleton represents ICU's datetime skeleton format
	// Skeleton DateFormat = "skeleton"
)

// parseDate attempts to parse the input at the given start position into a DateExpr
func (p *parser) parseDateTime(varName string, ctype string, nextChar rune, start, end int, ptr_input *[]rune) (Expression, int, error) {
	var result = DateTimeExpr{
		Key:  varName,
		Type: ctype,
	}

	format, _, cursor, err := readVar(start+1, end, ptr_input)
	if err != nil {
		return nil, cursor, errors.New("failed to parse date format")
	}

	switch string(format) {
	case Short:
		result.Format = Short
	case Medium:
		result.Format = Medium
	case Long:
		result.Format = Long
	case Full:
		result.Format = Full
	default:
		return nil, cursor, fmt.Errorf("InvalidDateFormat")
	}

	return &result, cursor, nil
}

func (f *formatter) formatDateTime(expr Expression, ptrOutput *bytes.Buffer, data map[string]any) error {
	if date, ok := expr.(*DateTimeExpr); ok {
		t, ok := data[date.Key].(time.Time)
		ctype := date.Type
		if !ok {
			return fmt.Errorf("InvalidArgType: want time.Time, got %T", t)
		}

		switch date.Format {
		case Short:
			if ctype == "date" {
				ptrOutput.WriteString(f.date.ShortDate(t))
			} else {
				ptrOutput.WriteString(f.date.ShortTime(t))
			}
		case Medium:
			if ctype == "date" {
				ptrOutput.WriteString(f.date.MediumDate(t))
			} else {
				ptrOutput.WriteString(f.date.MediumTime(t))
			}
		case Long:
			if ctype == "date" {
				ptrOutput.WriteString(f.date.LongDate(t))
			} else {
				ptrOutput.WriteString(f.date.LongTime(t))
			}
		case Full:
			if ctype == "date" {
				ptrOutput.WriteString(f.date.FullDate(t))
			} else {
				ptrOutput.WriteString(f.date.FullTime(t))
			}
		default:
			return fmt.Errorf("InvalidDateFormat")
		}
	} else {
		return fmt.Errorf("InvalidExprType: want DateTimeExpr, got %T", expr)
	}

	return nil
}

// Symbol represents a format symbol
type Symbol rune

const (
	Era        Symbol = 'G'
	Year       Symbol = 'y'
	ShortMonth Symbol = 'M'
	LongMonth  Symbol = 'L'
	DayOfMonth Symbol = 'd'
	DayOfWeek  Symbol = 'E'
	AmPmMarker Symbol = 'a'
	Hour112    Symbol = 'h'
	Hour023    Symbol = 'H'
	Hour011    Symbol = 'K'
	Hour124    Symbol = 'k'
	Minute     Symbol = 'm'
	Second     Symbol = 's'
	TimeZone   Symbol = 'Z'
)

// DateTimeSkeleton is an ICU datetime skeleton.
type DateTimeSkeleton struct {
}
