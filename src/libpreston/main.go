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

// Package libpreston provides the core functionality for the Preston CLI binary,
// and is the main entry point into all functionality. It provides various scanners
// and methods to verify the conformance of a given package.
package libpreston

import (
	"libpreston/license"
	"os"
	"path/filepath"
	"strings"
)

// TreeFunc is a simple callback type which takes a full path as a string to
// check and read.
type TreeFunc func(path string)

// TreeScanner is used to scan a source tree and do useful Things with said tree.
// Predominantly this is used for scanning and discovering licenses within the tree.
type TreeScanner struct {
	BaseDir      string // Base directory to actually scan
	callbacks    map[string]TreeFunc
	ignoredPaths []string             // List of paths to always ignore
	accum        *license.Accumulator // Detects license patterns
}

// NewTreeScanner will return a scanner for the given directory
func NewTreeScanner(basedir string) *TreeScanner {
	scanner := &TreeScanner{
		BaseDir:   basedir,
		callbacks: make(map[string]TreeFunc),
		ignoredPaths: []string{
			".git*", // Really no sense digging inside these
			"*.a",
			"*.so*",
		},
	}

	// Set up the license Accumulator
	scanner.accum = license.NewAccumulator()

	// Known license file names
	licenses := []string{
		"license*",
		"licence*",
		"copying*",
	}

	// Set up plain license files
	for _, l := range licenses {
		scanner.AddCallback(l, scanner.accum.ProcessPlainLicense)
	}

	return scanner
}

// AddCallback will register a callback for the given pattern. Note that
// all callbacks are lower-case.
func (t *TreeScanner) AddCallback(pattern string, callback TreeFunc) {
	t.callbacks[strings.ToLower(pattern)] = callback
}

// Scan will do the grunt work of actually _scanning_ the tree
func (t *TreeScanner) Scan() error {
	return filepath.Walk(t.BaseDir, t.walker)
}

// isIgnored will check if the basenamed file matches an ignored pattern
func (t *TreeScanner) isIgnored(path string) bool {
	base := strings.ToLower(filepath.Base(path))
	for _, p := range t.ignoredPaths {
		if b, _ := filepath.Match(p, base); b {
			return true
		}
	}
	return false
}

func (t *TreeScanner) fireCallbacks(path string) {
	base := strings.ToLower(filepath.Base(path))
	for pattern, callback := range t.callbacks {
		if b, _ := filepath.Match(pattern, base); b {
			callback(path)
		}
	}
}

// walker handles each item in the tree-walk, skipping "special" paths
func (t *TreeScanner) walker(path string, info os.FileInfo, err error) error {
	if info == nil {
		return nil
	}
	if t.isIgnored(path) {
		if info.IsDir() {
			return filepath.SkipDir
		}
		return nil
	}
	if info.IsDir() {
		return nil
	}
	t.fireCallbacks(path)
	return nil
}
