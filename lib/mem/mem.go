package mem

import (
	"fmt"

	m "github.com/shirou/gopsutil/mem"
	"github.com/zstyblik/go-facter/lib/common"
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// GetMemoryFacts gathers facts related to system memory
func GetMemoryFacts(f Facter) error {
	hostVMem, err := m.VirtualMemory()
	if err != nil {
		return err
	}
	memFree, unit, err := common.ConvertBytes(hostVMem.Free)
	if err != nil {
		return err
	}
	f.Add("memoryfree", fmt.Sprintf("%.2f %v", memFree, unit))

	memFreeMB, unit, err := common.ConvertBytesTo(hostVMem.Free, "MB")
	if err != nil {
		return err
	}
	f.Add("memoryfree_mb", fmt.Sprintf("%.2f %v", memFreeMB, unit))

	memTotal, unit, err := common.ConvertBytes(hostVMem.Total)
	if err != nil {
		return err
	}
	f.Add("memorysize", fmt.Sprintf("%.2f %v", memTotal, unit))

	memTotalMB, unit, err := common.ConvertBytesTo(hostVMem.Total, "MB")
	if err != nil {
		return err
	}
	f.Add("memorysize_mb", fmt.Sprintf("%.2f %v", memTotalMB, unit))

	hostSwapMem, err := m.SwapMemory()
	if err != nil {
		return err
	}

	swapFree, unit, err := common.ConvertBytes(hostSwapMem.Free)
	if err != nil {
		return err
	}
	f.Add("swapfree", fmt.Sprintf("%.2f %v", swapFree, unit))

	swapFreeMB, unit, err := common.ConvertBytesTo(hostSwapMem.Free, "MB")
	if err != nil {
		return err
	}
	f.Add("swapfree_mb", fmt.Sprintf("%.2f %v", swapFreeMB, unit))

	swapTotal, unit, err := common.ConvertBytes(hostSwapMem.Total)
	if err != nil {
		return err
	}
	f.Add("swapsize", fmt.Sprintf("%.2f %v", swapTotal, unit))

	swapTotalMB, unit, err := common.ConvertBytesTo(hostSwapMem.Total, "MB")
	if err != nil {
		return err
	}
	f.Add("swapsize_mb", fmt.Sprintf("%.2f %v", swapTotalMB, unit))

	return nil
}
