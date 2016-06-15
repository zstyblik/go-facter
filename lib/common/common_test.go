package common

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

type netmaskTPair struct {
	in  uint8
	out string
	err error
}

func TestConvertNetmask(t *testing.T) {
	testpairs := []netmaskTPair{
		{8, "255.0.0.0", nil},
		{17, "255.255.128.0", nil},
		{24, "255.255.255.0", nil},
		{28, "255.255.255.240", nil},
		{33, "", fmt.Errorf("Invalid Netmask given.")},
	}
	for _, pair := range testpairs {
		out, err := ConvertNetmask(pair.in)
		if err == nil && pair.err != nil {
			t.Fatalf("%v != %v", err, pair.err)
		}
		if pair.err != nil {
			// probably not safe to continue
			continue
		}
		if out != pair.out {
			t.Fatalf("%v != %v", out, pair.out)
		}
	}
}

func TestGetHostEtc(t *testing.T) {
	testValue := "test_value"
	err := os.Setenv("HOST_ETC", testValue)
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostEtc()
	if strings.Compare(value, testValue) != 0 {
		t.Fatalf("%v != %v", value, testValue)
	}
}

func TestGetHostEtcNotSet(t *testing.T) {
	expectedVal := "/etc"
	err := os.Unsetenv("HOST_ETC")
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostEtc()
	if strings.Compare(value, expectedVal) != 0 {
		t.Fatalf("%v != %v", value, expectedVal)
	}
}

func TestGetHostProc(t *testing.T) {
	testValue := "test_value"
	err := os.Setenv("HOST_PROC", testValue)
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostProc()
	if strings.Compare(value, testValue) != 0 {
		t.Fatalf("%v != %v", value, testValue)
	}
}

func TestGetHostProcNotSet(t *testing.T) {
	expectedVal := "/proc"
	err := os.Unsetenv("HOST_PROC")
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostProc()
	if strings.Compare(value, expectedVal) != 0 {
		t.Fatalf("%v != %v", value, expectedVal)
	}
}

func TestGetHostSys(t *testing.T) {
	testValue := "test_value"
	err := os.Setenv("HOST_SYS", testValue)
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostSys()
	if strings.Compare(value, testValue) != 0 {
		t.Fatalf("%v != %v", value, testValue)
	}
}

func TestGetHostSysNotSet(t *testing.T) {
	expectedVal := "/sys"
	err := os.Unsetenv("HOST_SYS")
	if err != nil {
		t.Fatalf("%v", err)
	}
	value := GetHostSys()
	if strings.Compare(value, expectedVal) != 0 {
		t.Fatalf("%v != %v", value, expectedVal)
	}
}
