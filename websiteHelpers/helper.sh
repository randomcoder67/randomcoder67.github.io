#!/usr/bin/env bash

# Website helpers shell script

if [[ "$1" == "--recent" ]]; then
	indexFile="../content/index.md"
	tempFile="temp.md"

	lineNum="$(grep -n "# Recent Posts" "${indexFile}" | grep -Eo '^[0-9]+')"
	recentFiles="$(go run main.go recentPosts.go --recent 10)"

	head -n "$lineNum" "${indexFile}" > "${tempFile}"

	echo "" >> "${tempFile}"
	echo "${recentFiles}" >> "${tempFile}"

	mv "${tempFile}" "${indexFile}"
fi
