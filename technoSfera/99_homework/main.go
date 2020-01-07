package main

import (
	"strconv"
)

func ReturnInt() int { return 1 }

func ReturnFloat() float32 { return 1.1 }

func ReturnIntArray() [3]int {
	return [3]int{1, 3, 4}
}

func ReturnIntSlice() []int {
	return []int{1, 2, 3}
}

func IntSliceToString(in []int) string {
	var res string
	for _, v := range in {
		res += strconv.Itoa(v)
	}
	return res
}

func MergeSlices(inF []float32, inI []int32) []int {
	var res = make([]int, len(inF))
	for i, v := range inF {
		res[i] = int(v)
	}
	for _, v := range inI {
		res = append(res, int(v))
	}

	return res
}

func GetMapValuesSortedByKey(in map[int]string) []string {
	var res []string
	var rawKeys []int
	for i := range in {
		rawKeys = append(rawKeys, i)
	}

	var ii = 1

	for i := 0; i < len(rawKeys); i++ {
		if rawKeys[ii] > rawKeys[i] {
			rawKeys[ii], rawKeys[i] = rawKeys[i], rawKeys[ii]
			i = 0
		}
		ii = i
	}

	for i := 0; i < len(in); i++ {
		key := rawKeys[i]
		res = append(res, in[key])
	}

	return res
}
