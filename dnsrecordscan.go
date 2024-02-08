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

// DNSRecordScanService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSRecordScanService] method
// instead.
type DNSRecordScanService struct {
	Options []option.RequestOption
}

// NewDNSRecordScanService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSRecordScanService(opts ...option.RequestOption) (r *DNSRecordScanService) {
	r = &DNSRecordScanService{}
	r.Options = opts
	return
}

// Scan for common DNS records on your domain and automatically add them to your
// zone. Useful if you haven't updated your nameservers yet.
func (r *DNSRecordScanService) DNSRecordsForAZoneScanDNSRecords(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/scan", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                                                   `json:"total_records_parsed"`
	JSON               dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON `json:"-"`
}

// dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON contains the JSON
// metadata for the struct [DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse]
type dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelope struct {
	Errors   []DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeSuccess `json:"success,required"`
	Timing  DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTiming  `json:"timing"`
	JSON    dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelope]
type dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrors struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrorsJSON contains
// the JSON metadata for the struct
// [DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrors]
type dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessages struct {
	Code    int64                                                                     `json:"code,required"`
	Message string                                                                    `json:"message,required"`
	JSON    dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessages]
type dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeSuccess bool

const (
	DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeSuccessTrue DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeSuccess = true
)

type DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                                                               `json:"start_time" format:"date-time"`
	JSON      dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTimingJSON `json:"-"`
}

// dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTimingJSON contains
// the JSON metadata for the struct
// [DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTiming]
type dnsRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordScanDNSRecordsForAZoneScanDNSRecordsResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
