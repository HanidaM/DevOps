
image_dir="/duplicate_files"

temp_dir=$(mktemp -d)

for image_path in "$image_dir"/*.{jpg}; do
    md5_hash=$(md5sum "$image_path" | awk '{print $1}')
    if [ ! -f "$temp_dir/$md5_hash" ]; then
        cp "$image_path" "$temp_dir/$md5_hash"
    else
        rm "$image_path"
    fi
done

mv "$temp_dir"/* "$image_dir"
rmdir "$temp_dir"

echo "Deduplication is done!"
