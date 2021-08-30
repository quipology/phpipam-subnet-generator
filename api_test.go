package main

import "testing"

func TestCheckResponseCode(t *testing.T) {
	c := CIDR{
		Name: "Test",
		Mask: 24,
	}
	m1 := map[string]interface{}{"code": float64(201)}
	m2 := map[string]interface{}{"code": float64(500)}
	_, err := checkResponseCode(c, m1)
	if err != nil {
		t.Errorf("Expected nil, but received %v\n", err)
	}
	_, err = checkResponseCode(c, m2)
	if err == nil {
		t.Errorf("Expected error, but received %v\n", err)
	}
}
