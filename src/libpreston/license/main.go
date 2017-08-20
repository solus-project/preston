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

package license

import (
	"fmt"
	"strings"
)

// Accumulator handles all of the potential license files within a repository
type Accumulator struct {
	hits       map[string]int
	transforms map[string]string // Bad name to good name mapping
	spdx       map[string]int    // SPDX Ids
	lomap      map[string]string // lower to real SPDX name
}

// NewAccumulator will return a new Accumulator ready for processing.
func NewAccumulator() *Accumulator {
	return &Accumulator{
		hits:       make(map[string]int),
		transforms: make(map[string]string),
		spdx:       make(map[string]int),
		lomap:      make(map[string]string),
	}
}

// pushLicenseFinal will push the license hit into the table without transforms
func (a *Accumulator) pushLicenseFinal(name string) {
	a.hits[a.lomap[name]] = 1
}

// isSPDX determines if the record is valid
func (a *Accumulator) isSPDX(name string) bool {
	if _, ok := a.spdx[name]; ok {
		return true
	}
	return false
}

func (a *Accumulator) pushTransform(name string) bool {
	if transform, ok := a.transforms[name]; ok {
		a.pushLicenseFinal(transform)
		return true
	}
	return false
}

// pushLicense will push the name into the table if possible, and attempt transforms
// on said name until one is hopefully found.
func (a *Accumulator) pushLicense(nameOrig string) {
	name := strings.TrimSpace(nameOrig)
	names := []string{
		strings.ToLower(name),
		strings.Replace(strings.ToLower(name), "_", "-", -1),
	}

	for _, n := range names {
		if a.isSPDX(n) {
			a.pushLicenseFinal(name)
		}
		if a.pushTransform(n) {
			return
		}
	}
}

// ProcessPlainLicense will handle LICENSE/LICENCE/COPYING files and determine
// the applicable license automatically.
func (a *Accumulator) ProcessPlainLicense(path string) {
	fmt.Printf("License file: %s\n", path)
}
