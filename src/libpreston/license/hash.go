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
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

// getCondensed will convert the file into an array of strings with empties
// and whitespace removed
func (a *Accumulator) getCondensed(path string) ([]string, error) {
	var ret []string
	i, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer i.Close()
	sc := bufio.NewScanner(i)
	for sc.Scan() {
		text := sc.Text()
		// Nuke whitespace
		text = strings.Replace(text, "\n", "", -1)
		text = strings.Replace(text, "\r", "", -1)
		text = strings.Replace(text, "\t", "", -1)
		text = strings.Replace(text, " ", "", -1)

		// Say no to empty lines
		if text == "" {
			continue
		}
		// Pop it back without a newline (slice is implicit \n join)
		ret = append(ret, text)
	}
	return ret, nil
}

// getHash will return the appropriate hash for the given set of lines
func (a *Accumulator) getHash(lines *[]string) (string, error) {
	joinedBlob := strings.Join(*lines, "\n")
	sr := strings.NewReader(joinedBlob)
	hash := sha256.New()
	if _, err := io.Copy(hash, sr); err != nil {
		return "", err
	}
	return hex.EncodeToString([]byte(hash.Sum(nil))), nil
}
