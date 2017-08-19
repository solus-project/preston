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

package source

import (
	"testing"
)

var (
	pspecTestFile = "testdata/pspec.xml"
)

func TestPspecParser(t *testing.T) {
	spkg, err := NewEopkgPackageLegacy(pspecTestFile)
	if err != nil {
		t.Fatalf("Failed to parse known file: %v", err)
	}
	if spkg.Name != "os-prober" {
		t.Fatalf("Expected 'os-prober', got '%s'", spkg.Name)
	}
	if len(spkg.License) != 1 {
		t.Fatalf("Invalid number of licenses")
	}
	if spkg.License[0] != "GPL-2.0" {
		t.Fatalf("Expected 'GPL-2.0', got '%s'", spkg.License[0])
	}
}
