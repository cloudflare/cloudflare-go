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

// ZoneLogpushValidateOriginService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneLogpushValidateOriginService] method instead.
type ZoneLogpushValidateOriginService struct {
	Options []option.RequestOption
}

// NewZoneLogpushValidateOriginService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneLogpushValidateOriginService(opts ...option.RequestOption) (r *ZoneLogpushValidateOriginService) {
	r = &ZoneLogpushValidateOriginService{}
	r.Options = opts
	return
}

// Validates logpull origin with logpull_options.
func (r *ZoneLogpushValidateOriginService) PostZonesZoneIdentifierLogpushValidateOrigin(ctx context.Context, zoneIdentifier string, body ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams, opts ...option.RequestOption) (res *ValidateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/logpush/validate/origin", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams struct {
	// Configuration string. It specifies things like requested fields and timestamp
	// formats. If migrating from the logpull api, copy the url (full url or just the
	// query string) of your call here, and logpush will keep on making this call for
	// you, setting start and end times appropriately.
	LogpullOptions param.Field[string] `json:"logpull_options,required" format:"uri-reference"`
}

func (r ZoneLogpushValidateOriginPostZonesZoneIdentifierLogpushValidateOriginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
