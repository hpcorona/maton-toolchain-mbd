package main

import (
	"strconv"
	"xml"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInnerXml(t xml.StartElement, p *xml.Parser, xpath string) bool {
	newPath := xpath + "/" + t.Name.Local
	
	if xpath != "" {
		idx := 1
	
		idx = counts[newPath] + 1
		counts[newPath] = idx
	
		newPath = newPath + "[" + strconv.Itoa(idx) + "]"
	}
	
	for i := 0; i < len(t.Attr); i++ {
		attrName := newPath + "/@" + t.Attr[i].Name.Local
		attrValue := t.Attr[i].Value
		
		values[attrName] = attrValue
	}
	
	for {
		nt, err := p.Token()
		if err == os.EOF {
			return true
		}
		if err != nil {
			fmt.Printf("Error: %s\n", err.String())
			return false
		}
		
		if n, ok := nt.(xml.EndElement); ok {
			if n.Name.Local == t.Name.Local {
				return true
			}
		}
		
		if s, ok := nt.(xml.StartElement); ok {
			r := parseInnerXml(s, p, newPath)
			if r == false {
				return false
			}
		}
	}
	
	return true
}

func parseXml(file string) {
	fmt.Printf("Processing: %s\n", file)
	
	out, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("Error: %s\n", err.String())
		os.Exit(1)
	}
	
	p := xml.NewParser(strings.NewReader(string(out)))
	
	for {
		t, err := p.Token()
		if err == os.EOF {
			break
		}
		
		if err != nil {
			fmt.Printf("Error: %s\n", err.String())
			os.Exit(1)
		}
		
		var se xml.StartElement
		se, ok := t.(xml.StartElement)
		
		if ok == false {
			continue
		}
		
		if parseInnerXml(se, p, "") == false {
			os.Exit(1)
		} else {
			break
		}
	}
}
