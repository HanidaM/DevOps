directory='/assignment1/task6/testdir'
find "$directory" -type f -exec du -h {} + | sort -rh | head -n 5
