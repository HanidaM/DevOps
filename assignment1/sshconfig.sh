
remote_server="kasper@172.16.33.241"
if [ ! -f ~/.ssh/id_rsa ]; then
    ssh-keygen -t rsa -b 4096
fi
ssh-copy-id "$remote_server"

echo "SSH confidgured!"
