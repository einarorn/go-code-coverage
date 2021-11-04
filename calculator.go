package calculator

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	} else if y < 0 {
		return -x + Multiply(x, y+1)
	} else if y > 0 {
		return x + Multiply(x, y - 1)
	}
	return 0
}
