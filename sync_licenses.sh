#!/bin/bash

fail_exit(){
    echo "$1"
    exit 1
}

git clone https://github.com/spdx/license-list.git --depth=1 || fail_exit "Failed to clone"

if [[ -e "licenses.spdx" ]]; then
    rm -v licenses.spdx
fi

pushd license-list

rm "Updating the SPDX Licenses.txt" -v

for i in *.txt ; do
        # Strip all whitespace from it due to many licenses being reflowed
        # Removes all newlines and whitespace
        tr -d '\t\n\r ' < $i > $i.tmp
        mv $i.tmp $i
        sum=`sha256sum "${i}"|cut -f 1 -d ' '`
        nom=`echo "$i" | sed 's@\.txt$@@'`
        echo -e "${sum}\t${nom}" >> ../licenses.spdx
done
popd
rm -rf license-list
