package main

import (
	"flag"

	"github.com/zstyblik/go-facter/lib/cpu"
	"github.com/zstyblik/go-facter/lib/disk"
	"github.com/zstyblik/go-facter/lib/facter"
	"github.com/zstyblik/go-facter/lib/formatter"
	"github.com/zstyblik/go-facter/lib/host"
	"github.com/zstyblik/go-facter/lib/mem"
	"github.com/zstyblik/go-facter/lib/net"
)

func main() {
	conf := facter.Config{}
	plainText := flag.Bool("plaintext", false,
		"Emit facts as key => value pairs")
	keyValue := flag.Bool("keyvalue", false,
		"Emit facts as key:value pairs")
	flag.Parse()

	if *plainText == true {
		conf.Formatter = formatter.NewFormatter()
	} else if *keyValue == true {
		conf.Formatter = formatter.NewKeyValueFormatter()
	} else {
		conf.Formatter = formatter.NewFormatter()
	}

	facter := facter.New(&conf)
	_ = cpu.GetCPUFacts(facter)
	_ = disk.GetDiskFacts(facter)
	_ = host.GetHostFacts(facter)
	_ = mem.GetMemoryFacts(facter)
	_ = net.GetNetFacts(facter)
	facter.Print()
}
