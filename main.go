package main

import (
	"IPSync/Cloudflare"
	"IPSync/Common"
	"IPSync/Twilio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// GetPublicIP fetches and returns the current internet-facing IP address
func GetPublicIP() (string, error) {
	resp, err := http.Get(os.Getenv("GET_IP_QUERY_URL"))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}

func main() {
	ipStore := "current_ip.txt"

	// Ensure the IP store file exists
	err := Common.CreateFileIfNotExist(ipStore)
	if err != nil {
		log.Fatalf("Failed to create IP store file: %v", err)
	}

	// Read the last known IP address
	lastKnownIp, err := Common.ReadFromFile(ipStore)
	if err != nil {
		log.Fatalf("Failed to read IP store file: %v", err)
	}

	// Get the current public IP address
	newIpAddress, err := GetPublicIP()
	if err != nil {
		log.Fatalf("Failed to get public IP address from %s: %v", os.Getenv("GET_IP_QUERY_URL"), err)
	}

	// Check if the IP address has changed
	if newIpAddress == lastKnownIp {
		log.Println("IP address has not changed.")
		return
	}

	log.Println("IP address has changed.")
	err = Common.WriteToFile(ipStore, newIpAddress)
	if err != nil {
		log.Fatalf("Failed to write new IP address to store file: %v", err)
	}

	// Update IP access list in Twilio
	err = Twilio.UpdateIpAccessList(newIpAddress)
	if err != nil {
		log.Fatalf("Failed to update IP access list in Twilio: %v", err)
	}

	// Update DNS record in Cloudflare
	err = Cloudflare.UpdateDNSRecord(os.Getenv("DOMAIN_NAME"), newIpAddress)
	if err != nil {
		log.Fatalf("Failed to update DNS record in Cloudflare: %v", err)
	}
}
