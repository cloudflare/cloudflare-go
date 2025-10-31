// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// LOADocumentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLOADocumentService] method instead.
type LOADocumentService struct {
	Options []option.RequestOption
}

// NewLOADocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLOADocumentService(opts ...option.RequestOption) (r *LOADocumentService) {
	r = &LOADocumentService{}
	r.Options = opts
	return
}

// Download specified LOA document under the account.
func (r *LOADocumentService) Get(ctx context.Context, loaDocumentID string, query LOADocumentGetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/pdf")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if loaDocumentID == "" {
		err = errors.New("missing required loa_document_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents/%s/download", query.AccountID, loaDocumentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LOADocumentGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}
