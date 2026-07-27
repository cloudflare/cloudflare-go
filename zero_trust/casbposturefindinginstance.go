// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// CasbPostureFindingInstanceService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCasbPostureFindingInstanceService] method instead.
type CasbPostureFindingInstanceService struct {
	Options []option.RequestOption
}

// NewCasbPostureFindingInstanceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewCasbPostureFindingInstanceService(opts ...option.RequestOption) (r *CasbPostureFindingInstanceService) {
	r = &CasbPostureFindingInstanceService{}
	r.Options = opts
	return
}

// Lists all security finding instances for a given security finding.
func (r *CasbPostureFindingInstanceService) List(ctx context.Context, findingID string, params CasbPostureFindingInstanceListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[CasbPostureFindingInstanceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/instances", params.AccountID, findingID)
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

// Lists all security finding instances for a given security finding.
func (r *CasbPostureFindingInstanceService) ListAutoPaging(ctx context.Context, findingID string, params CasbPostureFindingInstanceListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[CasbPostureFindingInstanceListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, findingID, params, opts...))
}

// Archive one or more finding instances.
func (r *CasbPostureFindingInstanceService) Archive(ctx context.Context, findingID string, params CasbPostureFindingInstanceArchiveParams, opts ...option.RequestOption) (res *CasbPostureFindingInstanceArchiveResponse, err error) {
	var env CasbPostureFindingInstanceArchiveResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/instances/archive", params.AccountID, findingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Creates a CSV export for Finding instances and accepts optional filters in the
// payload.
//
// The `storage_namespace_id` path parameter is derived from the finding ID by
// base64-decoding it (which yields `integration_id:finding_type_id`) and replacing
// the colon with a hyphen.
func (r *CasbPostureFindingInstanceService) Export(ctx context.Context, storageNamespaceID string, params CasbPostureFindingInstanceExportParams, opts ...option.RequestOption) (res *CasbPostureFindingInstanceExportResponse, err error) {
	var env CasbPostureFindingInstanceExportResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if storageNamespaceID == "" {
		err = errors.New("missing required storage_namespace_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/instances/export", params.AccountID, storageNamespaceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Gets a security Finding instance by id.
func (r *CasbPostureFindingInstanceService) Get(ctx context.Context, findingID string, instanceID string, query CasbPostureFindingInstanceGetParams, opts ...option.RequestOption) (res *CasbPostureFindingInstanceGetResponse, err error) {
	var env CasbPostureFindingInstanceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	if instanceID == "" {
		err = errors.New("missing required instance_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/instances/%s", query.AccountID, findingID, instanceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Remove the archive marking from one or more finding instances.
func (r *CasbPostureFindingInstanceService) Unarchive(ctx context.Context, findingID string, params CasbPostureFindingInstanceUnarchiveParams, opts ...option.RequestOption) (res *CasbPostureFindingInstanceUnarchiveResponse, err error) {
	var env CasbPostureFindingInstanceUnarchiveResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if findingID == "" {
		err = errors.New("missing required finding_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/data-security/posture/findings/%s/instances/unarchive", params.AccountID, findingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A specific instance of a security finding. In the API interface, we refer to the
// 'finding' table in our DB as finding instances, optimized for the p99 use case.
type CasbPostureFindingInstanceListResponse struct {
	// When this specific instance was identified.
	AfflictionDate time.Time `json:"affliction_date" api:"required" format:"date-time"`
	// Asset information including metadata and categorization.
	Asset CasbPostureFindingInstanceListResponseAsset `json:"asset" api:"required"`
	// DLP context information if this is a content finding.
	DLPContexts []CasbPostureFindingInstanceListResponseDLPContext `json:"dlp_contexts" api:"required"`
	// A list of the 10 most recent remediation jobs for this finding instance, ordered
	// by creation time (most recent first). The 'stale' field indicates whether the
	// remediation job was created before the finding instance's affliction_date (true)
	// or after it (false). If there has never been a remediation job for this finding
	// instance, this field will be an empty array.
	Remediations []CasbPostureFindingInstanceListResponseRemediation `json:"remediations" api:"required"`
	// The most recent webhook job invocation for each webhook configuration associated
	// with this finding instance. Each entry represents the latest job (any status)
	// per webhook config. The 'stale' field indicates whether the job was invoked
	// before the finding instance's current affliction_date. If no webhook jobs have
	// been created, this field will be an empty array.
	Webhooks []CasbPostureFindingInstanceListResponseWebhook `json:"webhooks" api:"required"`
	// Unique identifier for the finding instance.
	ID string `json:"id" format:"uuid"`
	// Whether this finding instance has been archived.
	IsArchived bool                                       `json:"is_archived"`
	JSON       casbPostureFindingInstanceListResponseJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingInstanceListResponse]
type casbPostureFindingInstanceListResponseJSON struct {
	AfflictionDate apijson.Field
	Asset          apijson.Field
	DLPContexts    apijson.Field
	Remediations   apijson.Field
	Webhooks       apijson.Field
	ID             apijson.Field
	IsArchived     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseJSON) RawJSON() string {
	return r.raw
}

// Asset information including metadata and categorization.
type CasbPostureFindingInstanceListResponseAsset struct {
	// Category information for an asset.
	Category CasbPostureFindingInstanceListResponseAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// The fields associated with the asset.
	Fields []CasbPostureFindingInstanceListResponseAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Unique identifier for the asset.
	ID string `json:"id" format:"uuid"`
	// Direct link to the asset.
	Link string                                          `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceListResponseAssetJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseAssetJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceListResponseAsset]
type casbPostureFindingInstanceListResponseAssetJSON struct {
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for an asset.
type CasbPostureFindingInstanceListResponseAssetCategory struct {
	// The specific service within the vendor the asset is part of (often none).
	// Example - AWS is the vendor, S3 is the service.
	Service string `json:"service" api:"required,nullable"`
	// The type of asset.
	Type string `json:"type" api:"required"`
	// The vendor the asset is part of.
	Vendor string `json:"vendor" api:"required"`
	// Unique identifier for the asset category.
	ID   string                                                  `json:"id" format:"uuid"`
	JSON casbPostureFindingInstanceListResponseAssetCategoryJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseAssetCategoryJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceListResponseAssetCategory]
type casbPostureFindingInstanceListResponseAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Additional field information for an asset.
type CasbPostureFindingInstanceListResponseAssetField struct {
	// The name of the field.
	Name string `json:"name" api:"required"`
	// The value of the field.
	Value string `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                               `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceListResponseAssetFieldJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseAssetFieldJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceListResponseAssetField]
type casbPostureFindingInstanceListResponseAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseAssetFieldJSON) RawJSON() string {
	return r.raw
}

// DLP context information for a finding.
type CasbPostureFindingInstanceListResponseDLPContext struct {
	// When the DLP context was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// DLP Entry IDs.
	EntryIDs []string `json:"entry_ids" api:"required" format:"uuid"`
	// DLP Profile ID.
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// When the DLP context was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Unique identifier for the DLP context.
	ID string `json:"id" format:"uuid"`
	// When the DLP context was deleted.
	Deleted time.Time `json:"deleted" api:"nullable" format:"date-time"`
	// DLP Right Boundary of match context.
	MatchContextMaxExtent int64 `json:"match_context_max_extent" api:"nullable"`
	// DLP Left Boundary of match context.
	MatchContextMinExtent int64 `json:"match_context_min_extent" api:"nullable"`
	// DLP Match context payload that matched the profile in question.
	MatchContextPayload map[string]interface{}                               `json:"match_context_payload" api:"nullable"`
	JSON                casbPostureFindingInstanceListResponseDLPContextJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseDLPContextJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceListResponseDLPContext]
type casbPostureFindingInstanceListResponseDLPContextJSON struct {
	Created               apijson.Field
	EntryIDs              apijson.Field
	ProfileID             apijson.Field
	Updated               apijson.Field
	ID                    apijson.Field
	Deleted               apijson.Field
	MatchContextMaxExtent apijson.Field
	MatchContextMinExtent apijson.Field
	MatchContextPayload   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseDLPContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseDLPContextJSON) RawJSON() string {
	return r.raw
}

// Summary information about a remediation job.
type CasbPostureFindingInstanceListResponseRemediation struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this remediation job is stale (created before the finding instance's
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Status of a remediation job.
	Status CasbPostureFindingInstanceListResponseRemediationsStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceListResponseRemediationJSON    `json:"-"`
}

// casbPostureFindingInstanceListResponseRemediationJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceListResponseRemediation]
type casbPostureFindingInstanceListResponseRemediationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseRemediationJSON) RawJSON() string {
	return r.raw
}

// Status of a remediation job.
type CasbPostureFindingInstanceListResponseRemediationsStatus string

const (
	CasbPostureFindingInstanceListResponseRemediationsStatusPending    CasbPostureFindingInstanceListResponseRemediationsStatus = "pending"
	CasbPostureFindingInstanceListResponseRemediationsStatusProcessing CasbPostureFindingInstanceListResponseRemediationsStatus = "processing"
	CasbPostureFindingInstanceListResponseRemediationsStatusCompleted  CasbPostureFindingInstanceListResponseRemediationsStatus = "completed"
	CasbPostureFindingInstanceListResponseRemediationsStatusFailed     CasbPostureFindingInstanceListResponseRemediationsStatus = "failed"
	CasbPostureFindingInstanceListResponseRemediationsStatusValidating CasbPostureFindingInstanceListResponseRemediationsStatus = "validating"
)

func (r CasbPostureFindingInstanceListResponseRemediationsStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceListResponseRemediationsStatusPending, CasbPostureFindingInstanceListResponseRemediationsStatusProcessing, CasbPostureFindingInstanceListResponseRemediationsStatusCompleted, CasbPostureFindingInstanceListResponseRemediationsStatusFailed, CasbPostureFindingInstanceListResponseRemediationsStatusValidating:
		return true
	}
	return false
}

// Summary of the most recent webhook job invocation for a specific webhook
// configuration.
type CasbPostureFindingInstanceListResponseWebhook struct {
	// The most recent webhook job for this webhook configuration.
	LatestJob CasbPostureFindingInstanceListResponseWebhooksLatestJob `json:"latest_job" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookID string `json:"webhook_id" api:"required" format:"uuid"`
	// Account-specified display label for the webhook configuration.
	WebhookLabel string                                            `json:"webhook_label" api:"required"`
	JSON         casbPostureFindingInstanceListResponseWebhookJSON `json:"-"`
}

// casbPostureFindingInstanceListResponseWebhookJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceListResponseWebhook]
type casbPostureFindingInstanceListResponseWebhookJSON struct {
	LatestJob    apijson.Field
	WebhookID    apijson.Field
	WebhookLabel apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseWebhookJSON) RawJSON() string {
	return r.raw
}

// The most recent webhook job for this webhook configuration.
type CasbPostureFindingInstanceListResponseWebhooksLatestJob struct {
	// Unique identifier for the webhook job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the webhook job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this webhook job is stale (created before the finding instance's current
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Current status of the webhook job.
	Status CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceListResponseWebhooksLatestJobJSON   `json:"-"`
}

// casbPostureFindingInstanceListResponseWebhooksLatestJobJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceListResponseWebhooksLatestJob]
type casbPostureFindingInstanceListResponseWebhooksLatestJobJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceListResponseWebhooksLatestJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceListResponseWebhooksLatestJobJSON) RawJSON() string {
	return r.raw
}

// Current status of the webhook job.
type CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus string

const (
	CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusPending    CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus = "pending"
	CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusProcessing CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus = "processing"
	CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusCompleted  CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus = "completed"
)

func (r CasbPostureFindingInstanceListResponseWebhooksLatestJobStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusPending, CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusProcessing, CasbPostureFindingInstanceListResponseWebhooksLatestJobStatusCompleted:
		return true
	}
	return false
}

// A specific instance of a security finding. In the API interface, we refer to the
// 'finding' table in our DB as finding instances, optimized for the p99 use case.
type CasbPostureFindingInstanceArchiveResponse struct {
	// When this specific instance was identified.
	AfflictionDate time.Time `json:"affliction_date" api:"required" format:"date-time"`
	// Asset information including metadata and categorization.
	Asset CasbPostureFindingInstanceArchiveResponseAsset `json:"asset" api:"required"`
	// DLP context information if this is a content finding.
	DLPContexts []CasbPostureFindingInstanceArchiveResponseDLPContext `json:"dlp_contexts" api:"required"`
	// A list of the 10 most recent remediation jobs for this finding instance, ordered
	// by creation time (most recent first). The 'stale' field indicates whether the
	// remediation job was created before the finding instance's affliction_date (true)
	// or after it (false). If there has never been a remediation job for this finding
	// instance, this field will be an empty array.
	Remediations []CasbPostureFindingInstanceArchiveResponseRemediation `json:"remediations" api:"required"`
	// The most recent webhook job invocation for each webhook configuration associated
	// with this finding instance. Each entry represents the latest job (any status)
	// per webhook config. The 'stale' field indicates whether the job was invoked
	// before the finding instance's current affliction_date. If no webhook jobs have
	// been created, this field will be an empty array.
	Webhooks []CasbPostureFindingInstanceArchiveResponseWebhook `json:"webhooks" api:"required"`
	// Unique identifier for the finding instance.
	ID string `json:"id" format:"uuid"`
	// Whether this finding instance has been archived.
	IsArchived bool                                          `json:"is_archived"`
	JSON       casbPostureFindingInstanceArchiveResponseJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingInstanceArchiveResponse]
type casbPostureFindingInstanceArchiveResponseJSON struct {
	AfflictionDate apijson.Field
	Asset          apijson.Field
	DLPContexts    apijson.Field
	Remediations   apijson.Field
	Webhooks       apijson.Field
	ID             apijson.Field
	IsArchived     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseJSON) RawJSON() string {
	return r.raw
}

// Asset information including metadata and categorization.
type CasbPostureFindingInstanceArchiveResponseAsset struct {
	// Category information for an asset.
	Category CasbPostureFindingInstanceArchiveResponseAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// The fields associated with the asset.
	Fields []CasbPostureFindingInstanceArchiveResponseAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Unique identifier for the asset.
	ID string `json:"id" format:"uuid"`
	// Direct link to the asset.
	Link string                                             `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceArchiveResponseAssetJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseAssetJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceArchiveResponseAsset]
type casbPostureFindingInstanceArchiveResponseAssetJSON struct {
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for an asset.
type CasbPostureFindingInstanceArchiveResponseAssetCategory struct {
	// The specific service within the vendor the asset is part of (often none).
	// Example - AWS is the vendor, S3 is the service.
	Service string `json:"service" api:"required,nullable"`
	// The type of asset.
	Type string `json:"type" api:"required"`
	// The vendor the asset is part of.
	Vendor string `json:"vendor" api:"required"`
	// Unique identifier for the asset category.
	ID   string                                                     `json:"id" format:"uuid"`
	JSON casbPostureFindingInstanceArchiveResponseAssetCategoryJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseAssetCategoryJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceArchiveResponseAssetCategory]
type casbPostureFindingInstanceArchiveResponseAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Additional field information for an asset.
type CasbPostureFindingInstanceArchiveResponseAssetField struct {
	// The name of the field.
	Name string `json:"name" api:"required"`
	// The value of the field.
	Value string `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                                  `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceArchiveResponseAssetFieldJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseAssetFieldJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceArchiveResponseAssetField]
type casbPostureFindingInstanceArchiveResponseAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseAssetFieldJSON) RawJSON() string {
	return r.raw
}

// DLP context information for a finding.
type CasbPostureFindingInstanceArchiveResponseDLPContext struct {
	// When the DLP context was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// DLP Entry IDs.
	EntryIDs []string `json:"entry_ids" api:"required" format:"uuid"`
	// DLP Profile ID.
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// When the DLP context was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Unique identifier for the DLP context.
	ID string `json:"id" format:"uuid"`
	// When the DLP context was deleted.
	Deleted time.Time `json:"deleted" api:"nullable" format:"date-time"`
	// DLP Right Boundary of match context.
	MatchContextMaxExtent int64 `json:"match_context_max_extent" api:"nullable"`
	// DLP Left Boundary of match context.
	MatchContextMinExtent int64 `json:"match_context_min_extent" api:"nullable"`
	// DLP Match context payload that matched the profile in question.
	MatchContextPayload map[string]interface{}                                  `json:"match_context_payload" api:"nullable"`
	JSON                casbPostureFindingInstanceArchiveResponseDLPContextJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseDLPContextJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceArchiveResponseDLPContext]
type casbPostureFindingInstanceArchiveResponseDLPContextJSON struct {
	Created               apijson.Field
	EntryIDs              apijson.Field
	ProfileID             apijson.Field
	Updated               apijson.Field
	ID                    apijson.Field
	Deleted               apijson.Field
	MatchContextMaxExtent apijson.Field
	MatchContextMinExtent apijson.Field
	MatchContextPayload   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseDLPContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseDLPContextJSON) RawJSON() string {
	return r.raw
}

// Summary information about a remediation job.
type CasbPostureFindingInstanceArchiveResponseRemediation struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this remediation job is stale (created before the finding instance's
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Status of a remediation job.
	Status CasbPostureFindingInstanceArchiveResponseRemediationsStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceArchiveResponseRemediationJSON    `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseRemediationJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceArchiveResponseRemediation]
type casbPostureFindingInstanceArchiveResponseRemediationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseRemediationJSON) RawJSON() string {
	return r.raw
}

// Status of a remediation job.
type CasbPostureFindingInstanceArchiveResponseRemediationsStatus string

const (
	CasbPostureFindingInstanceArchiveResponseRemediationsStatusPending    CasbPostureFindingInstanceArchiveResponseRemediationsStatus = "pending"
	CasbPostureFindingInstanceArchiveResponseRemediationsStatusProcessing CasbPostureFindingInstanceArchiveResponseRemediationsStatus = "processing"
	CasbPostureFindingInstanceArchiveResponseRemediationsStatusCompleted  CasbPostureFindingInstanceArchiveResponseRemediationsStatus = "completed"
	CasbPostureFindingInstanceArchiveResponseRemediationsStatusFailed     CasbPostureFindingInstanceArchiveResponseRemediationsStatus = "failed"
	CasbPostureFindingInstanceArchiveResponseRemediationsStatusValidating CasbPostureFindingInstanceArchiveResponseRemediationsStatus = "validating"
)

func (r CasbPostureFindingInstanceArchiveResponseRemediationsStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceArchiveResponseRemediationsStatusPending, CasbPostureFindingInstanceArchiveResponseRemediationsStatusProcessing, CasbPostureFindingInstanceArchiveResponseRemediationsStatusCompleted, CasbPostureFindingInstanceArchiveResponseRemediationsStatusFailed, CasbPostureFindingInstanceArchiveResponseRemediationsStatusValidating:
		return true
	}
	return false
}

// Summary of the most recent webhook job invocation for a specific webhook
// configuration.
type CasbPostureFindingInstanceArchiveResponseWebhook struct {
	// The most recent webhook job for this webhook configuration.
	LatestJob CasbPostureFindingInstanceArchiveResponseWebhooksLatestJob `json:"latest_job" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookID string `json:"webhook_id" api:"required" format:"uuid"`
	// Account-specified display label for the webhook configuration.
	WebhookLabel string                                               `json:"webhook_label" api:"required"`
	JSON         casbPostureFindingInstanceArchiveResponseWebhookJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseWebhookJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceArchiveResponseWebhook]
type casbPostureFindingInstanceArchiveResponseWebhookJSON struct {
	LatestJob    apijson.Field
	WebhookID    apijson.Field
	WebhookLabel apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseWebhookJSON) RawJSON() string {
	return r.raw
}

// The most recent webhook job for this webhook configuration.
type CasbPostureFindingInstanceArchiveResponseWebhooksLatestJob struct {
	// Unique identifier for the webhook job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the webhook job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this webhook job is stale (created before the finding instance's current
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Current status of the webhook job.
	Status CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceArchiveResponseWebhooksLatestJobJSON   `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseWebhooksLatestJobJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceArchiveResponseWebhooksLatestJob]
type casbPostureFindingInstanceArchiveResponseWebhooksLatestJobJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseWebhooksLatestJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseWebhooksLatestJobJSON) RawJSON() string {
	return r.raw
}

// Current status of the webhook job.
type CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus string

const (
	CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusPending    CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus = "pending"
	CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusProcessing CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus = "processing"
	CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusCompleted  CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus = "completed"
)

func (r CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusPending, CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusProcessing, CasbPostureFindingInstanceArchiveResponseWebhooksLatestJobStatusCompleted:
		return true
	}
	return false
}

// Information about an export job.
type CasbPostureFindingInstanceExportResponse struct {
	// Unique identifier for the export job.
	ID string `json:"id" api:"required" format:"uuid"`
	// Status of an export job.
	Status CasbPostureFindingInstanceExportResponseStatus `json:"status" api:"required"`
	// Type of export job.
	Type CasbPostureFindingInstanceExportResponseType `json:"type" api:"required"`
	// ID of the export-requesting user.
	UserID string `json:"user_id" api:"required"`
	// The URL by which the successfully created export can be downloaded by the end
	// users.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Contains information on errors which may have occurred during export creation.
	Errors string `json:"errors" api:"nullable"`
	// The base name of the file that is/was generated by the export job.
	FileName string `json:"file_name" api:"nullable"`
	// The full path of the file that is stored within external storage (currently R2).
	FilePath string                                       `json:"file_path" api:"nullable"`
	JSON     casbPostureFindingInstanceExportResponseJSON `json:"-"`
}

// casbPostureFindingInstanceExportResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingInstanceExportResponse]
type casbPostureFindingInstanceExportResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	UserID      apijson.Field
	DownloadURL apijson.Field
	Errors      apijson.Field
	FileName    apijson.Field
	FilePath    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseJSON) RawJSON() string {
	return r.raw
}

// Status of an export job.
type CasbPostureFindingInstanceExportResponseStatus string

const (
	CasbPostureFindingInstanceExportResponseStatusPending     CasbPostureFindingInstanceExportResponseStatus = "Pending"
	CasbPostureFindingInstanceExportResponseStatusSuccess     CasbPostureFindingInstanceExportResponseStatus = "Success"
	CasbPostureFindingInstanceExportResponseStatusFailure     CasbPostureFindingInstanceExportResponseStatus = "Failure"
	CasbPostureFindingInstanceExportResponseStatusRescheduled CasbPostureFindingInstanceExportResponseStatus = "Rescheduled"
	CasbPostureFindingInstanceExportResponseStatusInProgress  CasbPostureFindingInstanceExportResponseStatus = "In-Progress"
)

func (r CasbPostureFindingInstanceExportResponseStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceExportResponseStatusPending, CasbPostureFindingInstanceExportResponseStatusSuccess, CasbPostureFindingInstanceExportResponseStatusFailure, CasbPostureFindingInstanceExportResponseStatusRescheduled, CasbPostureFindingInstanceExportResponseStatusInProgress:
		return true
	}
	return false
}

// Type of export job.
type CasbPostureFindingInstanceExportResponseType string

const (
	CasbPostureFindingInstanceExportResponseTypeFinding         CasbPostureFindingInstanceExportResponseType = "finding"
	CasbPostureFindingInstanceExportResponseTypeFindingInstance CasbPostureFindingInstanceExportResponseType = "findingInstance"
	CasbPostureFindingInstanceExportResponseTypeContent         CasbPostureFindingInstanceExportResponseType = "content"
	CasbPostureFindingInstanceExportResponseTypeRemediationJob  CasbPostureFindingInstanceExportResponseType = "remediationJob"
)

func (r CasbPostureFindingInstanceExportResponseType) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceExportResponseTypeFinding, CasbPostureFindingInstanceExportResponseTypeFindingInstance, CasbPostureFindingInstanceExportResponseTypeContent, CasbPostureFindingInstanceExportResponseTypeRemediationJob:
		return true
	}
	return false
}

// A specific instance of a security finding. In the API interface, we refer to the
// 'finding' table in our DB as finding instances, optimized for the p99 use case.
type CasbPostureFindingInstanceGetResponse struct {
	// When this specific instance was identified.
	AfflictionDate time.Time `json:"affliction_date" api:"required" format:"date-time"`
	// Asset information including metadata and categorization.
	Asset CasbPostureFindingInstanceGetResponseAsset `json:"asset" api:"required"`
	// DLP context information if this is a content finding.
	DLPContexts []CasbPostureFindingInstanceGetResponseDLPContext `json:"dlp_contexts" api:"required"`
	// A list of the 10 most recent remediation jobs for this finding instance, ordered
	// by creation time (most recent first). The 'stale' field indicates whether the
	// remediation job was created before the finding instance's affliction_date (true)
	// or after it (false). If there has never been a remediation job for this finding
	// instance, this field will be an empty array.
	Remediations []CasbPostureFindingInstanceGetResponseRemediation `json:"remediations" api:"required"`
	// The most recent webhook job invocation for each webhook configuration associated
	// with this finding instance. Each entry represents the latest job (any status)
	// per webhook config. The 'stale' field indicates whether the job was invoked
	// before the finding instance's current affliction_date. If no webhook jobs have
	// been created, this field will be an empty array.
	Webhooks []CasbPostureFindingInstanceGetResponseWebhook `json:"webhooks" api:"required"`
	// Unique identifier for the finding instance.
	ID string `json:"id" format:"uuid"`
	// Whether this finding instance has been archived.
	IsArchived bool                                      `json:"is_archived"`
	JSON       casbPostureFindingInstanceGetResponseJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseJSON contains the JSON metadata for the
// struct [CasbPostureFindingInstanceGetResponse]
type casbPostureFindingInstanceGetResponseJSON struct {
	AfflictionDate apijson.Field
	Asset          apijson.Field
	DLPContexts    apijson.Field
	Remediations   apijson.Field
	Webhooks       apijson.Field
	ID             apijson.Field
	IsArchived     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseJSON) RawJSON() string {
	return r.raw
}

// Asset information including metadata and categorization.
type CasbPostureFindingInstanceGetResponseAsset struct {
	// Category information for an asset.
	Category CasbPostureFindingInstanceGetResponseAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// The fields associated with the asset.
	Fields []CasbPostureFindingInstanceGetResponseAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Unique identifier for the asset.
	ID string `json:"id" format:"uuid"`
	// Direct link to the asset.
	Link string                                         `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceGetResponseAssetJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseAssetJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceGetResponseAsset]
type casbPostureFindingInstanceGetResponseAssetJSON struct {
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for an asset.
type CasbPostureFindingInstanceGetResponseAssetCategory struct {
	// The specific service within the vendor the asset is part of (often none).
	// Example - AWS is the vendor, S3 is the service.
	Service string `json:"service" api:"required,nullable"`
	// The type of asset.
	Type string `json:"type" api:"required"`
	// The vendor the asset is part of.
	Vendor string `json:"vendor" api:"required"`
	// Unique identifier for the asset category.
	ID   string                                                 `json:"id" format:"uuid"`
	JSON casbPostureFindingInstanceGetResponseAssetCategoryJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseAssetCategoryJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceGetResponseAssetCategory]
type casbPostureFindingInstanceGetResponseAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Additional field information for an asset.
type CasbPostureFindingInstanceGetResponseAssetField struct {
	// The name of the field.
	Name string `json:"name" api:"required"`
	// The value of the field.
	Value string `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                              `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceGetResponseAssetFieldJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseAssetFieldJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceGetResponseAssetField]
type casbPostureFindingInstanceGetResponseAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseAssetFieldJSON) RawJSON() string {
	return r.raw
}

// DLP context information for a finding.
type CasbPostureFindingInstanceGetResponseDLPContext struct {
	// When the DLP context was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// DLP Entry IDs.
	EntryIDs []string `json:"entry_ids" api:"required" format:"uuid"`
	// DLP Profile ID.
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// When the DLP context was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Unique identifier for the DLP context.
	ID string `json:"id" format:"uuid"`
	// When the DLP context was deleted.
	Deleted time.Time `json:"deleted" api:"nullable" format:"date-time"`
	// DLP Right Boundary of match context.
	MatchContextMaxExtent int64 `json:"match_context_max_extent" api:"nullable"`
	// DLP Left Boundary of match context.
	MatchContextMinExtent int64 `json:"match_context_min_extent" api:"nullable"`
	// DLP Match context payload that matched the profile in question.
	MatchContextPayload map[string]interface{}                              `json:"match_context_payload" api:"nullable"`
	JSON                casbPostureFindingInstanceGetResponseDLPContextJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseDLPContextJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceGetResponseDLPContext]
type casbPostureFindingInstanceGetResponseDLPContextJSON struct {
	Created               apijson.Field
	EntryIDs              apijson.Field
	ProfileID             apijson.Field
	Updated               apijson.Field
	ID                    apijson.Field
	Deleted               apijson.Field
	MatchContextMaxExtent apijson.Field
	MatchContextMinExtent apijson.Field
	MatchContextPayload   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseDLPContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseDLPContextJSON) RawJSON() string {
	return r.raw
}

// Summary information about a remediation job.
type CasbPostureFindingInstanceGetResponseRemediation struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this remediation job is stale (created before the finding instance's
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Status of a remediation job.
	Status CasbPostureFindingInstanceGetResponseRemediationsStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceGetResponseRemediationJSON    `json:"-"`
}

// casbPostureFindingInstanceGetResponseRemediationJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceGetResponseRemediation]
type casbPostureFindingInstanceGetResponseRemediationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseRemediationJSON) RawJSON() string {
	return r.raw
}

// Status of a remediation job.
type CasbPostureFindingInstanceGetResponseRemediationsStatus string

const (
	CasbPostureFindingInstanceGetResponseRemediationsStatusPending    CasbPostureFindingInstanceGetResponseRemediationsStatus = "pending"
	CasbPostureFindingInstanceGetResponseRemediationsStatusProcessing CasbPostureFindingInstanceGetResponseRemediationsStatus = "processing"
	CasbPostureFindingInstanceGetResponseRemediationsStatusCompleted  CasbPostureFindingInstanceGetResponseRemediationsStatus = "completed"
	CasbPostureFindingInstanceGetResponseRemediationsStatusFailed     CasbPostureFindingInstanceGetResponseRemediationsStatus = "failed"
	CasbPostureFindingInstanceGetResponseRemediationsStatusValidating CasbPostureFindingInstanceGetResponseRemediationsStatus = "validating"
)

func (r CasbPostureFindingInstanceGetResponseRemediationsStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceGetResponseRemediationsStatusPending, CasbPostureFindingInstanceGetResponseRemediationsStatusProcessing, CasbPostureFindingInstanceGetResponseRemediationsStatusCompleted, CasbPostureFindingInstanceGetResponseRemediationsStatusFailed, CasbPostureFindingInstanceGetResponseRemediationsStatusValidating:
		return true
	}
	return false
}

// Summary of the most recent webhook job invocation for a specific webhook
// configuration.
type CasbPostureFindingInstanceGetResponseWebhook struct {
	// The most recent webhook job for this webhook configuration.
	LatestJob CasbPostureFindingInstanceGetResponseWebhooksLatestJob `json:"latest_job" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookID string `json:"webhook_id" api:"required" format:"uuid"`
	// Account-specified display label for the webhook configuration.
	WebhookLabel string                                           `json:"webhook_label" api:"required"`
	JSON         casbPostureFindingInstanceGetResponseWebhookJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseWebhookJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceGetResponseWebhook]
type casbPostureFindingInstanceGetResponseWebhookJSON struct {
	LatestJob    apijson.Field
	WebhookID    apijson.Field
	WebhookLabel apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseWebhookJSON) RawJSON() string {
	return r.raw
}

// The most recent webhook job for this webhook configuration.
type CasbPostureFindingInstanceGetResponseWebhooksLatestJob struct {
	// Unique identifier for the webhook job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the webhook job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this webhook job is stale (created before the finding instance's current
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Current status of the webhook job.
	Status CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceGetResponseWebhooksLatestJobJSON   `json:"-"`
}

// casbPostureFindingInstanceGetResponseWebhooksLatestJobJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceGetResponseWebhooksLatestJob]
type casbPostureFindingInstanceGetResponseWebhooksLatestJobJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseWebhooksLatestJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseWebhooksLatestJobJSON) RawJSON() string {
	return r.raw
}

// Current status of the webhook job.
type CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus string

const (
	CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusPending    CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus = "pending"
	CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusProcessing CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus = "processing"
	CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusCompleted  CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus = "completed"
)

func (r CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusPending, CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusProcessing, CasbPostureFindingInstanceGetResponseWebhooksLatestJobStatusCompleted:
		return true
	}
	return false
}

