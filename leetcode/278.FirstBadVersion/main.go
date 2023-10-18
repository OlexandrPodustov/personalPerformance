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
	lp := 1
	rp := n

	for lp < rp {
		firstBadVersion := lp + (rp-lp)/2
		bad := isBadVersion(firstBadVersion)

		if bad && isBadVersion(firstBadVersion-1) == false {
			return firstBadVersion
		}

		if bad {
			rp = firstBadVersion
		} else {
			lp = firstBadVersion + 1
		}
	}
	if lp == rp {
		return lp
	}

	return 0
}
