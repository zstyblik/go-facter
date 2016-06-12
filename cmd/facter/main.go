package main

import (
	"github.com/zstyblik/go-facter/lib/cpu"
	"github.com/zstyblik/go-facter/lib/facter"
	"github.com/zstyblik/go-facter/lib/host"
	"github.com/zstyblik/go-facter/lib/mem"
	"github.com/zstyblik/go-facter/lib/net"
)

func main() {
	facter := facter.New(nil)
	_ = cpu.GetCPUFacts(facter)
	_ = host.GetHostFacts(facter)
	_ = mem.GetMemoryFacts(facter)
	_ = net.GetNetFacts(facter)
	facter.Print()
}
