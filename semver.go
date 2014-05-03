package semver

import (
	"fmt"
	"regexp"
	"strings"
)

// Version represents a SemVer 2.0.0 version and implements some high-level
// methods to manage multiple versions.
type SemVer struct {
	Major  string // Backward-incompatible changes
	Minor  string // New functionality
	Patch  string // Bug fixes
	PreRel string // Optional pre-release tag
	Build  string // Optional build metadata
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
// the taken string. If sep is not present in subj, then "" is returned.
func takeR(subj *string, sep string) string {
	if !strings.Contains(*subj, sep) {
		return ""
	}
	parts := strings.Split(*subj, sep)
	last := len(parts) - 1
	*subj = strings.Join(parts[0:last], sep)
	return parts[last]
}
