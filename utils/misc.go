package utils

func FloydCycleDetection(sequence []int) (int, int) {
	var slow = 1
	var fast = 2
	var mu = 0
	var lambda = 1
	var sequenceTooShort = false
	for sequence[slow] != sequence[fast] && !sequenceTooShort {
		slow += 1
		fast += 2
		if slow > len(sequence) {
			sequenceTooShort = true
		}
	}
	if sequenceTooShort {
		return 0, 0
	} else {
		slow = 0
		for sequence[slow] != sequence[fast] {
			slow += 1
			fast += 1
			mu += 1
			if slow > len(sequence) {
				return 0, 0
			}
		}

		fast = slow + 1
		for sequence[slow] != sequence[fast] {
			fast += 1
			lambda += 1
		}
	}
	return lambda, mu
}
