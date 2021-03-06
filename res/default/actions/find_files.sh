#!/bin/bash

# Search files by name pattern (find)

if [ "$#" -lt 1 ]; then
    echo "Syntax: f <file_name_substring> [path]"
    exit 1
fi

path="."

if [ "$#" -gt 1 ]; then
	path="$2"
fi

find $path -name "*$1*"