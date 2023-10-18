package main

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(version int) bool {
	if version == 4 {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	if n == 1 {
		return 1
	}
	lp := 1
	rp := n
	firstBadVersion := n / 2

	for lp < rp {
		// fmt.Println("firstBadVersion, wasBadIndex", firstBadVersion,wasBadIndex)
		bad := isBadVersion(firstBadVersion)

		if bad && isBadVersion(firstBadVersion-1) == false {
			return firstBadVersion
		}

		if bad {
			rp = firstBadVersion
			firstBadVersion = lp/2 + firstBadVersion/2
		} else {
			lp = firstBadVersion
			firstBadVersion = firstBadVersion/2 + rp/2 + 1
		}
	}

	return 1
}
