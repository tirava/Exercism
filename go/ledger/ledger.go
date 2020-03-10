// Package ledger implements.
package ledger

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// Entry.
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

type records struct {
	i int
	s string
	e error
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	if (currency != "USD" && currency != "EUR") || (locale != "en-US" && locale != "nl-NL") {
		return "", errors.New("invalid format")
	}

	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		} else if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}

		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var s string
	if locale == "nl-NL" {
		s = "Datum      | Omschrijving              | Verandering\n"
	} else {
		s = "Date       | Description               | Change\n"
	}

	// Parallelism, always a great idea
	co := make(chan records)

	for i, et := range entriesCopy {

		go func(i int, entry Entry) {
			d1, d2, d3, d4, d5 := entry.Date[0:4], entry.Date[4], entry.Date[5:7], entry.Date[7], entry.Date[8:10]
			if len(entry.Date) != 10 || d2 != '-' || d4 != '-' {
				co <- records{e: errors.New("invalid date")}
			}

			//de := entry.Description
			var de string
			if len(entry.Description) > 25 {
				de = entry.Description[:22] + "..."
			} else {
				de = entry.Description + strings.Repeat(" ", 25-len(entry.Description))
			}

			var d string
			if locale == "nl-NL" {
				d = d5 + "-" + d3 + "-" + d1
			} else if locale == "en-US" {
				d = d3 + "/" + d5 + "/" + d1
			}
			negative := false
			cents := entry.Change
			if cents < 0 {
				cents = cents * -1
				negative = true
			}
			var a string
			if locale == "nl-NL" {
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- records{e: errors.New("")}
				}
				a += " "
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + "."
				}
				a = a[:len(a)-1]
				a += ","
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += "-"
				} else {
					a += " "
				}
			} else if locale == "en-US" {
				if negative {
					a += "("
				}
				if currency == "EUR" {
					a += "€"
				} else if currency == "USD" {
					a += "$"
				} else {
					co <- records{e: errors.New("")}
				}
				centsStr := strconv.Itoa(cents)
				switch len(centsStr) {
				case 1:
					centsStr = "00" + centsStr
				case 2:
					centsStr = "0" + centsStr
				}
				rest := centsStr[:len(centsStr)-2]
				var parts []string
				for len(rest) > 3 {
					parts = append(parts, rest[len(rest)-3:])
					rest = rest[:len(rest)-3]
				}
				if len(rest) > 0 {
					parts = append(parts, rest)
				}
				for i := len(parts) - 1; i >= 0; i-- {
					a += parts[i] + ","
				}
				a = a[:len(a)-1]
				a += "."
				a += centsStr[len(centsStr)-2:]
				if negative {
					a += ")"
				} else {
					a += " "
				}
			} else {
				co <- records{e: errors.New("")}
			}
			var al int
			for range a {
				al++
			}
			co <- records{i: i, s: d + strings.Repeat(" ", 10-len(d)) + " | " + de + " | " +
				strings.Repeat(" ", 13-al) + a + "\n"}
		}(i, et)
	}

	ss := make([]string, len(entriesCopy))

	for range entriesCopy {
		v := <-co
		if v.e != nil {
			return "", v.e
		}

		ss[v.i] = v.s
	}

	s += strings.Join(ss, "")

	return s, nil
}
