package formatter

import (
	"fmt"
	"sort"
	"strings"
)

// PlainTextFormatter prints-out facts in k=>v format
type PlainTextFormatter struct {
}

// NewFormatter returns new plain-text formatter
func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

func isMap(x interface{}) bool {
	t := fmt.Sprintf("%T", x)
	return strings.HasPrefix(t, "map[")
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
			pf.Print(v.(map[string]interface{}))
		} else {
			fmt.Printf("%v => %v\n", k, facts[k])
		}
	}
	return nil
}
