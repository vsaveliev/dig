package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// to do nothing
}

type Nameserver struct {
	Domain string
	IpV4   string
	IpV6   string
}

type Nameservers []Nameserver

func extractNameServersFromDig(domain string) (Nameservers, error) {
	cmd := exec.Command("dig", domain, "NS", "+short")
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, fmt.Errorf("Exec error: %s", err)
	}

	err = cmd.Start()
	if err != nil {
		return nil, fmt.Errorf("Exec error: %s", err)
	}
	defer cmd.Wait()

	nameservers := Nameservers{}

	scanner := bufio.NewScanner(out)
	for scanner.Scan() {
		nameserver := Nameserver{}

		s := scanner.Text()
		nameserver.Domain = s[:len(s)-1]

		if strings.HasSuffix(nameserver.Domain, "."+domain) {
			// it's a glue record = nameserver is subdomain of current domain
			// only in this case we need to get ip of nameserver
			ipv4, _ := exec.Command("dig", nameserver.Domain, "A", "+short").Output()
			nameserver.IpV4 = strings.TrimRight(string(ipv4), "\n")

			ipv6, _ := exec.Command("dig", nameserver.Domain, "AAAA", "+short").Output()
			nameserver.IpV6 = strings.TrimRight(string(ipv6), "\n")
		}

		nameservers = append(nameservers, nameserver)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nameservers, nil
}
