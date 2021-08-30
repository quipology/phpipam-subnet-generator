package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type apiCIDR struct {
	Mask        int    `json:"mask"`
	Description string `json:"description"`
}

const contentType = "application/json"

func addCIDR(c CIDR) {
	fmt.Printf("Attempting to create CIDR for '%v' with a '/%v' mask..\n", c.Name, c.Mask)
	h := make(http.Header)
	h.Add("phpipam-token", apiToken) // Add API token to header
	h.Add("Content-Type", contentType)
	a := apiCIDR{
		Description: c.Name,
		Mask:        c.Mask,
	}
	jsonBytes, err := json.Marshal(&a)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := strings.NewReader(string(jsonBytes))
	req, err := http.NewRequest("POST", fmt.Sprintf("%v/subnets/%v/first_subnet/%v/", apiURL, masterSubnetId, c.Mask), r) // Create new request
	if err != nil {
		fmt.Println(err) // Print error to Stdout
		return
	}
	req.Header = h // Update the request with custom header
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	http.DefaultClient.Transport = tr
	resp, err := http.DefaultClient.Do(req) // Execute the request
	if err != nil {
		fmt.Println(err) // Print error to Stdout
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body) // Read response body
	if err != nil {
		fmt.Println(err)
		return
	}
	responseMap := make(map[string]interface{})
	err = json.Unmarshal(body, &responseMap) // Unmarshal response
	if err != nil {
		fmt.Println(err)
		return
	}
	s, err := checkResponseCode(c, responseMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}

// Checks the response code
func checkResponseCode(c CIDR, r map[string]interface{}) (string, error) {
	switch r["code"] {
	case float64(200):
		return fmt.Sprintf("%v for '%v'", r["message"], c.Name), nil
	case float64(201):
		return fmt.Sprintf("%v %v for '%v'", r["data"], r["message"], c.Name), nil
	default:
		return "", fmt.Errorf("Some weird response code received: %v", r["code"])
	}
}
