package main

import (
	"fmt"
	"strings"
)

func main() {
	num := []uint64{27, 27, 27, 27, 27, 27, 27}
	for len(num) > 0 && num[0] != 1 {
		num = collatz(num)
		fmt.Println(hex(bin(num)))
	}
}

func collatz(in []uint64) []uint64 {
	if (in[0] & 1) != 0 {
		return up(in)
	}
	return down(in)
}

func up(in []uint64) []uint64 {
	carry := uint64(1)
	for k, w := range in {
		w += w << 1
		oflow := w >> 63
		w &= ^uint64(1 << 63)
		w += carry
		oflow += w >> 63
		w &= ^uint64(1 << 63)
		w >>= 1
		w |= (oflow & 1) << 62
		carry = oflow >> 1
		in[k] = w
	}
	if carry > 0 {
		if carry != 1 {
			panic("carry bug")
		}
		in = append(in, 1)
	}
	return in
}

func down(in []uint64) []uint64 {
	for k := 0; k < len(in)-1; k++ {
		in[k] = (in[k] >> 1) | ((in[k+1] & 1) << 62)
	}
	w := in[len(in)-1] >> 1
	if w == 0 {
		return in[:len(in)-1]
	}
	in[len(in)-1] = w
	return in
}

func hex(b []byte) string {
	digits := make([]byte, 0, (len(b)+3)/4)
	for len(b) > 0 {
		val := b[0]
		b = b[1:]
		if len(b) > 0 {
			val |= b[0] << 1
			b = b[1:]
			if len(b) > 0 {
				val |= b[0] << 2
				b = b[1:]
				if len(b) > 0 {
					val |= b[0] << 3
					b = b[1:]
				}
			}
		}
		digits = append(digits, val)
	}
	var s strings.Builder
	for k := len(digits) - 1; k >= 0; k-- {
		s.WriteByte("0123456789abcdef"[digits[k]])
	}
	return s.String()
}

func bin(in []uint64) []byte {
	var b []byte
	for _, w := range in {
		for k := 0; k < 63; k++ {
			if (w & (uint64(1) << k)) != 0 {
				b = append(b, 1)
			} else {
				b = append(b, 0)
			}
		}
	}
	for len(b) > 1 && b[len(b)-1] == 0 {
		b = b[0 : len(b)-1]
	}
	return b
}
