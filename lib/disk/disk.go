package disk

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strings"

	d "github.com/shirou/gopsutil/disk"
	"github.com/zstyblik/go-facter/lib/common"
)

var (
	reDevBlacklist = regexp.MustCompile("^(dm-[0-9]+|loop[0-9]+)$")
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// GetBlockDevices returns list of block devices
func getBlockDevices(all bool) ([]string, error) {
	blockDevs := []string{}
	targetDir := fmt.Sprintf("%v/block", common.GetHostSys())
	contents, err := ioutil.ReadDir(targetDir)
	if err != nil {
		return blockDevs, err
	}
	for _, v := range contents {
		if all == false {
			if reDevBlacklist.MatchString(v.Name()) {
				continue
			}
		}
		blockDevs = append(blockDevs, v.Name())
	}
	return blockDevs, nil
}

// GetDiskFacts gathers facts related to HDDs
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

	blockDevs, err := getBlockDevices(false)
	if err == nil {
		sort.Strings(blockDevs)
		f.Add("blockdevices", strings.Join(blockDevs, ","))
	}

	return nil
}
