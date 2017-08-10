package leap

const testVersion = 3

func IsLeapYear(in int) bool {
	// Write some code here to pass the test suite.
	if in%4 == 0 {
		if in%100 == 0 {
			if in%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}
