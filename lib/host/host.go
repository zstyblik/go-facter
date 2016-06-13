package host

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"

	h "github.com/shirou/gopsutil/host"
)

type Facter interface {
	Add(string, interface{})
}

// capitalize the first letter of given string
func capitalize(label string) string {
	firstLetter := strings.SplitN(label, "", 2)[0]
	return fmt.Sprintf("%v%v", strings.ToUpper(firstLetter),
		strings.TrimPrefix(label, firstLetter))
}

func guessArch(HWModel string) string {
	var arch string
	switch HWModel {
	case "x86_64":
		arch = "amd64"
		break
	default:
		arch = "unknown"
		break
	}
	return arch
}

// int8ToString converts [65]int8 in syscall.Utsname to string
func int8ToString(bs [65]int8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		if v < 0 {
			b[i] = byte(256 + int(v))
		} else {
			b[i] = byte(v)
		}
	}
	return strings.TrimRight(string(b), "\x00")
}

func GetHostFacts(f Facter) error {
	hostInfo, err := h.Info()
	if err != nil {
		return err
	}

	f.Add("fqdn", hostInfo.Hostname)
	splitted := strings.SplitN(hostInfo.Hostname, ".", 2)
	var hostname *string
	if len(splitted) > 1 {
		hostname = &splitted[0]
		f.Add("domain", splitted[1])
	} else {
		hostname = &hostInfo.Hostname
	}
	f.Add("hostname", *hostname)

	var is_virtual bool
	if hostInfo.VirtualizationRole == "host" {
		is_virtual = false
	} else {
		is_virtual = true
	}
	f.Add("is_virtual", is_virtual)

	f.Add("kernel", capitalize(hostInfo.OS))
	f.Add("operatingsystemrelease", hostInfo.PlatformVersion)
	f.Add("operatingsystem", capitalize(hostInfo.Platform))
	f.Add("osfamily", capitalize(hostInfo.PlatformFamily))
	f.Add("uptime_seconds", hostInfo.Uptime)
	f.Add("uptime_minutes", hostInfo.Uptime/60)
	f.Add("uptime_hours", hostInfo.Uptime/60/60)
	f.Add("uptime_days", hostInfo.Uptime/60/60/24)
	f.Add("uptime", fmt.Sprintf("%d days", hostInfo.Uptime/60/60/24))
	f.Add("virtual", hostInfo.VirtualizationSystem)

	envPath := os.Getenv("PATH")
	if envPath != "" {
		f.Add("path", envPath)
	}

	user, err := user.Current()
	if err == nil {
		f.Add("id", user.Username)
	} else {
		panic(err)
	}

	var uname syscall.Utsname
	err = syscall.Uname(&uname)
	if err == nil {
		kernelRelease := int8ToString(uname.Release)
		kernelVersion := strings.Split(kernelRelease, "-")[0]
		kvSplitted := strings.Split(kernelVersion, ".")
		f.Add("kernelrelease", kernelRelease)
		f.Add("kernelversion", kernelVersion)
		f.Add("kernelmajversion", strings.Join(kvSplitted[0:2], "."))

		hardwareModel := int8ToString(uname.Machine)
		f.Add("hardwaremodel", hardwareModel)
		f.Add("architecture", guessArch(hardwareModel))
	}

	return nil
}
