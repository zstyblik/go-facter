package formatter

import (
	"encoding/json"
	"fmt"
	"sort"
)

// PlainTextFormatter prints-out facts in k=>v format
type PlainTextFormatter struct {
}

// KeyValueFormatter prints-out facts in k:v format
type KeyValueFormatter struct {
}

// JSONFormatter prints-out facts in JSON format
type JSONFormatter struct {
}

// NewFormatter returns new plain-text formatter
func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

// Print prints-out facts in k=>v format
func (pf PlainTextFormatter) Print(facts map[string]interface{}) error {
	var keys []string
	for k := range facts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%v => %v\n", k, facts[k])
	}
	return nil
}

// NewKeyValueFormatter returns new key-value formatter
func NewKeyValueFormatter() *KeyValueFormatter {
	return &KeyValueFormatter{}
}

// Print prints-out facts in k:v format
func (kvf KeyValueFormatter) Print(facts map[string]interface{}) error {
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

// NewJSONFormatter returns new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// Print prints-out facts in JSON format
func (jf *JSONFormatter) Print(facts map[string]interface{}) error {
	b, err := json.MarshalIndent(facts, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", b)
	return nil
}
