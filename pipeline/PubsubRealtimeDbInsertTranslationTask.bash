#!/usr/bin/env bash


echo "
#####################################################################
#
#   Create soft links for testing
#
#####################################################################
"
ln -s ./functions/types.go types.go
ln -s ./functions/pubsub-realtime-db_test.go pubsub-realtime-db_test.go

ls -l types.go pubsub-realtime-db_test.go
