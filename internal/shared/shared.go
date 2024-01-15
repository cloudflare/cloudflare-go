// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
)

type VectorizeCreateIndexResponse struct {
	Config VectorizeCreateIndexResponseConfig `json:"config"`
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedOn interface{} `json:"created_on"`
	// Specifies the description of the index.
	Description string `json:"description"`
	// Specifies the timestamp the resource was modified as an ISO8601 string.
	ModifiedOn interface{}                      `json:"modified_on"`
	Name       string                           `json:"name"`
	JSON       vectorizeCreateIndexResponseJSON `json:"-"`
}

// vectorizeCreateIndexResponseJSON contains the JSON metadata for the struct
// [VectorizeCreateIndexResponse]
type vectorizeCreateIndexResponseJSON struct {
	Config      apijson.Field
	CreatedOn   apijson.Field
	Description apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeCreateIndexResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VectorizeCreateIndexResponseConfig struct {
	// Specifies the number of dimensions for the index
	Dimensions int64 `json:"dimensions,required"`
	// Specifies the type of metric to use calculating distance.
	Metric VectorizeCreateIndexResponseConfigMetric `json:"metric,required"`
	JSON   vectorizeCreateIndexResponseConfigJSON   `json:"-"`
}

// vectorizeCreateIndexResponseConfigJSON contains the JSON metadata for the struct
// [VectorizeCreateIndexResponseConfig]
type vectorizeCreateIndexResponseConfigJSON struct {
	Dimensions  apijson.Field
	Metric      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VectorizeCreateIndexResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Specifies the type of metric to use calculating distance.
type VectorizeCreateIndexResponseConfigMetric string

const (
	VectorizeCreateIndexResponseConfigMetricCosine     VectorizeCreateIndexResponseConfigMetric = "cosine"
	VectorizeCreateIndexResponseConfigMetricEuclidean  VectorizeCreateIndexResponseConfigMetric = "euclidean"
	VectorizeCreateIndexResponseConfigMetricDotProduct VectorizeCreateIndexResponseConfigMetric = "dot-product"
)

type WaitingRoomSettingsResponse struct {
	Result WaitingRoomSettingsResponseResult `json:"result"`
	JSON   waitingRoomSettingsResponseJSON   `json:"-"`
}

// waitingRoomSettingsResponseJSON contains the JSON metadata for the struct
// [WaitingRoomSettingsResponse]
type waitingRoomSettingsResponseJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WaitingRoomSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WaitingRoomSettingsResponseResult struct {
	// Whether to allow verified search engine crawlers to bypass all waiting rooms on
	// this zone. Verified search engine crawlers will not be tracked or counted by the
	// waiting room system, and will not appear in waiting room analytics.
	SearchEngineCrawlerBypass bool                                  `json:"search_engine_crawler_bypass,required"`
	JSON                      waitingRoomSettingsResponseResultJSON `json:"-"`
}

// waitingRoomSettingsResponseResultJSON contains the JSON metadata for the struct
// [WaitingRoomSettingsResponseResult]
type waitingRoomSettingsResponseResultJSON struct {
	SearchEngineCrawlerBypass apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *WaitingRoomSettingsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
