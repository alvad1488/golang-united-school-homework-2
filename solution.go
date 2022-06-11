package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import "math"

type sideCounter int

const (
	SidesTriangle sideCounter = 3
	SidesSquare   sideCounter = 4
	SidesCircle   sideCounter = 0
)

func Circle(radius float64) float64 {
	var res float64 = 0
	res = math.Pi * math.Pow(radius, 2)
	return res
}

func Square(sideLen float64) float64 {
	return math.Pow(sideLen, 2)
}

func Triangle(sideLen float64) float64 {
	return (math.Sqrt(3) * math.Pow(sideLen, 2)) / 4
}

func CalcSquare(sideLen float64, sidesNum sideCounter) float64 {

	var result float64 = 0

	switch sidesNum {
	case SidesTriangle:
		result = Triangle(sideLen)
	case SidesSquare:
		result = Square(sideLen)
	case SidesCircle:
		result = Circle(sideLen)
	default:
		result = 0
	}

	return result
}
