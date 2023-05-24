package formatter

import (
	"fmt"
	"sort"
	"strings"
)

// PlainTextFormatter prints-out facts in k=>v format
type PlainTextFormatter struct {
	// indent bool
}

// NewFormatter returns new plain-text formatter
func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{
		// indent: false,
	}
}

func isMap(x interface{}) bool {
	t := fmt.Sprintf("%T", x)
	return strings.HasPrefix(t, "map[")
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
		if isMap(v) {
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
		if isMap(v) {
			// if pf.indent {
			fmt.Printf("%v => {\n", k)
			pf.PrintIndent(v.(map[string]interface{}), "  ", "  ")
			fmt.Printf("}\n")
			// } else {
			// 	fmt.Printf("%v => { ", k)
			// 	pf.Print(v.(map[string]interface{}))
			// 	fmt.Printf(" }")
			// }
		} else {
			fmt.Printf("%v => %v\n", k, v)
		}
	}
	return nil
}
