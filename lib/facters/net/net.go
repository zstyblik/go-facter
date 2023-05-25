package net

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/KittenConnect/go-facter/lib/facter"
	"github.com/KittenConnect/go-facter/lib/facters/common"
	n "github.com/shirou/gopsutil/net"
)

var (
	pluginName = "net"
	reIPv4     = regexp.MustCompile(`^[0-9]+\.`)
)

func init() {
	facter.RegisterSafe(pluginName, []string{"interfaces", "macaddress_", "ipaddress_", "ipaddress6_", "netmask_", "mtu_", "ip_forward_", "ip6_forward_"}, GetNetFacts)
}

// GetNetFacts gathers network related facts
func GetNetFacts(f facter.IFacter) error {
	netIfaces, err := n.Interfaces()
	if err != nil {
		return fmt.Errorf("failed to retrieve network interfaces: %w", err)
	}

	var ifaces []string
	for _, v := range netIfaces {
		ifName := strings.ToLower(v.Name)
		ifaces = append(ifaces, ifName)
		f.Add(fmt.Sprintf("macaddress_%v", ifName), v.HardwareAddr)
		f.Add(fmt.Sprintf("mtu_%v", ifName), v.MTU)

		addr4idx, addr6idx := 0, 0
		for _, ipAddr := range v.Addrs {
			var labelIPAddr, labelNetmask string
			isIPv4 := reIPv4.MatchString(ipAddr.Addr)
			isFirstFound := false

			var netmaskBits uint64

			if isIPv4 {
				labelIPAddr = fmt.Sprintf("ipaddress_%v", ifName)
				labelNetmask = fmt.Sprintf("netmask_%v", ifName)

				if addr4idx == 0 {
					isFirstFound = true
				}
			} else {
				labelIPAddr = fmt.Sprintf("ipaddress6_%v", ifName)
				if addr6idx == 0 {
					isFirstFound = true
				}
			}

			splitted := strings.Split(ipAddr.Addr, "/")
			if len(splitted) > 0 && isFirstFound {
				f.Add(labelIPAddr, splitted[0])
			}

			if isIPv4 {
				if len(splitted) > 1 {
					netmaskBits, err = strconv.ParseUint(splitted[1], 10, 32)
					if err == nil {
						netmaskStr, err := common.ConvertNetmask(uint8(netmaskBits))
						if err == nil {
							if isFirstFound {
								f.Add(labelNetmask, netmaskStr)
							} else {
								f.Add(fmt.Sprintf("%s_%d", labelNetmask, addr4idx), netmaskStr)
							}
						}
					}
				}
				if !isFirstFound {
					f.Add(fmt.Sprintf("%s_%d", labelIPAddr, addr4idx), splitted[0])
				}
				addr4idx++
			} else {
				if !isFirstFound {
					f.Add(fmt.Sprintf("%s_%d", labelIPAddr, addr6idx), splitted[0])
				}
				addr6idx++
			}
		}

		// Check if IPv4 forwarding is enabled
		ipv4Forwarding, err := readProcSysNet(fmt.Sprintf("ipv4/conf/%s/forwarding", ifName))
		if err != nil {
			continue
			// return fmt.Errorf("failed to retrieve IPv4 forwarding status for interface %s: %w", ifName, err)
		}
		f.Add(fmt.Sprintf("ip_forward_%s", ifName), ipv4Forwarding)

		// Check if IPv6 forwarding is enabled
		ipv6Forwarding, err := readProcSysNet(fmt.Sprintf("ipv6/conf/%s/forwarding", ifName))
		if err != nil {
			continue
			// return fmt.Errorf("failed to retrieve IPv6 forwarding status for interface %s: %w", ifName, err)
		}
		f.Add(fmt.Sprintf("ip6_forward_%s", ifName), ipv6Forwarding)
	}

	if len(ifaces) > 0 {
		sort.Strings(ifaces)
		f.Add("interfaces", strings.Join(ifaces, ","))
	}

	return nil
}

// readProcSysNet reads the content of a file located in /proc/sys
func readProcSysNet(path string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/sys/net/%s", common.GetHostProc(), path))
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}
