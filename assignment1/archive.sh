#!/bin/bash

# Specify your ID and the original archive filename
your_id="21B031212"  # Replace with your actual ID
archive_filename="archive_1.zip"

# Create a temporary directory to work in
temp_dir=$(mktemp -d)

# Function to extract and process the archive
function process_archive() {
    local archive=$1
    local code_file=""

    # Unzip the archive into the temporary directory
    unzip -q "$archive" -d "$temp_dir"

    # Look for a text file with the code inside the extracted files
    code_file=$(find "$temp_dir" -type f -name "*.txt" -exec grep -l 'Code:' {} \;)

    if [ -n "$code_file" ]; then
        # Extract the code from the file
        old_code=$(grep 'Code:' "$code_file" | cut -d ':' -f 2 | tr -d '[:space:]')

        # Generate a new code by combining the old code with your ID
        new_code="CodeWord_${your_id}"

        # Replace the old code with the new one in the text file
        sed -i "s/$old_code/$new_code/g" "$code_file"

        # Recreate the original archive with the updated text file
        zip -q -j "$archive" "$code_file"
    else
        echo "Code file not found in $archive"
    fi
}

# Process the main archive
process_archive "$archive_filename"

# Clean up the temporary directory
rm -r "$temp_dir"

echo "Processing complete."
