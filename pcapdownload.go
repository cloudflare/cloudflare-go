// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// PCAPDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPCAPDownloadService] method
// instead.
type PCAPDownloadService struct {
	Options []option.RequestOption
}

// NewPCAPDownloadService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPCAPDownloadService(opts ...option.RequestOption) (r *PCAPDownloadService) {
	r = &PCAPDownloadService{}
	r.Options = opts
	return
}

// Download PCAP information into a file. Response is a binary PCAP file.
func (r *PCAPDownloadService) Get(ctx context.Context, accountID string, pcapID string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.tcpdump.pcap")}, opts...)
	path := fmt.Sprintf("accounts/%s/pcaps/%s/download", accountID, pcapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
