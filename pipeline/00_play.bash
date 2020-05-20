#!/usr/bin/env bash

cd /workspace/pipeline

filename=$(basename $0)
filebasename=${filename%.*}

versionfilename=$(ls ${filebasename}.version*)
echo "versionfilename: $versionfilename"
version=$(echo $versionfilename | cut -d . -f 2)
echo "version: $version"

cd /workspace/functions
pwd
sourceversion=$(find . -name "$version")
echo "sourceversion: $sourceversion"
sourcedir=$(dirname $sourceversion)
echo "sourcedir: $sourcedir"

#echo "
#####################################################################
#
#   Zip content from /workspace/functions/${sourcedir} to ${$version}.zip
#
#####################################################################
#"

cd ./$sourcedir
pwd
rm ${$version}.zip 2>/dev/null
zip -r ${$version}.zip ./* || exit 1