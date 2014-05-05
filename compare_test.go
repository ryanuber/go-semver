package semver

import (
	"strings"
	"testing"
)

func TestCompare_GreaterThan(t *testing.T) {
	v1, _ := NewFromString("2.2.3")
	v2, _ := NewFromString("1.2.3")

	if !Compare(v1, ">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.3.3")
	v2, _ = NewFromString("1.2.3")

	if !Compare(v1, ">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-5")
	v2, _ = NewFromString("1.2.3-4")

	if !Compare(v1, ">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}
}

func TestCompare_LessThan(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("2.2.3")

	if !Compare(v1, "<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.3.3")

	if !Compare(v1, "<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-4")
	v2, _ = NewFromString("1.2.3-5")

	if !Compare(v1, "<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}
}

func TestCompare_LessThanEqualTo(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("2.2.3")

	if !Compare(v1, "<=", v2) {
		t.Fatalf("%s should be <= %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.2.3")

	if !Compare(v1, "<=", v2) {
		t.Fatalf("%s should be <= %s", v1.String(), v2.String())
	}
}

func TestCompare_GreaterThanEqualTo(t *testing.T) {
	v1, _ := NewFromString("2.2.3")
	v2, _ := NewFromString("1.2.3")

	if !Compare(v1, ">=", v2) {
		t.Fatalf("%s should be >= %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.2.3")

	if !Compare(v1, ">=", v2) {
		t.Fatalf("%s should be >= %s", v1.String(), v2.String())
	}
}

func TestCompare_EqualTo(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3")

	if !Compare(v1, "=", v2) {
		t.Fatalf("%s should = %s", v1.String(), v2.String())
	}
}

func TestCompare_IgnoreBuildMetadata(t *testing.T) {
	v1, _ := NewFromString("1.2.3-4+5")
	v2, _ := NewFromString("1.2.3-4+6")

	if !Compare(v1, "=", v2) {
		t.Fatalf("%s should be equal to %s", v1.String(), v2.String())
	}
}

func TestCompare_FavorNormalOverPreRel(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3-1")

	if !Compare(v1, ">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-1")
	v2, _ = NewFromString("1.2.3")

	if !Compare(v1, "<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}
}

func TestCompare_UnknownOperator(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3")

	if Compare(v1, "!", v2) {
		t.Fatalf("Expected false (unknown operator)")
	}
}

func TestCompareString(t *testing.T) {
	v1, _ := NewFromString("1.2.3")

	res, err := CompareString(v1, "> 1.0.0, < 2.0.0")
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if !res {
		t.Fatalf("%s should be > 1.0.0, < 2.0.0", v1.String())
	}

	res, err = CompareString(v1, "< 1.0.0")
	if err != nil {
		t.Fatalf("err: %s")
	}
	if res {
		t.Fatalf("%s should NOT < 1.0.0")
	}
}

func TestCompareString_InvalidVersion(t *testing.T) {
	v1, _ := NewFromString("1.2.3")

	_, err := CompareString(v1, "> 2")
	if !strings.Contains(err.Error(), "too short") {
		t.Fatalf("Expected version too short error")
	}
}

func TestCompareString_InvalidComparison(t *testing.T) {
	v1, _ := NewFromString("1.2.3")

	_, err := CompareString(v1, "bad")
	if !strings.Contains(err.Error(), "invalid comparison") {
		t.Fatalf("Expected invalid comparison error")
	}
}
