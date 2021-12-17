package d16p1

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
	vsum := scanPackets(data, 0)
	return strconv.Itoa(int(vsum)), nil
}

func scanPackets(bs bitSlice, p int) uint64 {
	if bs.len()-p < 6 {
		return 0
	}
	vsum := bs.bits(p, p+3)
	typ := bs.bits(p+3, p+6)
	if typ == 4 {
		end := p + 6
		for bs.bit(end) == 1 {
			end += 5
		}
		end += 4
		if end+1 < bs.len() {
			vsum += scanPackets(bs, end+1)
		}
	} else {
		skip := 7 + 15
		if bs.bit(p+6) == 1 {
			skip = 7 + 11
		}
		vsum += scanPackets(bs, p+skip)
	}
	return vsum
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
