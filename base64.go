package base64

import (
	"fmt"
	)


func Xor(one, two string) string{
	size := len(one)
	out := ""
	for i := 0; i < size; i++ {
		l := HexToDecimal(byte(one[i]))
		r := HexToDecimal(byte(two[i]))
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
		val -= 7
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
		first := (HexToDecimal(byte(s[i])) << 4)
		second := HexToDecimal(byte(s[i+1]))
		out += string(first + second)
	}
	return out
}

func hexToBase64(s string) string {
	size := len(s)
	//padding := size % 3
	out := ""
	for i := 0; i < size; i += 6 {

		dec1 := byte( ( HexToDecimal( byte(s[i]) ) << 4 ) + (HexToDecimal(byte(s[i+ 1])) ))
		dec2 := byte( ( HexToDecimal( byte(s[i + 2]) ) << 4 ) + (HexToDecimal(byte(s[i + 3])) ))
		dec3 := byte( ( HexToDecimal( byte(s[i + 4]) ) << 4 ) + (HexToDecimal(byte(s[i + 5])) ))

		base1 := dec1 >> 2
		// 3 is a 00011 for masking
		base2 := ((dec1 << 6) >> 2) + (dec2 >> 4)
		base3 := ((dec2 << 4) >> 2) + (dec3 >> 6)
		// 63 is the all one mask required to get last 6 bites
		base4 := ((dec3 << 2) >> 2)
		
		for _, num := range []byte{base1, base2, base3, base4} {
			out += byteToBase64(num)
		}
	}
	return out
}

func byteToBase64(b byte) string {

	if b < 26 {
		return string(b + 65)
	} else if b < 52 {
		return string(b + 97 - 26)
	} else if b < 62 {
		return string(b + 48  - 52)
	} else if b == 62 {
		return "+"
	} else if b == 63 {
		return "/"
	} else {
		fmt.Println("error: num incorrectly parsed")
		return ""
	}
}
