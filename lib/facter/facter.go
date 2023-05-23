package facter

import (
	"fmt"
	"os"
	"reflect"
	"runtime"

	"github.com/KittenConnect/go-facter/lib/formatter"
)

type FetcherFunc func(IFacter) error

var (
	fetchers        = make(map[string]FetcherFunc, 0)
	registeredFacts = make(map[string]string, 0)
)

// IFacter interface
type IFacter interface {
	Add(string, interface{})
}

// Facter struct holds Facter-related attributes
type Facter struct {
	IFacter
	facts     map[string]interface{}
	formatter Formatter
}

// Config struct serves to pass Facter configuration
type Config struct {
	Formatter Formatter
}

// Formatter interface
type Formatter interface {
	Print(map[string]interface{}) error
}

func describeFunc(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Register a new facter Fetcher function
func Register(name string, f FetcherFunc) (err error) {
	value, ok := fetchers[name]
	if ok {
		err = fmt.Errorf("go-facter provider %s already defined as -> %s", name, describeFunc(value))
		fmt.Fprintf(os.Stderr, "%s(%s) failed for reason : %s\n", describeFunc(Register), describeFunc(f), err)
		return
	}

	fetchers[name] = f
	return nil
}

// RegisterSafe a new facter Fetcher function using safe methods
func RegisterSafe(name string, facts []string, f FetcherFunc) (err error) {
	// Ensure multiple providers dont override same fact
	for _, fact := range facts {
		fetcher, ok := registeredFacts[name]
		if ok {
			err = fmt.Errorf("go-facter facter %s already defined as -> %s", fact, describeFunc(fetchers[fetcher]))
			fmt.Fprintf(os.Stderr, "%s(%s) failed for reason : %s\n", describeFunc(Register), describeFunc(f), err)
			return
		}

		registeredFacts[fact] = name
	}

	// Ensure Provider not already defined
	value, ok := fetchers[name]
	if ok {
		err = fmt.Errorf("go-facter provider %s already defined as -> %s", name, describeFunc(value))
		fmt.Fprintf(os.Stderr, "%s(%s) failed for reason : %s\n", describeFunc(Register), describeFunc(f), err)
		return
	}

	// Register Provider
	fetchers[name] = f
	return nil
}

// New returns new instance of Facter
func New(userConf *Config) *Facter {
	var conf *Config
	if userConf != nil {
		conf = userConf
	} else {
		conf = &Config{
			Formatter: formatter.NewFormatter(),
		}
	}
	f := &Facter{
		facts:     make(map[string]interface{}),
		formatter: conf.Formatter,
	}
	f.Fetch()
	return f
}

// Add adds a fact
func (f *Facter) Fetch() *Facter {
	for _, fetcher := range fetchers {
		err := fetcher(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s() failed for reason : %s\n", describeFunc(fetcher), err)
		}
	}
	return f
}

// Add adds a fact
func (f *Facter) Add(k string, v interface{}) {
	f.facts[k] = v
}

// Delete deletes given fact
func (f *Facter) Delete(k string) {
	delete(f.facts, k)
}

// Get returns value of given fact, if it exists
func (f *Facter) Get(k string) (interface{}, bool) {
	value, ok := f.facts[k]
	return value, ok
}

// Print prints-out facts by calling formatter
func (f *Facter) Print() {
	f.formatter.Print(f.facts)
}
