//Package letter is pretty hard task when it comes to performance
package letter

import "sync"

const testVersion = 1

type FreqMap map[rune]int

//Frequency was implemented
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

//ConcurrentFrequency is new func implemented by me. Speed is low. :todo -revise.
func ConcurrentFrequency(in []string) (r FreqMap) {
	r = make(FreqMap)
	var c chan FreqMap = make(chan FreqMap, len(in))
	var rrr sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(len(in))

	for _, str := range in {
		go func(msg string) {
			c <- Frequency(msg)

			rrr.Lock()
			for key, val := range <-c {
				r[key] += val
			}
			rrr.Unlock()
			wg.Done()
		}(str)
	}

	wg.Wait()

	return r
}
