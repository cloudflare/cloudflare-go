package cloudflare_test

import (
	"fmt"

	cloudflare "github.com/cloudflare/cloudflare-go"
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
	records, err := api.ListDNSRecords(zoneID, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, r := range records.Result {
		fmt.Printf("%s: %s\n", r.Name, r.Content)
	}
}
