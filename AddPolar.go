package main

import "fmt"
import "math"
import "math/cmplx"

const (
	Pi        = math.Pi
	RadToDeg  = 180 / Pi
	DegToRad  = Pi / 180
	RadToGrad = 200 / Pi
	GradToDeg = Pi / 200
)

type Polar struct {
	radius float64
	angle  float64
}

func makePolar(r, a float64) Polar {
	return Polar{radius: r, angle: a}
}

func (p Polar) String() string {
	return fmt.Sprintf("{Radius: %.2f Angle: %.2f}", p.radius, p.angle)
}

func (p Polar) formatDegrees() string {
	return fmt.Sprintf("{Radius: %.2f Angle: %.2f degrees}", p.radius, p.angle*RadToDeg)
}

func addPolar(a, b Polar) Polar {

	// Add polar coordinates (5,140) and (3,70)
	x := cmplx.Rect(a.radius, a.angle)
	y := cmplx.Rect(b.radius, b.angle)
	radius, angle := cmplx.Polar(x + y)
	return Polar{radius, angle}

}

func main() {

	// flags processing
	// addpolar 5 140 3 70
	p1, p2 := Polar{5, 140 * DegToRad}, Polar{3, 70 * DegToRad}
	p3 := addPolar(p1, p2)
	println("Add (5,140deg) and (3,70deg)")
	println(p3.formatDegrees())
}
