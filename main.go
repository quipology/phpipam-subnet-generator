package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var (
	yamlFile       = "cidrs.yaml"                       // Stores YAML filename (default 'cidrs.yaml')
	apiToken       = os.Getenv("PHPIPAMTOKEN")          // Env variable that stores API token
	masterSubnetId = 231                                // phpipam master subnet ID
	apiURL         = "https://<some_domain>/api/netops" // Base API URL
)

type CIDR struct {
	Name string `yaml:"Name"`
	Mask int    `yaml:"Mask"`
}

// Checks for error
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}

// Processes program flags
func processFlags() {
	flag.StringVar(&yamlFile, "f", yamlFile, "Specify the YAML filename.")
	flag.StringVar(&apiURL, "u", apiURL, "API base URL.")
	flag.IntVar(&masterSubnetId, "m", masterSubnetId, "Master subnet id for nested subnet.")
	flag.Parse()
}

// Checks to see if API token available
func checkAPIToken(s string) error {
	if s == "" {
		return fmt.Errorf("PHPIPAMTOKEN env variable not found or empty")
	}
	return nil
}

// Checks to see if any CIDRs found in yaml file
func checkCIDRs(m map[string][]CIDR) error {
	if len(m["CIDRs"]) == 0 {
		return fmt.Errorf("No CIDRs found in '%v' file.", yamlFile)
	}
	return nil
}

func main() {
	processFlags()
	fmt.Println("Attempting to get PHPIPAM Token..")
	err := checkAPIToken(apiToken)
	checkError(err)
	fmt.Println("PHPIPAM Token found!")
	fmt.Printf("Attempting to open '%v'..\n", yamlFile)
	f, err := os.ReadFile(yamlFile)
	checkError(err)
	fmt.Printf("'%v' file successfully loaded!\n", yamlFile)
	cidrs := make(map[string][]CIDR)
	err = yaml.Unmarshal(f, &cidrs) // Unmarshal yaml file
	checkError(err)
	err = checkCIDRs(cidrs)
	checkError(err)
	for _, cidr := range cidrs["CIDRs"] {
		createSubnet(cidr)
		time.Sleep(time.Second) // Sleep a second after each API call
	}
	fmt.Println("Execution complete!")
}
