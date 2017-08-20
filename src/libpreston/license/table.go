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
	"strings"
)

func (a *Accumulator) insertTable(key, value string) {
	key = strings.ToLower(key)
	a.matchTable[key] = value
}

// Initialise the table with the key distinction points in each license text
func (a *Accumulator) initTable() {
	a.insertTable("AttributionAssuranceLicense", "AAL")
	a.insertTable("AcademicFreeLicenseVersion1.1", "AFL-1.1")
	a.insertTable("TheAcademicFreeLicensev.2.0", "AFL-2.0")
	a.insertTable("TheAcademicFreeLicensev.2.1", "AFL-2.1")
	a.insertTable("AcademicFreeLicense(“AFL”)v.3.0", "AFL-3.0")
	a.insertTable("AcademicFreeLicense(\"AFL\")v.3.0", "AFL-3.0")
	a.insertTable("AFFEROGENERALPUBLICLICENSEVersion 1", "AGPL-1.0")
	a.insertTable("GNUAFFEROGENERALPUBLICLICENSEVersion 3", "AGPL-3.0")
	a.insertTable("ADAPTIVEPUBLICLICENSEVersion1.0", "APL-1.0")
	a.insertTable("APPLEPUBLICSOURCELICENSEVersion1.0", "APSL-1.0")
	a.insertTable("APPLEPUBLICSOURCELICENSEVersion1.1", "APSL-1.1")
	a.insertTable("ApplePublicSourceLicenseVer.1.2", "APSL-1.2")
	a.insertTable("APPLEPUBLICSOURCELICENSEVersion2.0", "APSL-2.0")
	a.insertTable("APREAMBL.TEX,version1.10e", "Abstyles")
	a.insertTable("AladdinFreePublicLicense(Version 8", "Aladdin")
	a.insertTable("Copyright(c)1995-1999The Apache Group.All rights reserved.", "Apache-1.0")
	a.insertTable("ApacheLicenseVersion2.0", "Apache-2.0")
}

// pushTable tries to match the key parts of text input to headers in the table
func (a *Accumulator) pushTable(line string) bool {
	hit := false
	for k, v := range a.matchTable {
		if strings.Contains(line, k) {
			a.pushLicenseFinal(v)
			hit = true
		}
	}
	return hit
}
