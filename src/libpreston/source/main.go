//
// Copyright © 2017 Ikey Doherty <ikey@solus-project.com>
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
//

// Package source provides methods for dealing with source package specs
package source

import (
	"errors"
	"fmt"
	"path/filepath"
)

var (
	// ErrNotYetImplemented is a stub
	ErrNotYetImplemented = errors.New("Not yet implemented")
)

// A Package provides the basic parsing functionality to examine a package
// and determine licenses, metadata, etc.
type Package struct {
	Name    string   // Name of the package source
	License []string // Licenses from the source spec
}

// NewPackage will return an appropriate source package for the given
// path, and attempt to parse.
func NewPackage(path string) (*Package, error) {
	base := filepath.Base(path)
	switch base {
	case "package.yml":
		return NewEopkgPackage(base)
	case "pspec.xml":
		return NewEopkgPackageLegacy(base)
	// Potential: Add .spec, snapcraft, etc, if people are needing it
	default:
		return nil, fmt.Errorf("unknown package type: %v", path)
	}
}
