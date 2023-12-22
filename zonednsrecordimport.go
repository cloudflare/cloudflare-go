// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneDNSRecordImportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDNSRecordImportService]
// method instead.
type ZoneDNSRecordImportService struct {
	Options []option.RequestOption
}

// NewZoneDNSRecordImportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSRecordImportService(opts ...option.RequestOption) (r *ZoneDNSRecordImportService) {
	r = &ZoneDNSRecordImportService{}
	r.Options = opts
	return
}

// You can upload your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint. It assumes that cURL is called from a location with bind_config.txt
// (valid BIND config) present.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *ZoneDNSRecordImportService) DNSRecordsForAZoneImportDNSRecords(ctx context.Context, zoneIdentifier string, body ZoneDNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams, opts ...option.RequestOption) (res *DNSResponseImportScan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/import", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type DNSResponseImportScan struct {
	Errors   []DNSResponseImportScanError   `json:"errors"`
	Messages []DNSResponseImportScanMessage `json:"messages"`
	Result   DNSResponseImportScanResult    `json:"result"`
	// Whether the API call was successful
	Success DNSResponseImportScanSuccess `json:"success"`
	Timing  DNSResponseImportScanTiming  `json:"timing"`
	JSON    dnsResponseImportScanJSON    `json:"-"`
}

// dnsResponseImportScanJSON contains the JSON metadata for the struct
// [DNSResponseImportScan]
type dnsResponseImportScanJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseImportScan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseImportScanError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    dnsResponseImportScanErrorJSON `json:"-"`
}

// dnsResponseImportScanErrorJSON contains the JSON metadata for the struct
// [DNSResponseImportScanError]
type dnsResponseImportScanErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseImportScanError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseImportScanMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    dnsResponseImportScanMessageJSON `json:"-"`
}

// dnsResponseImportScanMessageJSON contains the JSON metadata for the struct
// [DNSResponseImportScanMessage]
type dnsResponseImportScanMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseImportScanMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSResponseImportScanResult struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                         `json:"total_records_parsed"`
	JSON               dnsResponseImportScanResultJSON `json:"-"`
}

// dnsResponseImportScanResultJSON contains the JSON metadata for the struct
// [DNSResponseImportScanResult]
type dnsResponseImportScanResultJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSResponseImportScanResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSResponseImportScanSuccess bool

const (
	DNSResponseImportScanSuccessTrue DNSResponseImportScanSuccess = true
)

type DNSResponseImportScanTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                       `json:"start_time" format:"date-time"`
	JSON      dnsResponseImportScanTimingJSON `json:"-"`
}

// dnsResponseImportScanTimingJSON contains the JSON metadata for the struct
// [DNSResponseImportScanTiming]
type dnsResponseImportScanTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSResponseImportScanTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams struct {
	// BIND config to upload.
	File param.Field[string] `json:"file,required"`
	// Whether or not proxiable records should receive the performance and security
	// benefits of Cloudflare.
	Proxied param.Field[bool] `json:"proxied"`
}

func (r ZoneDNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
