#!/usr/bin/env bash

cd /workspace/pipeline

filename=$(basename $0)
filebasename=${filename%.*}

versionfilename=$(ls ${filebasename}.version*)
version=$(echo $versionfilename | cut -d . -f 2)

cd /workspace/functions
sourceversion=$(find .. -name "$version")
sourcedir=$(dirname $sourceversion)

echo "
#####################################################################
#
#   Zip content from /workspace/functions/$sourcedir to ${$version}.zip
#
#####################################################################
"
cd ./$sourcedir
pwd
rm ${$version}.zip 2>/dev/null
zip -r ${$version}.zip ./* || exit 1