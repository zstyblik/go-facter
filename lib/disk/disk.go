package disk

import (
	"sort"
	"strings"

	d "github.com/shirou/gopsutil/disk"
)

type Facter interface {
	Add(string, interface{})
}

func GetDiskFacts(f Facter) error {
	partitions, err := d.Partitions(false)
	if err != nil {
		return err
	}

	// TODO - probably read from /proc/filesystems
	fstypes := make(map[string]struct{})
	for _, v := range partitions {
		fstypes[v.Fstype] = struct{}{}
	}
	fstypesStr := []string{}
	for k := range fstypes {
		fstypesStr = append(fstypesStr, k)
	}
	if len(fstypesStr) > 0 {
		sort.Strings(fstypesStr)
		f.Add("filesystems", strings.Join(fstypesStr, ","))
	}

	return nil
}
