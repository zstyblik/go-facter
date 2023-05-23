package routehOS

import (
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var reRADVDVersion = regexp.MustCompile(`Version: ([\d\.]+)`)

var reRADVDConfig = regexp.MustCompile(`\s+default\s+config\s+file\s+"(\w+)"`)
var reRADVDPidFile = regexp.MustCompile(`\s+default\s+pidfile\s+"(\w+)"`)
var reRADVDLogFile = regexp.MustCompile(`\s+default\s+logfile\s+"(\w+)"`)
var reRADVDSyslogFc = regexp.MustCompile(`\s+default\s+syslog\s+facility\s+(\d+)`)

func GetRADVDFacts(f facter.IFacter) error {
	cmd := exec.Command("radvd", "--version")

	// Run the command and retrieve the output
	output, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		return err
	}

	debug("RADVD => %v", reRADVDConfig.FindStringSubmatch(string(output)))
	debug("RADVD => %v", reRADVDPidFile.FindStringSubmatch(string(output)))
	debug("RADVD => %v", reRADVDLogFile.FindStringSubmatch(string(output)))
	debug("RADVD => %v", reRADVDSyslogFc.FindStringSubmatch(string(output)))

	f.Add("radvd_version", reRADVDVersion.FindStringSubmatch(string(output))[1])

	return nil
}
