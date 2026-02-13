// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ThreatEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventService] method instead.
type ThreatEventService struct {
	Options          []option.RequestOption
	Attackers        *ThreatEventAttackerService
	Categories       *ThreatEventCategoryService
	Countries        *ThreatEventCountryService
	Crons            *ThreatEventCronService
	Datasets         *ThreatEventDatasetService
	IndicatorTypes   *ThreatEventIndicatorTypeService
	Raw              *ThreatEventRawService
	Relate           *ThreatEventRelateService
	Tags             *ThreatEventTagService
	EventTags        *ThreatEventEventTagService
	TargetIndustries *ThreatEventTargetIndustryService
	Insights         *ThreatEventInsightService
}

// NewThreatEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewThreatEventService(opts ...option.RequestOption) (r *ThreatEventService) {
	r = &ThreatEventService{}
	r.Options = opts
	r.Attackers = NewThreatEventAttackerService(opts...)
	r.Categories = NewThreatEventCategoryService(opts...)
	r.Countries = NewThreatEventCountryService(opts...)
	r.Crons = NewThreatEventCronService(opts...)
	r.Datasets = NewThreatEventDatasetService(opts...)
	r.IndicatorTypes = NewThreatEventIndicatorTypeService(opts...)
	r.Raw = NewThreatEventRawService(opts...)
	r.Relate = NewThreatEventRelateService(opts...)
	r.Tags = NewThreatEventTagService(opts...)
	r.EventTags = NewThreatEventEventTagService(opts...)
	r.TargetIndustries = NewThreatEventTargetIndustryService(opts...)
	r.Insights = NewThreatEventInsightService(opts...)
	return
}

