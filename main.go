package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"./ipdb"
)

func main() {
	traceroute_target := parse_args()
	traceroute_out := traceroute(traceroute_target)
	print_traceroute_with_countries(traceroute_out)
}

func print_traceroute_with_countries(traceroute_out string) {
	ipdb.Init()
	ip_regex := regexp.MustCompile(`\((.*?)\)`)
	lines := strings.Split(traceroute_out, "\n")
	for i, line := range lines {
		if i == 0 {
			fmt.Println(line)
			continue
		}

		found_ip := ip_regex.FindStringSubmatch(line)
		if found_ip != nil {
			fmt.Printf("%s /%s/\n", line, ipdb.GetCountry(found_ip[1]))
		}
	}
}

func traceroute(domain_or_ip string) string {
	// Linux only (because it spawns `traceroute`)
	// Note: using --icmp in order to avoid firewalls blocking high UDP ports
	// Note: `traceroute --icmp` requires root privileges
	cmd := exec.Command("sudo", "traceroute", "--icmp", "-w 3", "-q 1", "-m 16", domain_or_ip)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func parse_args() string {
	trace_target := "example.com"
	if len(os.Args) == 2 {
		trace_target = os.Args[1]
	}
	return trace_target
}
