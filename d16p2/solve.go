package d16p2

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func Solve(input string) (string, error) {
	msg, err := hex.DecodeString(strings.Trim(input, "\n"))
	if err != nil {
		return "", err
	}
	data := bitSlice(msg)
	value, _ := calculate(data, 0)
	return strconv.Itoa(int(value)), nil
}

func calculate(bs bitSlice, p int) (uint64, int) {
	if bs.len()-p < 11 {
		return 0, bs.len() - 1
	}
	typ := bs.bits(p+3, p+6)
	if typ == 4 {
		prefixPos := p + 6
		lastPrefix := byte(1)
		val := uint64(0)
		for lastPrefix == 1 {
			val <<= 4
			val |= bs.bits(prefixPos+1, prefixPos+5)
			lastPrefix = bs.bit(prefixPos)
			prefixPos += 5
		}
		end := prefixPos - 1
		return val, end
	}
	subVals := []uint64{}
	var end int
	if bs.bit(p+6) == 0 {
		bitLength := bs.bits(p+7, p+22)
		limit := p + 22 + int(bitLength)
		for v, e := calculate(bs, p+22); e < limit; v, e = calculate(bs, e+1) {
			subVals = append(subVals, v)
			end = e
		}
	} else {
		numPackets := bs.bits(p+7, p+18)
		end = p + 17
		for i := 0; i < int(numPackets); i++ {
			v, e := calculate(bs, end+1)
			subVals = append(subVals, v)
			end = e
		}
	}
	var val uint64
	switch typ {
	case 0:
		for _, v := range subVals {
			val += v
		}
	case 1:
		val = 1
		for _, v := range subVals {
			val *= v
		}
	case 2:
		val = subVals[0]
		for _, v := range subVals {
			if v < val {
				val = v
			}
		}
	case 3:
		val = subVals[0]
		for _, v := range subVals {
			if v > val {
				val = v
			}
		}
	case 5:
		val = 0
		if subVals[0] > subVals[1] {
			val = 1
		}
	case 6:
		val = 0
		if subVals[0] < subVals[1] {
			val = 1
		}
	case 7:
		val = 0
		if subVals[0] == subVals[1] {
			val = 1
		}
	}
	return val, end
}

type bitSlice []byte

func (b bitSlice) len() int {
	return len(b) * 8
}

func (b bitSlice) bit(n int) byte {
	nbyte := n / 8
	nbit := 7 - n%8
	mask := byte(1) << nbit
	return (b[nbyte] & mask) >> nbit
}

func (b bitSlice) bits(start, end int) uint64 {
	bits := uint64(0)
	for pos, shift := (end-1)/8, 0; pos >= start/8; pos, shift = pos-1, shift+8 {
		bits |= uint64(b[pos]) << shift
	}
	mask := uint64((1 << (end - start)) - 1)
	shift := 7 - (end-1)%8
	return (bits >> shift) & mask
}
