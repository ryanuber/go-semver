package semver

import "testing"

func TestCompare_GreaterThan(t *testing.T) {
	// Major version difference
	v1, _ := NewFromString("2.2.3")
	v2, _ := NewFromString("1.2.3")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	// Minor version difference
	v1, _ = NewFromString("1.3.3")
	v2, _ = NewFromString("1.2.3")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	// Pre-release version comparison
	v1, _ = NewFromString("1.2.3-5")
	v2, _ = NewFromString("1.2.3-4")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}
}

func TestCompare_IgnoreBuildMetadata(t *testing.T) {
	v1, err := NewFromString("1.2.3-4+5")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	v2, err := NewFromString("1.2.3-4+6")
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if !v1.Compare("=", v2) {
		t.Fatalf("%s should be equal to %s", v1.String(), v2.String())
	}
}
