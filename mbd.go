package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
	"flag"
	"bytes"
	"encoding/binary"
	"io"
)

// Formato:
// MBD
// uint32[4] - Version y Adicional
// uint32 - sizes length
// uint32 - dictionary length
// uint32 - data length
// sizes {
//   uint32 - sizes_qty
//   type[sizes_qty] {
//     string - name
//     uint32 - qty
//   }
// }
// dictionary {
//   uint32 - dictionary_qty
//   type[dictionary_qty] {
//     string - name
//     uint32 - address (on data)
//   }
// }
// data {
//   * {
//     string - value (from dictionary)
//   }
// }

var counts map[string]int = make(map[string]int)
var values map[string]string = make(map[string]string)
var order binary.ByteOrder = binary.LittleEndian

func purge() {
	for k, _ := range counts {
		counts[k] = 0, false
	}
	
	for k, _ := range values {
		values[k] = "", false
	}
}

func writeInt(w io.Writer, i int) int {
	t := uint32(i)
	binary.Write(w, order, t)
	
	return 4
}

func writeString(w io.Writer, v string) int {
	bytes := []byte(v)
	
	writeInt(w, len(bytes))
	w.Write(bytes)

	return len(bytes) + 4
}

func writeBmd(file string) {
	fmt.Printf("Wring BMD: %s\n", file)
	
	sizes := bytes.NewBuffer([]byte(""))
	dictionary := bytes.NewBuffer([]byte(""))
	data := bytes.NewBuffer([]byte(""))
	
	writeInt(sizes, len(counts))
	for k, v := range counts {
		writeString(sizes, k)
		writeInt(sizes, v)
	}
	
	var address int = 0
	writeInt(dictionary, len(values))
	for k, v := range values {
		writeString(dictionary, k)
		writeInt(dictionary, address)
		
		w := writeString(data, v)
		
		address += w
	}
	
	all := bytes.NewBuffer([]byte("MBD"))
	writeInt(all, 1) // Version Mayor 1
	writeInt(all, 0) // Version Menor 0
	writeInt(all, 0) // Adicionales
	writeInt(all, 0) // Adicionales
	writeInt(all, sizes.Len())
	writeInt(all, dictionary.Len())
	writeInt(all, data.Len())
	all.Write(sizes.Bytes())
	all.Write(dictionary.Bytes())
	all.Write(data.Bytes())
	
	ioutil.WriteFile(file, all.Bytes(), 444)
	
	purge()
}

func main() {
	flag.Parse()
	
	if flag.NArg() == 0 {
		fmt.Printf("Usage:\n\tmdb <file.xml | file.lang>\n")
		os.Exit(1)
	}
	
	for i := 0; i < flag.NArg(); i++ {
		input := flag.Arg(i)
		output := input + ".mbd"
		
		idx := strings.LastIndex(input, ".")
		if idx >= 0 {
			output = input[0:idx] + ".mbd"
		}
		
		ext := input[idx+1:]
		
		if ext == "xml" || ext == "fnt" {
			parseXml(input)
		} else {
			parseLang(input)
		}
		writeBmd(output)
	}
}
