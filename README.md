# go-semver [![Build Status](https://travis-ci.org/ryanuber/go-semver.svg)](https://travis-ci.org/ryanuber/go-semver)

[Semantic Versioning](http://semver.org) implementation in go.

Goals:

* Complete enforcement of [SemVer 2.0](http://semver.org/spec/v2.0.0.html)
* Simple comparison syntax
* Version twiddle (~> comparator)

## Examples

Creating `*SemVer` objects:

```
// From a string
ver, err := semver.NewFromString("3.12.8")

// Pre-release
ver, err := semver.NewFromString("3.12.8-alpha1")

// Build metadata
ver, err := semver.NewFromString("3.12.8+727d41c")

// Pre-release + build metadata
ver, err := semver.NewFromString("3.12.8-alpha1+727d41c")

// From components
ver, err := semver.New("3", "12", "8", "alpha1", "727d41c")
```

Comparing `*SemVer` objects:

```
v1, err := semver.NewFromString("3.12.9")
if err != nil {
    return err
}

v2, err := semver.NewFromString("3.12.8-alpha1+727d41c")
if err != nil {
    return err
}

// Compare two semver objects
if semver.Compare(v1, ">", v2) {
    fmt.Println("greater than v2")
}

// Compare a semver object to a string
if semver.CompareString(v1, "> 3.2.8") {
    fmt.Println("greater than 3.2.8")
}

// Compare multiple strings
if semver.CompareString(v1, "> 3.2.8, < 4.0.0") {
    fmt.Println("greater than 3.2.8 and less than 4.0.0")
}
```
