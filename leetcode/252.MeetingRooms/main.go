package main

import "slices"

func canAttendMeetings(intervals [][]int) bool {
	slices.SortStableFunc[[][]int, []int](intervals, compareIntSlicesByLength)

	for i := 1; i < len(intervals); i++ {
		start, end := intervals[i][0], intervals[i][1]
		prev := i - 1
		prevStart, prevEnd := intervals[prev][0], intervals[prev][1]

		if (prevStart < start || prevStart < end) && prevEnd > start {
			return false
		}
	}

	return true
}

func compareIntSlicesByLength(a, b []int) int {
	if a[0] > b[0] {
		return -1
	}

	return 1
}
