#!/usr/bin/env bash

VERSION=0.2

mkdir -p builds

GOOS=linux   GOARCH=amd64  go build
tar czf builds/FhirMaker-${VERSION}-linux-amd64.tar.gz FhirMaker
rm FhirMaker
