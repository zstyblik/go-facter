package facter

import (
	"github.com/zstyblik/go-facter/lib/formatter"
)

type Facter struct {
	facts     map[string]interface{}
	formatter Formatter
}

type FacterConfig struct {
	Formatter Formatter
}

type Formatter interface {
	Print(map[string]interface{}) error
}

func New(userConf *FacterConfig) *Facter {
	var conf *FacterConfig
	if userConf != nil {
		conf = conf
	} else {
		conf = &FacterConfig{
			Formatter: formatter.NewFormatter(),
		}
	}
	f := &Facter{
		facts:     make(map[string]interface{}),
		formatter: conf.Formatter,
	}
	return f
}

func (f *Facter) Add(k string, v interface{}) {
	f.facts[k] = v
}

func (f *Facter) Print() {
	f.formatter.Print(f.facts)
}
