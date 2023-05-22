package cpu

import (
	"fmt"
	"strconv"

	"github.com/KittenConnect/go-facter/lib/facter"

	c "github.com/shirou/gopsutil/cpu"
)

var pluginName = "cpu"

func init() {
	err := facter.Register(pluginName, GetCPUFacts)
	if err != nil {
		fmt.Printf("Cannot register Facter %s : %v\n", pluginName, err)
	}
}

// GetCPUFacts gathers facts related to CPU
func GetCPUFacts(f facter.IFacter) error {
	totalCount, err := c.Counts(true)
	if err != nil {
		return err
	}
	f.Add("processorcount", totalCount)

	CPUs, err := c.Info()
	if err != nil {
		return err
	}
	physIDs := make(map[uint64]struct{})
	for _, v := range CPUs {
		physID, err := strconv.ParseUint(v.PhysicalID, 10, 32)
		if err == nil {
			physIDs[physID] = struct{}{}
		}
		f.Add(fmt.Sprintf("processor%v", v.CPU), v.ModelName)
	}
	f.Add("physicalprocessorcount", len(physIDs))
	return nil
}
