// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneLogpushOwnershipService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneLogpushOwnershipService]
// method instead.
type ZoneLogpushOwnershipService struct {
	Options   []option.RequestOption
	Validates *ZoneLogpushOwnershipValidateService
}

// NewZoneLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneLogpushOwnershipService(opts ...option.RequestOption) (r *ZoneLogpushOwnershipService) {
	r = &ZoneLogpushOwnershipService{}
	r.Options = opts
	r.Validates = NewZoneLogpushOwnershipValidateService(opts...)
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *ZoneLogpushOwnershipService) PostZonesZoneIdentifierLogpushOwnership(ctx context.Context, zoneIdentifier string, body ZoneLogpushOwnershipPostZonesZoneIdentifierLogpushOwnershipParams, opts ...option.RequestOption) (res *GetOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/ownership", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushOwnershipPostZonesZoneIdentifierLogpushOwnershipParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r ZoneLogpushOwnershipPostZonesZoneIdentifierLogpushOwnershipParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
