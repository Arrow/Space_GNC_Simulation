package util

import (
	"github.com/Arrow/unit"
)

type UnitVector3D struct {
	*unit.Unit
	V Vector3D
}

func (u UnitVector3D) Dot(v UnitVector3D) (s *unit.Quantity) {
	sU := unit.Copy(u)
	unit.Mul(sU, v)
		
	var sV float64
	for i := range u.V {
		sV += u.V[i] * v.V[i]
	}
	
	s = unit.CreateQuantity(sV, sU)
	return
}

func (u UnitVector3D) Add(v UnitVector3D) (s UnitVector3D) {
	if !unit.DimensionsMatch(u, v) {
		panic("Attempted to add the values of two units whose dimensions do not match.")
	}
	sU := unit.Copy(u)
	
	var sV Vector3D
	for i := range u.V {
		sV[i] = u.V[i] + v.V[i]
	}
	
	s = UnitVector3D{sU, sV}
	return
}

func (u UnitVector3D) Scale(f unit.Quantity) (s UnitVector3D) {
	sU := unit.Copy(u)
	unit.Mul(sU, &f)
		
	var sV Vector3D
	for i := range u.V {
		sV[i] = f.Value() * u.V[i]
	}
		
	s = UnitVector3D{sU, sV}
	return
}

func (u UnitVector3D) Cross(v UnitVector3D) (s UnitVector3D) {
	sU := unit.Copy(u)
	unit.Mul(sU, v)
	
	var sV Vector3D
	sV[0] = u.V[1]*v.V[2] - u.V[2]*v.V[1]
	sV[1] = u.V[2]*v.V[0] - u.V[0]*v.V[2]
	sV[2] = u.V[0]*v.V[1] - u.V[1]*v.V[0]
	
	s = UnitVector3D{sU, sV}
	return
}
