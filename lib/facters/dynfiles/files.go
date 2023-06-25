package dynfiles

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/KittenConnect/go-facter/lib/facter"
	"github.com/KittenConnect/go-facter/lib/parsers"
)


var (
	pluginName     = "dynamicfiles"
	// reDevBlacklist = regexp.MustCompile("^(dm-[0-9]+|loop[0-9]+)$")
)

func init() {
	facter.RegisterSafe(pluginName, []string{"files."}, GetParsedData)
}

func GetParsedData(f facter.IFacter) error {
	data := make(map[string]interface{})

	files, err := ioutil.ReadDir("/var/lib/facter")
	if err != nil {
		return fmt.Errorf("error reading directory: %s", err)
	}

	for _, file := range files {
		filePath := filepath.Join("/var/lib/facter", file.Name())
		ext := filepath.Ext(file.Name())

		key := strings.TrimSuffix(file.Name(), ext)

		switch strings.ToLower(ext) {
		case ".yaml", ".yml":
			parsedData, err := parsers.ParseYAMLFile(filePath)
			if err != nil {
				data[key] = fmt.Sprintf("error parsing YAML file: %s", err)
				continue
			}
			data[key] = parsedData

		case ".json":
			parsedData, err := parsers.ParseJSONFile(filePath)
			if err != nil {
				data[key] = fmt.Sprintf("error parsing JSON file: %s", err)
				continue
			}
			data[key] = parsedData

		case ".ini":
			parsedData, err := parsers.ParseINIFile(filePath)
			if err != nil {
				data[key] = fmt.Sprintf("error parsing INI file: %s", err)
				continue
			}
			data[key] = parsedData

		default:
			data[key] = fmt.Sprintf("unsupported file extension: %s", ext)
		}
	}

	f.Add("files", data)

	return nil
}