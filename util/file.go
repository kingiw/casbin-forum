// Copyright 2020 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"path"
	"strings"
)

var stdImageExt = []string{"png", "jpg", "gif", "jpeg"}

func FileExt(fileName string) string {
	ext := path.Ext(fileName)

	return strings.ReplaceAll(ext, ".", "")
}

// FileType return image or file
func FileType(fileName string) string {
	ext := FileExt(fileName)

	if ext == "" {
		return "file"
	}

	for _, v := range stdImageExt {
		if v == ext {
			return "image"
		}
	}
	return "file"
}
