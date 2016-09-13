package main

import (
	"testing"
	"fmt"
)

func TestExtactNameServers (t *testing.T) {
	nameservers, _ := extractNameServersFromDig("openprovider.nl")

	fmt.Printf("%#v", nameservers)
}
