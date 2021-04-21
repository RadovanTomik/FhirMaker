#!/usr/bin/env bash

VERSION=0.2

mkdir -p builds

GOOS=linux   GOARCH=amd64  go build
tar czf builds/fhir-maker-${VERSION}-linux-amd64.tar.gz fhir-maker
rm fhir-maker
