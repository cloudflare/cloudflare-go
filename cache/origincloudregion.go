// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// OriginCloudRegionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOriginCloudRegionService] method instead.
type OriginCloudRegionService struct {
	Options []option.RequestOption
}

// NewOriginCloudRegionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOriginCloudRegionService(opts ...option.RequestOption) (r *OriginCloudRegionService) {
	r = &OriginCloudRegionService{}
	r.Options = opts
	return
}

// Creates a new IP-to-cloud-region mapping or replaces the existing mapping for
// the specified IP. PUT is idempotent — calling it repeatedly with the same body
// produces the same result. The IP path parameter is normalized to canonical form
// (RFC 5952 for IPv6) before storage. The vendor and region are validated against
// the list from `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`.
// Returns 400 if the `origin_ip` in the body does not match the URL path
// parameter. Returns 403 (code 1164) when the zone has reached the limit of 3,500
// IP mappings.
func (r *OriginCloudRegionService) Update(ctx context.Context, originIP string, params OriginCloudRegionUpdateParams, opts ...option.RequestOption) (res *OriginCloudRegion, err error) {
	var env OriginCloudRegionUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if originIP == "" {
		err = errors.New("missing required origin_ip parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/%s", params.ZoneID, originIP)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns all IP-to-cloud-region mappings configured for the zone with pagination
// support. Each mapping tells Cloudflare which cloud vendor and region hosts the
// origin at that IP, enabling the edge to route via the nearest Tiered Cache
// upper-tier co-located with that cloud provider. Returns an empty array when no
// mappings exist.
func (r *OriginCloudRegionService) List(ctx context.Context, params OriginCloudRegionListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[OriginCloudRegion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns all IP-to-cloud-region mappings configured for the zone with pagination
// support. Each mapping tells Cloudflare which cloud vendor and region hosts the
// origin at that IP, enabling the edge to route via the nearest Tiered Cache
// upper-tier co-located with that cloud provider. Returns an empty array when no
// mappings exist.
func (r *OriginCloudRegionService) ListAutoPaging(ctx context.Context, params OriginCloudRegionListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[OriginCloudRegion] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Removes the cloud region mapping for a single origin IP address. The IP path
// parameter is normalized before lookup. Returns the deleted IP on success.
// Returns 404 if no mapping exists for the specified IP. When the last mapping for
// the zone is removed the underlying rule record is also deleted.
func (r *OriginCloudRegionService) Delete(ctx context.Context, originIP string, body OriginCloudRegionDeleteParams, opts ...option.RequestOption) (res *OriginCloudRegionDeleteResponse, err error) {
	var env OriginCloudRegionDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if originIP == "" {
		err = errors.New("missing required origin_ip parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/%s", body.ZoneID, originIP)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes up to 100 IP-to-cloud-region mappings in a single request. Each IP is
// validated independently — successfully deleted items are returned in the
// `succeeded` array and IPs that could not be found or are invalid are returned in
// the `failed` array.
func (r *OriginCloudRegionService) BulkDelete(ctx context.Context, body OriginCloudRegionBulkDeleteParams, opts ...option.RequestOption) (res *OriginCloudRegionBulkDeleteResponse, err error) {
	var env OriginCloudRegionBulkDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/batch", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Upserts up to 100 IP-to-cloud-region mappings in a single request. Items in the
// request body are created or replaced; mappings not included in the request body
// are preserved unchanged (this is a merge operation, not a full collection
// replacement). Each item is validated independently — valid items are applied and
// invalid items are returned in the `failed` array. The vendor and region for
// every item are validated against the list from
// `GET /zones/{zone_id}/origin/cloud_regions/supported_regions`.
func (r *OriginCloudRegionService) BulkUpdate(ctx context.Context, params OriginCloudRegionBulkUpdateParams, opts ...option.RequestOption) (res *OriginCloudRegionBulkUpdateResponse, err error) {
	var env OriginCloudRegionBulkUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/batch", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the cloud region mapping for a single origin IP address. The IP path
// parameter is normalized before lookup (RFC 5952 for IPv6). Returns 404 if the
// zone has no mappings or if the specified IP has no mapping.
func (r *OriginCloudRegionService) Get(ctx context.Context, originIP string, query OriginCloudRegionGetParams, opts ...option.RequestOption) (res *OriginCloudRegion, err error) {
	var env OriginCloudRegionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if originIP == "" {
		err = errors.New("missing required origin_ip parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/%s", query.ZoneID, originIP)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the cloud vendors and regions that are valid values for origin cloud
// region mappings. Each region includes the Tiered Cache upper-tier colocation
// codes that will be used for cache routing when a mapping targeting that region
// is active. Requires the zone to have Tiered Cache enabled.
func (r *OriginCloudRegionService) SupportedRegions(ctx context.Context, query OriginCloudRegionSupportedRegionsParams, opts ...option.RequestOption) (res *OriginCloudRegionSupportedRegionsResponse, err error) {
	var env OriginCloudRegionSupportedRegionsResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/origin/cloud_regions/supported_regions", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A single origin IP-to-cloud-region mapping.
type OriginCloudRegion struct {
	// The origin IP address (IPv4 or IPv6). Normalized to canonical form (RFC 5952 for
	// IPv6).
	OriginIP string `json:"origin_ip" api:"required"`
	// Cloud vendor region identifier.
	Region string `json:"region" api:"required"`
	// Cloud vendor hosting the origin.
	Vendor OriginCloudRegionVendor `json:"vendor" api:"required"`
	// Time this mapping was last modified.
	ModifiedOn time.Time             `json:"modified_on" format:"date-time"`
	JSON       originCloudRegionJSON `json:"-"`
}

// originCloudRegionJSON contains the JSON metadata for the struct
// [OriginCloudRegion]
type originCloudRegionJSON struct {
	OriginIP    apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionJSON) RawJSON() string {
	return r.raw
}

// Cloud vendor hosting the origin.
type OriginCloudRegionVendor string

const (
	OriginCloudRegionVendorAws   OriginCloudRegionVendor = "aws"
	OriginCloudRegionVendorAzure OriginCloudRegionVendor = "azure"
	OriginCloudRegionVendorGcp   OriginCloudRegionVendor = "gcp"
	OriginCloudRegionVendorOci   OriginCloudRegionVendor = "oci"
)

func (r OriginCloudRegionVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionVendorAws, OriginCloudRegionVendorAzure, OriginCloudRegionVendorGcp, OriginCloudRegionVendorOci:
		return true
	}
	return false
}

// Response result for a delete operation. Identifies the deleted mapping.
type OriginCloudRegionDeleteResponse struct {
	// The origin IP address whose mapping was deleted.
	OriginIP string                              `json:"origin_ip" api:"required"`
	JSON     originCloudRegionDeleteResponseJSON `json:"-"`
}

// originCloudRegionDeleteResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionDeleteResponse]
type originCloudRegionDeleteResponseJSON struct {
	OriginIP    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Response result for a batch origin cloud region operation.
type OriginCloudRegionBulkDeleteResponse struct {
	// Items that could not be applied, with error details.
	Failed []OriginCloudRegionBulkDeleteResponseFailed `json:"failed" api:"required"`
	// Items that were successfully applied.
	Succeeded []OriginCloudRegionBulkDeleteResponseSucceeded `json:"succeeded" api:"required"`
	JSON      originCloudRegionBulkDeleteResponseJSON        `json:"-"`
}

// originCloudRegionBulkDeleteResponseJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkDeleteResponse]
type originCloudRegionBulkDeleteResponseJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkDeleteResponseFailed struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin_ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items (the new value for
	// upsert, the deleted value for delete).
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items (the new value for upsert,
	// the deleted value for delete).
	Vendor string                                        `json:"vendor"`
	JSON   originCloudRegionBulkDeleteResponseFailedJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseFailedJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkDeleteResponseFailed]
type originCloudRegionBulkDeleteResponseFailedJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseFailedJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkDeleteResponseSucceeded struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin_ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items (the new value for
	// upsert, the deleted value for delete).
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items (the new value for upsert,
	// the deleted value for delete).
	Vendor string                                           `json:"vendor"`
	JSON   originCloudRegionBulkDeleteResponseSucceededJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseSucceededJSON contains the JSON metadata for
// the struct [OriginCloudRegionBulkDeleteResponseSucceeded]
type originCloudRegionBulkDeleteResponseSucceededJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseSucceeded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseSucceededJSON) RawJSON() string {
	return r.raw
}

// Response result for a batch origin cloud region operation.
type OriginCloudRegionBulkUpdateResponse struct {
	// Items that could not be applied, with error details.
	Failed []OriginCloudRegionBulkUpdateResponseFailed `json:"failed" api:"required"`
	// Items that were successfully applied.
	Succeeded []OriginCloudRegionBulkUpdateResponseSucceeded `json:"succeeded" api:"required"`
	JSON      originCloudRegionBulkUpdateResponseJSON        `json:"-"`
}

// originCloudRegionBulkUpdateResponseJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkUpdateResponse]
type originCloudRegionBulkUpdateResponseJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkUpdateResponseFailed struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin_ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items (the new value for
	// upsert, the deleted value for delete).
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items (the new value for upsert,
	// the deleted value for delete).
	Vendor string                                        `json:"vendor"`
	JSON   originCloudRegionBulkUpdateResponseFailedJSON `json:"-"`
}

// originCloudRegionBulkUpdateResponseFailedJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkUpdateResponseFailed]
type originCloudRegionBulkUpdateResponseFailedJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkUpdateResponseFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkUpdateResponseFailedJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkUpdateResponseSucceeded struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin_ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items (the new value for
	// upsert, the deleted value for delete).
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items (the new value for upsert,
	// the deleted value for delete).
	Vendor string                                           `json:"vendor"`
	JSON   originCloudRegionBulkUpdateResponseSucceededJSON `json:"-"`
}

// originCloudRegionBulkUpdateResponseSucceededJSON contains the JSON metadata for
// the struct [OriginCloudRegionBulkUpdateResponseSucceeded]
type originCloudRegionBulkUpdateResponseSucceededJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkUpdateResponseSucceeded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkUpdateResponseSucceededJSON) RawJSON() string {
	return r.raw
}

// Cloud vendors and their supported regions for origin cloud region mappings.
type OriginCloudRegionSupportedRegionsResponse struct {
	// Whether Cloudflare airport codes (IATA colo identifiers) were successfully
	// resolved for the `upper_tier_colos` field on each region. When `false`, the
	// `upper_tier_colos` arrays may be empty or incomplete.
	ObtainedCodes bool `json:"obtained_codes" api:"required"`
	// Map of vendor name to list of supported regions.
	Vendors map[string][]OriginCloudRegionSupportedRegionsResponseVendor `json:"vendors" api:"required"`
	JSON    originCloudRegionSupportedRegionsResponseJSON                `json:"-"`
}

// originCloudRegionSupportedRegionsResponseJSON contains the JSON metadata for the
// struct [OriginCloudRegionSupportedRegionsResponse]
type originCloudRegionSupportedRegionsResponseJSON struct {
	ObtainedCodes apijson.Field
	Vendors       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OriginCloudRegionSupportedRegionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionSupportedRegionsResponseJSON) RawJSON() string {
	return r.raw
}

// A single supported cloud region with associated Tiered Cache upper-tier
// colocations.
type OriginCloudRegionSupportedRegionsResponseVendor struct {
	// Cloud vendor region identifier.
	Name string `json:"name" api:"required"`
	// Cloudflare Tiered Cache upper-tier colocation codes co-located with this cloud
	// region. Requests from zones with a matching origin mapping will be routed
	// through these colos.
	UpperTierColos []string                                            `json:"upper_tier_colos" api:"required"`
	JSON           originCloudRegionSupportedRegionsResponseVendorJSON `json:"-"`
}

// originCloudRegionSupportedRegionsResponseVendorJSON contains the JSON metadata
// for the struct [OriginCloudRegionSupportedRegionsResponseVendor]
type originCloudRegionSupportedRegionsResponseVendorJSON struct {
	Name           apijson.Field
	UpperTierColos apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OriginCloudRegionSupportedRegionsResponseVendor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionSupportedRegionsResponseVendorJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Origin IP address (IPv4 or IPv6). For the single PUT endpoint
	// (`PUT /origin/cloud_regions/{origin_ip}`), this field must match the path
	// parameter or the request will be rejected with a 400 error. For the batch PUT
	// endpoint, this field identifies which mapping to upsert.
	OriginIP param.Field[string] `json:"origin_ip" api:"required"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor
	// as returned by the supported_regions endpoint.
	Region param.Field[string] `json:"region" api:"required"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors.
	Vendor param.Field[OriginCloudRegionUpdateParamsVendor] `json:"vendor" api:"required"`
}

func (r OriginCloudRegionUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud vendor hosting the origin. Must be one of the supported vendors.
type OriginCloudRegionUpdateParamsVendor string

const (
	OriginCloudRegionUpdateParamsVendorAws   OriginCloudRegionUpdateParamsVendor = "aws"
	OriginCloudRegionUpdateParamsVendorAzure OriginCloudRegionUpdateParamsVendor = "azure"
	OriginCloudRegionUpdateParamsVendorGcp   OriginCloudRegionUpdateParamsVendor = "gcp"
	OriginCloudRegionUpdateParamsVendorOci   OriginCloudRegionUpdateParamsVendor = "oci"
)

func (r OriginCloudRegionUpdateParamsVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionUpdateParamsVendorAws, OriginCloudRegionUpdateParamsVendorAzure, OriginCloudRegionUpdateParamsVendorGcp, OriginCloudRegionUpdateParamsVendorOci:
		return true
	}
	return false
}

type OriginCloudRegionUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// A single origin IP-to-cloud-region mapping.
	Result OriginCloudRegion                           `json:"result"`
	JSON   originCloudRegionUpdateResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionUpdateResponseEnvelope]
type originCloudRegionUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionUpdateResponseEnvelopeSuccess bool

const (
	OriginCloudRegionUpdateResponseEnvelopeSuccessTrue OriginCloudRegionUpdateResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Number of items per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [OriginCloudRegionListParams]'s query parameters as
// `url.Values`.
func (r OriginCloudRegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OriginCloudRegionDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginCloudRegionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a delete operation. Identifies the deleted mapping.
	Result OriginCloudRegionDeleteResponse             `json:"result"`
	JSON   originCloudRegionDeleteResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionDeleteResponseEnvelope]
type originCloudRegionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionDeleteResponseEnvelopeSuccess bool

const (
	OriginCloudRegionDeleteResponseEnvelopeSuccessTrue OriginCloudRegionDeleteResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionBulkDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginCloudRegionBulkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionBulkDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a batch origin cloud region operation.
	Result OriginCloudRegionBulkDeleteResponse             `json:"result"`
	JSON   originCloudRegionBulkDeleteResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginCloudRegionBulkDeleteResponseEnvelope]
type originCloudRegionBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionBulkDeleteResponseEnvelopeSuccess bool

const (
	OriginCloudRegionBulkDeleteResponseEnvelopeSuccessTrue OriginCloudRegionBulkDeleteResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionBulkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionBulkUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string]                     `path:"zone_id" api:"required"`
	Body   []OriginCloudRegionBulkUpdateParamsBody `json:"body" api:"required"`
}

func (r OriginCloudRegionBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or replacing an origin cloud region mapping.
type OriginCloudRegionBulkUpdateParamsBody struct {
	// Origin IP address (IPv4 or IPv6). For the single PUT endpoint
	// (`PUT /origin/cloud_regions/{origin_ip}`), this field must match the path
	// parameter or the request will be rejected with a 400 error. For the batch PUT
	// endpoint, this field identifies which mapping to upsert.
	OriginIP param.Field[string] `json:"origin_ip" api:"required"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor
	// as returned by the supported_regions endpoint.
	Region param.Field[string] `json:"region" api:"required"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors.
	Vendor param.Field[OriginCloudRegionBulkUpdateParamsBodyVendor] `json:"vendor" api:"required"`
}

func (r OriginCloudRegionBulkUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud vendor hosting the origin. Must be one of the supported vendors.
type OriginCloudRegionBulkUpdateParamsBodyVendor string

const (
	OriginCloudRegionBulkUpdateParamsBodyVendorAws   OriginCloudRegionBulkUpdateParamsBodyVendor = "aws"
	OriginCloudRegionBulkUpdateParamsBodyVendorAzure OriginCloudRegionBulkUpdateParamsBodyVendor = "azure"
	OriginCloudRegionBulkUpdateParamsBodyVendorGcp   OriginCloudRegionBulkUpdateParamsBodyVendor = "gcp"
	OriginCloudRegionBulkUpdateParamsBodyVendorOci   OriginCloudRegionBulkUpdateParamsBodyVendor = "oci"
)

func (r OriginCloudRegionBulkUpdateParamsBodyVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkUpdateParamsBodyVendorAws, OriginCloudRegionBulkUpdateParamsBodyVendorAzure, OriginCloudRegionBulkUpdateParamsBodyVendorGcp, OriginCloudRegionBulkUpdateParamsBodyVendorOci:
		return true
	}
	return false
}

type OriginCloudRegionBulkUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionBulkUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a batch origin cloud region operation.
	Result OriginCloudRegionBulkUpdateResponse             `json:"result"`
	JSON   originCloudRegionBulkUpdateResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionBulkUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [OriginCloudRegionBulkUpdateResponseEnvelope]
type originCloudRegionBulkUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionBulkUpdateResponseEnvelopeSuccess bool

const (
	OriginCloudRegionBulkUpdateResponseEnvelopeSuccessTrue OriginCloudRegionBulkUpdateResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionBulkUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginCloudRegionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// A single origin IP-to-cloud-region mapping.
	Result OriginCloudRegion                        `json:"result"`
	JSON   originCloudRegionGetResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionGetResponseEnvelope]
type originCloudRegionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionGetResponseEnvelopeSuccess bool

const (
	OriginCloudRegionGetResponseEnvelopeSuccessTrue OriginCloudRegionGetResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionSupportedRegionsParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginCloudRegionSupportedRegionsResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionSupportedRegionsResponseEnvelopeSuccess `json:"success" api:"required"`
	// Cloud vendors and their supported regions for origin cloud region mappings.
	Result OriginCloudRegionSupportedRegionsResponse             `json:"result"`
	JSON   originCloudRegionSupportedRegionsResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionSupportedRegionsResponseEnvelopeJSON contains the JSON metadata
// for the struct [OriginCloudRegionSupportedRegionsResponseEnvelope]
type originCloudRegionSupportedRegionsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionSupportedRegionsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionSupportedRegionsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionSupportedRegionsResponseEnvelopeSuccess bool

const (
	OriginCloudRegionSupportedRegionsResponseEnvelopeSuccessTrue OriginCloudRegionSupportedRegionsResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionSupportedRegionsResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionSupportedRegionsResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
