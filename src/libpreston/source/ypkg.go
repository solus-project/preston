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
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"strconv"
)

// ypkgParser deals with the peculiarities of the ypkg format.
// While it is YAML, it makes some stylistic loose choices to permit
// one or more methods for expressing the same thing, such as having
// a single string value where a list of strings is expected
type ypkgParser struct {
	indexedMap map[interface{}]interface{}

	PkgName    string
	PkgLicense []string
}

// convStringOnly will convert the input type back to a string, in case it had
// originally looked like a not-a-string in the YAML library. Note this will always
// lead to issues with version strings that look numeric in nature, i.e. with a
// single decimal place. To preserve it, use 'quoting' in the version field.
func (y *ypkgParser) convStringOnly(is interface{}) (string, error) {
	switch is.(type) {
	case string:
		return is.(string), nil
	case bool:
		return strconv.FormatBool(is.(bool)), nil
	case int:
		return fmt.Sprintf("%d", is.(int)), nil
	case float32:
		return fmt.Sprintf("%v", is.(float32)), nil
	case float64:
		return fmt.Sprintf("%v", is.(float64)), nil
	default:
		return "", fmt.Errorf("Not a string")
	}
}

// stringOnly expects only ONE string, and a slice, etc.
func (y *ypkgParser) stringOnly(key string) (string, error) {
	blob, ok := y.indexedMap[key]
	if !ok {
		return "", fmt.Errorf("unknown key: '%s'", key)
	}
	return y.convStringOnly(blob)
}

// oneOrMoreString will parse a field that is implicitly a list of strings,
// but may be shortened to a single string value for the sake of simplicity
// in building ypkg packages.
func (y *ypkgParser) oneOrMoreString(key string) ([]string, error) {
	blob, ok := y.indexedMap[key]
	if !ok {
		return nil, fmt.Errorf("unknown key: '%s'", key)
	}
	var ret []string
	switch blob.(type) {
	case []interface{}:
		bl := blob.([]interface{})
		for i := range bl {
			a, err := y.convStringOnly(bl[i])
			if err != nil {
				return nil, err
			}
			ret = append(ret, a)
		}
		return ret, nil
	case string:
		return []string{blob.(string)}, nil
	default:
		return nil, fmt.Errorf("'%v' is not oneOrMoreString: %v", key, blob)
	}
}

// Parse will attempt to check the main items in the package.yml
func (y *ypkgParser) parse() error {
	var err error
	if y.PkgName, err = y.stringOnly("name"); err != nil {
		return err
	}
	if y.PkgLicense, err = y.oneOrMoreString("license"); err != nil {
		return err
	}
	return nil
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
	if err := ret.parse(); err != nil {
		return nil, err
	}
	return ret, nil
}

// NewEopkgPackage will return a YPKG parsed source.Package
func NewEopkgPackage(path string) (*Package, error) {
	y, err := newYpkgParser(path)
	if err != nil {
		return nil, err
	}

	return &Package{
		Name:    y.PkgName,
		License: y.PkgLicense,
	}, nil
}
