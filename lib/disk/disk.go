package disk

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"

	d "github.com/shirou/gopsutil/v4/disk"
	"github.com/zstyblik/go-facter/lib/common"
)

var (
	reDevBlacklist = regexp.MustCompile("^(dm-[0-9]+|loop[0-9]+)$")
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// getBlockDevices returns list of block devices
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

// getBlockDeviceModel returns model of block device as reported by Linux
// kernel.
func getBlockDeviceModel(blockDevice string) (string, error) {
	modelFilename := fmt.Sprintf("%s/block/%s/device/model",
		common.GetHostSys(), blockDevice)
	model, err := ioutil.ReadFile(modelFilename)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", bytes.TrimSuffix(model, []byte("\n"))), nil
}

// getBlockDeviceSize returns size of block device as reported by Linux kernel
// multiplied by 512.
func getBlockDeviceSize(blockDevice string) (int64, error) {
	sizeFilename := fmt.Sprintf("%s/block/%s/size", common.GetHostSys(),
		blockDevice)
	size, err := ioutil.ReadFile(sizeFilename)
	if err != nil {
		return 0, err
	}
	sizeInt, err := strconv.ParseInt(fmt.Sprintf("%s",
		bytes.TrimSuffix(size, []byte("\n"))), 10, 64)
	if err != nil {
		return 0, err
	}
	return sizeInt * 512, nil
}

// getBlockDeviceVendor returns vendor of block device as reported by Linux
// kernel.
func getBlockDeviceVendor(blockDevice string) (string, error) {
	vendorFilename := fmt.Sprintf("%s/block/%s/device/vendor",
		common.GetHostSys(), blockDevice)
	vendor, err := ioutil.ReadFile(vendorFilename)
	if err != nil {
		return "", err
	}
	vendor = bytes.TrimSuffix(vendor, []byte("\n"))
	vendor = bytes.TrimRight(vendor, " ")
	return fmt.Sprintf("%s", vendor), nil
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
	if err != nil {
		return err
	}

	sort.Strings(blockDevs)
	f.Add("blockdevices", strings.Join(blockDevs, ","))
	for _, blockDevice := range blockDevs {
		size, err := getBlockDeviceSize(blockDevice)
		if err == nil {
			f.Add(fmt.Sprintf("blockdevice_%s_size", blockDevice), size)
		}

		model, err := getBlockDeviceModel(blockDevice)
		if err == nil {
			f.Add(fmt.Sprintf("blockdevice_%s_model", blockDevice), model)
		}

		vendor, err := getBlockDeviceVendor(blockDevice)
		if err == nil {
			f.Add(fmt.Sprintf("blockdevice_%s_vendor", blockDevice), vendor)
		}
	}

	return nil
}
