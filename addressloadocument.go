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
func (r *AddressLoaDocumentService) IPAddressManagementPrefixesUploadLoaDocument(ctx context.Context, accountID string, body AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams, opts ...option.RequestOption) (res *AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelope
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse struct {
	// Name of LOA document.
	Filename string                                                                     `json:"filename"`
	JSON     addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON `json:"-"`
}

// addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON
// contains the JSON metadata for the struct
// [AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse]
type addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams struct {
	// LOA document to upload.
	LoaDocument param.Field[string] `json:"loa_document,required"`
}

func (r AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelope struct {
	Errors   []AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessages `json:"messages,required"`
	Result   AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeSuccess `json:"success,required"`
	JSON    addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeJSON    `json:"-"`
}

// addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelope]
type addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrors struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrorsJSON `json:"-"`
}

// addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrors]
type addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessages struct {
	Code    int64                                                                                      `json:"code,required"`
	Message string                                                                                     `json:"message,required"`
	JSON    addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessagesJSON `json:"-"`
}

// addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessages]
type addressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeSuccess bool

const (
	AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeSuccessTrue AddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseEnvelopeSuccess = true
)
