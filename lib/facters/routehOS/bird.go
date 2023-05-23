package routehOS

import (
	"os/exec"
	"regexp"
	"strings"

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

	debug("%v", result)

	f.Add("bird_version", strings.Join(result, " "))

	return nil
}
