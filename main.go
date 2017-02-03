// extract email from burp logs
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/huxin/bexpo/utils"
)

func main() {
	burpLog := "/Users/huxin/code/burp/burp_jcyxy.logs"
	uniq := make(map[string]bool)

	file, err := os.Open(burpLog)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)

	cnt := 0

	for err == nil {
		line, err := utils.Readln(r)
		cnt++
		for _, email := range utils.FindEmails(line) {
			if uniq[email] == false {
				uniq[email] = true
				fmt.Println(email)
			}
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	fmt.Printf("Read %d lines, %d emails\n", cnt, len(uniq))
}
