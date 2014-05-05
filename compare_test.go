package semver

import "testing"

func TestCompare_GreaterThan(t *testing.T) {
	v1, _ := NewFromString("2.2.3")
	v2, _ := NewFromString("1.2.3")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.3.3")
	v2, _ = NewFromString("1.2.3")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-5")
	v2, _ = NewFromString("1.2.3-4")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}
}

func TestCompare_LessThan(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("2.2.3")

	if !v1.Compare("<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.3.3")

	if !v1.Compare("<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-4")
	v2, _ = NewFromString("1.2.3-5")

	if !v1.Compare("<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}
}

func TestCompare_LessThanEqualTo(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("2.2.3")

	if !v1.Compare("<=", v2) {
		t.Fatalf("%s should be <= %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.2.3")

	if !v1.Compare("<=", v2) {
		t.Fatalf("%s should be <= %s", v1.String(), v2.String())
	}
}

func TestCompare_GreaterThanEqualTo(t *testing.T) {
	v1, _ := NewFromString("2.2.3")
	v2, _ := NewFromString("1.2.3")

	if !v1.Compare(">=", v2) {
		t.Fatalf("%s should be >= %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3")
	v2, _ = NewFromString("1.2.3")

	if !v1.Compare(">=", v2) {
		t.Fatalf("%s should be >= %s", v1.String(), v2.String())
	}
}

func TestCompare_EqualTo(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3")

	if !v1.Compare("=", v2) {
		t.Fatalf("%s should = %s", v1.String(), v2.String())
	}
}

func TestCompare_IgnoreBuildMetadata(t *testing.T) {
	v1, _ := NewFromString("1.2.3-4+5")
	v2, _ := NewFromString("1.2.3-4+6")

	if !v1.Compare("=", v2) {
		t.Fatalf("%s should be equal to %s", v1.String(), v2.String())
	}
}

func TestCompare_FavorNormalOverPreRel(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3-1")

	if !v1.Compare(">", v2) {
		t.Fatalf("%s should be > %s", v1.String(), v2.String())
	}

	v1, _ = NewFromString("1.2.3-1")
	v2, _ = NewFromString("1.2.3")

	if !v1.Compare("<", v2) {
		t.Fatalf("%s should be < %s", v1.String(), v2.String())
	}
}

func TestCompare_UnknownOperator(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3")

	if v1.Compare("!", v2) {
		t.Fatalf("Expected false (unknown operator)")
	}
}

func TestCompare_TwiddlePreRelease(t *testing.T) {
	v1, _ := NewFromString("1.2.3-5")
	v2, _ := NewFromString("1.2.3-4")

	if !v1.Compare("~>", v2) {
		t.Fatalf("%s should be ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddlePatch(t *testing.T) {
	v1, _ := NewFromString("1.2.5")
	v2, _ := NewFromString("1.2.4")

	if !v1.Compare("~>", v2) {
		t.Fatalf("%s should be ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddleMinor(t *testing.T) {
	v1, _ := NewFromString("1.3.3")
	v2, _ := NewFromString("1.3.0")

	if !v1.Compare("~>", v2) {
		t.Fatalf("%s should be ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddleEqual(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.3")

	if !v1.Compare("~>", v2) {
		t.Fatalf("%s should be ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddlePatchFalse(t *testing.T) {
	v1, _ := NewFromString("1.2.3")
	v2, _ := NewFromString("1.2.4")

	if v1.Compare("~>", v2) {
		t.Fatalf("%s should NOT ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddleMinorFalse(t *testing.T) {
	v1, _ := NewFromString("1.2.0")
	v2, _ := NewFromString("1.3.0")

	if v1.Compare("~>", v2) {
		t.Fatalf("%s should NOT ~> %s", v1.String(), v2.String())
	}
}

func TestCompare_TwiddleMajorFalse(t *testing.T) {
	v1, _ := NewFromString("1.0.0")
	v2, _ := NewFromString("2.0.0")

	if v1.Compare("~>", v2) {
		t.Fatalf("%s should NOT ~> %s", v1.String(), v2.String())
	}
}
