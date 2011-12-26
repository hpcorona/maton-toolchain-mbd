package main

import (
	"fmt"
	"os"
  "strconv"
  "strings"
  "io/ioutil"
)

type Device struct {
  Name string
  Count int
  V1 float32
  V2 float32
}

type Entry struct {
  Name string
  Devices map[string]*Device
}

func parsePos(input string) []*Entry {
  entries := make([]*Entry, 0, 5000)
  var entry *Entry = nil

	fmt.Printf("Processing: %s\n", input)
	allText, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Printf("Error: %s\n", err.String())
		os.Exit(1)
	}

	curr := ""
	lines := strings.Split(string(allText), "\n")
	for i := 0; i < len(lines); i++ {
    fmt.Printf("--->%v\n", entry)
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

      fmt.Printf("\t%s\n", text[0:idx])
			val := strings.Trim(text[idx+1:], " \t")
			idxComma := strings.Index(val, ",")
			if idxComma < 0 {
        v, _ := strconv.Atof32(val)
        entry.Devices[text[0:idx]] = &Device { Name: text[0:idx], Count: 1, V1: v }
			} else {
				valX := strings.Trim(val[0:idxComma], " \t")
				valY := strings.Trim(val[idxComma+1:], " \t")

        valXf, _ := strconv.Atof32(valX)
        valYf, _ := strconv.Atof32(valY)
        entry.Devices[text[0:idx]] = &Device { Name: text[0:idx], Count: 2, V1: valXf, V2: valYf }
			}
		} else {
			curr = text[0:idx]
      entry = &Entry { Name: text[0:idx], Devices: make(map[string]*Device) }
      entries = append(entries, entry)
      fmt.Printf("Device: %s %v\n", curr, entry)
		}
	}

  return entries
}
