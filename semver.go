package semver

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	baseRe      *regexp.Regexp
	preReList   []*regexp.Regexp
	buildReList []*regexp.Regexp
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
	baseRe = regexp.MustCompile("^([1-9][0-9]+)|[0-9]$")
	preReList = []*regexp.Regexp{
		regexp.MustCompile("^[0-9]+[a-zA-Z-]+([0-9a-zA-Z-]+)?$"),
		regexp.MustCompile("^[1-9a-zA-Z]([0-9a-zA-Z-]+)?$"),
		baseRe,
	}
	buildReList = []*regexp.Regexp{
		regexp.MustCompile("^[0-9a-zA-Z-]+$"),
		baseRe,
	}
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

// BaseString returns the normal version number (sans pre-release/build)
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

	parts = strings.SplitN(patch, "+", 2)
	patch = parts[0]
	if len(parts) > 1 {
		build = parts[1]
	}

	parts = strings.SplitN(patch, "-", 2)
	patch = parts[0]
	if len(parts) > 1 {
		pre = parts[1]
	}

	return New(major, minor, patch, pre, build)
}

// matchAny simplifies iterating over a slice of regexp patterns and testing if
// any of them match a subject text. matchAny automatically splits subj by "."
// and then compares each component to the list of patterns.
func matchAny(patterns []*regexp.Regexp, subj string) bool {
	for _, subj := range strings.Split(subj, ".") {
		for _, pattern := range patterns {
			if pattern.MatchString(subj) {
				return true
			}
		}
	}
	return false
}

// verify is used to ensure that a semver object complies with the format
// defined by semver.org.
func (s *SemVer) verify() error {
	if !baseRe.MatchString(s.Major) {
		return fmt.Errorf("semver: invalid major version: %s", s.Major)
	}

	if !baseRe.MatchString(s.Minor) {
		return fmt.Errorf("semver: invalid minor version: %s", s.Minor)
	}

	if !baseRe.MatchString(s.Patch) {
		return fmt.Errorf("semver: invalid patch version: %s", s.Patch)
	}

	if s.PreRel != "" && !matchAny(preReList, s.PreRel) {
		return fmt.Errorf("semver: invalid pre-release: %s", s.PreRel)
	}

	if s.Build != "" && !matchAny(buildReList, s.Build) {
		return fmt.Errorf("semver: invalid build metadata: %s", s.Build)
	}

	return nil
}
