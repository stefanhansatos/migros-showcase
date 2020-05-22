#!/usr/bin/env bash


echo "
#####################################################################
#
#   Create soft links for testing
#
#####################################################################
"
ln -s ./functions/types.go types.go
ln -s ./functions/e2e-storage_test.go e2e-storage_test.go

ls -l types.go e2e-storage_test.go
