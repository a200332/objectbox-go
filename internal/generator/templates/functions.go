/*
 * Copyright 2018 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package templates

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"StringTitle": strings.Title,
	"StringCamel": func(s string) string {
		result := strings.Title(s)
		return strings.ToLower(result[0:1]) + result[1:]
	},
	"TypeIdentifier": func(s string) string {
		if strings.HasPrefix(s, "[]") {
			return strings.Title(s[2:]) + "Vector"
		} else {
			return strings.Title(s)
		}
	},
}
