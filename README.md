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

## Retype

``` go
func main() {
    // json.Unmarshal provided this src slice in the awesome but not always
    // useful []interface{} type
    src := []interface{}{
        "strings",
        "are",
        "things",
    }
    // let's retype it as a []string slice
    strs, ok := slice.Retype[string](src)
    // strs == []string{"strings", "are", "things"}
    // ok == true
}
```

``` go
func main() {
    // json.Unmarshal provided this src slice in the awesome but not always
    // useful []interface{} type
    src := []interface{}{
        "strings",
        "are",
        "things",
    }
    // let's retype it as a []string slice
    strs, ok := slice.Retype[string](src)
    // strs == []string{"strings", "are", "things"}
    // ok == true
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
