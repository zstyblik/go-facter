package common

import (
	"fmt"
	"math"
)

func ConvertBytes(in uint64, unit string) (float64, string, error) {
	var outUnit string
	out := in / 1024
	if out < 1 {
		return float64(in), unit, nil
	} else {
		switch unit {
		case "B":
			outUnit = "kB"
			break
		case "kB":
			outUnit = "MB"
			break
		case "MB":
			outUnit = "GB"
			break
		case "GB":
			outUnit = "TB"
			break
		default:
			return float64(in), unit, nil
		}
		return ConvertBytes(out, outUnit)
	}
}

func ConvertNetmask(in uint8) (string, error) {
	if in > 32 {
		return "", fmt.Errorf("Invalid Netmask given.")
	}
	octets := map[uint8]uint8{
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}
	var idx uint8 = 1
	for in > 0 && idx < 5 {
		if (in / 8) > 0 {
			in = in - 8
			octets[idx] = 255
		} else {
			mod := in % 8
			octets[idx] = 255 - uint8(math.Pow(2, float64(8-mod))) + 1
			in = 0
		}
		idx++
	}
	return fmt.Sprintf("%d.%d.%d.%d", octets[1], octets[2], octets[3],
		octets[4]), nil
}
