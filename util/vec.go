package util

type Vector3D [3]float64

func (u Vector3D) Dot(v Vector3D) (s float64) {

	for i := range u {
		s += u[i] * v[i]
	}

	return
}

func (u Vector3D) Add(v Vector3D) (s Vector3D) {
	for i := range u {
		s[i] = u[i] + v[i]
	}
	return
}

func (u Vector3D) Scale(f float64) (s Vector3D) {
	for i := range u {
		s[i] = f * u[i]
	}
	return
}

func (u Vector3D) Cross(v Vector3D) (s Vector3D) {
	s[0] = u[1]*v[2] - u[2]*v[1]
	s[1] = u[2]*v[0] - u[0]*v[2]
	s[2] = u[0]*v[1] - u[1]*v[0]
	return
}
