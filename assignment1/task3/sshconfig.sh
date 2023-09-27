
username="kasper"
remote_server="172.16.33.241"

if [ ! -f ~/.ssh/id_rsa ]; then
  ssh-keygen -t rsa -b 2048
fi

ssh-copy-id "$username@$remote_server"

ssh -i ~/.ssh/id_rsa "$username@$remote_server"
