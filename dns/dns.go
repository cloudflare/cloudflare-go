// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// DNSService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSService] method instead.
type DNSService struct {
	Options       []option.RequestOption
	DNSSEC        *DNSSECService
	Records       *RecordService
	Settings      *SettingService
	Analytics     *AnalyticsService
	ZoneTransfers *ZoneTransferService
}

// NewDNSService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSService(opts ...option.RequestOption) (r *DNSService) {
	r = &DNSService{}
	r.Options = opts
	r.DNSSEC = NewDNSSECService(opts...)
	r.Records = NewRecordService(opts...)
	r.Settings = NewSettingService(opts...)
	r.Analytics = NewAnalyticsService(opts...)
	r.ZoneTransfers = NewZoneTransferService(opts...)
	return
}

type DNSAnalyticsNominalMetric []interface{}

type DNSAnalyticsQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta DNSAnalyticsQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string              `json:"sort"`
	JSON dnsAnalyticsQueryJSON `json:"-"`
}

// dnsAnalyticsQueryJSON contains the JSON metadata for the struct
// [DNSAnalyticsQuery]
type dnsAnalyticsQueryJSON struct {
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

func (r *DNSAnalyticsQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsQueryJSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type DNSAnalyticsQueryTimeDelta string

const (
	DNSAnalyticsQueryTimeDeltaAll        DNSAnalyticsQueryTimeDelta = "all"
	DNSAnalyticsQueryTimeDeltaAuto       DNSAnalyticsQueryTimeDelta = "auto"
	DNSAnalyticsQueryTimeDeltaYear       DNSAnalyticsQueryTimeDelta = "year"
	DNSAnalyticsQueryTimeDeltaQuarter    DNSAnalyticsQueryTimeDelta = "quarter"
	DNSAnalyticsQueryTimeDeltaMonth      DNSAnalyticsQueryTimeDelta = "month"
	DNSAnalyticsQueryTimeDeltaWeek       DNSAnalyticsQueryTimeDelta = "week"
	DNSAnalyticsQueryTimeDeltaDay        DNSAnalyticsQueryTimeDelta = "day"
	DNSAnalyticsQueryTimeDeltaHour       DNSAnalyticsQueryTimeDelta = "hour"
	DNSAnalyticsQueryTimeDeltaDekaminute DNSAnalyticsQueryTimeDelta = "dekaminute"
	DNSAnalyticsQueryTimeDeltaMinute     DNSAnalyticsQueryTimeDelta = "minute"
)

func (r DNSAnalyticsQueryTimeDelta) IsKnown() bool {
	switch r {
	case DNSAnalyticsQueryTimeDeltaAll, DNSAnalyticsQueryTimeDeltaAuto, DNSAnalyticsQueryTimeDeltaYear, DNSAnalyticsQueryTimeDeltaQuarter, DNSAnalyticsQueryTimeDeltaMonth, DNSAnalyticsQueryTimeDeltaWeek, DNSAnalyticsQueryTimeDeltaDay, DNSAnalyticsQueryTimeDeltaHour, DNSAnalyticsQueryTimeDeltaDekaminute, DNSAnalyticsQueryTimeDeltaMinute:
		return true
	}
	return false
}
