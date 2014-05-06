# go-semver [![Build Status](https://travis-ci.org/ryanuber/go-semver.svg)](https://travis-ci.org/ryanuber/go-semver)

[Semantic Versioning](http://semver.org) implementation in go.

## API

```go
func Compare(v1 *SemVer, comp string, v2 *SemVer) bool
func CompareString(v1 *SemVer, in string) (bool, error)
type SemVer
    func New(major, minor, patch, preRel, build string) (*SemVer, error)
    func NewFromString(vstr string) (*SemVer, error)
    func (s *SemVer) BaseString() string
    func (s *SemVer) String() string
```

## Examples

```go
package main

import (
    "fmt"
    "github.com/ryanuber/go-semver"
)

func main() {
    v1, err := semver.NewFromString("3.12.9")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    v2, err := semver.NewFromString("3.12.8-alpha1+727d41c")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    // Compare two semver objects
    res := semver.Compare(v1, ">", v2)
    fmt.Printf("%s > %s: %#v\n", v1.String(), v2.String(), res)

    // Compare a semver object to a string
    res, err = semver.CompareString(v1, "> 3.2.8")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("%s > 3.2.8: %#v\n", v1.String(), res)

    // Compare multiple strings
    res, err = semver.CompareString(v1, "> 3.2.8, < 4.0.0")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Printf("%s > 3.2.8 and < 4.0.0: %#v\n", v1.String(), res)
}
```
