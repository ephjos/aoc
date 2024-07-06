#!/bin/bash

find . | entr -c -s 'go test -v'
