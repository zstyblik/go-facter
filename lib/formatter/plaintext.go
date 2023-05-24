package formatter

import (
	"fmt"
	"sort"
	"strings"
)

// PlainTextFormatter prints-out facts in k=>v format
type PlainTextFormatter struct {}

// NewFormatter returns new plain-text formatter
func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

// Print prints-out facts in k=>v format
func (pf PlainTextFormatter) PrintIndent(facts map[string]interface{}, prefix, indent string) error {
	var keys []string
	for k := range facts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := facts[k]
		if strings.HasPrefix(fmt.Sprintf("%T", v), "map[") {
			fmt.Printf("%s%v => {\n", prefix, k)
			pf.PrintIndent(v.(map[string]interface{}), prefix+indent, indent)
			fmt.Printf("%s}\n", prefix)
		} else {
			fmt.Printf("%s%v => %v\n", prefix, k, v)
		}
	}
	return nil
}

// Print prints-out facts in k=>v format
func (pf PlainTextFormatter) Print(facts map[string]interface{}) error {
	var keys []string
	for k := range facts {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		v := facts[k]
		if strings.HasPrefix(fmt.Sprintf("%T", v), "map[") {
			fmt.Printf("%v => {\n", k)
			pf.PrintIndent(v.(map[string]interface{}), "  ", "  ")
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v => %v\n", k, v)
		}
	}
	return nil
}
