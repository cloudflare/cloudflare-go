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

// ZoneLogpushValidateDestinationExistService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneLogpushValidateDestinationExistService] method instead.
type ZoneLogpushValidateDestinationExistService struct {
	Options []option.RequestOption
}

// NewZoneLogpushValidateDestinationExistService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneLogpushValidateDestinationExistService(opts ...option.RequestOption) (r *ZoneLogpushValidateDestinationExistService) {
	r = &ZoneLogpushValidateDestinationExistService{}
	r.Options = opts
	return
}

// Checks if there is an existing job with a destination.
func (r *ZoneLogpushValidateDestinationExistService) PostZonesZoneIdentifierLogpushValidateDestinationExists(ctx context.Context, zoneIdentifier string, body ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams, opts ...option.RequestOption) (res *DestinationExistsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/validate/destination/exists", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r ZoneLogpushValidateDestinationExistPostZonesZoneIdentifierLogpushValidateDestinationExistsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
