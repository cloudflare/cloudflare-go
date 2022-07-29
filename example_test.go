package cloudflare_test

import (
	"context"
	"fmt"

	cloudflare "github.com/js0cha/cloudflare-go-bm"
)

const (
	user   = "cloudflare@example.org"
	domain = "example.com"
	apiKey = "deadbeef"
)

func Example() {
	api, err := cloudflare.New("deadbeef", "cloudflare@example.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch the zone ID for zone example.org
	zoneID, err := api.ZoneIDByName("example.org")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Fetch all DNS records for example.org
	records, err := api.DNSRecords(context.Background(), zoneID, cloudflare.DNSRecord{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range records {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}
