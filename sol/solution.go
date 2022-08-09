package sol

func getSum(a int, b int) int {
	sum, carry := a, b
	for carry != 0 {
		temp := (sum & carry) << 1 // carry
		sum = sum ^ carry
		carry = temp
	}
	return sum
}
