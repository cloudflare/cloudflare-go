// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDNSService] method instead.
type DNSService struct {
	Options   []option.RequestOption
	Records   *RecordService
	Analytics *AnalyticsService
	Firewall  *FirewallService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.Records = NewRecordService(opts...)
	r.Analytics = NewAnalyticsService(opts...)
	r.Firewall = NewFirewallService(opts...)
	return
}

type UnnamedSchemaRef6595695ff25b0614667b25f66b7bbaba struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                                            `json:"metrics,required"`
	JSON    unnamedSchemaRef6595695ff25b0614667b25f66b7bbabaJSON `json:"-"`
}

// unnamedSchemaRef6595695ff25b0614667b25f66b7bbabaJSON contains the JSON metadata
// for the struct [UnnamedSchemaRef6595695ff25b0614667b25f66b7bbaba]
type unnamedSchemaRef6595695ff25b0614667b25f66b7bbabaJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef6595695ff25b0614667b25f66b7bbaba) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef6595695ff25b0614667b25f66b7bbabaJSON) RawJSON() string {
	return r.raw
}

type UnnamedSchemaRef65be9614de145bf4a58d0fddf46df7ca []interface{}

type UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2 struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                             `json:"sort"`
	JSON unnamedSchemaRef85b45d163202bbab7456da6b346d9fe2JSON `json:"-"`
}

// unnamedSchemaRef85b45d163202bbab7456da6b346d9fe2JSON contains the JSON metadata
// for the struct [UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2]
type unnamedSchemaRef85b45d163202bbab7456da6b346d9fe2JSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r unnamedSchemaRef85b45d163202bbab7456da6b346d9fe2JSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta string

const (
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaAll        UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "all"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaAuto       UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "auto"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaYear       UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "year"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaQuarter    UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "quarter"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaMonth      UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "month"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaWeek       UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "week"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaDay        UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "day"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaHour       UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "hour"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaDekaminute UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "dekaminute"
	UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaMinute     UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta = "minute"
)

func (r UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDelta) IsKnown() bool {
	switch r {
	case UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaAll, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaAuto, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaYear, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaQuarter, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaMonth, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaWeek, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaDay, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaHour, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaDekaminute, UnnamedSchemaRef85b45d163202bbab7456da6b346d9fe2TimeDeltaMinute:
		return true
	}
	return false
}
