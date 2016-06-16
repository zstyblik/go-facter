package facter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/zstyblik/go-facter/lib/formatter"
)

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
