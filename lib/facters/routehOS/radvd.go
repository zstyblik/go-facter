package routehOS

import (
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var reRADVDVersion = regexp.MustCompile(`Version: ([\d\.]+)$`)

func GetRADVDFacts(f facter.IFacter) error {
	cmd := exec.Command("radvd", "--version")

	// Run the command and retrieve the output
	output, err := cmd.CombinedOutput()
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		return err
	}

	result := reRADVDVersion.FindStringSubmatch(string(output))

	debug("RADVDVersionRegex.Match = %v\n", result)

	if len(result) > 0 {
		f.Add("radvd_version", result[1])
	}

	return nil
}
