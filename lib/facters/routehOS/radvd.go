package routehOS

import (
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var reRADVDVersion = regexp.MustCompile(`Version: ([\d\.]+)`)

var reRADVDConfig = regexp.MustCompile(`default[\s]+config[\s]+file[\s]+"([\/-_.\w]+)"`)
var reRADVDPidFile = regexp.MustCompile(`default[\s]+pidfile[\s]+"([\/-_.\w]+)"`)
var reRADVDLogFile = regexp.MustCompile(`default[\s]+logfile[\s]+"([\/-_.\w]+)"`)
var reRADVDSyslogFc = regexp.MustCompile(`default[\s]+syslog[\s]+facility[\s]+(\d+)`)

func GetRADVDFacts(f facter.IFacter) error {
	cmd := exec.Command("radvd", "--version")

	// Run the command and retrieve the output
	output, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		return err
	}

	out := string(output)
	debug("RADVD => %s\n", out)

	debug("RADVDConfig => %v\n", reRADVDConfig.FindStringSubmatch(out))
	debug("RADVDPidFile => %v\n", reRADVDPidFile.FindStringSubmatch(out))
	debug("RADVDLogFile => %v\n", reRADVDLogFile.FindStringSubmatch(out))
	debug("RADVDSyslogFc => %v\n", reRADVDSyslogFc.FindStringSubmatch(out))

	f.Add("radvd_version", reRADVDVersion.FindStringSubmatch(string(output))[1])

	return nil
}
