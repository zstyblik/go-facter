package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/KittenConnect/go-facter/lib/facter"
	"github.com/KittenConnect/go-facter/lib/formatter"

	_ "github.com/KittenConnect/go-facter/lib/facters/routehOS"
)

func main() {
	conf := facter.Config{}
	// ptFormat := flag.Bool("plaintext", false,
	// 	"Emit facts as key => value pairs")
	kvFormat := flag.Bool("keyvalue", false,
		"Emit facts as key:value pairs")
	jsonFormat := flag.Bool("json", false,
		"Emit facts as a JSON")
	yamlFormat := flag.Bool("yaml", false,
		"Emit facts as a YAML")
	flag.Parse()

	// if *ptFormat == true {
	// 	conf.Formatter = formatter.NewFormatter()
	// } else
	if *kvFormat {
		conf.Formatter = formatter.NewKeyValueFormatter()
	} else if *jsonFormat {
		conf.Formatter = formatter.NewJSONFormatter()
	} else if *yamlFormat {
		conf.Formatter = formatter.NewYAMLFormatter()
	} else {
		conf.Formatter = formatter.NewFormatter()
	}

	facter := facter.New(&conf)

	args := flag.Args()

	if len(args) >= 1 {
		out := make(map[string]interface{})
		for _, query := range args {
			value, ok := facter.Get(query)
			if !ok {
				fmt.Fprintf(os.Stderr, "%s not found\n", query)
				continue
			}
			out[query] = value
		}
		conf.Formatter.Print(out)
	} else {
		facter.Print()
	}
}
