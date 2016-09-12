package main

import (
	"fmt"
	"os/exec"
)

func extractNameServersFromDig(domain string) ([]map[string]string, error) {
//	nameservers = make([]map[string]string, 1)

	out, err := exec.Command("dig", domain, "NS", "+short").Output()
	if err != nil {
		return nil, fmt.Errorf("Exec error: %s", err)
	}

	fmt.Printf("%s", out)

	return nil, nil
}
