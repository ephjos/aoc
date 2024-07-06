#!/bin/bash

find . | entr -c -s 'go run .'
