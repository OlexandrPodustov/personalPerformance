package secret

const testVersion = 2

func Handshake(in uint) []string {
	var result []string
	if ok := in & 10000; ok == 0 {
		if ok := in & 1; ok != 0 {
			result = append(result, "wink")
		}
		if ok := in & 2; ok != 0 {
			result = append(result, "double blink")
		}
		if ok := in & 100; ok != 0 {
			result = append(result, "close your eyes")
		}
		if ok := in & 1000; ok != 0 {
			result = append(result, "jump")
		}
	} else {
		if ok := in & 1000; ok != 0 {
			result = append(result, "jump")
		}
		if ok := in & 100; ok != 0 {
			result = append(result, "close your eyes")
		}
		if ok := in & 2; ok != 0 {
			result = append(result, "double blink")
		}
		if ok := in & 1; ok != 0 {
			result = append(result, "wink")
		}
	}
	return result
}
