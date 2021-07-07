package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	url := "https://en.wikipedia.org/wiki/UEFA_Euro_2020"

	res, err := collectURLS(url)
	if err != nil {
		return log.Fatal(err)
	}
	// we need to print the results in descending order

	fmt.Println(res)
}

func collectURLS(url string) (map[string]int, error) {
	result := make(map[string]int)

	urls, err := getURLs(result, url)
	if err != nil {
		return nil, err
	}

	urls, err = getURLs(result, urls...)
	if err != nil {
		return nil, err
	}

	urls, err = getURLs(result, urls...)
	if err != nil {
		return nil, err
	}

	urls, err = getURLs(result, urls...)
	if err != nil {
		return nil, err
	}

	_, err = getURLs(result, urls...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

var href = regexp.MustCompile("")

// todo: add new input param like rootUrl
func getURLs(result map[string]int, urls ...string) ([]string, error) {
	urlRes := make([]string, 0)

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return nil, nil
		}

		rb, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		_ = rb
		// 1. look for href tag
		// href.

		// 2. check if it starts from /
		// 2.1 if yes prepend it with the root url

		newUrl := ""
		result[newUrl]++

		// urlRes = append(urlRes, s)
	}

	return urlRes, nil
}
