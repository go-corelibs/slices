[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/slices)
[![codecov](https://codecov.io/gh/go-corelibs/slices/graph/badge.svg?token=JCylkSZcov)](https://codecov.io/gh/go-corelibs/slices)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/slices)](https://goreportcard.com/report/github.com/go-corelibs/slices)

# slices - go slice type utilities

slices is a package doing all sorts of things with slices of stuff.

# Installation

``` shell
> go get github.com/go-corelibs/slices@latest
```

# Examples

## Cut

``` go
func main() {
    slice := []string{"Before", "things", "|", "After", "things"}
    before, after, found := Cut(slice, []string{"|"})
    // found == true
    // before == []string{"Before", "things"}
    // after == []string{"After", "things"}
}
```

## Carve, CarveString

``` go
func main() {
    // slices.Carve is a generic list carving function which is most easily
    // demonstrated with a slice of runes yet useful for any comparable slice
    slice := []rune(`This is STARTa sliceEND of runes`)
    before, middle, after, found := Carve(slice, []rune("START"), []rune("END"))
    // found == true
    // before == []rune("This is ")
    // middle == []rune("a slice")
    // after == []rune(" of runes")

    // slices.CarveString is a convenient wrapper around slices.Carve geared
    // for use with ~string types, making the above slices.Carve example look
    // like this:
    s := `This is STARTa a stringEND of characters`
    b, m, a, ok := CarveString(s, "START", "END")
    // ok == true
    // before == "This is "
    // middle == "a string"
    // after == " of characters"
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2023 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
