package cloudflare_test

import (
	"context"
	"fmt"
	"log"

	cloudflare "github.com/cloudflare/cloudflare-go"
)

var exampleNewRateLimit = cloudflare.RateLimit{
	Description: "test",
	Match: cloudflare.RateLimitTrafficMatcher{
		Request: cloudflare.RateLimitRequestMatcher{
			URLPattern: "exampledomain.com/test-rate-limit",
		},
	},
	Threshold: 0,
	Period:    0,
	Action: cloudflare.RateLimitAction{
		Mode:    "ban",
		Timeout: 60,
	},
	Correlate: &cloudflare.RateLimitCorrelate{
		By: "nat",
	},
}

func ExampleAPI_CreateRateLimit() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimit, err := api.CreateRateLimit(context.Background(), zoneID, exampleNewRateLimit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimit)
}

func ExampleAPI_ListRateLimits() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimits, _, err := api.ListRateLimits(context.Background(), cloudflare.ZoneIdentifier(zoneID), cloudflare.RateLimitListParams{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
	for _, r := range rateLimits {
		fmt.Printf("%+v\n", r)
	}
}

func ExampleAPI_RateLimit() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimits, err := api.RateLimit(context.Background(), cloudflare.ZoneIdentifier(zoneID), "my_rate_limit_id")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
}

func ExampleAPI_DeleteRateLimit() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	err = api.DeleteRateLimit(context.Background(), zoneID, "my_rate_limit_id")
	if err != nil {
		log.Fatal(err)
	}
}
