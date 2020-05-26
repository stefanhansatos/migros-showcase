#!/usr/bin/env bash


echo "
#####################################################################
#
#   Create soft links for testing
#
#####################################################################
"
ln -s ./functions/types.go types.go
ln -s ./functions/pubsub-storage_test.go pubsub-storage_test.go

ls -l types.go pubsub-storage_test.go
