package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var (
	yamlFile       = "cidrs.yaml"                       // Stores YAML filename (default 'cidrs.yaml')
	apiToken       = os.Getenv("PHPIPAMTOKEN")          // Env variable that stores API token
	sectionId      = "1"                                // phpipam section ID
	masterSubnetId = "231"                              // phpipam master subnet ID
	apiURL         = "https://<some-domain>/api/netops" // Base API URL
)

type CIDR struct {
	Name string `yaml:"Name"`
	Mask int    `yaml:"Mask"`
}

// Checks for error
func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// Processes program flags
func processFlags() {
	flag.StringVar(&yamlFile, "f", yamlFile, "Specify the YAML filename.")
	flag.StringVar(&sectionId, "s", sectionId, "Section identifier - mandatory on add method.")
	flag.StringVar(&masterSubnetId, "m", masterSubnetId, "Master subnet id for nested subnet.")
	flag.StringVar(&apiURL, "u", apiURL, "API base URL.")
	flag.Parse()
}

// Checks to see if API token available
func checkAPIToken(s string) error {
	if s == "" {
		// return fmt.Errorf("PHPIPAMTOKEN env variable not found or empty")
		return fmt.Errorf("PHPIPAMTOKEN env variable not found or empty")
	}
	return nil
}

// Checks sectionId and masterSubnetId flags to ensure a number was passed in
func checkIDFlags(s, m string) error {
	msg := "Passed in Section Identifier or Master Subnet Id is not a number."
	if _, err := strconv.Atoi(s); err != nil {
		return fmt.Errorf(msg)
	} else if _, err = strconv.Atoi(m); err != nil {
		return fmt.Errorf(msg)
	}
	return nil
}

// Checks to see if any CIDRs found in yaml file
func checkCIDRs(m map[string][]CIDR) error {
	if len(m) == 0 {
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
	err = checkIDFlags(sectionId, masterSubnetId)
	checkError(err)
	fmt.Printf("Attempting to open '%v'..\n", yamlFile)
	f, err := os.ReadFile(yamlFile)
	checkError(err)
	fmt.Printf("'%v' file successfully loaded!\n", yamlFile)
	cidrs := make(map[string][]CIDR)
	err = yaml.Unmarshal(f, cidrs) // Unmarshal yaml file
	checkError(err)
	err = checkCIDRs(cidrs)
	checkError(err)
	for _, cidr := range cidrs["CIDRs"] {
		addCIDR(cidr)
		time.Sleep(time.Second) // Sleep a second after each API call
	}
	fmt.Println("Execution complete!")
}
