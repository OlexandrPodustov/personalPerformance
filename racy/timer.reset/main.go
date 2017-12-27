package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//func main() {
//	start := time.Now()
//	var l sync.Mutex
//	var t *time.Timer
//	l.Lock()
//	t = time.AfterFunc(randomDuration(), func() {
//		fmt.Println(time.Now().Sub(start))
//		l.Lock()
//		t.Reset(randomDuration())
//		l.Unlock()
//	})
//	l.Unlock()
//	time.Sleep(5 * time.Second)
//}

//func main() {
//	start := time.Now()
//	reset := make(chan bool)
//	var t *time.Timer
//	t = time.AfterFunc(randomDuration(), func() {
//		fmt.Println(time.Now().Sub(start))
//		reset <- true
//	})
//	for time.Since(start) < 5*time.Second {
//		<-reset
//		t.Reset(randomDuration())
//	}
//}

//func main() {
//	c := make(chan bool)
//	m := make(map[string]string)
//	go func() {
//		m["1"] = "a" // First conflicting access.
//		c <- true
//	}()
//	m["2"] = "b" // Second conflicting access.
//	<-c
//	for k, v := range m {
//		fmt.Println(k, v)
//	}
//}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(ii int) {
			fmt.Println(ii) // Not the 'i' you are looking for.
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
