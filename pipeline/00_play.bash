#!/usr/bin/env bash

cd /workspace/pipeline

filename=$(basename $0)
filebasename=${filename%.*}

versionfilename=$(ls ${filebasename}.version*)
version=$(echo $versionfilename | cut -d . -f 2)

echo "
#####################################################################
#
#   Retrieved deployment version $version
#
#####################################################################
"

cd /workspace/functions
sourceversion=$(find . -name "$version")
sourcedir=$(dirname $sourceversion)

echo "
#####################################################################
#
#   Move Go's vendor directory to /workspace/functions/${sourcedir}
#
#####################################################################
"
mv /workspace/vendor /workspace/functions || exit 1

echo "
#####################################################################
#
#   Zip content from /workspace/functions/${sourcedir} to ${version}.zip
#
#####################################################################
"
cd ./$sourcedir
rm ${version}.zip 2>/dev/null
zip -r ${version}.zip ./* || exit 1