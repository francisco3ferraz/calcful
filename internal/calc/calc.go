package calc

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func add[T numeric](values ...T) T {
	var sum T

	for _, value := range values {
		sum += value
	}

	return sum
}

func sub[T numeric](values ...T) T {
	var diff T

	for i, value := range values {
		if i == 0 {
			diff = value
		} else {
			diff -= value
		}
	}

	return diff
}

func mul[T numeric](values ...T) T {
	var product T

	for i, value := range values {
		if i == 0 {
			product = value
		} else {
			product *= value
		}
	}

	return product
}

func div[T numeric](values ...T) T {
	var quotient T

	for i, value := range values {
		if i == 0 {
			quotient = value
		} else {
			quotient /= value
		}
	}

	return quotient
}
