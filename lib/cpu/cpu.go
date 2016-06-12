package cpu

import (
	"fmt"
	"strconv"

	c "github.com/shirou/gopsutil/cpu"
)

type Facter interface {
	Add(string, interface{})
}

func GetCPUFacts(f Facter) error {
	var physCount uint64
	totalCount, err := c.Counts(true)
	if err != nil {
		return err
	}
	f.Add("processorcount", totalCount)

	CPUs, err := c.Info()
	if err != nil {
		return err
	}
	for _, v := range CPUs {
		physID, err := strconv.ParseUint(v.PhysicalID, 10, 32)
		if err == nil {
			physCount += physID
		}
		f.Add(fmt.Sprintf("processor%v", v.CPU), v.ModelName)
	}
	f.Add("physicalprocessorcount", physCount+1)
	return nil
}
