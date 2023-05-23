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
	// debug("RADVD => %s\n", out)

	f.Add("radvd_version", reRADVDVersion.FindStringSubmatch(out)[1])
	f.Add("radvd_conf_file", reRADVDConfig.FindStringSubmatch(out)[1])
	f.Add("radvd_pid_file", reRADVDPidFile.FindStringSubmatch(out)[1])
	f.Add("radvd_log_file", reRADVDLogFile.FindStringSubmatch(out)[1])
	f.Add("radvd_log_facility", reRADVDSyslogFc.FindStringSubmatch(out)[1])

	return nil
}
