package main

import (
	"fmt"
	"strings"
	"time"
)

var testLines = [5]string{
	"HAUS AM LÜTZOWPLATZ",
	"Amy Ball: ‘Women’",
	"Performance: Friday, Aug. 17; 7–10pm",
	"Installation: Aug. 18–19, 2018",
	"Lützowplatz 9, 10785 Berlin, click here for map",
}

func splitText(s, sep string) (string, string) {
	x := strings.Split(s, sep)

	switch len(x) {
	case 0:
		return "", ""
	case 1:
		return x[0], ""
	default:
		return x[0], x[1]
	}
}

type Parser interface {
	Exec(line string)
}

type Parsers struct {
	parsers []Parser
}

func parse(line string, p Parser) {
	p.Exec(line)
}

type Date struct {
	date  time.Time
	begin time.Time
	end   time.Time
}

func (d Date) Get() time.Time {
	return d.date
}

func (d Date) Exec(line string) {
	_, text := splitText(line, ":")

	if text == "" {
		return
	}

	datePart, timePart := splitText(text, ";")

	if datePart != "" {
		date, err := time.Parse("Monday, Jan. 2", strings.Trim(datePart, " "))

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("DATE: %v\n", date)
			d.date = date
		}
	}

	if timePart != "" {
		timePart = strings.Trim(timePart, " ")
		beginPart, endPart := splitText(timePart, "-")

		begin, err := time.Parse("3", beginPart)
		end, err := time.Parse("3pm", endPart)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("BEGIN: %v\n", begin)
			fmt.Printf("END: %v\n", end)
			d.begin = begin
			d.end = end
		}
	}
}

func main() {
	date := Date{}

	parsers := []Parser{
		date,
	}

	for _, line := range testLines {
		for _, parser := range parsers {
			parse(line, parser)
		}
	}
}

// type EventTypeParser struct {
// 	eventType string
// }

// func (e EventTypeParser) Get() string {
// 	return e.eventType
// }

// func (e EventTypeParser) Exec(text string) {
// 	types := [8]string{
// 		"opening",
// 		"panel discussion",
// 		"event",
// 		"performance",
// 		"installation",
// 		"exhibition",
// 		"film-screening",
// 		"talk",
// 	}

// 	var eventType, tempDate string
// 	var date time.Time
// 	var err error

// 	for _, knownType := range types {
// 		if strings.Contains(strings.ToLower(text), knownType) {
// 			ok = splitText(text, ":")
// 			eventType = knownType

// 			date, err = parseDate(tempDate)

// 			if err != nil {
// 				fmt.Sprint(err)
// 			}
// 		}
// 	}
// }

// type LineParser func() string

// func (l LineParser) Exec(line string) (string, ok bool) {
// 	date, err := time.Parse("Saturday, Aug. 18; 3–9pm", strings.Trim(text, " "))

// 	if err != nil {
// 		fmt.Println(err)
// 		return time.Time{}, true
// 	}

// 	return date, true
// }
