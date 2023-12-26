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

import (
	"github.com/iancoleman/strcase"
)

func ToCamels(list []string) (camel []string) {
	for _, value := range list {
		camel = append(camel, strcase.ToCamel(value))
	}
	return
}

func ToKebabs(list []string) (kebabs []string) {
	for _, value := range list {
		kebabs = append(kebabs, strcase.ToKebab(value))
	}
	return
}

func ToScreamingKebabs(list []string) (screamingKebabs []string) {
	for _, value := range list {
		screamingKebabs = append(screamingKebabs, strcase.ToScreamingKebab(value))
	}
	return
}

func ToSnakes(list []string) (snakes []string) {
	for _, value := range list {
		snakes = append(snakes, strcase.ToSnake(value))
	}
	return
}

func ToScreamingSnakes(list []string) (screamingSnakes []string) {
	for _, value := range list {
		screamingSnakes = append(screamingSnakes, strcase.ToScreamingSnake(value))
	}
	return
}