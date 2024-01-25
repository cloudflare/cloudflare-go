// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogpushService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogpushService] method
// instead.
type ZoneLogpushService struct {
	Options    []option.RequestOption
	Datasets   *ZoneLogpushDatasetService
	Edges      *ZoneLogpushEdgeService
	Jobs       *ZoneLogpushJobService
	Ownerships *ZoneLogpushOwnershipService
	Validates  *ZoneLogpushValidateService
}

// NewZoneLogpushService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneLogpushService(opts ...option.RequestOption) (r *ZoneLogpushService) {
	r = &ZoneLogpushService{}
	r.Options = opts
	r.Datasets = NewZoneLogpushDatasetService(opts...)
	r.Edges = NewZoneLogpushEdgeService(opts...)
	r.Jobs = NewZoneLogpushJobService(opts...)
	r.Ownerships = NewZoneLogpushOwnershipService(opts...)
	r.Validates = NewZoneLogpushValidateService(opts...)
	return
}
