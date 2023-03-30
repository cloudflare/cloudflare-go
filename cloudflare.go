package cloudflare

import (
	"os"

	"github.com/cloudflare/cloudflare-sdk-go/options"
	"github.com/cloudflare/cloudflare-sdk-go/services"
)

type Cloudflare struct {
	Options []options.RequestOption
	Zones   *services.ZoneService
}

// NewCloudflare generates a new client with the default options read from the
// environment ("CLOUDFLARE_API_KEY"). The options passed in as arguments are
// applied after these default arguments, and all options will be passed down to
// the services and requests that this client makes.
func NewCloudflare(opts ...options.RequestOption) (r *Cloudflare) {
	defaults := []options.RequestOption{options.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CLOUDFLARE_API_KEY"); ok {
		defaults = append(defaults, options.WithAPIKey(o))
	}
	opts = append(defaults, opts...)

	r = &Cloudflare{Options: opts}

	r.Zones = services.NewZoneService(opts...)

	return
}
