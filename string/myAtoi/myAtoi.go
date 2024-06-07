package myatoi

import "math"

func myAtoi(str string) int {
	idx := 0
	for ; idx < len(str); idx++ {
		if str[idx] != ' ' {
			break
		}
	}
	if len(str)-idx == 0 {
		return 0
	}

	positive := true
	boundary := math.MaxInt32
	if str[idx] == '-' {
		positive = false
		boundary++
		idx++
	} else if str[idx] == '+' {
		positive = true
		idx++
	}

	result := 0

	for ; idx < len(str); idx++ {
		if str[idx] >= '0' && str[idx] <= '9' {
			if result*10 > boundary {
				result = boundary
				break
			}
			result = result*10 + int(str[idx]-'0')
			if result > boundary {
				result = boundary
			}
		} else {
			break
		}
	}
	if !positive {
		return -result
	} else {
		return result
	}

}
