package base64

import (
	//"fmt"
	)


func Xor(one, two string) string{
	size := len(one)
	out := ""
	for i := 0; i < size; i++ {
		l := hexToDecimal(byte(one[i]))
		r := hexToDecimal(byte(two[i]))
		out += decimalToHex(byte(l ^ r))
	}
	return string(out)
}

func decimalToHex(i byte) string {
	val := i
	out := ""
	if val < 10 {
		out += string(val+48)
	} else if val < 16 {
		val += 55
		out += string(val)
	} else {
		out += "z"
	}
	return out
}

func HexToDecimal(s byte) int {
	val := s
	if val > 96 {
		val -= 32
	}
	if val > 57 {
		val -= 8
	}
	val -= 48
	return int(val)
}

func hexToAscii(s string) string {
	size := len(s)
	str := s
	out := ""
	if size % 2 == 1 {
		str = "0" + str
		size++
	}
	for i := 0; i < size; i += 2 {
		first := (hexToDecimal(byte(s[i])) << 4)
		second := hexToDecimal(byte(s[i+1]))
		out += string(first + second)
	}
	return out
}
