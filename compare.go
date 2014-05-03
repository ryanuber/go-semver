package semver

// compare is used to in dependency resolution to compare two versions of
// software. The return value is an integer, either -1, 0, or 1. The result may
// be compared against '0' with standard operators to determine greater than,
// less than, etc.
func (v1 *SemVer) compare(v2 *SemVer) int {
	if v1.String() == v2.String() {
		return 0
	}
	return vcomp(v1.parts(), v2.parts())
}

// vcomp is a recursive function which will compare two slices of version
// components and return an integer representing which is greater. The result
// is intended to be compared against integer 0 using standard operators.
func vcomp(v1, v2 []string) int {
	switch {
	case len(v1) == 0 && len(v2) == 0:
		return 0
	case v1[0] > v2[0]:
		return 1
	case v1[0] < v2[0]:
		return -1
	default:
		return vcomp(v1[1:], v2[1:])
	}
}

// greaterThan determines if one version is larger than another
func (v1 *SemVer) greaterThan(v2 *SemVer) bool {
	return v1.compare(v2) > 0
}

// lessThan is the inverse of the Greater than function
func (v1 *SemVer) lessThan(v2 *SemVer) bool {
	return v1.compare(v2) < 0
}

// equalTo determines if two versions are equivalent.
func (v1 *SemVer) equalTo(v2 *SemVer) bool {
	return v1.compare(v2) == 0
}

// twiddleCompare will perform a pessemistic version comparison (~> comparator)
// to allow implicit semantic versioning trust between versioned artifacts.
func twiddleCompare(v1, v2 []string) bool {
	last := len(v2)
	for i := len(v2) - 1; i >= 0; i-- {
		if v2[i] == "" || v2[i] == "0" {
			last = i
		}
		if v2[i] == "" {
			continue
		}
		break
	}

	if vcomp(v1[:last-1], v2[:last-1]) != 0 {
		return false
	}
	if vcomp(v1[last-1:], v2[last-1:]) < 0 {
		return false
	}
	return true
}

// Compare will take a human-comprehensible string comparator and execute
// the proper comparison function agains v2.
func (v1 *SemVer) Compare(comparator string, v2 *SemVer) bool {
	switch comparator {
	case ">":
		return v1.greaterThan(v2)
	case "<":
		return v1.lessThan(v2)
	case "=":
		return v1.equalTo(v2)
	case ">=":
		return v1.greaterThan(v2) || v1.equalTo(v2)
	case "<=":
		return v1.lessThan(v2) || v1.equalTo(v2)
	case "~>":
		return twiddleCompare(v1.parts(), v2.parts())
	default:
		return false
	}
}
