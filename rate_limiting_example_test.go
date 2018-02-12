package cloudflare_test

import (
	"fmt"
	cloudflare "github.com/cloudflare/cloudflare-go"
	"log"
)

const (
	apiKey = "deadbeef"
	user   = "test@example.org"
	domain = "example.com"
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

	rateLimit, err := api.CreateRateLimit(zoneID, exampleNewRateLimit)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimit)
}

func ExampleAPI_ListRateLimitsAll() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimits, err := api.ListRateLimits(zoneID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", rateLimits)
	for _, r := range rateLimits {
		fmt.Printf("%+v\n", r)
	}
}

func ExampleAPI_GetRateLimit() {
	api, err := cloudflare.New(apiKey, user)
	if err != nil {
		log.Fatal(err)
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		log.Fatal(err)
	}

	rateLimits, err := api.RateLimit(zoneID, "my_rate_limit_id")
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

	err = api.DeleteRateLimit(zoneID, "my_rate_limit_id")
	if err != nil {
		log.Fatal(err)
	}
}
