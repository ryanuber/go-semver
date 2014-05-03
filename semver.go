package semver

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	baseRe    *regexp.Regexp
	extReList []*regexp.Regexp
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

func init() {
	baseRe = regexp.MustCompile("^(([1-9][0-9]+)|[0-9])?$")
	extReList = []*regexp.Regexp{
		regexp.MustCompile("^([1-9a-zA-Z]([0-9a-zA-Z-]+)?)?$"),
		regexp.MustCompile("^([1-9]([0-9]+)?)?$"),
	}
}

// String will return a flat string representing the semantic version
func (s *SemVer) String() string {
	res := s.BaseString()
	if s.PreRel != "" {
		res = fmt.Sprintf("%s-%s", res, s.PreRel)
	}
	if s.Build != "" {
		res = fmt.Sprintf("%s+%s", res, s.Build)
	}
	return res
}

// BaseString will return the base version number (sans pre-release and build)
// as a formatted string.
func (s *SemVer) BaseString() string {
	return fmt.Sprintf("%s.%s.%s", s.Major, s.Minor, s.Patch)
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
	var major, minor, patch, pre, build string

	parts := strings.SplitN(vstr, ".", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("semver: version too short: %s", vstr)
	}
	major, minor, patch = parts[0], parts[1], parts[2]

	parts = strings.SplitN(patch, "-", 2)
	patch = parts[0]
	if len(parts) > 1 {
		pre = parts[1]
	}

	parts = strings.SplitN(pre, "+", 2)
	pre = parts[0]
	if len(parts) > 1 {
		build = parts[1]
	}

	return New(major, minor, patch, pre, build)
}

// matchAny simplifies iterating over a slice of regexp patterns and testing if
// any of them match a subject text.
func matchAny(patterns []*regexp.Regexp, subj string) bool {
	for _, pattern := range patterns {
		if pattern.MatchString(subj) {
			return true
		}
	}
	return false
}

// verify is used to ensure that a semver object complies with the format
// defined by semver.org.
func (s *SemVer) verify() error {
	if !(baseRe.MatchString(s.Major) &&
		baseRe.MatchString(s.Minor) &&
		baseRe.MatchString(s.Patch)) {
		return fmt.Errorf("semver: invalid base version: %s", s.BaseString())
	}

	for _, subj := range strings.Split(s.PreRel, ".") {
		if !matchAny(extReList, subj) {
			return fmt.Errorf("semver: invalid pre-release tag: %s", s.PreRel)
		}
	}

	for _, subj := range strings.Split(s.Build, ".") {
		if !matchAny(extReList, subj) {
			return fmt.Errorf("semver: invalid build metadata: %s", s.Build)
		}
	}

	return nil
}