// A specific instance of a security finding. In the API interface, we refer to the
// 'finding' table in our DB as finding instances, optimized for the p99 use case.
type CasbPostureFindingInstanceUnarchiveResponse struct {
	// When this specific instance was identified.
	AfflictionDate time.Time `json:"affliction_date" api:"required" format:"date-time"`
	// Asset information including metadata and categorization.
	Asset CasbPostureFindingInstanceUnarchiveResponseAsset `json:"asset" api:"required"`
	// DLP context information if this is a content finding.
	DLPContexts []CasbPostureFindingInstanceUnarchiveResponseDLPContext `json:"dlp_contexts" api:"required"`
	// A list of the 10 most recent remediation jobs for this finding instance, ordered
	// by creation time (most recent first). The 'stale' field indicates whether the
	// remediation job was created before the finding instance's affliction_date (true)
	// or after it (false). If there has never been a remediation job for this finding
	// instance, this field will be an empty array.
	Remediations []CasbPostureFindingInstanceUnarchiveResponseRemediation `json:"remediations" api:"required"`
	// The most recent webhook job invocation for each webhook configuration associated
	// with this finding instance. Each entry represents the latest job (any status)
	// per webhook config. The 'stale' field indicates whether the job was invoked
	// before the finding instance's current affliction_date. If no webhook jobs have
	// been created, this field will be an empty array.
	Webhooks []CasbPostureFindingInstanceUnarchiveResponseWebhook `json:"webhooks" api:"required"`
	// Unique identifier for the finding instance.
	ID string `json:"id" format:"uuid"`
	// Whether this finding instance has been archived.
	IsArchived bool                                            `json:"is_archived"`
	JSON       casbPostureFindingInstanceUnarchiveResponseJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceUnarchiveResponse]
