package main

import "fmt"
import "math"
import "math/cmplx"

const (
    Pi = math.Pi
    RadToDeg = 180/Pi
    DegToRad = Pi/180
    RadToGrad = 200/Pi
    GradToDeg = Pi/200
)

type Polar struct {
	r float64
	theta float64
}

func makePolar(radius, angle float64) Polar {
	return Polar{ r: radius, theta: angle }
}

func (p Polar) String() string {
	return fmt.Sprintf("{r: %.2f theta: %.2f}", p.r, p.theta)
}

func main() {

	// Add polar coordinates (5,140) and (3,70)
	x := cmplx.Rect(5,140*DegToRad)
	y := cmplx.Rect(3, 70*DegToRad)
	r, theta := cmplx.Polar(x + y)
	p := Polar{r, theta}
	fmt.Println(p)

}

