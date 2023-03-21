package opconvert

import "testing"

func TestUint(t *testing.T) {
	funcDo1 := func(num interface{}) {
		n1_b := PutToBytes(num)
		t.Logf("out :\n%d - %b:\n%d - %b", num, num, n1_b, n1_b)

		// n1_int := BytesToUint64(n1_b)
		// n1_int := BytesToUint8(n1_b)
		// t.Logf("out :\n%d - %b:\n%d - %b", num, num, n1_int, n1_int)
	}

	var n1 int = 104
	funcDo1(n1)

	var n2 int = -104
	funcDo1(n2)

	// var un2 uint8 = 4
	// funcDo1(un2)

	// var un3 uint16 = 4
	// funcDo1(un3)
}
