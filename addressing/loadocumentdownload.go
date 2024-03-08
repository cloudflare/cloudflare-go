// File generated from our OpenAPI spec by Stainless.

package addressing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// LOADocumentDownloadService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLOADocumentDownloadService]
// method instead.
type LOADocumentDownloadService struct {
	Options []option.RequestOption
}

// NewLOADocumentDownloadService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLOADocumentDownloadService(opts ...option.RequestOption) (r *LOADocumentDownloadService) {
	r = &LOADocumentDownloadService{}
	r.Options = opts
	return
}

// Download specified LOA document under the account.
func (r *LOADocumentDownloadService) Get(ctx context.Context, loaDocumentID string, query LOADocumentDownloadGetParams, opts ...option.RequestOption) (res *LOADocumentDownloadGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", query.AccountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type LOADocumentDownloadGetResponse = interface{}

type LOADocumentDownloadGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
