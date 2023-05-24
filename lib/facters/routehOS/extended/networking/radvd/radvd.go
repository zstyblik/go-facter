package routehOS

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var (
	PluginName = "radvd"

	reRADVDVersion  = regexp.MustCompile(`Version: ([\d\.]+)`)
	reRADVDConfig   = regexp.MustCompile(`default[\s]+config[\s]+file[\s]+"([\/-_.\w]+)"`)
	reRADVDPidFile  = regexp.MustCompile(`default[\s]+pidfile[\s]+"([\/-_.\w]+)"`)
	reRADVDLogFile  = regexp.MustCompile(`default[\s]+logfile[\s]+"([\/-_.\w]+)"`)
	reRADVDSyslogFc = regexp.MustCompile(`default[\s]+syslog[\s]+facility[\s]+(\d+)`)
)

func init() {
	facter.RegisterSafe(PluginName, []string{"radvd."}, GetRADVDFacts)
}

func GetRADVDFacts(f facter.IFacter) error {
	cmd := exec.Command("radvd", "--version")

	// Run the command and retrieve the output
	output, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		return err
	}

	out := string(output)
	// debug("RADVD => %s\n", out)

	result := reRADVDVersion.FindStringSubmatch(out)

	if len(result) < 1 {
		return fmt.Errorf("cannot find [ %s ] in RADvd version output", reRADVDVersion.String())
	}

	f.Add("radvd", map[string]interface{}{
		"version":      result[1],
		"conf_file":    reRADVDConfig.FindStringSubmatch(out)[1],
		"pid_file":     reRADVDPidFile.FindStringSubmatch(out)[1],
		"log_file":     reRADVDLogFile.FindStringSubmatch(out)[1],
		"log_facility": reRADVDSyslogFc.FindStringSubmatch(out)[1],
	})

	return nil
}
