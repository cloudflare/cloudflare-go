// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/shared"
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

// Adds a single IP-to-cloud-region mapping for the zone. The IP must be a valid
// IPv4 or IPv6 address and is normalized to canonical form before storage (RFC
// 5952 for IPv6). Returns 400 (code 1145) if a mapping for that IP already exists
// — use PATCH to update an existing entry. The vendor and region are validated
// against the list from
// `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.
func (r *OriginCloudRegionService) New(ctx context.Context, params OriginCloudRegionNewParams, opts ...option.RequestOption) (res *OriginCloudRegionNewResponse, err error) {
	var env OriginCloudRegionNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns all IP-to-cloud-region mappings configured for the zone. Each mapping
// tells Cloudflare which cloud vendor and region hosts the origin at that IP,
// enabling the edge to route via the nearest Tiered Cache upper-tier co-located
// with that cloud provider. Returns an empty array when no mappings exist.
func (r *OriginCloudRegionService) List(ctx context.Context, query OriginCloudRegionListParams, opts ...option.RequestOption) (res *OriginCloudRegionListResponse, err error) {
	var env OriginCloudRegionListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Removes the cloud region mapping for a single origin IP address. The IP path
// parameter is normalized before lookup. Returns the deleted entry on success.
// Returns 404 (code 1163) if no mapping exists for the specified IP. When the last
// mapping for the zone is removed the underlying rule record is also deleted.
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
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions/%s", body.ZoneID, originIP)
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
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions/batch", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Adds or updates up to 100 IP-to-cloud-region mappings in a single request. Each
// item is validated independently — valid items are applied and invalid items are
// returned in the `failed` array. The vendor and region for every item are
// validated against the list from
// `GET /zones/{zone_id}/cache/origin_cloud_regions/supported_regions`.
func (r *OriginCloudRegionService) BulkEdit(ctx context.Context, params OriginCloudRegionBulkEditParams, opts ...option.RequestOption) (res *OriginCloudRegionBulkEditResponse, err error) {
	var env OriginCloudRegionBulkEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions/batch", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Adds or updates a single IP-to-cloud-region mapping for the zone. Unlike POST,
// this operation is idempotent — if a mapping for the IP already exists it is
// overwritten. Returns the complete updated list of all mappings for the zone.
// Returns 403 (code 1164) when the zone has reached the limit of 3,500 IP
// mappings.
func (r *OriginCloudRegionService) Edit(ctx context.Context, params OriginCloudRegionEditParams, opts ...option.RequestOption) (res *OriginCloudRegionEditResponse, err error) {
	var env OriginCloudRegionEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Returns the cloud region mapping for a single origin IP address. The IP path
// parameter is normalized before lookup (RFC 5952 for IPv6). Returns 404
// (code 1142) if the zone has no mappings or if the specified IP has no mapping.
func (r *OriginCloudRegionService) Get(ctx context.Context, originIP string, query OriginCloudRegionGetParams, opts ...option.RequestOption) (res *OriginCloudRegionGetResponse, err error) {
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
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions/%s", query.ZoneID, originIP)
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
	path := fmt.Sprintf("zones/%s/cache/origin_cloud_regions/supported_regions", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A single origin IP-to-cloud-region mapping.
type OriginCloudRegion struct {
	// The origin IP address (IPv4 or IPv6, canonicalized).
	OriginIP string `json:"origin-ip" api:"required"`
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

// Response result for a single origin cloud region mapping.
type OriginCloudRegionNewResponse struct {
	ID OriginCloudRegionNewResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool `json:"editable" api:"required"`
	// A single origin IP-to-cloud-region mapping.
	Value OriginCloudRegion `json:"value" api:"required"`
	// Time the mapping was last modified.
	ModifiedOn time.Time                        `json:"modified_on" format:"date-time"`
	JSON       originCloudRegionNewResponseJSON `json:"-"`
}

// originCloudRegionNewResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionNewResponse]
type originCloudRegionNewResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionNewResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionNewResponseID string

const (
	OriginCloudRegionNewResponseIDOriginPublicCloudRegion OriginCloudRegionNewResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionNewResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionNewResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

// Response result for a list of origin cloud region mappings.
type OriginCloudRegionListResponse struct {
	ID OriginCloudRegionListResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool                `json:"editable" api:"required"`
	Value    []OriginCloudRegion `json:"value" api:"required"`
	// Time the mapping set was last modified. Null when no mappings exist.
	ModifiedOn time.Time                         `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originCloudRegionListResponseJSON `json:"-"`
}

// originCloudRegionListResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionListResponse]
type originCloudRegionListResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionListResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionListResponseID string

const (
	OriginCloudRegionListResponseIDOriginPublicCloudRegion OriginCloudRegionListResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionListResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionListResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

// Response result for a single origin cloud region mapping.
type OriginCloudRegionDeleteResponse struct {
	ID OriginCloudRegionDeleteResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool `json:"editable" api:"required"`
	// A single origin IP-to-cloud-region mapping.
	Value OriginCloudRegion `json:"value" api:"required"`
	// Time the mapping was last modified.
	ModifiedOn time.Time                           `json:"modified_on" format:"date-time"`
	JSON       originCloudRegionDeleteResponseJSON `json:"-"`
}

// originCloudRegionDeleteResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionDeleteResponse]
type originCloudRegionDeleteResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionDeleteResponseID string

const (
	OriginCloudRegionDeleteResponseIDOriginPublicCloudRegion OriginCloudRegionDeleteResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionDeleteResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionDeleteResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

// Response result for a batch origin cloud region operation.
type OriginCloudRegionBulkDeleteResponse struct {
	ID OriginCloudRegionBulkDeleteResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool                                     `json:"editable" api:"required"`
	Value    OriginCloudRegionBulkDeleteResponseValue `json:"value" api:"required"`
	// Time the mapping set was last modified. Null when no items were successfully
	// applied.
	ModifiedOn time.Time                               `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originCloudRegionBulkDeleteResponseJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkDeleteResponse]
type originCloudRegionBulkDeleteResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionBulkDeleteResponseID string

const (
	OriginCloudRegionBulkDeleteResponseIDOriginPublicCloudRegion OriginCloudRegionBulkDeleteResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionBulkDeleteResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkDeleteResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

type OriginCloudRegionBulkDeleteResponseValue struct {
	// Items that could not be applied, with error details.
	Failed []OriginCloudRegionBulkDeleteResponseValueFailed `json:"failed" api:"required"`
	// Items that were successfully applied.
	Succeeded []OriginCloudRegionBulkDeleteResponseValueSucceeded `json:"succeeded" api:"required"`
	JSON      originCloudRegionBulkDeleteResponseValueJSON        `json:"-"`
}

// originCloudRegionBulkDeleteResponseValueJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkDeleteResponseValue]
type originCloudRegionBulkDeleteResponseValueJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseValueJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkDeleteResponseValueFailed struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin-ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items for patch operations.
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items for patch operations.
	Vendor string                                             `json:"vendor"`
	JSON   originCloudRegionBulkDeleteResponseValueFailedJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseValueFailedJSON contains the JSON metadata
// for the struct [OriginCloudRegionBulkDeleteResponseValueFailed]
type originCloudRegionBulkDeleteResponseValueFailedJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseValueFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseValueFailedJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkDeleteResponseValueSucceeded struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin-ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items for patch operations.
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items for patch operations.
	Vendor string                                                `json:"vendor"`
	JSON   originCloudRegionBulkDeleteResponseValueSucceededJSON `json:"-"`
}

// originCloudRegionBulkDeleteResponseValueSucceededJSON contains the JSON metadata
// for the struct [OriginCloudRegionBulkDeleteResponseValueSucceeded]
type originCloudRegionBulkDeleteResponseValueSucceededJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkDeleteResponseValueSucceeded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkDeleteResponseValueSucceededJSON) RawJSON() string {
	return r.raw
}

// Response result for a batch origin cloud region operation.
type OriginCloudRegionBulkEditResponse struct {
	ID OriginCloudRegionBulkEditResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool                                   `json:"editable" api:"required"`
	Value    OriginCloudRegionBulkEditResponseValue `json:"value" api:"required"`
	// Time the mapping set was last modified. Null when no items were successfully
	// applied.
	ModifiedOn time.Time                             `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originCloudRegionBulkEditResponseJSON `json:"-"`
}

// originCloudRegionBulkEditResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionBulkEditResponse]
type originCloudRegionBulkEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkEditResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionBulkEditResponseID string

const (
	OriginCloudRegionBulkEditResponseIDOriginPublicCloudRegion OriginCloudRegionBulkEditResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionBulkEditResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkEditResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

type OriginCloudRegionBulkEditResponseValue struct {
	// Items that could not be applied, with error details.
	Failed []OriginCloudRegionBulkEditResponseValueFailed `json:"failed" api:"required"`
	// Items that were successfully applied.
	Succeeded []OriginCloudRegionBulkEditResponseValueSucceeded `json:"succeeded" api:"required"`
	JSON      originCloudRegionBulkEditResponseValueJSON        `json:"-"`
}

// originCloudRegionBulkEditResponseValueJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkEditResponseValue]
type originCloudRegionBulkEditResponseValueJSON struct {
	Failed      apijson.Field
	Succeeded   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkEditResponseValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkEditResponseValueJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkEditResponseValueFailed struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin-ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items for patch operations.
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items for patch operations.
	Vendor string                                           `json:"vendor"`
	JSON   originCloudRegionBulkEditResponseValueFailedJSON `json:"-"`
}

// originCloudRegionBulkEditResponseValueFailedJSON contains the JSON metadata for
// the struct [OriginCloudRegionBulkEditResponseValueFailed]
type originCloudRegionBulkEditResponseValueFailedJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkEditResponseValueFailed) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkEditResponseValueFailedJSON) RawJSON() string {
	return r.raw
}

// Result for a single item in a batch operation.
type OriginCloudRegionBulkEditResponseValueSucceeded struct {
	// The origin IP address for this item.
	OriginIP string `json:"origin-ip" api:"required"`
	// Error message explaining why the item failed. Present only on failed items.
	Error string `json:"error"`
	// Cloud vendor region identifier. Present on succeeded items for patch operations.
	Region string `json:"region"`
	// Cloud vendor identifier. Present on succeeded items for patch operations.
	Vendor string                                              `json:"vendor"`
	JSON   originCloudRegionBulkEditResponseValueSucceededJSON `json:"-"`
}

// originCloudRegionBulkEditResponseValueSucceededJSON contains the JSON metadata
// for the struct [OriginCloudRegionBulkEditResponseValueSucceeded]
type originCloudRegionBulkEditResponseValueSucceededJSON struct {
	OriginIP    apijson.Field
	Error       apijson.Field
	Region      apijson.Field
	Vendor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkEditResponseValueSucceeded) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkEditResponseValueSucceededJSON) RawJSON() string {
	return r.raw
}

// Response result for a list of origin cloud region mappings.
type OriginCloudRegionEditResponse struct {
	ID OriginCloudRegionEditResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool                `json:"editable" api:"required"`
	Value    []OriginCloudRegion `json:"value" api:"required"`
	// Time the mapping set was last modified. Null when no mappings exist.
	ModifiedOn time.Time                         `json:"modified_on" api:"nullable" format:"date-time"`
	JSON       originCloudRegionEditResponseJSON `json:"-"`
}

// originCloudRegionEditResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionEditResponse]
type originCloudRegionEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionEditResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionEditResponseID string

const (
	OriginCloudRegionEditResponseIDOriginPublicCloudRegion OriginCloudRegionEditResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionEditResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionEditResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
}

// Response result for a single origin cloud region mapping.
type OriginCloudRegionGetResponse struct {
	ID OriginCloudRegionGetResponseID `json:"id" api:"required"`
	// Whether the setting can be modified by the current user.
	Editable bool `json:"editable" api:"required"`
	// A single origin IP-to-cloud-region mapping.
	Value OriginCloudRegion `json:"value" api:"required"`
	// Time the mapping was last modified.
	ModifiedOn time.Time                        `json:"modified_on" format:"date-time"`
	JSON       originCloudRegionGetResponseJSON `json:"-"`
}

// originCloudRegionGetResponseJSON contains the JSON metadata for the struct
// [OriginCloudRegionGetResponse]
type originCloudRegionGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionGetResponseJSON) RawJSON() string {
	return r.raw
}

type OriginCloudRegionGetResponseID string

const (
	OriginCloudRegionGetResponseIDOriginPublicCloudRegion OriginCloudRegionGetResponseID = "origin_public_cloud_region"
)

func (r OriginCloudRegionGetResponseID) IsKnown() bool {
	switch r {
	case OriginCloudRegionGetResponseIDOriginPublicCloudRegion:
		return true
	}
	return false
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

type OriginCloudRegionNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Origin IP address (IPv4 or IPv6). Normalized to canonical form before storage
	// (RFC 5952 for IPv6).
	IP param.Field[string] `json:"ip" api:"required"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor
	// as returned by the supported_regions endpoint.
	Region param.Field[string] `json:"region" api:"required"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors.
	Vendor param.Field[OriginCloudRegionNewParamsVendor] `json:"vendor" api:"required"`
}

func (r OriginCloudRegionNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud vendor hosting the origin. Must be one of the supported vendors.
type OriginCloudRegionNewParamsVendor string

const (
	OriginCloudRegionNewParamsVendorAws   OriginCloudRegionNewParamsVendor = "aws"
	OriginCloudRegionNewParamsVendorAzure OriginCloudRegionNewParamsVendor = "azure"
	OriginCloudRegionNewParamsVendorGcp   OriginCloudRegionNewParamsVendor = "gcp"
	OriginCloudRegionNewParamsVendorOci   OriginCloudRegionNewParamsVendor = "oci"
)

func (r OriginCloudRegionNewParamsVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionNewParamsVendorAws, OriginCloudRegionNewParamsVendorAzure, OriginCloudRegionNewParamsVendorGcp, OriginCloudRegionNewParamsVendorOci:
		return true
	}
	return false
}

type OriginCloudRegionNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionNewResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a single origin cloud region mapping.
	Result OriginCloudRegionNewResponse             `json:"result"`
	JSON   originCloudRegionNewResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionNewResponseEnvelope]
type originCloudRegionNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionNewResponseEnvelopeSuccess bool

const (
	OriginCloudRegionNewResponseEnvelopeSuccessTrue OriginCloudRegionNewResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type OriginCloudRegionListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionListResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a list of origin cloud region mappings.
	Result OriginCloudRegionListResponse             `json:"result"`
	JSON   originCloudRegionListResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionListResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionListResponseEnvelope]
type originCloudRegionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionListResponseEnvelopeSuccess bool

const (
	OriginCloudRegionListResponseEnvelopeSuccessTrue OriginCloudRegionListResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
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
	// Response result for a single origin cloud region mapping.
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

type OriginCloudRegionBulkEditParams struct {
	// Identifier.
	ZoneID param.Field[string]                   `path:"zone_id" api:"required"`
	Body   []OriginCloudRegionBulkEditParamsBody `json:"body" api:"required"`
}

func (r OriginCloudRegionBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Request body for creating or updating an origin cloud region mapping.
type OriginCloudRegionBulkEditParamsBody struct {
	// Origin IP address (IPv4 or IPv6). Normalized to canonical form before storage
	// (RFC 5952 for IPv6).
	IP param.Field[string] `json:"ip" api:"required"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor
	// as returned by the supported_regions endpoint.
	Region param.Field[string] `json:"region" api:"required"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors.
	Vendor param.Field[OriginCloudRegionBulkEditParamsBodyVendor] `json:"vendor" api:"required"`
}

func (r OriginCloudRegionBulkEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud vendor hosting the origin. Must be one of the supported vendors.
type OriginCloudRegionBulkEditParamsBodyVendor string

const (
	OriginCloudRegionBulkEditParamsBodyVendorAws   OriginCloudRegionBulkEditParamsBodyVendor = "aws"
	OriginCloudRegionBulkEditParamsBodyVendorAzure OriginCloudRegionBulkEditParamsBodyVendor = "azure"
	OriginCloudRegionBulkEditParamsBodyVendorGcp   OriginCloudRegionBulkEditParamsBodyVendor = "gcp"
	OriginCloudRegionBulkEditParamsBodyVendorOci   OriginCloudRegionBulkEditParamsBodyVendor = "oci"
)

func (r OriginCloudRegionBulkEditParamsBodyVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkEditParamsBodyVendorAws, OriginCloudRegionBulkEditParamsBodyVendorAzure, OriginCloudRegionBulkEditParamsBodyVendorGcp, OriginCloudRegionBulkEditParamsBodyVendorOci:
		return true
	}
	return false
}

type OriginCloudRegionBulkEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionBulkEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a batch origin cloud region operation.
	Result OriginCloudRegionBulkEditResponse             `json:"result"`
	JSON   originCloudRegionBulkEditResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionBulkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionBulkEditResponseEnvelope]
type originCloudRegionBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionBulkEditResponseEnvelopeSuccess bool

const (
	OriginCloudRegionBulkEditResponseEnvelopeSuccessTrue OriginCloudRegionBulkEditResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OriginCloudRegionEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Origin IP address (IPv4 or IPv6). Normalized to canonical form before storage
	// (RFC 5952 for IPv6).
	IP param.Field[string] `json:"ip" api:"required"`
	// Cloud vendor region identifier. Must be a valid region for the specified vendor
	// as returned by the supported_regions endpoint.
	Region param.Field[string] `json:"region" api:"required"`
	// Cloud vendor hosting the origin. Must be one of the supported vendors.
	Vendor param.Field[OriginCloudRegionEditParamsVendor] `json:"vendor" api:"required"`
}

func (r OriginCloudRegionEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Cloud vendor hosting the origin. Must be one of the supported vendors.
type OriginCloudRegionEditParamsVendor string

const (
	OriginCloudRegionEditParamsVendorAws   OriginCloudRegionEditParamsVendor = "aws"
	OriginCloudRegionEditParamsVendorAzure OriginCloudRegionEditParamsVendor = "azure"
	OriginCloudRegionEditParamsVendorGcp   OriginCloudRegionEditParamsVendor = "gcp"
	OriginCloudRegionEditParamsVendorOci   OriginCloudRegionEditParamsVendor = "oci"
)

func (r OriginCloudRegionEditParamsVendor) IsKnown() bool {
	switch r {
	case OriginCloudRegionEditParamsVendorAws, OriginCloudRegionEditParamsVendorAzure, OriginCloudRegionEditParamsVendorGcp, OriginCloudRegionEditParamsVendorOci:
		return true
	}
	return false
}

type OriginCloudRegionEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success OriginCloudRegionEditResponseEnvelopeSuccess `json:"success" api:"required"`
	// Response result for a list of origin cloud region mappings.
	Result OriginCloudRegionEditResponse             `json:"result"`
	JSON   originCloudRegionEditResponseEnvelopeJSON `json:"-"`
}

// originCloudRegionEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [OriginCloudRegionEditResponseEnvelope]
type originCloudRegionEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OriginCloudRegionEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r originCloudRegionEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OriginCloudRegionEditResponseEnvelopeSuccess bool

const (
	OriginCloudRegionEditResponseEnvelopeSuccessTrue OriginCloudRegionEditResponseEnvelopeSuccess = true
)

func (r OriginCloudRegionEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OriginCloudRegionEditResponseEnvelopeSuccessTrue:
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
	// Response result for a single origin cloud region mapping.
	Result OriginCloudRegionGetResponse             `json:"result"`
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
