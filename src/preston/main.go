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

package main

import (
	"fmt"
	"libpreston/source"
	"os"
	"strings"
)

func main() {
	spkg, err := source.NewPackage("src/libpreston/source/testdata/package.yml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing package: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Got source package: %v\n", spkg.Name)
	fmt.Printf("License(s): %s\n", strings.Join(spkg.License, ", "))
}
