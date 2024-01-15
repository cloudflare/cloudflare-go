// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneDNSRecordExportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDNSRecordExportService]
// method instead.
type ZoneDNSRecordExportService struct {
	Options []option.RequestOption
}

// NewZoneDNSRecordExportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSRecordExportService(opts ...option.RequestOption) (r *ZoneDNSRecordExportService) {
	r = &ZoneDNSRecordExportService{}
	r.Options = opts
	return
}

// You can export your
// [BIND config](https://en.wikipedia.org/wiki/Zone_file "Zone file") through this
// endpoint.
//
// See
// [the documentation](https://developers.cloudflare.com/dns/manage-dns-records/how-to/import-and-export/ "Import and export records")
// for more information.
func (r *ZoneDNSRecordExportService) List(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	path := fmt.Sprintf("zones/%s/dns_records/export", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
