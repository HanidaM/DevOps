
allowedIP=""

if [ -z "$allowedIP" ]; then
  echo "SET IP"
  exit 1
fi

iptables -F

iptables -P INPUT DROP

iptables -A INPUT -s "$allowedIP" -j ACCEPT

iptables-save > /etc/iptables/rules.v4

echo "Firewall rules have been configured. Only allowing access from $allowedIP."
