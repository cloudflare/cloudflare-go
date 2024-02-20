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
func (r *URLNormalizationService) Get(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *URLNormalizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the URL normalization settings.
func (r *URLNormalizationService) Replace(ctx context.Context, zoneID string, body URLNormalizationReplaceParams, opts ...option.RequestOption) (res *URLNormalizationReplaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/url_normalization", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type URLNormalizationGetResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                          `json:"type"`
	JSON urlNormalizationGetResponseJSON `json:"-"`
}

// urlNormalizationGetResponseJSON contains the JSON metadata for the struct
// [URLNormalizationGetResponse]
type urlNormalizationGetResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLNormalizationReplaceResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                              `json:"type"`
	JSON urlNormalizationReplaceResponseJSON `json:"-"`
}

// urlNormalizationReplaceResponseJSON contains the JSON metadata for the struct
// [URLNormalizationReplaceResponse]
type urlNormalizationReplaceResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationReplaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type URLNormalizationReplaceParams struct {
	// The scope of the URL normalization.
	Scope param.Field[string] `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[string] `json:"type"`
}

func (r URLNormalizationReplaceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
