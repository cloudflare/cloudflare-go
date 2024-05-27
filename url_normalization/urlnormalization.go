// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_normalization

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// URLNormalizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLNormalizationService] method instead.
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

// Updates the URL normalization settings.
func (r *URLNormalizationService) Update(ctx context.Context, params URLNormalizationUpdateParams, opts ...option.RequestOption) (res *URLNormalizationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Fetches the current URL normalization settings.
func (r *URLNormalizationService) Get(ctx context.Context, query URLNormalizationGetParams, opts ...option.RequestOption) (res *URLNormalizationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/url_normalization", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type URLNormalizationUpdateResponse struct {
	// The scope of the URL normalization.
	Scope string `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type string                             `json:"type"`
	JSON urlNormalizationUpdateResponseJSON `json:"-"`
}

// urlNormalizationUpdateResponseJSON contains the JSON metadata for the struct
// [URLNormalizationUpdateResponse]
type urlNormalizationUpdateResponseJSON struct {
	Scope       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLNormalizationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlNormalizationUpdateResponseJSON) RawJSON() string {
	return r.raw
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

func (r urlNormalizationGetResponseJSON) RawJSON() string {
	return r.raw
}

type URLNormalizationUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The scope of the URL normalization.
	Scope param.Field[string] `json:"scope"`
	// The type of URL normalization performed by Cloudflare.
	Type param.Field[string] `json:"type"`
}

func (r URLNormalizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type URLNormalizationGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}
