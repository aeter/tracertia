Enhances traceroute with country info. Linux only, currently.

Requires root privileges due to the traceroute flags used.

```
$ go run main.go 192.168.0.1
[sudo] password for user:
traceroute to 192.168.0.1 (192.168.0.1), 16 hops max, 60 byte packets
 1  gateway (192.168.0.1)  5.243 ms /None/
```
