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
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

// ypkgParser deals with the peculiarities of the ypkg format.
// While it is YAML, it makes some stylistic loose choices to permit
// one or more methods for expressing the same thing, such as having
// a single string value where a list of strings is expected
type ypkgParser struct {
	indexedMap map[interface{}]interface{}
}

// newYpkgParser returns a new parser which will attempt to load the YAML
// file into an indexed map
func newYpkgParser(path string) (*ypkgParser, error) {
	blob, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	ret := &ypkgParser{
		indexedMap: make(map[interface{}]interface{}),
	}
	if err = yaml.Unmarshal([]byte(blob), &ret.indexedMap); err != nil {
		return nil, err
	}
	return ret, nil
}

// NewEopkgPackage will return a YPKG parsed source.Package
func NewEopkgPackage(path string) (*Package, error) {
	_, err := newYpkgParser(path)
	if err != nil {
		return nil, err
	}

	return nil, ErrNotYetImplemented
}
