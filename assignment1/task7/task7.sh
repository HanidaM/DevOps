username="task5user"

directory="/home/logs"

setfacl -d -m u:"$username":rwX "$directory"

echo "Default RW permissions have been set for $username in $directory."