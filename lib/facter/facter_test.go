package facter

import (
	"testing"

	"github.com/zstyblik/go-facter/lib/formatter"
)

func TestNewNilConf(t *testing.T) {
	facter := New(nil)
	if facter == nil {
		t.Fail()
	}
}

func TestNewConf(t *testing.T) {
	conf := Config{
		Formatter: formatter.NewFormatter(),
	}
	facter := New(&conf)
	if facter == nil {
		t.Fail()
	}
}
