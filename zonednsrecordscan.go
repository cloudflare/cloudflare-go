// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneDNSRecordScanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDNSRecordScanService] method
// instead.
type ZoneDNSRecordScanService struct {
	Options []option.RequestOption
}

// NewZoneDNSRecordScanService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSRecordScanService(opts ...option.RequestOption) (r *ZoneDNSRecordScanService) {
	r = &ZoneDNSRecordScanService{}
	r.Options = opts
	return
}

// Scan for common DNS records on your domain and automatically add them to your
// zone. Useful if you haven't updated your nameservers yet.
func (r *ZoneDNSRecordScanService) DNSRecordsForAZoneScanDNSRecords(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/scan", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse struct {
	Errors   []ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseError   `json:"errors"`
	Messages []ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessage `json:"messages"`
	Result   ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseSuccess `json:"success"`
	Timing  ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTiming  `json:"timing"`
	JSON    zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON    `json:"-"`
}

// zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON contains the JSON
// metadata for the struct
// [ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse]
type zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseErrorJSON `json:"-"`
}

// zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseErrorJSON contains the
// JSON metadata for the struct
// [ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseError]
type zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessageJSON `json:"-"`
}

// zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessageJSON contains
// the JSON metadata for the struct
// [ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessage]
type zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResult struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                                                             `json:"total_records_parsed"`
	JSON               zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResultJSON `json:"-"`
}

// zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResultJSON contains the
// JSON metadata for the struct
// [ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResult]
type zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResultJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseSuccess bool

const (
	ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseSuccessTrue ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseSuccess = true
)

type ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                                                           `json:"start_time" format:"date-time"`
	JSON      zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTimingJSON `json:"-"`
}

// zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTimingJSON contains the
// JSON metadata for the struct
// [ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTiming]
type zoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