type casbPostureFindingInstanceUnarchiveResponseJSON struct {
	AfflictionDate apijson.Field
	Asset          apijson.Field
	DLPContexts    apijson.Field
	Remediations   apijson.Field
	Webhooks       apijson.Field
	ID             apijson.Field
	IsArchived     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseJSON) RawJSON() string {
	return r.raw
}

// Asset information including metadata and categorization.
type CasbPostureFindingInstanceUnarchiveResponseAsset struct {
	// Category information for an asset.
	Category CasbPostureFindingInstanceUnarchiveResponseAssetCategory `json:"category" api:"required"`
	// External identifier from the source system.
	ExternalID string `json:"external_id" api:"required"`
	// The fields associated with the asset.
	Fields []CasbPostureFindingInstanceUnarchiveResponseAssetField `json:"fields" api:"required"`
	// Human-readable name of the asset.
	Name string `json:"name" api:"required"`
	// Unique identifier for the asset.
	ID string `json:"id" format:"uuid"`
	// Direct link to the asset.
	Link string                                               `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceUnarchiveResponseAssetJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseAssetJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceUnarchiveResponseAsset]
type casbPostureFindingInstanceUnarchiveResponseAssetJSON struct {
	Category    apijson.Field
	ExternalID  apijson.Field
	Fields      apijson.Field
	Name        apijson.Field
	ID          apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseAsset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseAssetJSON) RawJSON() string {
	return r.raw
}

// Category information for an asset.
type CasbPostureFindingInstanceUnarchiveResponseAssetCategory struct {
	// The specific service within the vendor the asset is part of (often none).
	// Example - AWS is the vendor, S3 is the service.
	Service string `json:"service" api:"required,nullable"`
	// The type of asset.
	Type string `json:"type" api:"required"`
	// The vendor the asset is part of.
	Vendor string `json:"vendor" api:"required"`
	// Unique identifier for the asset category.
	ID   string                                                       `json:"id" format:"uuid"`
	JSON casbPostureFindingInstanceUnarchiveResponseAssetCategoryJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseAssetCategoryJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseAssetCategory]
type casbPostureFindingInstanceUnarchiveResponseAssetCategoryJSON struct {
	Service     apijson.Field
	Type        apijson.Field
	Vendor      apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseAssetCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseAssetCategoryJSON) RawJSON() string {
	return r.raw
}

// Additional field information for an asset.
type CasbPostureFindingInstanceUnarchiveResponseAssetField struct {
	// The name of the field.
	Name string `json:"name" api:"required"`
	// The value of the field.
	Value string `json:"value" api:"required"`
	// Optional link associated with the field.
	Link string                                                    `json:"link" api:"nullable" format:"uri"`
	JSON casbPostureFindingInstanceUnarchiveResponseAssetFieldJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseAssetFieldJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceUnarchiveResponseAssetField]
type casbPostureFindingInstanceUnarchiveResponseAssetFieldJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	Link        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseAssetField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseAssetFieldJSON) RawJSON() string {
	return r.raw
}

// DLP context information for a finding.
type CasbPostureFindingInstanceUnarchiveResponseDLPContext struct {
	// When the DLP context was created.
	Created time.Time `json:"created" api:"required" format:"date-time"`
	// DLP Entry IDs.
	EntryIDs []string `json:"entry_ids" api:"required" format:"uuid"`
	// DLP Profile ID.
	ProfileID string `json:"profile_id" api:"required" format:"uuid"`
	// When the DLP context was last updated.
	Updated time.Time `json:"updated" api:"required" format:"date-time"`
	// Unique identifier for the DLP context.
	ID string `json:"id" format:"uuid"`
	// When the DLP context was deleted.
	Deleted time.Time `json:"deleted" api:"nullable" format:"date-time"`
	// DLP Right Boundary of match context.
	MatchContextMaxExtent int64 `json:"match_context_max_extent" api:"nullable"`
	// DLP Left Boundary of match context.
	MatchContextMinExtent int64 `json:"match_context_min_extent" api:"nullable"`
	// DLP Match context payload that matched the profile in question.
	MatchContextPayload map[string]interface{}                                    `json:"match_context_payload" api:"nullable"`
	JSON                casbPostureFindingInstanceUnarchiveResponseDLPContextJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseDLPContextJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceUnarchiveResponseDLPContext]
type casbPostureFindingInstanceUnarchiveResponseDLPContextJSON struct {
	Created               apijson.Field
	EntryIDs              apijson.Field
	ProfileID             apijson.Field
	Updated               apijson.Field
	ID                    apijson.Field
	Deleted               apijson.Field
	MatchContextMaxExtent apijson.Field
	MatchContextMinExtent apijson.Field
	MatchContextPayload   apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseDLPContext) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseDLPContextJSON) RawJSON() string {
	return r.raw
}

// Summary information about a remediation job.
type CasbPostureFindingInstanceUnarchiveResponseRemediation struct {
	// Unique identifier for the remediation job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the remediation job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this remediation job is stale (created before the finding instance's
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Status of a remediation job.
	Status CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceUnarchiveResponseRemediationJSON    `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseRemediationJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceUnarchiveResponseRemediation]
