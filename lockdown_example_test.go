package cloudflare_test

import (
	"fmt"
	"log"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func ExampleAPI_ListZoneLockdowns_all() {
	api, err := cloudflare.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch all Zone Lockdown rules for a zone, by page.
	rules, err := api.ListZoneLockdowns(zoneID, 1)
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range rules.Result {
		fmt.Printf("%s: %s\n", strings.Join(r.URLs, ", "), r.Configurations)
	}
}

func ExampleAPI_CreateZoneLockdown() {
	api, err := cloudflare.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName("example.org")
	if err != nil {
		log.Fatal(err)
	}

	newZoneLockdown := cloudflare.ZoneLockdown{
		Description: "Test Zone Lockdown Rule",
		URLs: []string{
			"*.example.org/test",
		},
		Configurations: []cloudflare.ZoneLockdownConfig{
			cloudflare.ZoneLockdownConfig{
				Target: "ip",
				Value:  "127.0.0.1",
			},
		},
		Paused:   false,
		Priority: 1,
	}

	response, err := api.CreateZoneLockdown(zoneID, newZoneLockdown)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response: ", response)
}
