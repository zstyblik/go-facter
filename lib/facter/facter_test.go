package facter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/KittenConnect/go-facter/lib/formatter"
)

func TestFacter(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	f := New(nil)
	if f == nil {
		t.Fail()
	}
	f.Add(testKey, testValue)
	value, ok := f.Get(testKey)
	if ok == false || strings.Compare(fmt.Sprintf("%v", value), testValue) != 0 {
		t.Fatalf("Failed to get K/V: %v:%v:%v", testKey, value, ok)
	}
	f.Delete(testKey)
	value, ok = f.Get(testKey)
	if ok != false {
		t.Fatalf("Got %v, value %v", ok, value)
	}
}

type FakeFormatter struct {
	facts map[string]interface{}
}

func (ff *FakeFormatter) Get(k string) (interface{}, bool) {
	val, ok := ff.facts[k]
	return val, ok
}

func NewFakeFormatter() *FakeFormatter {
	f := FakeFormatter{}
	f.facts = make(map[string]interface{})
	return &f
}

func (ff *FakeFormatter) Print(facts map[string]interface{}) error {
	for k, v := range facts {
		ff.facts[k] = v
	}

	return nil
}

func TestFacterFormatter(t *testing.T) {
	testKey := "test"
	testValue := "test_value"
	ff := NewFakeFormatter()
	conf := &Config{
		Formatter: ff,
	}
	f := New(conf)
	if f == nil {
		t.Fatal()
	}
	f.Add(testKey, testValue)
	f.Print()
	val, ok := ff.Get(testKey)
	if ok != true {
		t.Fatal()
	}
	if strings.Compare(fmt.Sprintf("%v", val), testValue) != 0 {
		t.Fatal()
	}
}

func TestNewNilConf(t *testing.T) {
	f := New(nil)
	if f == nil {
		t.Fail()
	}
}

func TestNewConf(t *testing.T) {
	conf := Config{
		Formatter: formatter.NewFormatter(),
	}
	f := New(&conf)
	if f == nil {
		t.Fail()
	}
}
