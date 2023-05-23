package main

import (
	"flag"

	"github.com/KittenConnect/go-facter/lib/facter"
	"github.com/KittenConnect/go-facter/lib/formatter"
)

func main() {
	conf := facter.Config{}
	ptFormat := flag.Bool("plaintext", false,
		"Emit facts as key => value pairs")
	kvFormat := flag.Bool("keyvalue", false,
		"Emit facts as key:value pairs")
	jsonFormat := flag.Bool("json", false,
		"Emit facts as a JSON")
	yamlFormat := flag.Bool("yaml", false,
		"Emit facts as a YAML")
	flag.Parse()

	if *ptFormat == true {
		conf.Formatter = formatter.NewFormatter()
	} else if *kvFormat == true {
		conf.Formatter = formatter.NewKeyValueFormatter()
	} else if *jsonFormat == true {
		conf.Formatter = formatter.NewJSONFormatter()
	} else if *yamlFormat == true {
		conf.Formatter = formatter.NewYAMLFormatter()
	} else {
		conf.Formatter = formatter.NewFormatter()
	}

	facter := facter.New(&conf)
	// _ = cpu.GetCPUFacts(facter)
	// _ = disk.GetDiskFacts(facter)
	// _ = host.GetHostFacts(facter)
	// _ = mem.GetMemoryFacts(facter)
	// _ = net.GetNetFacts(facter)
	facter.Print()
}
