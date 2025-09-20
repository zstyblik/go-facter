package mem

import (
	"fmt"

	m "github.com/shirou/gopsutil/v4/mem"
	"github.com/zstyblik/go-facter/lib/common"
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// addMemoryUnits will convert a memory fact into "fact_mb" and "fact_bytes"
func addMemoryUnits(f Facter, label string, memory uint64) error {
	units := map[string]string{
		"MB": "mb",
		"B":  "bytes",
	}

	for unit, unitLabel := range units {
		factLabel := fmt.Sprintf("%s_%s", label, unitLabel)
		convertedMemory, _, err := common.ConvertBytesTo(memory, unit)
		if err != nil {
			return err
		}

		f.Add(factLabel, fmt.Sprintf("%.2f", convertedMemory))
	}

	return nil
}

// GetMemoryFacts gathers facts related to system memory
func GetMemoryFacts(f Facter) error {
	// Get the virtual memory from gopsutil
	hostVMem, err := m.VirtualMemory()
	if err != nil {
		return err
	}

	// Add a memoryfree fact
	memFree, unit, err := common.ConvertBytes(hostVMem.Free)
	if err != nil {
		return err
	}
	f.Add("memoryfree", fmt.Sprintf("%.2f %v", memFree, unit))

	// Add memoryfree_mb and memoryfree_bytes facts
	err = addMemoryUnits(f, "memoryfree", hostVMem.Free)
	if err != nil {
		return err
	}

	// Add a memorytotal fact
	memTotal, unit, err := common.ConvertBytes(hostVMem.Total)
	if err != nil {
		return err
	}
	f.Add("memorysize", fmt.Sprintf("%.2f %v", memTotal, unit))

	// add memorytotal_mb and memorytotal_bytes facts
	err = addMemoryUnits(f, "memorysize", hostVMem.Total)
	if err != nil {
		return err
	}

	// Get the swap information from gopsutil
	hostSwapMem, err := m.SwapMemory()
	if err != nil {
		return err
	}

	// Add a swapfree fact
	swapFree, unit, err := common.ConvertBytes(hostSwapMem.Free)
	if err != nil {
		return err
	}
	f.Add("swapfree", fmt.Sprintf("%.2f %v", swapFree, unit))

	// Add swapfree_mb and swapfree_bytes facts
	err = addMemoryUnits(f, "swapfree", hostSwapMem.Free)
	if err != nil {
		return err
	}

	// Add a swapsize fact
	swapTotal, unit, err := common.ConvertBytes(hostSwapMem.Total)
	if err != nil {
		return err
	}
	f.Add("swapsize", fmt.Sprintf("%.2f %v", swapTotal, unit))

	// Add swapsize_mb and swapsize_bytes facts
	err = addMemoryUnits(f, "swapsize", hostSwapMem.Total)
	if err != nil {
		return err
	}

	return nil
}
