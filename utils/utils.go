package utils

import (
	"bufio"
	"regexp"
)

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix = true
		err      error
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// regex match emails
var emailRegex = regexp.MustCompile("([A-Za-z0-9._%+-]+@[A-Za-z0-9._%+-]+\\.[A-Za-z0-9._%+-]+)")

// FindEmails returns a list of extracted email addresses
func FindEmails(s string) []string {

	res := emailRegex.FindAllStringSubmatch(s, -1)

	emails := make([]string, len(res))
	for i, r := range res {
		emails[i] = r[1]
	}

	return emails
}
