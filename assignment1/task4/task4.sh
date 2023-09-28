log_dir="/home/logs"

mkdir -p "$log_dir"

cat <<EOF > "$log_dir/myecomproject_logrotate.conf"
$log_dir/*.log {
    size 1M
    rotate 3
    missingok
    notifempty
    compress
    create 0644 $(whoami) $(id -gn)
    postrotate
        /bin/echo "\$(date) - Log rotated" >> /home/log_rotates.log
    endscript
}
EOF

touch /home/log_rotates.log

logrotate -s /var/lib/logrotate/status "$log_dir/myecomproject_logrotate.conf"

(crontab -l ; echo "0 0 * * * /usr/sbin/logrotate -f $log_dir/myecomproject_logrotate.conf") | crontab -

echo "Log rotation setup complete."