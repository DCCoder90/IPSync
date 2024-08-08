package Cloudflare

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"log"
	"os"
	"time"
)

// UpdateDNSRecord updates the DNS A record with a new IP address for a specified domain
func UpdateDNSRecord(domain string, newIP string) error {
	// Initialize the Cloudflare API
	api, err := cloudflare.NewWithAPIToken(os.Getenv("CLOUDFLARE_API_KEY"))
	if err != nil {
		return fmt.Errorf("failed to create Cloudflare API client: %w", err)
	}

	// Retrieve the zone ID using the domain name
	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		return fmt.Errorf("failed to get zone ID for domain %s: %w", domain, err)
	}
	log.Printf("Zone ID: %s\n", zoneID)

	// List DNS A records for the specified zone
	records, _, err := api.ListDNSRecords(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.ListDNSRecordsParams{Type: "A"})
	if err != nil {
		return fmt.Errorf("failed to list DNS records: %w", err)
	}

	// Check if there are any A records
	if len(records) == 0 {
		return fmt.Errorf("no A records found for domain: %s", domain)
	}

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	comment := fmt.Sprintf("Updated by IP Updater - %s", currentTime)

	// Update the A records with the new IP address
	for _, record := range records {
		rc := &cloudflare.ResourceContainer{
			Level:      cloudflare.ZoneRouteLevel,
			Identifier: zoneID,
			Type:       cloudflare.ZoneType,
		}
		params := cloudflare.UpdateDNSRecordParams{
			ID:      record.ID,
			Content: newIP,
			Comment: &comment,
		}

		_, err := api.UpdateDNSRecord(context.Background(), rc, params)
		if err != nil {
			return fmt.Errorf("failed to update DNS record for %s: %w", record.Name, err)
		}
		log.Printf("Updated DNS record %s to IP %s\n", record.Name, newIP)
	}

	return nil
}
