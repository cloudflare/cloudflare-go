// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"os"

	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the cloudflare API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options []option.RequestOption
	Zones   *ZoneService
	AI      *AIService
}

// NewClient generates a new client with the default option read from the
// environment (CLOUDFLARE_API_KEY, CLOUDFLARE_EMAIL). The option passed in as
// arguments are applied after these default arguments, and all option will be
// passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("CLOUDFLARE_EMAIL"); ok {
		defaults = append(defaults, option.WithEmail(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Zones = NewZoneService(opts...)
	r.AI = NewAIService(opts...)

	return
}
