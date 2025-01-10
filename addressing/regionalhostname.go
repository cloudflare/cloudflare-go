// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// RegionalHostnameService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalHostnameService] method instead.
type RegionalHostnameService struct {
	Options []option.RequestOption
	Regions *RegionalHostnameRegionService
}

// NewRegionalHostnameService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalHostnameService(opts ...option.RequestOption) (r *RegionalHostnameService) {
	r = &RegionalHostnameService{}
	r.Options = opts
	r.Regions = NewRegionalHostnameRegionService(opts...)
	return
}

// Create a new Regional Hostname entry. Cloudflare will only use data centers that
// are physically located within the chosen region to decrypt and service HTTPS
// traffic. Learn more about
// [Regional Services](https://developers.cloudflare.com/data-localization/regional-services/get-started/).
func (r *RegionalHostnameService) New(ctx context.Context, params RegionalHostnameNewParams, opts ...option.RequestOption) (res *RegionalHostnameNewResponse, err error) {
	var env RegionalHostnameNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Regional Hostnames within a zone.
func (r *RegionalHostnameService) List(ctx context.Context, query RegionalHostnameListParams, opts ...option.RequestOption) (res *pagination.SinglePage[RegionalHostnameListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all Regional Hostnames within a zone.
func (r *RegionalHostnameService) ListAutoPaging(ctx context.Context, query RegionalHostnameListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[RegionalHostnameListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete the region configuration for a specific Regional Hostname.
func (r *RegionalHostnameService) Delete(ctx context.Context, hostname string, body RegionalHostnameDeleteParams, opts ...option.RequestOption) (res *RegionalHostnameDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", body.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Update the configuration for a specific Regional Hostname. Only the region_key
// of a hostname is mutable.
func (r *RegionalHostnameService) Edit(ctx context.Context, hostname string, params RegionalHostnameEditParams, opts ...option.RequestOption) (res *RegionalHostnameEditResponse, err error) {
	var env RegionalHostnameEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", params.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the configuration for a specific Regional Hostname, within a zone.
func (r *RegionalHostnameService) Get(ctx context.Context, hostname string, query RegionalHostnameGetParams, opts ...option.RequestOption) (res *RegionalHostnameGetResponse, err error) {
	var env RegionalHostnameGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/addressing/regional_hostnames/%s", query.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RegionalHostnameNewResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string                          `json:"region_key,required"`
	JSON      regionalHostnameNewResponseJSON `json:"-"`
}

// regionalHostnameNewResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameNewResponse]
type regionalHostnameNewResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameListResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string                           `json:"region_key,required"`
	JSON      regionalHostnameListResponseJSON `json:"-"`
}

// regionalHostnameListResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameListResponse]
type regionalHostnameListResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameListResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameDeleteResponse struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalHostnameDeleteResponseSuccess `json:"success,required"`
	JSON    regionalHostnameDeleteResponseJSON    `json:"-"`
}

// regionalHostnameDeleteResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameDeleteResponse]
type regionalHostnameDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalHostnameDeleteResponseSuccess bool

const (
	RegionalHostnameDeleteResponseSuccessTrue RegionalHostnameDeleteResponseSuccess = true
)

func (r RegionalHostnameDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameEditResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string                           `json:"region_key,required"`
	JSON      regionalHostnameEditResponseJSON `json:"-"`
}

// regionalHostnameEditResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameEditResponse]
type regionalHostnameEditResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameGetResponse struct {
	// When the regional hostname was created
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname string `json:"hostname,required"`
	// Identifying key for the region
	RegionKey string                          `json:"region_key,required"`
	JSON      regionalHostnameGetResponseJSON `json:"-"`
}

// regionalHostnameGetResponseJSON contains the JSON metadata for the struct
// [RegionalHostnameGetResponse]
type regionalHostnameGetResponseJSON struct {
	CreatedOn   apijson.Field
	Hostname    apijson.Field
	RegionKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseJSON) RawJSON() string {
	return r.raw
}

type RegionalHostnameNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// DNS hostname to be regionalized, must be a subdomain of the zone. Wildcards are
	// supported for one level, e.g `*.example.com`
	Hostname param.Field[string] `json:"hostname,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
}

func (r RegionalHostnameNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalHostnameNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalHostnameNewResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameNewResponse                `json:"result"`
	JSON    regionalHostnameNewResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameNewResponseEnvelope]
type regionalHostnameNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalHostnameNewResponseEnvelopeSuccess bool

const (
	RegionalHostnameNewResponseEnvelopeSuccessTrue RegionalHostnameNewResponseEnvelopeSuccess = true
)

func (r RegionalHostnameNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Identifying key for the region
	RegionKey param.Field[string] `json:"region_key,required"`
}

func (r RegionalHostnameEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RegionalHostnameEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalHostnameEditResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameEditResponse                `json:"result"`
	JSON    regionalHostnameEditResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameEditResponseEnvelope]
type regionalHostnameEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalHostnameEditResponseEnvelopeSuccess bool

const (
	RegionalHostnameEditResponseEnvelopeSuccessTrue RegionalHostnameEditResponseEnvelopeSuccess = true
)

func (r RegionalHostnameEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalHostnameGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalHostnameGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalHostnameGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalHostnameGetResponse                `json:"result"`
	JSON    regionalHostnameGetResponseEnvelopeJSON    `json:"-"`
}

// regionalHostnameGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalHostnameGetResponseEnvelope]
type regionalHostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalHostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalHostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalHostnameGetResponseEnvelopeSuccess bool

const (
	RegionalHostnameGetResponseEnvelopeSuccessTrue RegionalHostnameGetResponseEnvelopeSuccess = true
)

func (r RegionalHostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalHostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
