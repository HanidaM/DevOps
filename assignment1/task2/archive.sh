
file_name="archive_"
path="word.txt"
i=1
unzip "$file_name$i.zip"

rm -f "$file_name$i.zip"

while [ $i -gt 0 ]
do
    if [ -f "$file_name$((i + 1)).zip" ]
    then
        rm -f "empty.txt"
        echo "Extracting all the files now..."
        unzip "$file_name$((i+1)).zip"
        echo "Done!"
        i=$((i + 1))
        echo "$i"
        rm -f "$file_name$i.zip"
    else
        break
    fi
done


rm -f "$file_name$i.zip"
cat word.txt

echo "_21B031186" >> word.txt

cat word.txt


tar -cf "$file_name$i.zip" empty.txt word.txt

while [ $i -gt 1 ]
do
    tar -cf "$file_name$((i-1)).zip" empty.txt  "archive_$i.zip"
    rm -f "$file_name$i.zip"
    i=$((i - 1))
done

rm -f "empty.txt"
rm -f "word.txt"

echo "done"