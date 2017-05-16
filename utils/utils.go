package utils

import (
	"bufio"
	"regexp"
	"strings"
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

// HTML2Text to text, find index of the email and do 100 bytes before and 100 bytes later
func HTML2Text(html string) string {
	//TODO: implement
	return html
}

// FindEmailContext finds context surrounding the email
func FindEmailContext(s string, size int) map[string]string {
	b := []byte(s)
	res := emailRegex.FindAllSubmatchIndex(b, -1)
	ret := make(map[string]string)

	for _, r := range res {
		pos := r[1]
		start := pos - size
		if start < 0 {
			start = 0
		}
		end := pos + 20
		if end > len(b) {
			end = len(b)
		}
		context := strings.Replace(string(b[start:end]), "\n", " ", -1)
		for _, e := range FindEmails(context) {
			ret[e] += context
		}
	}
	return ret
}

var phoneNumberRegex = regexp.MustCompile("0551")

func findPhoneNumberContext(s string, size int) (ret []string) {
	b := []byte(s)
	res := phoneNumberRegex.FindAllSubmatchIndex(b, -1)

	for _, r := range res {
		pos := r[1]
		start := pos - size
		if start < 0 {
			start = 0
		}
		end := pos + size
		if end > len(b) {
			end = len(b)
		}
		context := strings.Replace(string(b[start:end]), "\n", " ", -1)
		ret = append(ret, context)
	}

	return ret
}
