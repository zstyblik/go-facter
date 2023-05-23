package routehOS

import (
	"fmt"
	"os"

	"github.com/KittenConnect/go-facter/lib/facter"
)

var PluginName = "routehOS"

var fetcherFuncs = []facter.FetcherFunc{
	GetFacts,
	GetRADVDFacts,
	GetBirdFacts,
}

func init() {
	facter.RegisterSafe(PluginName, []string{"foo", "bird_", "radvd", "radvd_"}, GetAllFacts)
}

func debug(f string, v ...any) {
	fmt.Fprintf(os.Stderr, f, v...)
}

// GetAllFacts gathers all facts related to KittenConnect's RoutehOS Appliance
func GetAllFacts(f facter.IFacter) (e error) {
	for _, fetcherFunc := range fetcherFuncs {
		err := fetcherFunc(f)
		if err != nil {
			debug("go-facter/routehOS(%s) failed: %s\n", fetcherFunc.Describe(), err)
			e = err
			continue
		}
		debug("go-facter/routehOS(%s) fetched\n", fetcherFunc.Describe())
	}
	return
}

// GetFacts gathers facts related to KittenConnect's RoutehOS Appliance
func GetFacts(f facter.IFacter) error {
	f.Add("foo", map[string]interface{}{
		"version": "1.0.2",
		"bar":     "dfjkgdfgdfg",
		"baz":     "dfkgnldgdfsg",
		"test":    "gfdsgdfsgdsg",
		"tos":     "dfjkngbkdfngdfs",
		"toz": map[string]interface{}{
			"beep": "boop",
		},
	})

	return nil
}
