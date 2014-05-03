package semver

import (
	"fmt"
	"regexp"
	"strings"
)

// Version represents a semantic version of a crate package.
type SemVer struct {
	// The major version. This is an integer that increments when a backward-
	// incompatible change is made to a project's API.
	Major string

	// The minor version. Used to indicate feature addition.
	Minor string

	// The patch number. This is typically used to indicate bug fixes that
	// enhance existing functionality without changing or adding any API's.
	Patch string

	// The (optional) pre-release version from point 9 of semver 2.0.
	PreRel string

	// A string of optional build metadata per point 10 of semver 2.0.
	Build string
}

// String will return a flat string representing the semantic version
func (s *SemVer) String() string {
	res := fmt.Sprintf("%s.%s.%s", s.Major, s.Minor, s.Patch)
	if s.PreRel != "" {
		res = fmt.Sprintf("%s-%s", res, s.PreRel)
	}
	if s.Build != "" {
		res = fmt.Sprintf("%s+%s", res, s.Build)
	}
	return res
}

// parts will return all version components as a slice of strings.
func (s *SemVer) parts() []string {
	return []string{s.Major, s.Minor, s.Patch, s.PreRel, s.Build}
}

// New will create a new semantic versioning object from a flat
// string and populate all struct fields.
func New(vstr string) (*SemVer, error) {
	s := &SemVer{
		Build:  takeR(&vstr, "+"),
		PreRel: takeR(&vstr, "-"),
		Patch:  takeR(&vstr, "."),
		Minor:  takeR(&vstr, "."),
		Major:  vstr,
	}
	if err := s.verify(); err != nil {
		return nil, err
	}
	return s, nil
}

// verify is used to ensure that a semver object complies with the format
// defined by semver.org.
func (s *SemVer) verify() error {
	baseRe, err := regexp.Compile("^[0-9]+$")
	extRe, err := regexp.Compile("^([0-9a-zA-Z-]+)?$")
	if err != nil {
		return err
	}

	if !(baseRe.MatchString(s.Major) &&
		baseRe.MatchString(s.Minor) &&
		baseRe.MatchString(s.Patch) &&
		extRe.MatchString(s.PreRel) &&
		extRe.MatchString(s.Build)) {

		return fmt.Errorf("semver: invalid version: %s", s.String())
	}

	return nil
}

// takeR will take all characters in a string from the right side of the subject
// until sep is encountered. The subject will be pruned in-place of both sep and
// the taken string.
func takeR(subj *string, sep string) string {
	parts := strings.Split(*subj, sep)
	l := len(parts)
	if l == 1 {
		return ""
	}
	*subj = strings.Join(parts[0:l-1], sep)
	return parts[l-1]
}

// Compare is used to in dependency resolution to compare two versions of
// software. The return value is an integer, either -1, 0, or 1. The result may
// be compared against '0' with standard operators to determine greater than,
// less than, etc.
func (v1 *SemVer) Compare(v2 *SemVer) int {
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

// GreaterThan determines if one version is larger than another
func (v1 *SemVer) GreaterThan(v2 *SemVer) bool {
	return v1.Compare(v2) > 0
}

// LessThan is the inverse of the Greater than function
func (v1 *SemVer) LessThan(v2 *SemVer) bool {
	return v1.Compare(v2) < 0
}

// EqualTo determines if two versions are equivalent.
func (v1 *SemVer) EqualTo(v2 *SemVer) bool {
	return v1.Compare(v2) == 0
}
