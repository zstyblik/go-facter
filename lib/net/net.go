package net

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	n "github.com/shirou/gopsutil/v4/net"
	"github.com/zstyblik/go-facter/lib/common"
)

var (
	reIPv4 = regexp.MustCompile("^[0-9]+\\.")
)

// Facter interface
type Facter interface {
	Add(string, interface{})
}

// GetNetFacts gathers network related facts
func GetNetFacts(f Facter) error {
	netIfaces, err := n.Interfaces()
	if err != nil {
		return err
	}

	var ifaces []string
	for _, v := range netIfaces {
		ifName := strings.ToLower(v.Name)
		ifaces = append(ifaces, ifName)
		if v.HardwareAddr != "" {
			f.Add(fmt.Sprintf("macaddress_%v", ifName), v.HardwareAddr)
		}
		f.Add(fmt.Sprintf("mtu_%v", ifName), v.MTU)
		addr4idx := (-1)
		addr6idx := (-1)
		for _, ipAddr := range v.Addrs {
			var labelIPAddr string
			var labelNetmask string
			if reIPv4.MatchString(ipAddr.Addr) {
				if addr4idx < 0 {
					labelIPAddr = fmt.Sprintf("ipaddress_%v", ifName)
					labelNetmask = fmt.Sprintf("netmask_%v", ifName)
				} else {
					labelIPAddr = fmt.Sprintf("ipaddress_%v_%d", ifName,
						addr4idx)
					labelNetmask = fmt.Sprintf("netmask_%v_%d", ifName,
						addr4idx)
				}
				addr4idx++
			} else {
				if addr6idx < 0 {
					labelIPAddr = fmt.Sprintf("ipaddress6_%v", ifName)
				} else {
					labelIPAddr = fmt.Sprintf("ipaddress6_%v_%d", ifName,
						addr6idx)
				}
				addr6idx++
			}
			splitted := strings.Split(ipAddr.Addr, "/")
			f.Add(labelIPAddr, splitted[0])
			if len(splitted) > 1 && reIPv4.MatchString(ipAddr.Addr) {
				netmaskBits, err := strconv.ParseUint(splitted[1], 10, 32)
				if err != nil {
					// TODO
					continue
				}
				netmaskStr, err := common.ConvertNetmask(uint8(netmaskBits))
				if err != nil {
					// TODO
					continue
				}
				f.Add(labelNetmask, netmaskStr)
			}
		}
	}
	if len(ifaces) > 0 {
		sort.Strings(ifaces)
		f.Add("interfaces", strings.Join(ifaces, ","))
	}

	return nil
}
