// Copyright (c) 2023  The Go-Curses Authors
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

// Copy creates a new slice (array) from the given slice
func Copy[V interface{}, S ~[]V](slice S) (copied S) {
	copied = make(S, len(slice))
	copy(copied, slice)
	return
}

// Truncate creates a new slice (array), of specified length, from the given slice
func Truncate[V interface{}, S ~[]V](slice S, length int) (truncated S) {
	truncated = make(S, length)
	copy(truncated, slice)
	return
}

// Insert creates a new slice (array) from the given slice, with additional values inserted at the given index
func Insert[V interface{}, S ~[]V](slice S, at int, values ...V) (modified S) {
	before := slice[:at]
	after := slice[at:]
	modified = make(S, 0)
	modified = append(modified, before...)
	modified = append(modified, values...)
	modified = append(modified, after...)
	return
}

// Prune removes all instances of the specified values from a copy of the given slice
func Prune[V comparable, S ~[]V](slice S, values ...V) (modified S) {
	modified = make(S, 0)
	for _, v := range slice {
		if !Within(v, values) {
			modified = append(modified, v)
		}
	}
	return
}

// Remove creates a new slice (array) from the given slice, with the specified index removed
func Remove[V interface{}, S ~[]V](slice S, at int) (modified S) {
	modified = make(S, 0)
	if at >= 0 && at < len(slice) {
		modified = append(modified, slice[:at]...)
		modified = append(modified, slice[at+1:]...)
	} else {
		modified = append(modified, slice...)
	}
	return
}

// Push appends the given value to a new copy of the given slice
func Push[V interface{}, S ~[]V](slice S, values ...V) (modified S) {
	modified = append(Copy(slice), values...)
	return
}

// Pop removes the last value from a Copy of the slice and returns it
func Pop[V interface{}, S ~[]V](slice S) (modified S, value V) {
	if last := len(slice) - 1; last > -1 {
		value = slice[last]
		modified = Truncate(slice, last)
	}
	return
}

// Shift prepends the given value to a new copy of the given slice
func Shift[V interface{}, S ~[]V](slice S, values ...V) (modified S) {
	modified = make(S, 0)
	modified = append(modified, values...)
	modified = append(modified, slice...)
	return
}

// Unshift removes the first value from a Copy of the slice and returns it
func Unshift[V interface{}, S ~[]V](slice S) (modified S, value V) {
	if len(slice) > 0 {
		value = slice[0]
		modified = Copy(slice[1:])
	}
	return
}

// IndexOf returns the first index matching the value given
func IndexOf[V comparable, S ~[]V](slice S, value V) (index int) {
	index = -1
	for idx, v := range slice {
		if v == value {
			index = idx
			return
		}
	}
	return
}

// IndexesOf returns a list of all indexes matching the value given
func IndexesOf[V comparable, S ~[]V](slice S, value V) (indexes []int) {
	for idx, v := range slice {
		if v == value {
			indexes = append(indexes, idx)
		}
	}
	return
}

// Present returns true if the search value is present in any of the other values given
func Present[V comparable](search V, others ...V) (present bool) {
	present = Within(search, others)
	return
}

// Within return true if the search value is present in any of the other slices of V given
func Within[V comparable, S ~[]V](search V, others ...S) (present bool) {
	for _, other := range others {
		for _, value := range other {
			if present = search == value; present {
				return
			}
		}
	}
	return
}

// AnyWithin returns true if any of the values in the source given are present in any of the other slices given
func AnyWithin[V comparable, S ~[]V](src S, others ...S) (present bool) {
	for _, search := range src {
		if present = Within(search, others...); present {
			return
		}
	}
	return
}

// Equal returns true if all the slices given have the same values
func Equal[V comparable, S ~[]V](src S, others ...S) (same bool) {
	srcLen := len(src)
	for _, other := range others {
		if srcLen != len(other) {
			return
		}
	}
	// they're all the same length so StartsWith is a valid comparison for equality
	same = StartsWith(src, others...)
	return
}

// StartsWith returns true if the other slices given start with the same values as the src
func StartsWith[V comparable, S ~[]V](src S, others ...S) (same bool) {
	srcLen := len(src)
	for _, other := range others {
		if srcLen > len(other) {
			// not enough in other
			return
		}
	}
	for idx, v := range src {
		for _, other := range others {
			if v != other[idx] {
				return
			}
		}
	}
	same = true
	return
}

// Append returns a new slice appended with only values not within the src slice
func Append[V comparable, S ~[]V](src S, values ...V) (modified S) {
	unique := make(map[V]struct{})
	for _, v := range src {
		unique[v] = struct{}{}
	}
	modified = append(modified, src...)
	for _, v := range values {
		if _, present := unique[v]; !present {
			unique[v] = struct{}{}
			modified = append(modified, v)
		}
	}
	return
}

// Merge returns a new slice with the new values found within others appended to the src slice
func Merge[V comparable, S ~[]V](src S, others ...S) (modified S) {
	unique := make(map[V]struct{})
	for _, v := range src {
		unique[v] = struct{}{}
	}
	modified = append(modified, src...)
	for _, other := range others {
		for _, v := range other {
			if _, present := unique[v]; !present {
				unique[v] = struct{}{}
				modified = append(modified, v)
			}
		}
	}
	return
}

// Unique returns a new slice with duplicate values omitted, maintaining order
func Unique[V comparable, S ~[]V](src S) (unique S) {
	lookup := make(map[V]struct{})
	for _, this := range src {
		if _, present := lookup[this]; present {
			continue
		}
		lookup[this] = struct{}{}
		unique = append(unique, this)
	}
	return
}

// DuplicateCounts returns a mapping of values and their respective counts, distinct values not included
func DuplicateCounts[V comparable, S ~[]V](src S) (counts map[V]int) {
	counts = make(map[V]int)
	for _, this := range src {
		counts[this] += 1
	}
	for _, this := range src {
		if count, ok := counts[this]; ok && count == 1 {
			delete(counts, this)
		}
	}
	return
}

// Cut is the slices version of strings.Cut
func Cut[V comparable, S ~[]V](src, sep S) (before, after S, found bool) {
	var count int
	if count = len(sep); count == 0 {
		before = src
		return
	}
	for idx, item := range src {
		if found = item == sep[0]; found {
			for jdx, other := range sep {
				if found = src[idx+jdx] == other; !found {
					break
				}
			}
			if found {
				before = src[:idx]
				after = src[idx+count:]
				return
			}
		}
	}
	before = src
	return
}
