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

// AddressingLOADocumentDownloadService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAddressingLOADocumentDownloadService] method instead.
type AddressingLOADocumentDownloadService struct {
	Options []option.RequestOption
}

// NewAddressingLOADocumentDownloadService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAddressingLOADocumentDownloadService(opts ...option.RequestOption) (r *AddressingLOADocumentDownloadService) {
	r = &AddressingLOADocumentDownloadService{}
	r.Options = opts
	return
}

// Download specified LOA document under the account.
func (r *AddressingLOADocumentDownloadService) Get(ctx context.Context, loaDocumentID string, query AddressingLOADocumentDownloadGetParams, opts ...option.RequestOption) (res *AddressingLOADocumentDownloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", query.AccountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AddressingLOADocumentDownloadGetResponse = interface{}

type AddressingLOADocumentDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