type casbPostureFindingInstanceUnarchiveResponseRemediationJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseRemediation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseRemediationJSON) RawJSON() string {
	return r.raw
}

// Status of a remediation job.
type CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus string

const (
	CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusPending    CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus = "pending"
	CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusProcessing CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus = "processing"
	CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusCompleted  CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus = "completed"
	CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusFailed     CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus = "failed"
	CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusValidating CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus = "validating"
)

func (r CasbPostureFindingInstanceUnarchiveResponseRemediationsStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusPending, CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusProcessing, CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusCompleted, CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusFailed, CasbPostureFindingInstanceUnarchiveResponseRemediationsStatusValidating:
		return true
	}
	return false
}

// Summary of the most recent webhook job invocation for a specific webhook
// configuration.
type CasbPostureFindingInstanceUnarchiveResponseWebhook struct {
	// The most recent webhook job for this webhook configuration.
	LatestJob CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJob `json:"latest_job" api:"required"`
	// Unique identifier for the webhook configuration.
	WebhookID string `json:"webhook_id" api:"required" format:"uuid"`
	// Account-specified display label for the webhook configuration.
	WebhookLabel string                                                 `json:"webhook_label" api:"required"`
	JSON         casbPostureFindingInstanceUnarchiveResponseWebhookJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseWebhookJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceUnarchiveResponseWebhook]
