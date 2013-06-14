package util

func Average(data []float64) (avg float64) {
	l := len(data)
	if l == 0 {
		//error
		return 0
	}
	acc := float64(0)
	for i, _ := range data {
		acc += data[i]
	}
	avg = acc / float64(l)
	return avg
}
