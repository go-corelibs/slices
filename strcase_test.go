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

var (
	srcInput           = []string{"strings-are-things", "THEY_CAN_BE"}
	srcCamels          = []string{"StringsAreThings", "TheyCanBe"}
	srcKebabs          = []string{"strings-are-things", "they-can-be"}
	srcScreamingKebabs = []string{"STRINGS-ARE-THINGS", "THEY-CAN-BE"}
	srcSnakes          = []string{"strings_are_things", "they_can_be"}
	srcScreamingSnakes = []string{"STRINGS_ARE_THINGS", "THEY_CAN_BE"}
)

func TestStrcase(t *testing.T) {
	Convey("ToCamels", t, func() {
		camel := ToCamels(srcInput)
		So(camel, ShouldEqual, srcCamels)
	})
	Convey("ToKebabs", t, func() {
		kebabs := ToKebabs(srcInput)
		So(kebabs, ShouldEqual, srcKebabs)
	})
	Convey("ToScreamingKebabs", t, func() {
		kebabs := ToScreamingKebabs(srcInput)
		So(kebabs, ShouldEqual, srcScreamingKebabs)
	})
	Convey("ToSnakes", t, func() {
		kebabs := ToSnakes(srcInput)
		So(kebabs, ShouldEqual, srcSnakes)
	})
	Convey("ToScreamingSnakes", t, func() {
		kebabs := ToScreamingSnakes(srcInput)
		So(kebabs, ShouldEqual, srcScreamingSnakes)
	})
}