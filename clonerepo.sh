#!/bin/bash
read -p "Enter the file path: " file
feature="features"
while IFS= read -r file; do
	git clone "$file"
	name=$(basename "$file" .git)
	cd "$name"
	git checkout -b "$feature"
	cd ..
done < "$file"
