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
	domainOrIP := parseArgs()
	routes := traceroute(domainOrIP)
	printCountries(routes)
}

func printCountries(tracerouteOut string) {
	ipRegex := regexp.MustCompile(`\((.*?)\)`)
	lines := strings.Split(tracerouteOut, "\n")
	for _, line := range lines {
		ip := ipRegex.FindStringSubmatch(line)
		if ip != nil {
			fmt.Printf("%s /%s/\n", line, ipdb.GetCountry(ip[1]))
		} else {
			fmt.Println(line)
		}
	}
}

func traceroute(domainOrIP string) string {
	// Linux only (because it spawns `traceroute`)
	// Note: using --icmp in order to avoid firewalls blocking high UDP ports
	// Note: `traceroute --icmp` requires root privileges
	cmd := exec.Command("sudo", "traceroute", "--icmp", "-w 3", "-q 1", "-m 16", domainOrIP)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}

func parseArgs() string {
	target := "example.com"
	if len(os.Args) == 2 {
		target = os.Args[1]
	}
	return target
}
