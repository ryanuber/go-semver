package semver

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	parts := []string{"1", "2", "3", "4", "5"}

	ver, err := New(parts[0], parts[1], parts[2], parts[3], parts[4])
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !reflect.DeepEqual(parts, ver.parts()) {
		t.Fatalf("bad: %v", parts)
	}
}

func TestNewFromString(t *testing.T) {
	ver, err := NewFromString("1.2.3")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "",
		Build:  "",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestNewFromString_PrereleaseOnly(t *testing.T) {
	ver, err := NewFromString("1.2.3-4.0a")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "4.0a",
		Build:  "",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestNewFromString_BuildOnly(t *testing.T) {
	ver, err := NewFromString("1.2.3+4.0a")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "",
		Build:  "4.0a",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestNewFromString_PrereleaseAndBuild(t *testing.T) {
	ver, err := NewFromString("1.2.3-alpha5+4.0a")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "alpha5",
		Build:  "4.0a",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestNewFromString_PrerelaseWithPoints(t *testing.T) {
	ver, err := NewFromString("1.2.3-alpha5.3.1")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "alpha5.3.1",
		Build:  "",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestNewFromString_BuildWithPoints(t *testing.T) {
	ver, err := NewFromString("1.2.3+5.3.1a")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "",
		Build:  "5.3.1a",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}

func TestBadBaseVersion(t *testing.T) {
	_, err := NewFromString("1a.2.3")
	if err == nil || !strings.Contains(err.Error(), "invalid major") {
		t.Fatalf("Expected major version error")
	}

	_, err = NewFromString("1.2a.3")
	if err == nil || !strings.Contains(err.Error(), "invalid minor") {
		t.Fatalf("Expected minor version error")
	}

	_, err = NewFromString("1.2.3a")
	if err == nil || !strings.Contains(err.Error(), "invalid patch") {
		t.Fatalf("Expected patch version error")
	}
}

func TestBadPreRelease(t *testing.T) {
	_, err := NewFromString("1.2.3-1_a")
	if err == nil || !strings.Contains(err.Error(), "invalid pre-release") {
		t.Fatalf("Expected pre-release error")
	}
}

func TestBadBuildMetadata(t *testing.T) {
	_, err := NewFromString("1.2.3+1_a")
	if err == nil || !strings.Contains(err.Error(), "invalid build") {
		t.Fatalf("Expected build metadata error")
	}
}
