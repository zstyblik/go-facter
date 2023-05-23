package routehOS

import (
	"fmt"
	"os"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var PluginName = "routehOS"

var fetcherFuncs = []facter.FetcherFunc{GetFacts, GetBirdFacts}

func init() {
	facter.RegisterSafe(PluginName, []string{"foo", "bird_"}, GetAllFacts)
}

// GetAllFacts gathers all facts related to KittenConnect's RoutehOS Appliance
func GetAllFacts(f facter.IFacter) (e error) {
	for _, fetcherFunc := range fetcherFuncs {
		err := fetcherFunc(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "go-facter/routehOS(%s) failed: %s\n", fetcherFunc.Describe(), err)
			e = err
			continue
		}
		fmt.Fprintf(os.Stderr, "go-facter/routehOS(%s) fetched\n", fetcherFunc.Describe())
	}
	return
}

// GetFacts gathers facts related to KittenConnect's RoutehOS Appliance
func GetFacts(f facter.IFacter) error {
	f.Add("foo", "bar")

	return nil
}
