package functional

func adder() func(int) int {
	var sum int
	return func(v int) int {
		sum += v
		return sum
	}
}

type Adder func(int) (int, Adder)

func adder2(base int) Adder {
	return func(v int) (int, Adder) {
		return base + v, adder2(base + v)
	}
}
