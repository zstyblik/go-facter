package routehOS

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var (
	PluginName = "bird"

	reBirdVersion = regexp.MustCompile(`BIRD version ([\d\.]+)`)
	reBirdConfig  = regexp.MustCompile(`[\s]+-c[\s]+<config-file>[\w\s]+instead[\w\s]+\n[\s]+([\/-_.\w]+)`)
)

func init() {
	facter.RegisterSafe(PluginName, []string{"bird."}, GetBirdFacts)
}

func GetBirdFacts(f facter.IFacter) error {
	cmd := exec.Command("bird", "--version")

	// Run the command and retrieve the output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	versionOut := string(output)

	result := reBirdVersion.FindStringSubmatch(versionOut)

	if len(result) < 1 {
		return fmt.Errorf("cannot find [ %s ] in BIRD version output", reBirdVersion.String())
	}

	cmd = exec.Command("bird", "--help")

	// Run the command and retrieve the output
	output, err = cmd.CombinedOutput()
	if err != nil {
		return err
	}

	helpOut := string(output)

	f.Add("bird", map[string]interface{}{
		"version":   result[1],
		"conf_file": reBirdConfig.FindStringSubmatch(helpOut)[1],

		// "pid_file":     reRADVDPidFile.FindStringSubmatch(out)[1],
		// "log_file":     reRADVDLogFile.FindStringSubmatch(out)[1],
		// "log_facility": reRADVDSyslogFc.FindStringSubmatch(out)[1],
	})

	return nil
}
