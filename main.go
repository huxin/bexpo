// extract email from burp logs
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/huxin/bexpo/utils"
	"github.com/jaytaylor/html2text"
)

func main() {

	burpLog := "/Users/huxin/code/burp/burp_jcyxy.logs"
	emailContext := make(map[string]string)

	file, err := os.Open(burpLog)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	cnt := 0
	htmlLines := []string{}
	inHTML := false

	for err == nil {
		line, err := utils.Readln(r)
		cnt++
		if strings.Contains(line, "<html") {
			inHTML = true
		}
		if inHTML {
			htmlLines = append(htmlLines, line)
		}

		if strings.Contains(line, "</html>") {
			inHTML = false
			html := strings.Join(htmlLines, "\n")
			text, err := html2text.FromString(html)
			if err != nil {
				fmt.Println("Error:", err)
				text = html
			}
			c := utils.FindEmailContext(text, 100)
			for email, context := range c {
				if emailContext[email] == "" {
					emailContext[email] = context
					fmt.Println(email, context)
				}
			}

			// reset
			htmlLines = []string{}
		}

		for _, email := range utils.FindEmails(line) {
			if _, ok := emailContext[email]; !ok {
				emailContext[email] = ""
			}
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Printf("Read %d lines, %d emails\n", cnt, len(emailContext))

	outf, err := os.Create("email_context.txt")
	if err != nil {
		log.Fatal(err)
	}
	for email, context := range emailContext {
		outf.Write([]byte(email + " ` " + context + "\n"))
	}
	outf.Close()
}
