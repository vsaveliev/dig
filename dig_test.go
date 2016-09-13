package main

import (
	"testing"
)

func TestExtactNameServers(t *testing.T) {
	nameservers, err := extractNameServersFromDig("openprovider.nl")

	if err != nil {
		t.Fatalf("Extracting error: %s", err)
	}

	if nameservers == nil {
		t.Fatalf("Nameservers are empty")
	}
}
