#!/usr/bin/env bash

which bash || exit 1

env

ls -l /workspace/pipeline

filename=$(basename $0)
echo $filename
filebasename=${filename%.*}
echo $filebasename
ls ${filebasename}.version*

versionfilename=$(ls ${filebasename}.version*)
echo $versionfilename

version=$(echo $versionfilename | cut -d . -f 2)
echo $version

