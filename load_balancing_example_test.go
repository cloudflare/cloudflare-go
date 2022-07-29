package cloudflare_test

import (
	context "context"
	"fmt"
	"log"

	cloudflare "github.com/js0cha/cloudflare-go-bm"
)

func ExampleAPI_ListLoadBalancers() {
	// Construct a new API object.
	api, err := cloudflare.New("deadbeef", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID.
	id, err := api.ZoneIDByName("example.com") // Assuming example.com exists in your Cloudflare account
	if err != nil {
		log.Fatal(err)
	}

	// List LBs configured in zone.
	lbList, err := api.ListLoadBalancers(context.Background(), id)
	if err != nil {
		log.Fatal(err)
	}

	for _, lb := range lbList {
		fmt.Println(lb)
	}
}

func ExampleAPI_PoolHealthDetails() {
	// Construct a new API object.
	api, err := cloudflare.New("deadbeef", "test@example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Fetch pool health details.
	healthInfo, err := api.PoolHealthDetails(context.Background(), "example-pool-id")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(healthInfo)
}
