package main

import (
	"fmt"
)

type storage struct {
	version int

	data map[string]*node
}

type node struct {
	value   int
	version int
	next    *node
}

const (
	valueDoesNotExist = "<NULL>"
)

func main() {
	ns := storage{
		version: 0,
		data:    map[string]*node{},
	}
	key1 := "key1"
	key2 := "key2"
	ns.Put(key1, 5)
	ns.Put(key2, 6)

	ns.GetUnversioned(key1)
	ns.GetVersioned(key1, 1)
	ns.GetVersioned(key2, 2)
	ns.Put(key1, 7)

	ns.GetVersioned(key1, 1)
	ns.GetVersioned(key1, 2)
	ns.GetVersioned(key1, 3)
	ns.GetUnversioned("key4")
	ns.GetVersioned(key1, 4)
	ns.GetVersioned(key2, 1)
}

func (s *storage) Put(key string, value int) {
	s.version++
	if _, ok := s.data[key]; !ok {
		s.data[key] = &node{
			value:   value,
			version: s.version,
			next:    nil,
		}
	} else {
		curNode := s.data[key]

		for {
			if curNode.next == nil {
				curNode.next = &node{
					value:   value,
					version: s.version,
					next:    nil,
				}
				break
			}
		}

		s.data[key] = curNode
	}

	fmt.Printf("PUT(#%d) %s = %d\n", s.version, key, value)
}

func (s *storage) GetUnversioned(key string) int {
	currentValue, ok := s.data[key]
	if !ok {
		fmt.Printf("GET %s = %s\n", key, valueDoesNotExist)

		return -1
	}
	fmt.Printf("GET %s = %d\n", key, currentValue.value)

	return currentValue.value
}

func (s *storage) GetVersioned(key string, version int) int {
	result := -1

	currentValue, ok := s.data[key]
	if !ok {
		fmt.Printf("GET %s(#%d) = %s\n", key, version, valueDoesNotExist)

		return result
	}

	minVersion := currentValue.version
	tempCurrentValue := currentValue
	for currentValue != nil {
		if minVersion > currentValue.version {
			minVersion = currentValue.version
		}

		if currentValue.version == version {
			result = currentValue.value
			break
		}
		currentValue = currentValue.next
	}
	if result == -1 && version >= minVersion {
		result = tempCurrentValue.value
	}
	if result == -1 {
		fmt.Printf("GET %s(#%d) = %s\n", key, version, valueDoesNotExist)

		return result
	}

	fmt.Printf("GET %s(#%d) = %d\n", key, version, result)

	return result
}
