package semver

import (
	"fmt"
	"regexp"
	"strings"
)

// Version represents a SemVer 2.0.0 version and implements some high-level
// methods to manage multiple versions.
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

// New creates a new semver object from individual version components.
func New(major, minor, patch, preRel, build string) (*SemVer, error) {
	s := &SemVer{
		Major:  major,
		Minor:  minor,
		Patch:  patch,
		PreRel: preRel,
		Build:  build,
	}
	if err := s.verify(); err != nil {
		return nil, err
	}
	return s, nil
}

// New will create a new semantic versioning object from a flat
// string and populate all struct fields.
func NewFromString(vstr string) (*SemVer, error) {
	build := takeR(&vstr, "+")
	preRe := takeR(&vstr, "-")
	patch := takeR(&vstr, ".")
	minor := takeR(&vstr, ".")
	major := vstr
	return New(major, minor, patch, preRe, build)
}

// verify is used to ensure that a semver object complies with the format
// defined by semver.org.
func (s *SemVer) verify() error {
	baseRe, err := regexp.Compile("^[0-9]+$")
	extRe, err := regexp.Compile("^([0-9a-zA-Z-]+)?$")
	if err != nil {
		return err
	}

	if !(baseRe.MatchString(s.Major) && baseRe.MatchString(s.Minor) &&
		baseRe.MatchString(s.Patch) && extRe.MatchString(s.PreRel) &&
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
