
your_id="21B031212"
archive_filename="archive_1.zip"

temp_dir=$(mktemp -d)

function process_archive() {
    local archive=$1
    local code_file=""

    unzip -q "$archive" -d "$temp_dir"

    code_file=$(find "$temp_dir" -type f -name "*.txt" -exec grep -l 'Code:' {} \;)

    if [ -n "$code_file" ]; then
        old_code=$(grep 'Code:' "$code_file" | cut -d ':' -f 2 | tr -d '[:space:]')
        new_code="CodeWord_${your_id}"
        sed -i "s/$old_code/$new_code/g" "$code_file"
        zip -q -j "$archive" "$code_file"
    else
        echo "Code file not found in $archive"
    fi
}

process_archive "$archive_filename"

rm -r "$temp_dir"

echo "Processing complete."
