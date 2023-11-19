package cmath

func Add(elements ...int) (sum int) {
	for i := 0; i < len(elements); i++ {
		sum += elements[i]
	}

	return
}
