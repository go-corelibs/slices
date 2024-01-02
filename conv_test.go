// Copyright (c) 2024  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slices

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRetype(t *testing.T) {
	Convey("Valid casting", t, func() {
		src := []interface{}{"strings", "are", "things"}
		strs, ok := Retype[string](src)
		So(ok, ShouldEqual, true)
		So(strs, ShouldEqual, []string{"strings", "are", "things"})
	})
	Convey("Invalid casting", t, func() {
		src := []interface{}{"strings", "are", "things"}
		strs, ok := Retype[int](src)
		So(ok, ShouldEqual, false)
		So(strs, ShouldEqual, []int(nil))
	})
}