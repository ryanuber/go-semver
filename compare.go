package semver

import (
	"fmt"
	"strings"
)

// Compare executes a Comparison and returns its result
func Compare(v1 *SemVer, comp string, v2 *SemVer) bool {
	switch comp {
	case ">":
		return v1.compare(v2) > 0
	case "<":
		return v1.compare(v2) < 0
	case "=":
		return v1.compare(v2) == 0
	case ">=":
		return v1.compare(v2) >= 0
	case "<=":
		return v1.compare(v2) <= 0
	default:
		return false
	}
}

// CompareString will accept both the comparator and the version to use for
// comparison as a series of comma-delimited strings. If any comparison is
// false, it is returned immediately before any further processing is done.
// Otherwise, if all comparisons succeed, true is returned.
func CompareString(v1 *SemVer, in string) (bool, error) {
	for _, cond := range strings.Split(in, ",") {
		cond = strings.TrimSpace(cond)
		parts := strings.Split(cond, " ")
		if len(parts) != 2 {
			return false, fmt.Errorf("semver: invalid comparison: %s", cond)
		}

		v2, err := NewFromString(parts[1])
		if err != nil {
			return false, err
		}

		// Bail with false result at first false comparison
		if !Compare(v1, parts[0], v2) {
			return false, nil
		}
	}

	return true, nil
}

// compare is used to compare two versions of software. The return value is an
// integer, either -1, 0, or 1. The result may be compared against '0' with
// standard operators to determine greater than, less than, etc.
func (v1 *SemVer) compare(v2 *SemVer) int {
	// Return quickly if versions match
	if v1.String() == v2.String() {
		return 0
	}

	// Normal versions are favored over pre-releases.
	if v1.BaseString() == v2.BaseString() {
		if v1.PreRel == "" && v2.PreRel != "" {
			return 1
		}
		if v1.PreRel != "" && v2.PreRel == "" {
			return -1
		}
	}

	partsA := v1.parts()
	partsB := v2.parts()

	// Exclude build metadata during version comparison
	return vcomp(partsA[:len(partsA)-1], partsB[:len(partsB)-1])
}

// vcomp is a recursive function which will compare two slices of version
// components and return an integer representing which is greater. The result
// is intended to be compared against integer 0 using standard operators.
func vcomp(v1, v2 []string) int {
	switch {
	case len(v1) == 0 && len(v2) == 0:
		return 0
	case isNumeric(v1[0]) && len(v1[0]) > len(v2[0]):
		return 1
	case isNumeric(v1[0]) && len(v1[0]) < len(v2[0]):
		return -1
	case v1[0] > v2[0]:
		return 1
	case v1[0] < v2[0]:
		return -1
	}
	return vcomp(v1[1:], v2[1:])
}
