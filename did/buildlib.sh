#!/bin/bash
go build -ldflags "-s -w" -o didlib.a -buildmode='c-archive'
