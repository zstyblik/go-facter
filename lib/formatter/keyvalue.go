package formatter

import (
	"fmt"
	"sort"

	"github.com/nqd/flat"
)

// KeyValueFormatter prints-out facts in k:v format
type KeyValueFormatter struct {
}

// NewKeyValueFormatter returns new key-value formatter
func NewKeyValueFormatter() *KeyValueFormatter {
	return &KeyValueFormatter{}
}

// Print prints-out facts in k:v format
func (kvf KeyValueFormatter) Print(facts map[string]interface{}) error {
	facts, err := flat.Flatten(facts, nil)
	if err != nil {
		return err
	}
	var keys []string
	for k := range facts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v: %v\n", k, facts[k])
	}
	return nil
}
