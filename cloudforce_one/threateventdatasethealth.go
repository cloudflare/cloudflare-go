// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// ThreatEventDatasetHealthService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventDatasetHealthService] method instead.
type ThreatEventDatasetHealthService struct {
	Options []option.RequestOption
}

// NewThreatEventDatasetHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatEventDatasetHealthService(opts ...option.RequestOption) (r *ThreatEventDatasetHealthService) {
	r = &ThreatEventDatasetHealthService{}
	r.Options = opts
	return
}

// Benchmark Durable Object warmup
func (r *ThreatEventDatasetHealthService) Get(ctx context.Context, datasetID string, query ThreatEventDatasetHealthGetParams, opts ...option.RequestOption) (res *ThreatEventDatasetHealthGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if datasetID == "" {
		err = errors.New("missing required dataset_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/dataset/%s/health", query.AccountID, datasetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ThreatEventDatasetHealthGetResponse struct {
	Properties ThreatEventDatasetHealthGetResponseProperties `json:"properties,required"`
	Type       string                                        `json:"type,required"`
	JSON       threatEventDatasetHealthGetResponseJSON       `json:"-"`
}

// threatEventDatasetHealthGetResponseJSON contains the JSON metadata for the
// struct [ThreatEventDatasetHealthGetResponse]
type threatEventDatasetHealthGetResponseJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponseProperties struct {
	DurationMs     ThreatEventDatasetHealthGetResponsePropertiesDurationMs     `json:"durationMs,required"`
	Ok             ThreatEventDatasetHealthGetResponsePropertiesOk             `json:"ok,required"`
	Shards         ThreatEventDatasetHealthGetResponsePropertiesShards         `json:"shards,required"`
	TotalShards    ThreatEventDatasetHealthGetResponsePropertiesTotalShards    `json:"totalShards,required"`
	TotalSizeBytes ThreatEventDatasetHealthGetResponsePropertiesTotalSizeBytes `json:"totalSizeBytes,required"`
	TotalSizeMB    ThreatEventDatasetHealthGetResponsePropertiesTotalSizeMB    `json:"totalSizeMB,required"`
	JSON           threatEventDatasetHealthGetResponsePropertiesJSON           `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesJSON contains the JSON metadata for
// the struct [ThreatEventDatasetHealthGetResponseProperties]
type threatEventDatasetHealthGetResponsePropertiesJSON struct {
	DurationMs     apijson.Field
	Ok             apijson.Field
	Shards         apijson.Field
	TotalShards    apijson.Field
	TotalSizeBytes apijson.Field
	TotalSizeMB    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesDurationMs struct {
	Type string                                                      `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesDurationMsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesDurationMsJSON contains the JSON
// metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesDurationMs]
type threatEventDatasetHealthGetResponsePropertiesDurationMsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesDurationMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesDurationMsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesOk struct {
	Type string                                              `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesOkJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesOkJSON contains the JSON metadata
// for the struct [ThreatEventDatasetHealthGetResponsePropertiesOk]
type threatEventDatasetHealthGetResponsePropertiesOkJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesOk) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesOkJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShards struct {
	Items ThreatEventDatasetHealthGetResponsePropertiesShardsItems `json:"items,required"`
	Type  string                                                   `json:"type,required"`
	JSON  threatEventDatasetHealthGetResponsePropertiesShardsJSON  `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsJSON contains the JSON
// metadata for the struct [ThreatEventDatasetHealthGetResponsePropertiesShards]
type threatEventDatasetHealthGetResponsePropertiesShardsJSON struct {
	Items       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShards) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItems struct {
	Properties ThreatEventDatasetHealthGetResponsePropertiesShardsItemsProperties `json:"properties,required"`
	Type       string                                                             `json:"type,required"`
	JSON       threatEventDatasetHealthGetResponsePropertiesShardsItemsJSON       `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsJSON contains the JSON
// metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItems]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsJSON struct {
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItems) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsProperties struct {
	DatasetID     ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetID     `json:"datasetId,required"`
	Date          ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDate          `json:"date,required"`
	HealthCheckMs ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMs `json:"healthCheckMs,required"`
	PageCount     ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCount     `json:"pageCount,required"`
	PageSize      ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSize      `json:"pageSize,required"`
	SizeBytes     ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytes     `json:"sizeBytes,required"`
	SizeMB        ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMB        `json:"sizeMB,required"`
	StartupMs     ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMs     `json:"startupMs,required"`
	TableStats    ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStats    `json:"tableStats,required"`
	TimedOut      ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOut      `json:"timedOut,required"`
	TotalMs       ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMs       `json:"totalMs,required"`
	JSON          threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesJSON          `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesJSON contains
// the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsProperties]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesJSON struct {
	DatasetID     apijson.Field
	Date          apijson.Field
	HealthCheckMs apijson.Field
	PageCount     apijson.Field
	PageSize      apijson.Field
	SizeBytes     apijson.Field
	SizeMB        apijson.Field
	StartupMs     apijson.Field
	TableStats    apijson.Field
	TimedOut      apijson.Field
	TotalMs       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetID struct {
	Type string                                                                          `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetIDJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetIDJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetID]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetIDJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDatasetIDJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDate struct {
	Type string                                                                     `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDateJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDateJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDate]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDateJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesDateJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMs struct {
	Type string                                                                              `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMsJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMs]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesHealthCheckMsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCount struct {
	Type string                                                                          `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCountJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCountJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCount]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCountJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageCountJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSize struct {
	Type string                                                                         `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSizeJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSizeJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSize]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSizeJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSize) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesPageSizeJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytes struct {
	Type string                                                                          `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytesJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytesJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytes]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeBytesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMB struct {
	Type string                                                                       `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMBJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMBJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMB]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMBJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesSizeMBJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMs struct {
	Type string                                                                          `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMsJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMs]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesStartupMsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStats struct {
	AdditionalProperties ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalProperties `json:"additionalProperties,required"`
	Type                 string                                                                                           `json:"type,required"`
	JSON                 threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsJSON                 `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStats]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsJSON struct {
	AdditionalProperties apijson.Field
	Type                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalProperties struct {
	Type string                                                                                               `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalPropertiesJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalPropertiesJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalProperties]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalPropertiesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTableStatsAdditionalPropertiesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOut struct {
	Type string                                                                         `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOutJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOutJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOut]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOutJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOut) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTimedOutJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMs struct {
	Type string                                                                        `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMsJSON
// contains the JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMs]
type threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesShardsItemsPropertiesTotalMsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesTotalShards struct {
	Type string                                                       `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesTotalShardsJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesTotalShardsJSON contains the JSON
// metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesTotalShards]
type threatEventDatasetHealthGetResponsePropertiesTotalShardsJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesTotalShards) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesTotalShardsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesTotalSizeBytes struct {
	Type string                                                          `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesTotalSizeBytesJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesTotalSizeBytesJSON contains the
// JSON metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesTotalSizeBytes]
type threatEventDatasetHealthGetResponsePropertiesTotalSizeBytesJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesTotalSizeBytes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesTotalSizeBytesJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetResponsePropertiesTotalSizeMB struct {
	Type string                                                       `json:"type,required"`
	JSON threatEventDatasetHealthGetResponsePropertiesTotalSizeMBJSON `json:"-"`
}

// threatEventDatasetHealthGetResponsePropertiesTotalSizeMBJSON contains the JSON
// metadata for the struct
// [ThreatEventDatasetHealthGetResponsePropertiesTotalSizeMB]
type threatEventDatasetHealthGetResponsePropertiesTotalSizeMBJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventDatasetHealthGetResponsePropertiesTotalSizeMB) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventDatasetHealthGetResponsePropertiesTotalSizeMBJSON) RawJSON() string {
	return r.raw
}

type ThreatEventDatasetHealthGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
