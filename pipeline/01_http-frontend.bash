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
#   Check /workspace/functions/${sourcedir}/${version}.zip
#
#####################################################################
"
ls -l /workspace/functions/${sourcedir}/${version}.zip || exit 1
