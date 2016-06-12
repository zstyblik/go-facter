package formatter

import (
	"fmt"
)

type PlainTextFormatter struct {
}

type KeyValueFormatter struct {
}

func NewFormatter() *PlainTextFormatter {
	return &PlainTextFormatter{}
}

func (pf PlainTextFormatter) Print(facts map[string]interface{}) error {
	// TODO - sort by key first
	for k, v := range facts {
		fmt.Printf("%v => %v\n", k, v)
	}
	return nil
}

func NewKeyValueFormatter() *KeyValueFormatter {
	return &KeyValueFormatter{}
}

func (jf KeyValueFormatter) Print(facts map[string]interface{}) error {
	// TODO - sort by key first
	for k, v := range facts {
		fmt.Printf("%v:%v\n", k, v)
	}
	return nil
}
