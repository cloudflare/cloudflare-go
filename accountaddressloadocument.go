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

// AccountAddressLoaDocumentService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAddressLoaDocumentService] method instead.
type AccountAddressLoaDocumentService struct {
	Options   []option.RequestOption
	Downloads *AccountAddressLoaDocumentDownloadService
}

// NewAccountAddressLoaDocumentService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressLoaDocumentService(opts ...option.RequestOption) (r *AccountAddressLoaDocumentService) {
	r = &AccountAddressLoaDocumentService{}
	r.Options = opts
	r.Downloads = NewAccountAddressLoaDocumentDownloadService(opts...)
	return
}

// Submit LOA document (pdf format) under the account.
func (r *AccountAddressLoaDocumentService) IPAddressManagementPrefixesUploadLoaDocument(ctx context.Context, accountIdentifier string, body AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams, opts ...option.RequestOption) (res *AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse struct {
	Errors   []AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseError   `json:"errors"`
	Messages []AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessage `json:"messages"`
	Result   AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseSuccess `json:"success"`
	JSON    accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON    `json:"-"`
}

// accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON
// contains the JSON metadata for the struct
// [AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse]
type accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseError struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseErrorJSON `json:"-"`
}

// accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseErrorJSON
// contains the JSON metadata for the struct
// [AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseError]
type accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessage struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessageJSON `json:"-"`
}

// accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessageJSON
// contains the JSON metadata for the struct
// [AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessage]
type accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResult struct {
	// Name of LOA document.
	Filename string                                                                                  `json:"filename"`
	JSON     accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResultJSON `json:"-"`
}

// accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResultJSON
// contains the JSON metadata for the struct
// [AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResult]
type accountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResultJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseSuccess bool

const (
	AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseSuccessTrue AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentResponseSuccess = true
)

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams struct {
	// LOA document to upload.
	LoaDocument param.Field[string] `json:"loa_document,required"`
}

func (r AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
