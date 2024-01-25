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

// ZoneURLNormalizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneURLNormalizationService]
// method instead.
type ZoneURLNormalizationService struct {
	Options []option.RequestOption
}

// NewZoneURLNormalizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneURLNormalizationService(opts ...option.RequestOption) (r *ZoneURLNormalizationService) {
	r = &ZoneURLNormalizationService{}
	r.Options = opts
	return
}

// Fetches the current URL normalization settings.
func (r *ZoneURLNormalizationService) URLNormalizationGetURLNormalizationSettings(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *ZoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL normalization settings.
func (r *ZoneURLNormalizationService) URLNormalizationUpdateURLNormalizationSettings(ctx context.Context, zoneID string, body ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams, opts ...option.RequestOption) (res *ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                                                                      `json:"type"`
	JSON zoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON `json:"-"`
}

// zoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON
// contains the JSON metadata for the struct
// [ZoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponse]
type zoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationURLNormalizationGetURLNormalizationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                                                                         `json:"type"`
	JSON zoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON `json:"-"`
}

// zoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON
// contains the JSON metadata for the struct
// [ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse]
type zoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams struct {
	// The scope of the URL normalization.
	Scope param.Field[string] `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[string] `json:"type"`
}

func (r ZoneURLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