// To create a dataset, see the
// [`Create Dataset`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/create/)
// endpoint. When `datasetId` parameter is unspecified, it will be created in a
// default dataset named `Cloudforce One Threat Events`.
func (r *ThreatEventService) New(ctx context.Context, params ThreatEventNewParams, opts ...option.RequestOption) (res *ThreatEventNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.PathAccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create", params.PathAccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// When `datasetId` is unspecified, events will be listed from the
// `Cloudforce One Threat Events` dataset. To list existing datasets (and their
// IDs), use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint). Also, must provide query parameters.
func (r *ThreatEventService) List(ctx context.Context, params ThreatEventListParams, opts ...option.RequestOption) (res *[]ThreatEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// The `datasetId` parameter must be defined. To list existing datasets (and their
// IDs) in your account, use the
// [`List Datasets`](https://developers.cloudflare.com/api/resources/cloudforce_one/subresources/threat_events/subresources/datasets/methods/list/)
// endpoint.
func (r *ThreatEventService) BulkNew(ctx context.Context, params ThreatEventBulkNewParams, opts ...option.RequestOption) (res *ThreatEventBulkNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/create/bulk", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates an event
func (r *ThreatEventService) Edit(ctx context.Context, eventID string, params ThreatEventEditParams, opts ...option.RequestOption) (res *ThreatEventEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", params.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// This Method is deprecated. Please use
// /events/dataset/:dataset_id/events/:event_id instead.
//
// Deprecated: deprecated
func (r *ThreatEventService) Get(ctx context.Context, eventID string, query ThreatEventGetParams, opts ...option.RequestOption) (res *ThreatEventGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if eventID == "" {
		err = errors.New("missing required event_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/%s", query.AccountID, eventID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventNewResponse struct {
	Attacker        string                     `json:"attacker,required"`
	AttackerCountry string                     `json:"attackerCountry,required"`
	Category        string                     `json:"category,required"`
	DatasetID       string                     `json:"datasetId,required"`
	Date            string                     `json:"date,required"`
	Event           string                     `json:"event,required"`
	HasChildren     bool                       `json:"hasChildren,required"`
	Indicator       string                     `json:"indicator,required"`
	IndicatorType   string                     `json:"indicatorType,required"`
	IndicatorTypeID float64                    `json:"indicatorTypeId,required"`
	KillChain       float64                    `json:"killChain,required"`
	MitreAttack     []string                   `json:"mitreAttack,required"`
	NumReferenced   float64                    `json:"numReferenced,required"`
	NumReferences   float64                    `json:"numReferences,required"`
	RawID           string                     `json:"rawId,required"`
	Referenced      []string                   `json:"referenced,required"`
	ReferencedIDs   []float64                  `json:"referencedIds,required"`
	References      []string                   `json:"references,required"`
	ReferencesIDs   []float64                  `json:"referencesIds,required"`
	Tags            []string                   `json:"tags,required"`
	TargetCountry   string                     `json:"targetCountry,required"`
	TargetIndustry  string                     `json:"targetIndustry,required"`
	TLP             string                     `json:"tlp,required"`
	UUID            string                     `json:"uuid,required"`
	Insight         string                     `json:"insight"`
	ReleasabilityID string                     `json:"releasabilityId"`
	JSON            threatEventNewResponseJSON `json:"-"`
}

// threatEventNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventNewResponse]
type threatEventNewResponseJSON struct {
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	DatasetID       apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	HasChildren     apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventListResponse struct {
	Attacker        string                      `json:"attacker,required"`
	AttackerCountry string                      `json:"attackerCountry,required"`
	Category        string                      `json:"category,required"`
	DatasetID       string                      `json:"datasetId,required"`
	Date            string                      `json:"date,required"`
	Event           string                      `json:"event,required"`
	HasChildren     bool                        `json:"hasChildren,required"`
	Indicator       string                      `json:"indicator,required"`
	IndicatorType   string                      `json:"indicatorType,required"`
	IndicatorTypeID float64                     `json:"indicatorTypeId,required"`
	KillChain       float64                     `json:"killChain,required"`
	MitreAttack     []string                    `json:"mitreAttack,required"`
	NumReferenced   float64                     `json:"numReferenced,required"`
	NumReferences   float64                     `json:"numReferences,required"`
	RawID           string                      `json:"rawId,required"`
	Referenced      []string                    `json:"referenced,required"`
	ReferencedIDs   []float64                   `json:"referencedIds,required"`
	References      []string                    `json:"references,required"`
	ReferencesIDs   []float64                   `json:"referencesIds,required"`
	Tags            []string                    `json:"tags,required"`
	TargetCountry   string                      `json:"targetCountry,required"`
	TargetIndustry  string                      `json:"targetIndustry,required"`
	TLP             string                      `json:"tlp,required"`
	UUID            string                      `json:"uuid,required"`
	Insight         string                      `json:"insight"`
	ReleasabilityID string                      `json:"releasabilityId"`
	JSON            threatEventListResponseJSON `json:"-"`
}

// threatEventListResponseJSON contains the JSON metadata for the struct
// [ThreatEventListResponse]
type threatEventListResponseJSON struct {
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	DatasetID       apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	HasChildren     apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventListResponseJSON) RawJSON() string {
	return r.raw
}

// Detailed result of bulk event creation with auto-tag management
type ThreatEventBulkNewResponse struct {
	// Number of events created
	CreatedEventsCount float64 `json:"createdEventsCount,required"`
	// Number of new tags created in SoT
	CreatedTagsCount float64 `json:"createdTagsCount,required"`
	// Number of errors encountered
	ErrorCount float64 `json:"errorCount,required"`
	// Number of indicators queued for async processing
	QueuedIndicatorsCount float64 `json:"queuedIndicatorsCount,required"`
	// Correlation ID for async indicator processing
	CreateBulkEventsRequestID string `json:"createBulkEventsRequestId" format:"uuid"`
	// Array of created events with UUIDs and shard locations. Only present when
	// includeCreatedEvents=true
	CreatedEvents []ThreatEventBulkNewResponseCreatedEvent `json:"createdEvents"`
	// Array of error details
	Errors []ThreatEventBulkNewResponseError `json:"errors"`
	JSON   threatEventBulkNewResponseJSON    `json:"-"`
}

// threatEventBulkNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventBulkNewResponse]
type threatEventBulkNewResponseJSON struct {
	CreatedEventsCount        apijson.Field
	CreatedTagsCount          apijson.Field
	ErrorCount                apijson.Field
	QueuedIndicatorsCount     apijson.Field
	CreateBulkEventsRequestID apijson.Field
	CreatedEvents             apijson.Field
	Errors                    apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *ThreatEventBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventBulkNewResponseCreatedEvent struct {
	// Original index in the input data array
	EventIndex float64 `json:"eventIndex,required"`
	// Dataset ID of the shard where the event was created
	ShardID string `json:"shardId,required"`
	// UUID of the created event
	UUID string                                     `json:"uuid,required" format:"uuid"`
	JSON threatEventBulkNewResponseCreatedEventJSON `json:"-"`
}

// threatEventBulkNewResponseCreatedEventJSON contains the JSON metadata for the
// struct [ThreatEventBulkNewResponseCreatedEvent]
type threatEventBulkNewResponseCreatedEventJSON struct {
	EventIndex  apijson.Field
	ShardID     apijson.Field
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventBulkNewResponseCreatedEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventBulkNewResponseCreatedEventJSON) RawJSON() string {
	return r.raw
}

type ThreatEventBulkNewResponseError struct {
	// Error message
	Error string `json:"error,required"`
	// Index of the event that caused the error
	EventIndex float64                             `json:"eventIndex,required"`
	JSON       threatEventBulkNewResponseErrorJSON `json:"-"`
}

// threatEventBulkNewResponseErrorJSON contains the JSON metadata for the struct
// [ThreatEventBulkNewResponseError]
type threatEventBulkNewResponseErrorJSON struct {
	Error       apijson.Field
	EventIndex  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventBulkNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventBulkNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type ThreatEventEditResponse struct {
	Attacker        string                      `json:"attacker,required"`
	AttackerCountry string                      `json:"attackerCountry,required"`
	Category        string                      `json:"category,required"`
	DatasetID       string                      `json:"datasetId,required"`
	Date            string                      `json:"date,required"`
	Event           string                      `json:"event,required"`
	HasChildren     bool                        `json:"hasChildren,required"`
	Indicator       string                      `json:"indicator,required"`
	IndicatorType   string                      `json:"indicatorType,required"`
	IndicatorTypeID float64                     `json:"indicatorTypeId,required"`
	KillChain       float64                     `json:"killChain,required"`
	MitreAttack     []string                    `json:"mitreAttack,required"`
	NumReferenced   float64                     `json:"numReferenced,required"`
	NumReferences   float64                     `json:"numReferences,required"`
	RawID           string                      `json:"rawId,required"`
	Referenced      []string                    `json:"referenced,required"`
	ReferencedIDs   []float64                   `json:"referencedIds,required"`
	References      []string                    `json:"references,required"`
	ReferencesIDs   []float64                   `json:"referencesIds,required"`
	Tags            []string                    `json:"tags,required"`
	TargetCountry   string                      `json:"targetCountry,required"`
	TargetIndustry  string                      `json:"targetIndustry,required"`
	TLP             string                      `json:"tlp,required"`
	UUID            string                      `json:"uuid,required"`
	Insight         string                      `json:"insight"`
	ReleasabilityID string                      `json:"releasabilityId"`
	JSON            threatEventEditResponseJSON `json:"-"`
}

// threatEventEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventEditResponse]
type threatEventEditResponseJSON struct {
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	DatasetID       apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	HasChildren     apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventGetResponse struct {
	Attacker        string                     `json:"attacker,required"`
	AttackerCountry string                     `json:"attackerCountry,required"`
	Category        string                     `json:"category,required"`
	DatasetID       string                     `json:"datasetId,required"`
	Date            string                     `json:"date,required"`
	Event           string                     `json:"event,required"`
	HasChildren     bool                       `json:"hasChildren,required"`
	Indicator       string                     `json:"indicator,required"`
	IndicatorType   string                     `json:"indicatorType,required"`
	IndicatorTypeID float64                    `json:"indicatorTypeId,required"`
	KillChain       float64                    `json:"killChain,required"`
	MitreAttack     []string                   `json:"mitreAttack,required"`
	NumReferenced   float64                    `json:"numReferenced,required"`
	NumReferences   float64                    `json:"numReferences,required"`
	RawID           string                     `json:"rawId,required"`
	Referenced      []string                   `json:"referenced,required"`
	ReferencedIDs   []float64                  `json:"referencedIds,required"`
	References      []string                   `json:"references,required"`
	ReferencesIDs   []float64                  `json:"referencesIds,required"`
	Tags            []string                   `json:"tags,required"`
	TargetCountry   string                     `json:"targetCountry,required"`
	TargetIndustry  string                     `json:"targetIndustry,required"`
	TLP             string                     `json:"tlp,required"`
	UUID            string                     `json:"uuid,required"`
	Insight         string                     `json:"insight"`
	ReleasabilityID string                     `json:"releasabilityId"`
	JSON            threatEventGetResponseJSON `json:"-"`
}

// threatEventGetResponseJSON contains the JSON metadata for the struct
// [ThreatEventGetResponse]
type threatEventGetResponseJSON struct {
	Attacker        apijson.Field
	AttackerCountry apijson.Field
	Category        apijson.Field
	DatasetID       apijson.Field
	Date            apijson.Field
	Event           apijson.Field
	HasChildren     apijson.Field
	Indicator       apijson.Field
	IndicatorType   apijson.Field
	IndicatorTypeID apijson.Field
	KillChain       apijson.Field
	MitreAttack     apijson.Field
	NumReferenced   apijson.Field
	NumReferences   apijson.Field
	RawID           apijson.Field
	Referenced      apijson.Field
	ReferencedIDs   apijson.Field
	References      apijson.Field
	ReferencesIDs   apijson.Field
	Tags            apijson.Field
	TargetCountry   apijson.Field
	TargetIndustry  apijson.Field
	TLP             apijson.Field
	UUID            apijson.Field
	Insight         apijson.Field
	ReleasabilityID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ThreatEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventNewParams struct {
	// Account ID.
	PathAccountID   param.Field[string]                  `path:"account_id,required"`
	Category        param.Field[string]                  `json:"category,required"`
	Date            param.Field[time.Time]               `json:"date,required" format:"date-time"`
	Event           param.Field[string]                  `json:"event,required"`
	Raw             param.Field[ThreatEventNewParamsRaw] `json:"raw,required"`
	TLP             param.Field[string]                  `json:"tlp,required"`
	BodyAccountID   param.Field[float64]                 `json:"accountId"`
	Attacker        param.Field[string]                  `json:"attacker"`
	AttackerCountry param.Field[string]                  `json:"attackerCountry"`
	DatasetID       param.Field[string]                  `json:"datasetId"`
	Indicator       param.Field[string]                  `json:"indicator"`
	// Array of indicators for this event. Supports multiple indicators per event for
	// complex scenarios.
	Indicators     param.Field[[]ThreatEventNewParamsIndicator] `json:"indicators"`
	IndicatorType  param.Field[string]                          `json:"indicatorType"`
	Insight        param.Field[string]                          `json:"insight"`
	Tags           param.Field[[]string]                        `json:"tags"`
	TargetCountry  param.Field[string]                          `json:"targetCountry"`
	TargetIndustry param.Field[string]                          `json:"targetIndustry"`
}

func (r ThreatEventNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventNewParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data,required"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventNewParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventNewParamsIndicator struct {
	// The type of indicator (e.g., DOMAIN, IP, JA3, HASH)
	IndicatorType param.Field[string] `json:"indicatorType,required"`
	// The indicator value (e.g., domain name, IP address, hash)
	Value param.Field[string] `json:"value,required"`
}

func (r ThreatEventNewParamsIndicator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// Cursor for pagination. When provided, filters are embedded in the cursor so you
	// only need to pass cursor and pageSize. Returned in the previous response's
	// result_info.cursor field. Use cursor-based pagination for deep pagination
	// (beyond 100,000 records) or for optimal performance.
	Cursor       param.Field[string]                      `query:"cursor"`
	DatasetID    param.Field[[]string]                    `query:"datasetId"`
	ForceRefresh param.Field[bool]                        `query:"forceRefresh"`
	Format       param.Field[ThreatEventListParamsFormat] `query:"format"`
	Order        param.Field[ThreatEventListParamsOrder]  `query:"order"`
	OrderBy      param.Field[string]                      `query:"orderBy"`
	// Page number (1-indexed) for offset-based pagination. Limited to offset of
	// 100,000 records. For deep pagination, use cursor-based pagination instead.
	Page param.Field[float64] `query:"page"`
	// Number of results per page. Maximum 25,000.
	PageSize param.Field[float64]                       `query:"pageSize"`
	Search   param.Field[[]ThreatEventListParamsSearch] `query:"search"`
}

// URLQuery serializes [ThreatEventListParams]'s query parameters as `url.Values`.
func (r ThreatEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventListParamsFormat string

const (
	ThreatEventListParamsFormatJson  ThreatEventListParamsFormat = "json"
	ThreatEventListParamsFormatStix2 ThreatEventListParamsFormat = "stix2"
)

func (r ThreatEventListParamsFormat) IsKnown() bool {
	switch r {
	case ThreatEventListParamsFormatJson, ThreatEventListParamsFormatStix2:
		return true
	}
	return false
}

type ThreatEventListParamsOrder string

const (
	ThreatEventListParamsOrderAsc  ThreatEventListParamsOrder = "asc"
	ThreatEventListParamsOrderDesc ThreatEventListParamsOrder = "desc"
)

func (r ThreatEventListParamsOrder) IsKnown() bool {
	switch r {
	case ThreatEventListParamsOrderAsc, ThreatEventListParamsOrderDesc:
		return true
	}
	return false
}

type ThreatEventListParamsSearch struct {
	Field param.Field[string]                                `query:"field"`
	Op    param.Field[ThreatEventListParamsSearchOp]         `query:"op"`
	Value param.Field[ThreatEventListParamsSearchValueUnion] `query:"value"`
}

// URLQuery serializes [ThreatEventListParamsSearch]'s query parameters as
// `url.Values`.
func (r ThreatEventListParamsSearch) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventListParamsSearchOp string

const (
	ThreatEventListParamsSearchOpEquals     ThreatEventListParamsSearchOp = "equals"
	ThreatEventListParamsSearchOpNot        ThreatEventListParamsSearchOp = "not"
	ThreatEventListParamsSearchOpGt         ThreatEventListParamsSearchOp = "gt"
	ThreatEventListParamsSearchOpGte        ThreatEventListParamsSearchOp = "gte"
	ThreatEventListParamsSearchOpLt         ThreatEventListParamsSearchOp = "lt"
	ThreatEventListParamsSearchOpLte        ThreatEventListParamsSearchOp = "lte"
	ThreatEventListParamsSearchOpLike       ThreatEventListParamsSearchOp = "like"
	ThreatEventListParamsSearchOpContains   ThreatEventListParamsSearchOp = "contains"
	ThreatEventListParamsSearchOpStartsWith ThreatEventListParamsSearchOp = "startsWith"
	ThreatEventListParamsSearchOpEndsWith   ThreatEventListParamsSearchOp = "endsWith"
	ThreatEventListParamsSearchOpIn         ThreatEventListParamsSearchOp = "in"
	ThreatEventListParamsSearchOpFind       ThreatEventListParamsSearchOp = "find"
)

func (r ThreatEventListParamsSearchOp) IsKnown() bool {
	switch r {
	case ThreatEventListParamsSearchOpEquals, ThreatEventListParamsSearchOpNot, ThreatEventListParamsSearchOpGt, ThreatEventListParamsSearchOpGte, ThreatEventListParamsSearchOpLt, ThreatEventListParamsSearchOpLte, ThreatEventListParamsSearchOpLike, ThreatEventListParamsSearchOpContains, ThreatEventListParamsSearchOpStartsWith, ThreatEventListParamsSearchOpEndsWith, ThreatEventListParamsSearchOpIn, ThreatEventListParamsSearchOpFind:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat],
// [cloudforce_one.ThreatEventListParamsSearchValueArray].
type ThreatEventListParamsSearchValueUnion interface {
	ImplementsThreatEventListParamsSearchValueUnion()
}

type ThreatEventListParamsSearchValueArray []ThreatEventListParamsSearchValueArrayItemUnion

func (r ThreatEventListParamsSearchValueArray) ImplementsThreatEventListParamsSearchValueUnion() {}

// Satisfied by [shared.UnionString], [shared.UnionFloat].
type ThreatEventListParamsSearchValueArrayItemUnion interface {
	ImplementsThreatEventListParamsSearchValueArrayItemUnion()
}

type ThreatEventBulkNewParams struct {
	// Account ID.
	AccountID param.Field[string]                         `path:"account_id,required"`
	Data      param.Field[[]ThreatEventBulkNewParamsData] `json:"data,required"`
	DatasetID param.Field[string]                         `json:"datasetId,required"`
	// When true, response includes array of created event UUIDs and shard IDs. Useful
	// for tracking which events were created and where.
	IncludeCreatedEvents param.Field[bool] `json:"includeCreatedEvents"`
}

func (r ThreatEventBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventBulkNewParamsData struct {
	Category        param.Field[string]                          `json:"category,required"`
	Date            param.Field[time.Time]                       `json:"date,required" format:"date-time"`
	Event           param.Field[string]                          `json:"event,required"`
	Raw             param.Field[ThreatEventBulkNewParamsDataRaw] `json:"raw,required"`
	TLP             param.Field[string]                          `json:"tlp,required"`
	AccountID       param.Field[float64]                         `json:"accountId"`
	Attacker        param.Field[string]                          `json:"attacker"`
	AttackerCountry param.Field[string]                          `json:"attackerCountry"`
	DatasetID       param.Field[string]                          `json:"datasetId"`
	Indicator       param.Field[string]                          `json:"indicator"`
	// Array of indicators for this event. Supports multiple indicators per event for
	// complex scenarios.
	Indicators     param.Field[[]ThreatEventBulkNewParamsDataIndicator] `json:"indicators"`
	IndicatorType  param.Field[string]                                  `json:"indicatorType"`
	Insight        param.Field[string]                                  `json:"insight"`
	Tags           param.Field[[]string]                                `json:"tags"`
	TargetCountry  param.Field[string]                                  `json:"targetCountry"`
	TargetIndustry param.Field[string]                                  `json:"targetIndustry"`
}

func (r ThreatEventBulkNewParamsData) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventBulkNewParamsDataRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data,required"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventBulkNewParamsDataRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventBulkNewParamsDataIndicator struct {
	// The type of indicator (e.g., DOMAIN, IP, JA3, HASH)
	IndicatorType param.Field[string] `json:"indicatorType,required"`
	// The indicator value (e.g., domain name, IP address, hash)
	Value param.Field[string] `json:"value,required"`
}

func (r ThreatEventBulkNewParamsDataIndicator) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventEditParams struct {
	// Account ID.
	AccountID       param.Field[string]                   `path:"account_id,required"`
	Attacker        param.Field[string]                   `json:"attacker"`
	AttackerCountry param.Field[string]                   `json:"attackerCountry"`
	Category        param.Field[string]                   `json:"category"`
	CreatedAt       param.Field[time.Time]                `json:"createdAt" format:"date-time"`
	DatasetID       param.Field[string]                   `json:"datasetId"`
	Date            param.Field[time.Time]                `json:"date" format:"date-time"`
	Event           param.Field[string]                   `json:"event"`
	Indicator       param.Field[string]                   `json:"indicator"`
	IndicatorType   param.Field[string]                   `json:"indicatorType"`
	Insight         param.Field[string]                   `json:"insight"`
	Raw             param.Field[ThreatEventEditParamsRaw] `json:"raw"`
	TargetCountry   param.Field[string]                   `json:"targetCountry"`
	TargetIndustry  param.Field[string]                   `json:"targetIndustry"`
	TLP             param.Field[string]                   `json:"tlp"`
}

func (r ThreatEventEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventEditParamsRaw struct {
	Data   param.Field[map[string]interface{}] `json:"data"`
	Source param.Field[string]                 `json:"source"`
	TLP    param.Field[string]                 `json:"tlp"`
}

func (r ThreatEventEditParamsRaw) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
