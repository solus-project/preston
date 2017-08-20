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
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Accumulator handles all of the potential license files within a repository
type Accumulator struct {
	hits       map[string]int
	transforms map[string]string // Bad name to good name mapping
	hashes     map[string]string // Simple hash comparisons
	spdx       map[string]int    // SPDX Ids
	lomap      map[string]string // lower to real SPDX name
}

// NewAccumulator will return a new Accumulator ready for processing.
func NewAccumulator() (*Accumulator, error) {
	accum := &Accumulator{
		hits:       make(map[string]int),
		transforms: make(map[string]string),
		hashes:     make(map[string]string),
		spdx:       make(map[string]int),
		lomap:      make(map[string]string),
	}
	if err := accum.loadAssets(); err != nil {
		return nil, err
	}
	return accum, nil
}

// loadAssets will populate the description tables from the main SPDX table
// file. We'll need to eventually add a better path the SPDX file.
func (a *Accumulator) loadAssets() error {
	i, err := os.Open("licenses.spdx")
	if err != nil {
		return err
	}
	defer i.Close()
	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, "\t")
		if len(splits) != 2 {
			return fmt.Errorf("malformed line in licenses.spdx")
		}
		hash := splits[0]
		nom := splits[1]
		a.spdx[nom] = 1
		a.hashes[hash] = nom
		a.lomap[strings.ToLower(nom)] = nom
	}
	return nil
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
	lines, err := a.getCondensed(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get lines for %s: %v\n", path, err)
		return
	}
	hash, err := a.getHash(&lines)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get hash for %s: %v\n", path, err)
		return
	}

	if a.pushHash(hash) {
		fmt.Printf("Got hash match!\n")
		return
	}

	fmt.Printf("License file: %s %s\n", path, hash)

}
