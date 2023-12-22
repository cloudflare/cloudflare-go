// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

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
func (r *ZoneDNSRecordScanService) DNSRecordsForAZoneScanDNSRecords(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *DNSResponseImportScan, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_records/scan", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}
