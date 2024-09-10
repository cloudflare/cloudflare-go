// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/cloudflare/cloudflare-go/internal/apiform"
	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/cloudflare/cloudflare-go/shared"
)

// LOADocumentService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLOADocumentService] method instead.
type LOADocumentService struct {
	Options   []option.RequestOption
	Downloads *LOADocumentDownloadService
}

// NewLOADocumentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLOADocumentService(opts ...option.RequestOption) (r *LOADocumentService) {
	r = &LOADocumentService{}
	r.Options = opts
	r.Downloads = NewLOADocumentDownloadService(opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *LOADocumentService) New(ctx context.Context, params LOADocumentNewParams, opts ...option.RequestOption) (res *LOADocumentNewResponse, err error) {
	var env LOADocumentNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LOADocumentNewResponse struct {
	// Name of LOA document.
	Filename string                     `json:"filename"`
	JSON     loaDocumentNewResponseJSON `json:"-"`
}

// loaDocumentNewResponseJSON contains the JSON metadata for the struct
// [LOADocumentNewResponse]
type loaDocumentNewResponseJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseJSON) RawJSON() string {
	return r.raw
}

type LOADocumentNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// LOA document to upload.
	LOADocument param.Field[string] `json:"loa_document,required"`
}

func (r LOADocumentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type LOADocumentNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success LOADocumentNewResponseEnvelopeSuccess `json:"success,required"`
	Result  LOADocumentNewResponse                `json:"result"`
	JSON    loaDocumentNewResponseEnvelopeJSON    `json:"-"`
}

// loaDocumentNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [LOADocumentNewResponseEnvelope]
type loaDocumentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LOADocumentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loaDocumentNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LOADocumentNewResponseEnvelopeSuccess bool

const (
	LOADocumentNewResponseEnvelopeSuccessTrue LOADocumentNewResponseEnvelopeSuccess = true
)

func (r LOADocumentNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LOADocumentNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
