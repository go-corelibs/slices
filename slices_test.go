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

func TestSlices(t *testing.T) {

	Convey("Copy", t, func() {
		src := []interface{}{"strings", "are", "things"}
		dupe := Copy(src)
		So(dupe, ShouldEqual, src)
	})

	Convey("Truncate", t, func() {
		src := []interface{}{"strings", "are", "things"}
		dupe := Truncate(src, 1)
		So(dupe, ShouldEqual, []interface{}{"strings"})
	})

	Convey("Insert", t, func() {
		src := []interface{}{"strings", "things"}
		dupe := Insert(src, 1, "are")
		So(dupe, ShouldEqual, []interface{}{"strings", "are", "things"})
	})

	Convey("Insert", t, func() {
		src := []interface{}{"are", "strings", "are", "things", "are"}
		dupe := Prune(src, "are")
		So(dupe, ShouldEqual, []interface{}{"strings", "things"})
	})

	Convey("Remove", t, func() {
		src := []interface{}{"strings", "are", "things", "are"}
		dupe := Remove(src, 3)
		So(dupe, ShouldEqual, []interface{}{"strings", "are", "things"})
		more := Remove(src, 4)
		So(more, ShouldEqual, []interface{}{"strings", "are", "things", "are"})
	})

	Convey("Push", t, func() {
		src := []interface{}{"strings", "are"}
		dupe := Push(src, "things")
		So(dupe, ShouldEqual, []interface{}{"strings", "are", "things"})
	})

	Convey("Pop", t, func() {
		src := []interface{}{"strings", "are", "things"}
		dupe, value := Pop(src)
		So(dupe, ShouldEqual, []interface{}{"strings", "are"})
		So(value, ShouldEqual, "things")
	})

	Convey("Shift", t, func() {
		src := []interface{}{"are", "things"}
		dupe := Shift(src, "strings")
		So(dupe, ShouldEqual, []interface{}{"strings", "are", "things"})
	})

	Convey("Unshift", t, func() {
		src := []interface{}{"strings", "are", "things"}
		dupe, value := Unshift(src)
		So(dupe, ShouldEqual, []interface{}{"are", "things"})
		So(value, ShouldEqual, "strings")
	})

	Convey("IndexOf", t, func() {
		src := []interface{}{"strings", "are", "things"}
		idx := IndexOf(src, "are")
		So(idx, ShouldEqual, 1)
		idx = IndexOf(src, "nope")
		So(idx, ShouldEqual, -1)
	})

	Convey("IndexesOf", t, func() {
		src := []interface{}{"strings", "are", "things", "are", "they"}
		idxs := IndexesOf(src, "are")
		So(idxs, ShouldEqual, []int{1, 3})
	})

	Convey("Present, Within", t, func() {
		present := Present("are", "strings", "are", "things")
		So(present, ShouldEqual, true)
		present = Within("maybe", []string{"one", "two"}, []string{"maybe", "they", "are"})
		So(present, ShouldEqual, true)
		present = Within("nope", []string{"one", "two"}, []string{"maybe", "they", "are"})
		So(present, ShouldEqual, false)
	})

	Convey("AnyWithin", t, func() {
		present := AnyWithin([]string{"are", "they"}, []string{"one", "two"}, []string{"maybe", "they", "are"})
		So(present, ShouldEqual, true)
		present = AnyWithin([]string{"nope", "not"}, []string{"one", "two"}, []string{"maybe", "they", "are"})
		So(present, ShouldEqual, false)
	})

	Convey("Equal", t, func() {
		equal := Equal([]string{"one", "two"}, []string{"one", "two"})
		So(equal, ShouldEqual, true)
		equal = Equal([]string{"one", "two"}, []string{"one", "two"}, []string{"nope"})
		So(equal, ShouldEqual, false)
	})

	Convey("StartsWith", t, func() {
		ok := StartsWith([]string{"one", "two"}, []string{"one", "two", "many"})
		So(ok, ShouldEqual, true)
		ok = StartsWith([]string{"one", "many"}, []string{"one", "two", "many"})
		So(ok, ShouldEqual, false)
		ok = StartsWith([]string{"one", "many"}, []string{"one"})
		So(ok, ShouldEqual, false)
	})

	Convey("Append", t, func() {
		modified := Append([]string{"one", "two"}, "two", "many")
		So(modified, ShouldEqual, []string{"one", "two", "many"})
	})

	Convey("Merge", t, func() {
		modified := Merge([]string{"one", "two"}, []string{"two", "many"}, []string{"many", "more"})
		So(modified, ShouldEqual, []string{"one", "two", "many", "more"})
	})

	Convey("Unique", t, func() {
		modified := Unique([]string{"one", "two", "two", "many", "many", "more"})
		So(modified, ShouldEqual, []string{"one", "two", "many", "more"})
	})

	Convey("DuplicateCounts", t, func() {
		dupes := DuplicateCounts([]string{"one", "two", "two", "many", "many", "more"})
		So(dupes, ShouldEqual, map[string]int{
			"two":  2,
			"many": 2,
		})
	})

	Convey("Cut", t, func() {
		before, after, found := Cut(
			[]string{"one", "two", "many", "cooks", "in", "the", "kitchen"},
			[]string{"two", "many", "cooks"},
		)
		So(found, ShouldEqual, true)
		So(before, ShouldEqual, []string{"one"})
		So(after, ShouldEqual, []string{"in", "the", "kitchen"})
		before, after, found = Cut([]string{"hello", "world"}, []string{})
		So(found, ShouldEqual, false)
		So(before, ShouldEqual, []string{"hello", "world"})
		So(after, ShouldEqual, []string(nil))
		before, after, found = Cut(
			[]string{"one", "two", "many", "cooks", "in", "the", "kitchen"},
			[]string{"two", "more", "cooks"},
		)
		So(found, ShouldEqual, false)
		So(before, ShouldEqual, []string{"one", "two", "many", "cooks", "in", "the", "kitchen"})
		So(after, ShouldEqual, []string(nil))
	})
}
