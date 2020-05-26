#!/usr/bin/env bash


echo "
#####################################################################
#
#   Create soft links for testing
#
#####################################################################
"
ln -s ./functions/types.go types.go
ln -s ./functions/http-frontend_test.go http-frontend_test.go

ls -l types.go http-frontend_test.go
