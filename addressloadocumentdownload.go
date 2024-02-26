// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressLOADocumentDownloadService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAddressLOADocumentDownloadService] method instead.
type AddressLOADocumentDownloadService struct {
	Options []option.RequestOption
}

// NewAddressLOADocumentDownloadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressLOADocumentDownloadService(opts ...option.RequestOption) (r *AddressLOADocumentDownloadService) {
	r = &AddressLOADocumentDownloadService{}
	r.Options = opts
	return
}

// Download specified LOA document under the account.
func (r *AddressLOADocumentDownloadService) List(ctx context.Context, loaDocumentID string, query AddressLOADocumentDownloadListParams, opts ...option.RequestOption) (res *AddressLOADocumentDownloadListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", query.AccountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AddressLOADocumentDownloadListResponse = interface{}

type AddressLOADocumentDownloadListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
