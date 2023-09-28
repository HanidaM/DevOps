
image_directory="duplicate_files"

declare -A image_hashes

for file in "$image_directory"/*; do
  
  if [ -f "$file" ]; then
   
    md5_hash=$(md5sum "$file" | awk '{print $1}')

  
    if [ -n "${image_hashes[$md5_hash]}" ]; then
     
      echo "Removing duplicate: $file"
      rm "$file"
    else
     
      image_hashes[$md5_hash]=1
    fi
  fi
done

echo "done"