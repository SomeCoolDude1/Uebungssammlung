package main

// TODO: implement the function computeGrade
func computeGrade(gotPoints float32, maxPoints float32) float32 {
	grade := gotPoints/maxPoints*5 + 1
	return float32(grade)
}

func main() {
	// TODO: call the function computeGrade
	print(computeGrade(17.5, 28.0))  // 4.125
	print(computeGrade(14.0, 28.0))  // 3.5
	print(computeGrade(20.45, 28.0)) // 4.66
}
