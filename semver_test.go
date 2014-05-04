package semver

import (
	"reflect"
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
