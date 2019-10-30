Enhances traceroute with country info. Linux only, currently.

Requires root privileges due to the traceroute flags used (using ICMP packets in
order to avoid firewalls blocking high UDP ports)

```
$ go run main.go <ip_or_domain>
```
