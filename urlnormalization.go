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

// URLNormalizationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewURLNormalizationService] method
// instead.
type URLNormalizationService struct {
	Options []option.RequestOption
}

// NewURLNormalizationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewURLNormalizationService(opts ...option.RequestOption) (r *URLNormalizationService) {
	r = &URLNormalizationService{}
	r.Options = opts
	return
}

// Fetches the current URL normalization settings.
func (r *URLNormalizationService) URLNormalizationGetURLNormalizationSettings(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *URLNormalizationURLNormalizationGetURLNormalizationSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL normalization settings.
func (r *URLNormalizationService) URLNormalizationUpdateURLNormalizationSettings(ctx context.Context, zoneID string, body URLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams, opts ...option.RequestOption) (res *URLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type URLNormalizationURLNormalizationGetURLNormalizationSettingsResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                                                                  `json:"type"`
	JSON urlNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON `json:"-"`
}

// urlNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON contains
// the JSON metadata for the struct
// [URLNormalizationURLNormalizationGetURLNormalizationSettingsResponse]
type urlNormalizationURLNormalizationGetURLNormalizationSettingsResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationURLNormalizationGetURLNormalizationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                                                                     `json:"type"`
	JSON urlNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON `json:"-"`
}

// urlNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON
// contains the JSON metadata for the struct
// [URLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse]
type urlNormalizationURLNormalizationUpdateURLNormalizationSettingsResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationURLNormalizationUpdateURLNormalizationSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams struct {
	// The scope of the URL normalization.
	Scope param.Field[string] `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[string] `json:"type"`
}

func (r URLNormalizationURLNormalizationUpdateURLNormalizationSettingsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
