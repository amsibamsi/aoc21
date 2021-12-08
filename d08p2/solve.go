package d08p2

import (
	"strconv"
	"strings"

	"github.com/amsibamsi/aoc21"
)

func Solve(input string) (string, error) {
	sum := 0
	for _, line := range aoc21.ToLines(input) {
		parts := strings.Split(line, " | ")
		output := strings.Fields(parts[1])
		samples := append(strings.Fields(parts[0]), output...)
		mapping := []uint8{127, 127, 127, 127, 127, 127, 127}
		for _, digit := range samples {
			signal := uint8(0)
			for _, seg := range digit {
				signal |= 1 << (uint8(seg) - 97)
			}
			mask := uint8(127)
			mask2 := uint8(127)
			switch len(digit) {
			case 2:
				mask = 0b0100100
				mask2 = ^mask
			case 3:
				mask = 0b0100101
				mask2 = ^mask
			case 4:
				mask = 0b0101110
				mask2 = ^mask
			case 5:
				mask2 = 0b0110110
			case 6:
				mask2 = 0b0011100
			}
			for i := 0; i < 7; i++ {
				if signal&(1<<i) != 0 {
					mapping[i] &= mask
				}
				if ^signal&(1<<i) != 0 {
					mapping[i] &= mask2
				}
			}
		}
		for i := 0; i < 7; i++ {
			m := mapping[i]
			if m == 1 || m == 2 || m == 4 || m == 8 || m == 16 || m == 32 || m == 64 {
				for j := 0; j < 7; j++ {
					if j != i {
						mapping[j] &= ^m
					}
				}
			}
		}
		out := 0
		for _, digit := range output {
			num := uint8(0)
			signal := uint8(0)
			for _, seg := range digit {
				signal |= 1 << (uint8(seg) - 97)
			}
			for i := 0; i < 7; i++ {
				if signal&(1<<i) != 0 {
					num |= mapping[i]
				}
			}
			out = out * 10
			switch num {
			case 0b0100100:
				out += 1
			case 0b1011101:
				out += 2
			case 0b1101101:
				out += 3
			case 0b0101110:
				out += 4
			case 0b1101011:
				out += 5
			case 0b1111011:
				out += 6
			case 0b0100101:
				out += 7
			case 0b1111111:
				out += 8
			case 0b1101111:
				out += 9
			}
		}
		sum += out
	}
	return strconv.Itoa(sum), nil
}
