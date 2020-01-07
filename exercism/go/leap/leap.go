package leap

const testVersion = 3

func IsLeapYear(in int) bool {
	if in%4 == 0 {
		if in%100 == 0 {
			return in%400 == 0
		}
		return true
	}
	return false
}
