package cpu

import (
	"fmt"
	"strconv"

	c "github.com/shirou/gopsutil/cpu"
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// GetCPUFacts gathers facts related to CPU
func GetCPUFacts(f Facter) error {
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
