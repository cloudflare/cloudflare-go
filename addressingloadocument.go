// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AddressingLOADocumentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingLOADocumentService]
// method instead.
type AddressingLOADocumentService struct {
	Options   []option.RequestOption
	Downloads *AddressingLOADocumentDownloadService
}

// NewAddressingLOADocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingLOADocumentService(opts ...option.RequestOption) (r *AddressingLOADocumentService) {
	r = &AddressingLOADocumentService{}
	r.Options = opts
	r.Downloads = NewAddressingLOADocumentDownloadService(opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *AddressingLOADocumentService) New(ctx context.Context, params AddressingLOADocumentNewParams, opts ...option.RequestOption) (res *AddressingLOADocumentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressingLOADocumentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressingLOADocumentNewResponse struct {
	// Name of LOA document.
	Filename string                               `json:"filename"`
	JSON     addressingLOADocumentNewResponseJSON `json:"-"`
}

// addressingLOADocumentNewResponseJSON contains the JSON metadata for the struct
// [AddressingLOADocumentNewResponse]
type addressingLOADocumentNewResponseJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingLOADocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingLOADocumentNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// LOA document to upload.
	LOADocument param.Field[string] `json:"loa_document,required"`
}

func (r AddressingLOADocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressingLOADocumentNewResponseEnvelope struct {
	Errors   []AddressingLOADocumentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressingLOADocumentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressingLOADocumentNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressingLOADocumentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressingLOADocumentNewResponseEnvelopeJSON    `json:"-"`
}

// addressingLOADocumentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressingLOADocumentNewResponseEnvelope]
type addressingLOADocumentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingLOADocumentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingLOADocumentNewResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    addressingLOADocumentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressingLOADocumentNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AddressingLOADocumentNewResponseEnvelopeErrors]
type addressingLOADocumentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingLOADocumentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingLOADocumentNewResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    addressingLOADocumentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressingLOADocumentNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AddressingLOADocumentNewResponseEnvelopeMessages]
type addressingLOADocumentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingLOADocumentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingLOADocumentNewResponseEnvelopeSuccess bool

const (
	AddressingLOADocumentNewResponseEnvelopeSuccessTrue AddressingLOADocumentNewResponseEnvelopeSuccess = true
)
