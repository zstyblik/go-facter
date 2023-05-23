package formatter

import (
	"fmt"

	y "gopkg.in/yaml.v3"
)

// YAMLFormatter prints-out facts in JSON format
type YAMLFormatter struct {
}

// NewYAMLFormatter returns new JSON formatter
func NewYAMLFormatter() *YAMLFormatter {
	return &YAMLFormatter{}
}

// Print prints-out facts in JSON format
func (jf *YAMLFormatter) Print(facts map[string]interface{}) error {
	b, err := y.Marshal(facts)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", b)
	return nil
}
