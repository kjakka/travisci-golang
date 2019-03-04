#!/bin/bash

go build -o testproject -i -v -ldflags="-X main.version=$(git describe --always --long)" .
