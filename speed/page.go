// File generated from our OpenAPI spec by Stainless.

package speed

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PageService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPageService] method instead.
type PageService struct {
	Options []option.RequestOption
}

// NewPageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPageService(opts ...option.RequestOption) (r *PageService) {
	r = &PageService{}
	r.Options = opts
	return
}

// Lists all webpages which have been tested.
func (r *PageService) List(ctx context.Context, query PageListParams, opts ...option.RequestOption) (res *[]PageListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageListResponseEnvelope
	path := fmt.Sprintf("zones/%s/speed_api/pages", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PageListResponse struct {
	// A test region with a label.
	Region PageListResponseRegion `json:"region"`
	// The frequency of the test.
	ScheduleFrequency PageListResponseScheduleFrequency `json:"scheduleFrequency"`
	Tests             []ObservatoryPageTest             `json:"tests"`
	// A URL.
	URL  string               `json:"url"`
	JSON pageListResponseJSON `json:"-"`
}

// pageListResponseJSON contains the JSON metadata for the struct
// [PageListResponse]
type pageListResponseJSON struct {
	Region            apijson.Field
	ScheduleFrequency apijson.Field
	Tests             apijson.Field
	URL               apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PageListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseJSON) RawJSON() string {
	return r.raw
}

// A test region with a label.
type PageListResponseRegion struct {
	Label string `json:"label"`
	// A test region.
	Value PageListResponseRegionValue `json:"value"`
	JSON  pageListResponseRegionJSON  `json:"-"`
}

// pageListResponseRegionJSON contains the JSON metadata for the struct
// [PageListResponseRegion]
type pageListResponseRegionJSON struct {
	Label       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageListResponseRegion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseRegionJSON) RawJSON() string {
	return r.raw
}

// A test region.
type PageListResponseRegionValue string

const (
	PageListResponseRegionValueAsiaEast1           PageListResponseRegionValue = "asia-east1"
	PageListResponseRegionValueAsiaNortheast1      PageListResponseRegionValue = "asia-northeast1"
	PageListResponseRegionValueAsiaNortheast2      PageListResponseRegionValue = "asia-northeast2"
	PageListResponseRegionValueAsiaSouth1          PageListResponseRegionValue = "asia-south1"
	PageListResponseRegionValueAsiaSoutheast1      PageListResponseRegionValue = "asia-southeast1"
	PageListResponseRegionValueAustraliaSoutheast1 PageListResponseRegionValue = "australia-southeast1"
	PageListResponseRegionValueEuropeNorth1        PageListResponseRegionValue = "europe-north1"
	PageListResponseRegionValueEuropeSouthwest1    PageListResponseRegionValue = "europe-southwest1"
	PageListResponseRegionValueEuropeWest1         PageListResponseRegionValue = "europe-west1"
	PageListResponseRegionValueEuropeWest2         PageListResponseRegionValue = "europe-west2"
	PageListResponseRegionValueEuropeWest3         PageListResponseRegionValue = "europe-west3"
	PageListResponseRegionValueEuropeWest4         PageListResponseRegionValue = "europe-west4"
	PageListResponseRegionValueEuropeWest8         PageListResponseRegionValue = "europe-west8"
	PageListResponseRegionValueEuropeWest9         PageListResponseRegionValue = "europe-west9"
	PageListResponseRegionValueMeWest1             PageListResponseRegionValue = "me-west1"
	PageListResponseRegionValueSouthamericaEast1   PageListResponseRegionValue = "southamerica-east1"
	PageListResponseRegionValueUsCentral1          PageListResponseRegionValue = "us-central1"
	PageListResponseRegionValueUsEast1             PageListResponseRegionValue = "us-east1"
	PageListResponseRegionValueUsEast4             PageListResponseRegionValue = "us-east4"
	PageListResponseRegionValueUsSouth1            PageListResponseRegionValue = "us-south1"
	PageListResponseRegionValueUsWest1             PageListResponseRegionValue = "us-west1"
)

// The frequency of the test.
type PageListResponseScheduleFrequency string

const (
	PageListResponseScheduleFrequencyDaily  PageListResponseScheduleFrequency = "DAILY"
	PageListResponseScheduleFrequencyWeekly PageListResponseScheduleFrequency = "WEEKLY"
)

type PageListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageListResponseEnvelope struct {
	Result []PageListResponse           `json:"result"`
	JSON   pageListResponseEnvelopeJSON `json:"-"`
}

// pageListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageListResponseEnvelope]
type pageListResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
