package hilbert

import (
	"fmt"
	"testing"
)


func TestGrayCode(t *testing.T) {
	const in6, out6 = 6, 5
	const in15, out15 = 15, 8

	if x6 := GrayCode(in6); x6 != out6 {
		t.Errorf("GrayCode(%v) = %v, want %v", in6, x6, out6)
	}
	if x15 := GrayCode(in15); x15 != out15 {
		t.Errorf("GrayCode(%v) = %v, want %v", in6, x15, out6)
	}
}

func TestGrayCodeInverse(t *testing.T) {
	const in3, out3 = 3, 2
	const in12, out12 = 12, 8

	if x3 := GrayCodeInverse(in3); x3 != out3 {
		t.Errorf("GrayCodeInverse(%v) = %v, want %v", in3, x3, out3)
	}
	if x12 := GrayCodeInverse(in12); x12 != out12 {
		t.Errorf("GrayCodeInverse(%v) = %v, want %v", in12, x12, out12)
	}
}

func TestDirectionN3(t *testing.T) {
	for i := uint(0); i < 8; i++ {
		fmt.Printf("Direction(%d): %03b\n", i, Direction(i, 3))
	}
}

func TestDirectionN2(t *testing.T) {
	for i := uint(0); i < 4; i++ {
		fmt.Printf("Direction(%d): %02b\n", i, Direction(i, 2))
	}
}

func TestEntryN3(t *testing.T) {
	for i := uint(0); i < 8; i++ {
		fmt.Printf("Entry(%d): %03b\n", i, Entry(i, 3))
	}
}

func TestEntryN2(t *testing.T) {
	for i := uint(0); i < 4; i++ {
		fmt.Printf("Entry(%d): %02b\n", i, Entry(i, 2))
	}
} 

func TestRRot(t *testing.T) {
	var numbits uint = 3
	var i uint
	for i = 0; i < 8; i++ {
		fmt.Printf("%03b : %03b : %03b : %03b\n", i, RRot(i, 1, numbits), RRot(i, 2, numbits), RRot(i, 3, numbits))
	}

	numbits = 4
	for i = 0; i < 15; i++ {
		fmt.Printf("%04b : %04b : %04b : %04b %04b\n", i, RRot(i, 1, numbits), RRot(i, 2, numbits), RRot(i, 3, numbits), RRot(i, 4, numbits))	
	}

	const in, out = 6, 5
	if x := RRot(in, 2, 3); x != out {
		t.Errorf("RRot(%v) = %v, want %v", in, x, out)
	}
}

func TestSubCubeTransform(t *testing.T) {
	res := SubCubeTransform(1, 3, 0)
	fmt.Println("SubCubeTransform(1): ", res)

	res = SubCubeTransform(2, 3, 0)
	fmt.Println("SubCubeTransform(2): ", res)

	res = SubCubeTransform(3, 0, 1)
	fmt.Println("SubCubeTransform(3): ", res)

}

func TestHilbertIndex(t *testing.T) {
	res := HilbertIdx(3, 3, &Point3{ 0, 1, 0})
	fmt.Println(res)
}



