#!/bin/bash

[[ -z "${BASH_SOURCE}" ]] && echo "Cannot locate self." && exit 1

cd "${BASH_SOURCE%/*}"
echo 'cleaning old generated code...'
rm -rf ./clients/*
rm -rf ./commands/*
echo 'building client generator...'
go get google.golang.org/api/google-api-go-generator
echo 'building gcloud_apis_gen...'
go get ./gcloud_apis_gen || exit 1
echo 'generating command source for gcloud_apis...'
gcloud_apis_gen --discovery-dir ./discovery_docs --clients-dir ./clients --commands-dir ./commands || exit
echo 'stripping buggy generated expected imports'
for f in clients/*/*/*.go; do
  tmpgo=$(dirname $f)/tmp.go
  sed -e 's/^\(package .*\) \/\/.*/\1/' < $f > $tmpgo
  cp $tmpgo $f
  rm $tmpgo
done
