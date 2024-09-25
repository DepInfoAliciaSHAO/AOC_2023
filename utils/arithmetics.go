package utils

func PGCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func PPCM(array []int) int {
	var res = array[0]
	for i := 1; i < len(array); i++ {
		res = (array[i] * res) / PGCD(array[i], res)
	}
	return res
}

func Abs(n int) int {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func Abs64(n int64) int64 {
	if n < 0 {
		return -n
	} else {
		return n
	}
}

func Max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
