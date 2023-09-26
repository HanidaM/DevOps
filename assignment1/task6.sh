directory='C:\Users\Bauyrzhan\Desktop'
find "$directory" -type f -exec du -h {} + | sort -rh | head -n 5
