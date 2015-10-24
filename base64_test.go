package base64

import (
	"testing"
	//"fmt"
)

func Test_hexToDecimal(t *testing.T) {
	
	testVal := 13
	testChar := "d"
	res := HexToDecimal(byte(testChar[0]))
	//fmt.Printf("bytes were %d bytes", len(res))
	if int(res) != int(testVal) {
		t.Errorf("%d is not %d", int(res), int(testVal))
	}
}

func Test_Xor(t *testing.T) {
	res := Xor("6141", "2020")
	if res != "4161" {
		t.Errorf("%s was not %s", "4161", res)
	}

}

func Test_HexToAscii(t *testing.T) {
	testval := "4120422043"
	testRes := "A B C"
	res := hexToAscii(testval)
	if res != testRes {
		t.Errorf("%s isn't %s", testval, testRes)
	}
}

func Test_HexToBase64(t *testing.T) {
	testval := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	testRes := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	res := hexToBase64(testval)
	if res != testRes {
		t.Errorf("%s isn't %s", res, testRes)
	}
}
