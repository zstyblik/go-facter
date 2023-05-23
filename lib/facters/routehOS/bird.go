package routehOS

import (
	"os/exec"
	"regexp"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var reBirdVersion = regexp.MustCompile(`BIRD version ([\d\.]+)`)

func GetBirdFacts(f facter.IFacter) error {
	cmd := exec.Command("bird", "--version")

	// Run the command and retrieve the output
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	result := reBirdVersion.FindStringSubmatch(string(output))

	if len(result) > 0 {
		f.Add("bird_version", result[1])
	}

	return nil
}
