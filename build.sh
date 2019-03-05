#!/bin/bash

go build -o webapp_spremraj -i -v -ldflags="-X main.version=$(git describe --always --long)" .
