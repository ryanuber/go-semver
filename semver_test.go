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
	// All components present
	ver, err := NewFromString("1.2.3-4+5")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "4",
		Build:  "5",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}

	// Build number present, pre-release absent
	ver, err = NewFromString("1.2.3+4")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected = &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "",
		Build:  "4",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}

	// Pre-release present, build number absent
	ver, err = NewFromString("1.2.3-4")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected = &SemVer{
		Major:  "1",
		Minor:  "2",
		Patch:  "3",
		PreRel: "4",
		Build:  "",
	}

	if !reflect.DeepEqual(ver, expected) {
		t.Fatalf("bad: %#v", ver)
	}
}
