package hilbert

import (
	"fmt"
	"math"
)


func GrayCode(i uint) uint {
	return i ^ (i >> 1)
}

func GrayCodeInverse(g uint) uint {
	m := uint(math.Ceil(math.Log2(float64(g + 1))))
	i, j := g, uint(1)
	for (j < m) {
		i = (i ^ (g >> j))
		j += 1
	}
	return i
}

func Direction(i, n uint) uint {
	if (float64(i) > math.Pow(2, float64(n)) - 1) || (i < 0) {
		panic("input out of range")
	}

	switch {
	case i == 0: 
		return 0
	case i%2 == 0:
		return GrayCodeInverse(i-1) % n
	case i%2 == 1:
		return GrayCodeInverse(i) % n
	}

	return 0
}

func Entry(i, n uint) uint {
	if (float64(i) > math.Pow(2, float64(n)) - 1) || (i < 0) {
		panic("input out of range")
	}

	switch i {
		case 0: return 0
		default: return GrayCode(uint(2 * math.Floor(float64( (i - 1) / 2 ) )))
	}
	return 0
}

func SubCubeTransform(in, e, d uint) uint {
	inxore := in ^ e
	res := RRot(inxore, d + 1, 2)
	return res
}


func RRot(in, rot, numbits uint) uint {
	// MOD rot amt by numbits?
	res := (in >> rot) | (in << (numbits - rot)) 
	
	var mask uint = 0
	var i uint
	for i = 0; i < numbits; i++ {
		mask = mask << 1
		mask |= 1
	}
	return res & mask
}


type Point3 struct {
	X	uint
	Y	uint
	Z uint
}

func HilbertIdx(n, m uint, p *Point3) uint {
	var h, e, d uint = 0, 0, 0
	var l uint

	for i := int(m-1); i>=0; i-- {
		l = p.Z & uint(math.Pow(2, float64(i))) 
		l <<= 1
		l |= p.Y & uint(math.Pow(2, float64(i)))
		l <<= 1
		l |= p.X & uint(math.Pow(2, float64(i)))

		fmt.Printf("i: %v, l: %08b\n", i, l)

	}

	fmt.Println(h, e, d)
	return 0
}



