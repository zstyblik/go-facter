package formatter

import (
	j "encoding/json"
	"fmt"
)

// JSONFormatter prints-out facts in JSON format
type JSONFormatter struct {
}

// NewJSONFormatter returns new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// Print prints-out facts in JSON format
func (jf *JSONFormatter) Print(facts map[string]interface{}) error {
	b, err := j.MarshalIndent(facts, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", b)
	return nil
}
