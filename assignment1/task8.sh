
if [ $# -ne 2 ]; then
  echo "Usage: $0 <input_file> <new_job_name>"
  exit 1
fi

input_file="$1"
new_job_name="$2"

if [ ! -f "$input_file" ]; then
  echo "Error: The file you want to change ('$input_file') doesn't exist."
  exit 1
fi

sed -i "s/job \".*\"/job \"$new_job_name\"/" "$input_file"

echo "The job name in '$input_file' has been changed to '$new_job_name'."
