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

// AddressLoaDocumentService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressLoaDocumentService] method
// instead.
type AddressLoaDocumentService struct {
	Options   []option.RequestOption
	Downloads *AddressLoaDocumentDownloadService
}

// NewAddressLoaDocumentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressLoaDocumentService(opts ...option.RequestOption) (r *AddressLoaDocumentService) {
	r = &AddressLoaDocumentService{}
	r.Options = opts
	r.Downloads = NewAddressLoaDocumentDownloadService(opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *AddressLoaDocumentService) New(ctx context.Context, accountID string, body AddressLoaDocumentNewParams, opts ...option.RequestOption) (res *AddressLoaDocumentNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressLoaDocumentNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressLoaDocumentNewResponse struct {
	// Name of LOA document.
	Filename string                            `json:"filename"`
	JSON     addressLoaDocumentNewResponseJSON `json:"-"`
}

// addressLoaDocumentNewResponseJSON contains the JSON metadata for the struct
// [AddressLoaDocumentNewResponse]
type addressLoaDocumentNewResponseJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentNewParams struct {
	// LOA document to upload.
	LoaDocument param.Field[string] `json:"loa_document,required"`
}

func (r AddressLoaDocumentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressLoaDocumentNewResponseEnvelope struct {
	Errors   []AddressLoaDocumentNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressLoaDocumentNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressLoaDocumentNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressLoaDocumentNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressLoaDocumentNewResponseEnvelopeJSON    `json:"-"`
}

// addressLoaDocumentNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AddressLoaDocumentNewResponseEnvelope]
type addressLoaDocumentNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    addressLoaDocumentNewResponseEnvelopeErrorsJSON `json:"-"`
}

// addressLoaDocumentNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AddressLoaDocumentNewResponseEnvelopeErrors]
type addressLoaDocumentNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    addressLoaDocumentNewResponseEnvelopeMessagesJSON `json:"-"`
}

// addressLoaDocumentNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AddressLoaDocumentNewResponseEnvelopeMessages]
type addressLoaDocumentNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressLoaDocumentNewResponseEnvelopeSuccess bool

const (
	AddressLoaDocumentNewResponseEnvelopeSuccessTrue AddressLoaDocumentNewResponseEnvelopeSuccess = true
)
