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
	a.insertTable("TheArtisticLicense2.0", "Artistic-2.0")
	a.insertTable("BoostSoftwareLicenseVersion1.0", "BSL-1.0")
	a.insertTable("shscriptareCopyright(c)GMV1991", "Bahyph")
	a.insertTable("builtontopofXy-pic byMichaelBarr", "Barr")
	a.insertTable("THEBEER-WARELICENSE", "Beerware")
	a.insertTable("bytheFreeSoftwareFoundationinversion2.2ofBison", "Bison-exception-2.2")
	a.insertTable("(CDDL)Version1.0", "CDDL-1.0")
	a.insertTable("(CDDL)Version1.1", "CDDL-1.1")
	a.insertTable("GNUGENERALPUBLICLICENSEVersion1", "GPL-1.0")
	a.insertTable("GNUGENERALPUBLICLICENSEVersion2", "GPL-2.0")
	a.insertTable("GNUGENERALPUBLICLICENSEVersion3", "GPL-3.0")
	a.insertTable("GNULIBRARYGENERALPUBLICLICENSEVersion2", "LGPL-2.0")
	a.insertTable("GNULESSERGENERALPUBLICLICENSEVersion2.1", "LGPL-2.1")
	a.insertTable("GNULESSERGENERALPUBLICLICENSEVersion3", "LGPL-3.0")
	a.insertTable("libpngnoticesisprovidedforyourconvenience", "Libpng")
	a.insertTable("thatisbuiltusingGNULibtool,youmayincludethisfile", "Libtool-exception")
	a.insertTable("MITLicense", "MIT")
	a.insertTable("MOZILLAPUBLICLICENSEVersion1.0", "MPL-1.0")
	a.insertTable("MozillaPublicLicenseVersion1.1", "MPL-1.1")
	a.insertTable("OpenSSLLicense", "OpenSSL")
	a.insertTable(",ThePostgreSQLGlobalDevelopmentGroup", "PostgreSQL")
	a.insertTable("PYTHONSOFTWAREFOUNDATIONLICENSEVERSION2", "Python-2.0")
	a.insertTable("VIMLICENSE", "Vim")
	a.insertTable("X11License", "X11")
	a.insertTable("zlibLicense", "Zlib")
	a.insertTable("DanielStenberg,<daniel@haxx.se>", "curl")
	a.insertTable("Version1.0.5of10December2007", "bzip2-1.0.5")
	a.insertTable("libbzip2version1.0.6of6September2010", "bzip2-1.0.6")
	a.insertTable("SAMLEFFLERORSILICONGRAPHICS", "libtiff")
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
