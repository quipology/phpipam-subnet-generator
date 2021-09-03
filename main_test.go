package main

import "testing"

func TestCheckAPIToken(t *testing.T) {
	err := checkAPIToken("Test Func")
	if err != nil {
		t.Error("Expected nil, but received error.")
	}
}

func TestCheckCIDRs(t *testing.T) {
	m := make(map[string][]CIDR)
	err := checkCIDRs(m)
	if err == nil {
		t.Errorf("Expected a length of 0, but received a length of %v\n", len(m))
	}
}
