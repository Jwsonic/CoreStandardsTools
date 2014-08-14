#!/usr/bin/env bash

go-bindata -o="./bindata/bindata.go" -pkg="bindata" data/
