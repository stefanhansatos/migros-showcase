#!/usr/bin/env bash


echo "
#####################################################################
#
#   Create soft links for testing
#
#####################################################################
"
ln -s ./functions/types.go types.go
ln -s ./functions/pubsub-bigquery_test.go pubsub-bigquery_test.go

ls -l types.go pubsub-bigquery_test.go
