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

// AddressLOADocumentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressLOADocumentService] method
// instead.
type AddressLOADocumentService struct {
	Options   []option.RequestOption
	Downloads *AddressLOADocumentDownloadService
}

// NewAddressLOADocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressLOADocumentService(opts ...option.RequestOption) (r *AddressLOADocumentService) {
	r = &AddressLOADocumentService{}
	r.Options = opts
	r.Downloads = NewAddressLOADocumentDownloadService(opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *AddressLOADocumentService) New(ctx context.Context, accountID string, body AddressLOADocumentNewParams, opts ...option.RequestOption) (res *AddressLOADocumentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressLOADocumentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressLOADocumentNewResponse struct {
	// Name of LOA document.
	Filename string                            `json:"filename"`
	JSON     addressLOADocumentNewResponseJSON `json:"-"`
}

// addressLOADocumentNewResponseJSON contains the JSON metadata for the struct
// [AddressLOADocumentNewResponse]
type addressLOADocumentNewResponseJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLOADocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLOADocumentNewParams struct {
	// LOA document to upload.
	LOADocument param.Field[string] `json:"loa_document,required"`
}

func (r AddressLOADocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressLOADocumentNewResponseEnvelope struct {
	Errors   []AddressLOADocumentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressLOADocumentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressLOADocumentNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressLOADocumentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressLOADocumentNewResponseEnvelopeJSON    `json:"-"`
}

// addressLOADocumentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressLOADocumentNewResponseEnvelope]
type addressLOADocumentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLOADocumentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLOADocumentNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressLOADocumentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressLOADocumentNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressLOADocumentNewResponseEnvelopeErrors]
type addressLOADocumentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLOADocumentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLOADocumentNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressLOADocumentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressLOADocumentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressLOADocumentNewResponseEnvelopeMessages]
type addressLOADocumentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLOADocumentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressLOADocumentNewResponseEnvelopeSuccess bool

const (
	AddressLOADocumentNewResponseEnvelopeSuccessTrue AddressLOADocumentNewResponseEnvelopeSuccess = true
)
