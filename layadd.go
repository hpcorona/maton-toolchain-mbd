package main

import (
	"fmt"
	"os"
	"flag"
  "strconv"
  "strings"
  "io/ioutil"
)

func toStr(v float32) string {
  vk := fmt.Sprintf("%f", v)
  vs := strings.TrimRight(vk, "0")
  return strings.TrimRight(vs, ".")
}

func main() {
	flag.Parse()

	if flag.NArg() < 5 {
		fmt.Printf("Usage:\n\tlayadd <file.pos> <dev_origin> <dev_dest> <scale> <output.pos>\n")
		os.Exit(1)
	}

  input := flag.Arg(0)
  devo := flag.Arg(1)
  devd := flag.Arg(2)
  scale, _ := strconv.Atof32(flag.Arg(3))
  output := flag.Arg(4)

  fmt.Printf("Adding layout '%s' based on '%s' * %f from file '%s' into a new file '%s'\n", devd, devo, scale, input, output)

  entries := parsePos(input)
  for i := 0; i < len(entries); i++ {
    entry := entries[i]

    devi := entry.Devices[devo]
    if devi.Count == 1 {
      devk := &Device { Name: devd, Count: 1, V1: devi.V1 * scale }
      entry.Devices[devd] = devk
    } else {
      devk := &Device { Name: devd, Count: 2, V1: devi.V1 * scale, V2: devi.V2 * scale }
      entry.Devices[devd] = devk
    }
  }

  data := ""
  for i := 0; i < len(entries); i++ {
    entry := entries[i]

    data += entry.Name + ":\n"
    for k, v := range entry.Devices {
      if v.Count == 1 {
        data += fmt.Sprintf("\t%s:\t\t%s\n", k, toStr(v.V1))
      } else {
        data += fmt.Sprintf("\t%s:\t\t%s,%s\n", k, toStr(v.V1), toStr(v.V2))
      }
    }
    data += "\n"
  }

  ioutil.WriteFile(output, []byte(data), 444)
  fmt.Printf("Done.")
}
