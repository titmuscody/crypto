package base64

import (
	"testing"
	//"fmt"
)

func Test_hexToDecimal(t *testing.T) {
	
	testVal := 2
	testChar := "2"
	res := hexToDecimal(byte(testChar[0]))
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
	res := HexToAscii(testval)
	if res != testRes {
		t.Errorf("%s isn't %s", testval, testRes)
	}
}
