package Twilio

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"log"
	"os"
	"time"
)

// UpdateIpAccessList updates the IP address in the Twilio SIP IP Access Control List
func UpdateIpAccessList(address string) error {
	client := twilio.NewRestClient()

	// Retrieve the SIP IP Access Control List SID from environment variables
	ipAccessControlListSid := os.Getenv("TWILIO_IP_LIST_SID")
	if ipAccessControlListSid == "" {
		return fmt.Errorf("TWILIO_IP_LIST_SID environment variable is not set")
	}

	// Prepare the parameters for updating the SIP IP Address
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	friendlyName := fmt.Sprintf("Updated by IP Updater - %s", currentTime)
	cidrPrefixLength := 32
	params := &openapi.UpdateSipIpAddressParams{
		IpAddress:        &address,
		FriendlyName:     &friendlyName,
		CidrPrefixLength: &cidrPrefixLength,
	}

	// Retrieve the list of SIP IP addresses
	ipAddresses, err := client.Api.ListSipIpAddress(ipAccessControlListSid, nil)
	if err != nil {
		return fmt.Errorf("failed to list SIP IP addresses: %v", err)
	}

	// Check if there are any IP addresses and get the first one
	if len(ipAddresses) == 0 {
		return fmt.Errorf("no SIP IP addresses found")
	}
	latestIpAddress := ipAddresses[0]

	// Update the SIP IP Address
	resp, err := client.Api.UpdateSipIpAddress(ipAccessControlListSid, *latestIpAddress.Sid, params)
	if err != nil {
		return fmt.Errorf("failed to update SIP IP Address: %v", err)
	}

	log.Printf("Updated SIP IP Address: %s, Friendly Name: %s\n", *resp.IpAddress, *resp.FriendlyName)
	return nil
}
