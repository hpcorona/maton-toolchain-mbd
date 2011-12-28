package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
)

func parsePos(input string) {
	fmt.Printf("Processing: %s\n", input)
	allText, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Printf("Error: %s\n", err.String())
		os.Exit(1)
	}
	
	curr := ""
	lines := strings.Split(string(allText), "\n")
	for i := 0; i < len(lines); i++ {
		text := strings.Trim(lines[i], " \t")
		if text == "" {
			continue
		}
		
		idx := strings.Index(text, ":")
		if idx < 0 {
			continue
		}
		
		if lines[i][0] == ' ' || lines[i][0] == '\t' {
			if curr == "" {
				continue
			}
			
			path := curr + "/@" + text[0:idx]
			val := strings.Trim(text[idx+1:], " \t")
			idxComma := strings.Index(val, ",")
			if idxComma < 0 {
				values[path] = val
			} else {
				valX := strings.Trim(val[0:idxComma], " \t")
				valY := strings.Trim(val[idxComma+1:], " \t")
				values[curr + "/" + text[0:idx] + "/@x"] = valX
				values[curr + "/" + text[0:idx] + "/@y"] = valY
			}
		} else {
			curr = text[0:idx]
		}
	}
}
