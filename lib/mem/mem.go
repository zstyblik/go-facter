package mem

import (
	"fmt"

	m "github.com/shirou/gopsutil/mem"
	"github.com/zstyblik/go-facter/lib/common"
)

type Facter interface {
	Add(string, interface{})
}

func GetMemoryFacts(f Facter) error {
	hostVMem, err := m.VirtualMemory()
	if err != nil {
		return err
	}
	memFree, unit, err := common.ConvertBytes(hostVMem.Free, "B")
	if err != nil {
		return err
	}
	f.Add("memoryfree", fmt.Sprintf("%.2f %v", memFree, unit))

	memTotal, unit, err := common.ConvertBytes(hostVMem.Total, "B")
	if err != nil {
		return err
	}
	f.Add("memorysize", fmt.Sprintf("%.2f %v", memTotal, unit))

	hostSwapMem, err := m.SwapMemory()
	if err != nil {
		return err
	}
	swapFree, unit, err := common.ConvertBytes(hostSwapMem.Free, "B")
	if err != nil {
		return err
	}
	f.Add("swapfree", fmt.Sprintf("%.2f %v", swapFree, unit))

	swapTotal, unit, err := common.ConvertBytes(hostSwapMem.Total, "B")
	if err != nil {
		return err
	}
	f.Add("swapsize", fmt.Sprintf("%.2f %v", swapTotal, unit))

	return nil
}
