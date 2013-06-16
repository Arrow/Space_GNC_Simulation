package main //dynamics

import (
	"fmt"
	"github.com/Arrow/Space_GNC_Simulation/util"
	"math"
	"strings"
)

const (
	EarthRadii = 6370   // km
	EarthMu    = 398600 // km3 / s2
//	EarthMuInv = 1 / EarthMu
)

var EarthMuInv float64

type StdOrbElems struct {
	A          float64
	Ecc        float64
	Inc        float64
	LongAscend float64
	ArgOfPeri  float64
	MeanAnom0  float64
}

func (s StdOrbElems) String() (str string) {
	return strings.Join([]string{
		fmt.Sprintf("  a                           =  %8.0f km", s.A),
		fmt.Sprintf("e                           =      %5.3f", s.Ecc),
		fmt.Sprintf("Longitude of Ascending Node =   %7.2f Deg", s.LongAscend*180/math.Pi),
		fmt.Sprintf("Inclination                 =   %7.2f Deg", s.Inc*180/math.Pi),
		fmt.Sprintf("Argument of Perigee         =   %7.2f Deg", s.ArgOfPeri*180/math.Pi),
		fmt.Sprintf("Mean Anomaly                =   %7.2f Deg", s.MeanAnom0*180/math.Pi)},
		"\n  ")
}

func init() {
	EarthMuInv = 1 / float64(EarthMu)
	return
}

func NewStdElems(r, r_dot util.Vector3D) (s StdOrbElems) {
	r0 := math.Sqrt(r.Dot(r))
	v02 := r_dot.Dot(r_dot)
	s.A = 2/r0 - v02*EarthMuInv
	s.A = 1 / s.A

	h := r.Cross(r_dot)
	fmt.Println(h, r, r_dot)
	//	fmt.Println(EarthMu, EarthMuInv)
	fmt.Println(r0, math.Sqrt(v02))
	c := r_dot.Cross(h).Add(r.Scale(-1 * EarthMu / r0))
	fmt.Println(c, r.Scale(-1*EarthMu/r0), r_dot.Cross(h))
	//	fmt.Println(math.Sqrt(c.Dot(c)) / EarthMu)
	s.Ecc = math.Sqrt(c.Dot(c)) / EarthMu

	h_mag := math.Sqrt(h.Dot(h))
	ie := c.Scale(1 / (EarthMu * s.Ecc))
	ih := h.Scale(1 / h_mag)
	ip := ih.Cross(ie)
	s.LongAscend = math.Atan2(ih[0], -ih[1])
	s.Inc = math.Acos(ih[2])
	s.ArgOfPeri = math.Atan2(ie[2], ip[2])

	sig0 := r.Dot(r_dot) / math.Sqrt(EarthMu)
	E0 := math.Atan2(sig0/math.Sqrt(s.A), 1-r0/s.A)
	s.MeanAnom0 = E0 - s.Ecc*math.Sin(E0)
	return
}

func main() {
	r := util.Vector3D{6488, 0, 0}
	r_dot := util.Vector3D{-0.12, 3.3, 7.8}
	s := NewStdElems(r, r_dot)
	fmt.Println(s)
	return
}
