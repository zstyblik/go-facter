package formatter

import (
	"fmt"
	"sort"
)

type PlainTextFormatter struct {
}

type KeyValueFormatter struct {
}

func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

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

func NewKeyValueFormatter() *KeyValueFormatter {
	return &KeyValueFormatter{}
}

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