type casbPostureFindingInstanceUnarchiveResponseWebhookJSON struct {
	LatestJob    apijson.Field
	WebhookID    apijson.Field
	WebhookLabel apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseWebhook) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseWebhookJSON) RawJSON() string {
	return r.raw
}

// The most recent webhook job for this webhook configuration.
type CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJob struct {
	// Unique identifier for the webhook job.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the webhook job was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Whether this webhook job is stale (created before the finding instance's current
	// affliction_date).
	Stale bool `json:"stale" api:"required"`
	// Current status of the webhook job.
	Status CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus `json:"status" api:"required"`
	JSON   casbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobJSON   `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJob]
type casbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Stale       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJob) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobJSON) RawJSON() string {
	return r.raw
}

// Current status of the webhook job.
type CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus string

const (
	CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusPending    CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus = "pending"
	CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusProcessing CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus = "processing"
	CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusCompleted  CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus = "completed"
)

func (r CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusPending, CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusProcessing, CasbPostureFindingInstanceUnarchiveResponseWebhooksLatestJobStatusCompleted:
		return true
	}
	return false
}

type CasbPostureFindingInstanceListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Archived
	Archived param.Field[bool] `query:"archived"`
	// Filter finding instances by an array of asset IDs. Supports multiple
	// comma-separated values.
	AssetIDs param.Field[[]string] `query:"asset_ids" format:"uuid"`
	// A cursor for pagination. Obtained from the `result_info.cursor` field of a
	// previous response.
	Cursor param.Field[string] `query:"cursor"`
	// Direction to order results.
	Direction param.Field[CasbPostureFindingInstanceListParamsDirection] `query:"direction"`
	// Filter finding instances by an array of finding instance IDs. Supports multiple
	// comma-separated values.
	FindingInstanceIDs param.Field[[]string] `query:"finding_instance_ids" format:"uuid"`
	// Filter to view findings that occurred on or before the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxAfflictionDate param.Field[time.Time] `query:"max_affliction_date" format:"date-time"`
	// Filter to view findings that occurred on or after the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinAfflictionDate param.Field[time.Time] `query:"min_affliction_date" format:"date-time"`
	// Which field to use when ordering the Finding's instances. When ordering by
	// 'remediation.status', only the most recent non-stale remediation job is
	// considered. Stale jobs (created before the instance's affliction_date) are
	// treated as having no status for ordering purposes.
	Order param.Field[CasbPostureFindingInstanceListParamsOrder] `query:"order"`
	// A page number within the paginated result set.
	Page param.Field[int64] `query:"page"`
	// Number of results to return per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter finding instances by most recent remediation job status. Supports
	// multiple comma-separated values. Use 'none' to filter instances with no
	// remediation jobs or instances where the most recent job is stale. Note: Stale
	// jobs (created before the instance's affliction_date) are ignored for filtering
	// purposes, but are still included in the 'remediations' array with stale=true.
	RemediationStatuses param.Field[[]CasbPostureFindingInstanceListParamsRemediationStatus] `query:"remediation_statuses"`
	// A search term.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [CasbPostureFindingInstanceListParams]'s query parameters as
// `url.Values`.
func (r CasbPostureFindingInstanceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type CasbPostureFindingInstanceListParamsDirection string

const (
	CasbPostureFindingInstanceListParamsDirectionAsc  CasbPostureFindingInstanceListParamsDirection = "asc"
	CasbPostureFindingInstanceListParamsDirectionDesc CasbPostureFindingInstanceListParamsDirection = "desc"
)

func (r CasbPostureFindingInstanceListParamsDirection) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceListParamsDirectionAsc, CasbPostureFindingInstanceListParamsDirectionDesc:
		return true
	}
	return false
}

// Which field to use when ordering the Finding's instances. When ordering by
// 'remediation.status', only the most recent non-stale remediation job is
// considered. Stale jobs (created before the instance's affliction_date) are
// treated as having no status for ordering purposes.
type CasbPostureFindingInstanceListParamsOrder string

const (
	CasbPostureFindingInstanceListParamsOrderAfflictionDate    CasbPostureFindingInstanceListParamsOrder = "affliction_date"
	CasbPostureFindingInstanceListParamsOrderAssetName         CasbPostureFindingInstanceListParamsOrder = "asset.name"
	CasbPostureFindingInstanceListParamsOrderRemediationStatus CasbPostureFindingInstanceListParamsOrder = "remediation.status"
)

func (r CasbPostureFindingInstanceListParamsOrder) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceListParamsOrderAfflictionDate, CasbPostureFindingInstanceListParamsOrderAssetName, CasbPostureFindingInstanceListParamsOrderRemediationStatus:
		return true
	}
	return false
}

type CasbPostureFindingInstanceListParamsRemediationStatus string

const (
	CasbPostureFindingInstanceListParamsRemediationStatusNone       CasbPostureFindingInstanceListParamsRemediationStatus = "none"
	CasbPostureFindingInstanceListParamsRemediationStatusPending    CasbPostureFindingInstanceListParamsRemediationStatus = "pending"
	CasbPostureFindingInstanceListParamsRemediationStatusProcessing CasbPostureFindingInstanceListParamsRemediationStatus = "processing"
	CasbPostureFindingInstanceListParamsRemediationStatusValidating CasbPostureFindingInstanceListParamsRemediationStatus = "validating"
	CasbPostureFindingInstanceListParamsRemediationStatusCompleted  CasbPostureFindingInstanceListParamsRemediationStatus = "completed"
	CasbPostureFindingInstanceListParamsRemediationStatusFailed     CasbPostureFindingInstanceListParamsRemediationStatus = "failed"
)

func (r CasbPostureFindingInstanceListParamsRemediationStatus) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceListParamsRemediationStatusNone, CasbPostureFindingInstanceListParamsRemediationStatusPending, CasbPostureFindingInstanceListParamsRemediationStatusProcessing, CasbPostureFindingInstanceListParamsRemediationStatusValidating, CasbPostureFindingInstanceListParamsRemediationStatusCompleted, CasbPostureFindingInstanceListParamsRemediationStatusFailed:
		return true
	}
	return false
}

type CasbPostureFindingInstanceArchiveParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A list of finding instance IDs to pass along.
	CheckInstances param.Field[[]string] `json:"check_instances" api:"required" format:"uuid"`
}

func (r CasbPostureFindingInstanceArchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureFindingInstanceArchiveResponseEnvelope struct {
	Errors   []CasbPostureFindingInstanceArchiveResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingInstanceArchiveResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// A specific instance of a security finding. In the API interface, we refer to the
	// 'finding' table in our DB as finding instances, optimized for the p99 use case.
	Result CasbPostureFindingInstanceArchiveResponse             `json:"result"`
	JSON   casbPostureFindingInstanceArchiveResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceArchiveResponseEnvelope]
type casbPostureFindingInstanceArchiveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceArchiveResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                        `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceArchiveResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingInstanceArchiveResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceArchiveResponseEnvelopeErrors]
type casbPostureFindingInstanceArchiveResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceArchiveResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                            `json:"pointer"`
	JSON    casbPostureFindingInstanceArchiveResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceArchiveResponseEnvelopeErrorsSource]
type casbPostureFindingInstanceArchiveResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceArchiveResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                          `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceArchiveResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingInstanceArchiveResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceArchiveResponseEnvelopeMessages]
type casbPostureFindingInstanceArchiveResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceArchiveResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                              `json:"pointer"`
	JSON    casbPostureFindingInstanceArchiveResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingInstanceArchiveResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceArchiveResponseEnvelopeMessagesSource]
type casbPostureFindingInstanceArchiveResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceArchiveResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceArchiveResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceExportParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter for archived status.
	Archived param.Field[bool] `json:"archived"`
	// Filter to view findings that occurred on or before the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MaxAfflictionDate param.Field[time.Time] `json:"max_affliction_date" format:"date-time"`
	// Filter to view findings that occurred on or after the affliction date. Can be a
	// date-time in ISO 8601 format or an epoch timestamp.
	MinAfflictionDate param.Field[time.Time] `json:"min_affliction_date" format:"date-time"`
	// Ordering specifications for the export.
	Orders param.Field[[]CasbPostureFindingInstanceExportParamsOrder] `json:"orders"`
	// A search term.
	Search param.Field[string] `json:"search"`
}

func (r CasbPostureFindingInstanceExportParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Order specification for finding instance exports.
type CasbPostureFindingInstanceExportParamsOrder struct {
	// Sort direction.
	Direction param.Field[CasbPostureFindingInstanceExportParamsOrdersDirection] `json:"direction" api:"required"`
	// Which field to use when ordering the finding instances.
	Name param.Field[CasbPostureFindingInstanceExportParamsOrdersName] `json:"name" api:"required"`
}

func (r CasbPostureFindingInstanceExportParamsOrder) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction.
type CasbPostureFindingInstanceExportParamsOrdersDirection string

const (
	CasbPostureFindingInstanceExportParamsOrdersDirectionAsc  CasbPostureFindingInstanceExportParamsOrdersDirection = "asc"
	CasbPostureFindingInstanceExportParamsOrdersDirectionDesc CasbPostureFindingInstanceExportParamsOrdersDirection = "desc"
)

func (r CasbPostureFindingInstanceExportParamsOrdersDirection) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceExportParamsOrdersDirectionAsc, CasbPostureFindingInstanceExportParamsOrdersDirectionDesc:
		return true
	}
	return false
}

// Which field to use when ordering the finding instances.
type CasbPostureFindingInstanceExportParamsOrdersName string

const (
	CasbPostureFindingInstanceExportParamsOrdersNameAssetName      CasbPostureFindingInstanceExportParamsOrdersName = "asset.name"
	CasbPostureFindingInstanceExportParamsOrdersNameAfflictionDate CasbPostureFindingInstanceExportParamsOrdersName = "affliction_date"
)

func (r CasbPostureFindingInstanceExportParamsOrdersName) IsKnown() bool {
	switch r {
	case CasbPostureFindingInstanceExportParamsOrdersNameAssetName, CasbPostureFindingInstanceExportParamsOrdersNameAfflictionDate:
		return true
	}
	return false
}

// Common response structure for all API endpoints.
type CasbPostureFindingInstanceExportResponseEnvelope struct {
	Errors   []CasbPostureFindingInstanceExportResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingInstanceExportResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// Information about an export job.
	Result CasbPostureFindingInstanceExportResponse             `json:"result"`
	JSON   casbPostureFindingInstanceExportResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingInstanceExportResponseEnvelopeJSON contains the JSON metadata
// for the struct [CasbPostureFindingInstanceExportResponseEnvelope]
type casbPostureFindingInstanceExportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceExportResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                       `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceExportResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingInstanceExportResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingInstanceExportResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceExportResponseEnvelopeErrors]
type casbPostureFindingInstanceExportResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceExportResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                           `json:"pointer"`
	JSON    casbPostureFindingInstanceExportResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingInstanceExportResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceExportResponseEnvelopeErrorsSource]
type casbPostureFindingInstanceExportResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceExportResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                         `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceExportResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingInstanceExportResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingInstanceExportResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceExportResponseEnvelopeMessages]
type casbPostureFindingInstanceExportResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceExportResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                             `json:"pointer"`
	JSON    casbPostureFindingInstanceExportResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingInstanceExportResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceExportResponseEnvelopeMessagesSource]
type casbPostureFindingInstanceExportResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceExportResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceExportResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

// Common response structure for all API endpoints.
type CasbPostureFindingInstanceGetResponseEnvelope struct {
	Errors   []CasbPostureFindingInstanceGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingInstanceGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// A specific instance of a security finding. In the API interface, we refer to the
	// 'finding' table in our DB as finding instances, optimized for the p99 use case.
	Result CasbPostureFindingInstanceGetResponse             `json:"result"`
	JSON   casbPostureFindingInstanceGetResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [CasbPostureFindingInstanceGetResponseEnvelope]
type casbPostureFindingInstanceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceGetResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                    `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingInstanceGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingInstanceGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceGetResponseEnvelopeErrors]
type casbPostureFindingInstanceGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceGetResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                        `json:"pointer"`
	JSON    casbPostureFindingInstanceGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceGetResponseEnvelopeErrorsSource]
type casbPostureFindingInstanceGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceGetResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                      `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingInstanceGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingInstanceGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceGetResponseEnvelopeMessages]
type casbPostureFindingInstanceGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceGetResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                          `json:"pointer"`
	JSON    casbPostureFindingInstanceGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingInstanceGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceGetResponseEnvelopeMessagesSource]
type casbPostureFindingInstanceGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceUnarchiveParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// A list of finding instance IDs to pass along.
	CheckInstances param.Field[[]string] `json:"check_instances" api:"required" format:"uuid"`
}

func (r CasbPostureFindingInstanceUnarchiveParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Common response structure for all API endpoints.
type CasbPostureFindingInstanceUnarchiveResponseEnvelope struct {
	Errors   []CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success bool `json:"success" api:"required"`
	// A specific instance of a security finding. In the API interface, we refer to the
	// 'finding' table in our DB as finding instances, optimized for the p99 use case.
	Result CasbPostureFindingInstanceUnarchiveResponse             `json:"result"`
	JSON   casbPostureFindingInstanceUnarchiveResponseEnvelopeJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseEnvelopeJSON contains the JSON
// metadata for the struct [CasbPostureFindingInstanceUnarchiveResponseEnvelope]
type casbPostureFindingInstanceUnarchiveResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrors struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                          `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSource `json:"source"`
	JSON             casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsJSON   `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrors]
type casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                              `json:"pointer"`
	JSON    casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSource]
type casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessages struct {
	// Error or message code.
	Code int64 `json:"code" api:"required"`
	// Human-readable message.
	Message string `json:"message" api:"required"`
	// Link to relevant documentation.
	DocumentationURL string                                                            `json:"documentation_url" format:"uri"`
	Source           CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSource `json:"source"`
	JSON             casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesJSON   `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessages]
type casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSource struct {
	// JSON pointer to the source of the error.
	Pointer string                                                                `json:"pointer"`
	JSON    casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSource]
type casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CasbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r casbPostureFindingInstanceUnarchiveResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}
