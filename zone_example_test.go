package cloudflare_test

import (
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

func ExampleAPI_ListZones_all() {
	api, err := cloudflare.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch a slice of all zones available to this account.
	zones, err := api.ListZones(1)
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones.Result {
		fmt.Println(z.Name)
	}
}

func ExampleAPI_ListZones_filter() {
	api, err := cloudflare.New("deadbeef", "test@example.org")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch a slice of all zones available to this account.
	zones, err := api.ListZones(1)
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones.Result {
		fmt.Println(z.Name)
	}
}
