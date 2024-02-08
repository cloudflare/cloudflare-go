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

// DNSRecordImportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSRecordImportService] method
// instead.
type DNSRecordImportService struct {
	Options []option.RequestOption
}

// NewDNSRecordImportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSRecordImportService(opts ...option.RequestOption) (r *DNSRecordImportService) {
	r = &DNSRecordImportService{}
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
func (r *DNSRecordImportService) DNSRecordsForAZoneImportDNSRecords(ctx context.Context, zoneID string, body DNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams, opts ...option.RequestOption) (res *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_records/import", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse struct {
	// Number of DNS records added.
	RecsAdded float64 `json:"recs_added"`
	// Total number of DNS records parsed.
	TotalRecordsParsed float64                                                       `json:"total_records_parsed"`
	JSON               dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseJSON `json:"-"`
}

// dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseJSON contains the JSON
// metadata for the struct
// [DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse]
type dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseJSON struct {
	RecsAdded          apijson.Field
	TotalRecordsParsed apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams struct {
	// BIND config to import.
	//
	// **Tip:** When using cURL, a file can be uploaded using
	// `--form 'file=@bind_config.txt'`.
	File param.Field[string] `json:"file,required"`
	// Whether or not proxiable records should receive the performance and security
	// benefits of Cloudflare.
	//
	// The value should be either `true` or `false`.
	Proxied param.Field[string] `json:"proxied"`
}

func (r DNSRecordImportDNSRecordsForAZoneImportDNSRecordsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelope struct {
	Errors   []DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeSuccess `json:"success,required"`
	Timing  DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTiming  `json:"timing"`
	JSON    dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeJSON    `json:"-"`
}

// dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeJSON contains
// the JSON metadata for the struct
// [DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelope]
type dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	Timing      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrors struct {
	Code    int64                                                                       `json:"code,required"`
	Message string                                                                      `json:"message,required"`
	JSON    dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrors]
type dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessages struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessages]
type dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeSuccess bool

const (
	DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeSuccessTrue DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeSuccess = true
)

type DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTiming struct {
	// When the file parsing ended.
	EndTime time.Time `json:"end_time" format:"date-time"`
	// Processing time of the file in seconds.
	ProcessTime float64 `json:"process_time"`
	// When the file parsing started.
	StartTime time.Time                                                                   `json:"start_time" format:"date-time"`
	JSON      dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTimingJSON `json:"-"`
}

// dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTimingJSON
// contains the JSON metadata for the struct
// [DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTiming]
type dnsRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTimingJSON struct {
	EndTime     apijson.Field
	ProcessTime apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSRecordImportDNSRecordsForAZoneImportDNSRecordsResponseEnvelopeTiming) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
