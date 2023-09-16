// Pair Coding Task A
// -------------------------
// The task is to implement a command that is going to be poor-man-clone of the `wget` command.
// It should take 1 argument - the url of a file to download, and save it to local disk.
// Each 1 second it should log current progress so far in bytes.
// Example usage:

// ```
// $ gowget http://releases.ubuntu.com/18.04.3/ubuntu-18.04.3-desktop-amd64.iso
// Downloaded 123 bytes ...
// Downloaded 456 bytes ...
// Downloaded 789 bytes ...
// ...
// ```

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	resp, err := http.Get("http://releases.ubuntu.com/18.04.3/ubuntu-18.04.3-desktop-amd64.iso")
	if err != nil {
		log.Fatal(err)
		return
	}

	file, err := os.Create("wget-result.iso")
	if err != nil {
		log.Fatal(err)
		return
	}

	cr := CountingReader{
		R:     resp.Body,
		total: 0,
	}

	chd := make(chan struct{})
	go func(chd chan struct{}) {
		for {
			fmt.Printf("Downloaded %v bytes\n", cr.total)

			select {
			case <-chd:
				break
			default:
			}

			time.Sleep(1 * time.Second)
		}
	}(chd)

	if _, err := io.Copy(file, &cr); err != nil {
		log.Fatal(err)
		return
	}
	chd <- struct{}{}
}

type CountingReader struct {
	R     io.Reader // underlying reader
	total int
}

func (rp *CountingReader) Read(p []byte) (n int, err error) {
	n, err = rp.R.Read(p)
	if err != nil {
		return 0, err
	}
	rp.total += n

	return
}
