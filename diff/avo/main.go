package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

const basePath = "/api"
const storagePath = "storage"

type storage map[string][]byte

var keyValue = struct {
	sync.RWMutex
	m storage
}{
	m: make(map[string][]byte),
}

func main() {
	setupRoutes(basePath)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(apiBasePath string) {
	hkv := http.HandlerFunc(handleOne)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, storagePath), hkv)
}

func handleOne(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleOne")

	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", storagePath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	key := urlPathSegments[len(urlPathSegments)-1]
	fmt.Println("key", key)

	switch r.Method {
	case http.MethodGet:
		fmt.Println("handleOne getOne")

		element := getOne(key)
		if element == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if _, err := w.Write(element); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case http.MethodPut:
		fmt.Println("handleOne put one")

		bts, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("handleOne put one bts", bts)

		upsertOne(key, bts)

	case http.MethodDelete:
		fmt.Println("handleOne removeOne")

		removeOne(key)

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getOne(k string) []byte {
	keyValue.RLock()
	defer keyValue.RUnlock()

	return keyValue.m[k]
}

func upsertOne(k string, v []byte) {
	keyValue.Lock()
	defer keyValue.Unlock()

	keyValue.m[k] = v
}

func removeOne(key string) {
	keyValue.Lock()
	defer keyValue.Unlock()

	delete(keyValue.m, key)
}
