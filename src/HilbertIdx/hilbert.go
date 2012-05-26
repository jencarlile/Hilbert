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
	// n = dimension
	// m = num bits
	var h, e, d uint = 0, 0, 0
	//var l uint

	//fmt.Printf("x: %03b, y: %03b, z: %03b\n", p.X, p.Y, p.Z)

	for i := int(m-1); i>=0; i-- {
		var l uint
		var mask uint = uint(math.Pow(2, float64(i)))
		//fmt.Printf("Mask: %03b\n", mask)
		bz := p.Z & mask
		bz >>= uint(i)
		//fmt.Printf("bz: %03b\n", bz)
		l |= bz << 2
		//fmt.Printf("l: %03b\n", l)

		by := p.Y & mask
		by >>= uint(i)
		//fmt.Printf("by: %03b\n", by)
		l |= by << 1
		//fmt.Printf("l: %03b\n", l)

		bx := p.X & mask
		bx >>= uint(i)
		//fmt.Printf("bx: %03b\n", bx)
		l |= bx
		//fmt.Printf("l: %03b\n", l)
		fmt.Printf("i: %v, l: %03b\n", i, l)

		l = SubCubeTransform(l, e, d)
		fmt.Printf("SubCubeTransform: %03b\n", l)

		w := GrayCodeInverse(l)
		fmt.Printf("GCI: %v\n", w)

		e = e ^ RRot(Entry(w, n), d+1, m)
		d = d + Direction(w, n) + 1%n
		h = (h << n) | w
	}

	//fmt.Println(h, e, d)
	return h
}


// func HilbertIdxInverse(n, m, h uint) Point3 {
// 	var e, d uint = 0, 0
// 	var p Point3{0, 0, 0} //??

// 	for i = int(m-1); i>=0; i-- {

// 		// w := [ bit(h, i*n + n-1) ... bit(h, i*n)]  n is dim, m is numbits, max h is 2^mn
// 		l := GrayCode(w)
// 		// l = Inverse SubCubeTransform(l)
// 		for j := 0; j<n; j++ {
// 			//bit(p-subj, i) = bit(l, j)
// 		}
// 		// ALT
// 		// bit(p.X, i) = bit(l, 0);
// 		// bit(p.Y, i) = bit(l, 1);
// 		// bit(p.Z, i) = bit(l, 2);

// 		e = e ^ (RRot(Entry(w, n), d+1, m))
// 		d = d + Direction(w) + 1%n

// 	}

// 	return p
// }








































