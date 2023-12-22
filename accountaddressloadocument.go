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
func (r *AccountAddressLoaDocumentService) IPAddressManagementPrefixesUploadLoaDocument(ctx context.Context, accountIdentifier string, body AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams, opts ...option.RequestOption) (res *LoaUploadResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/loa_documents", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type LoaUploadResponse struct {
	Errors   []LoaUploadResponseError   `json:"errors"`
	Messages []LoaUploadResponseMessage `json:"messages"`
	Result   LoaUploadResponseResult    `json:"result"`
	// Whether the API call was successful
	Success LoaUploadResponseSuccess `json:"success"`
	JSON    loaUploadResponseJSON    `json:"-"`
}

// loaUploadResponseJSON contains the JSON metadata for the struct
// [LoaUploadResponse]
type loaUploadResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoaUploadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoaUploadResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    loaUploadResponseErrorJSON `json:"-"`
}

// loaUploadResponseErrorJSON contains the JSON metadata for the struct
// [LoaUploadResponseError]
type loaUploadResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoaUploadResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoaUploadResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    loaUploadResponseMessageJSON `json:"-"`
}

// loaUploadResponseMessageJSON contains the JSON metadata for the struct
// [LoaUploadResponseMessage]
type loaUploadResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoaUploadResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoaUploadResponseResult struct {
	// Name of LOA document.
	Filename string                      `json:"filename"`
	JSON     loaUploadResponseResultJSON `json:"-"`
}

// loaUploadResponseResultJSON contains the JSON metadata for the struct
// [LoaUploadResponseResult]
type loaUploadResponseResultJSON struct {
	Filename    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoaUploadResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoaUploadResponseSuccess bool

const (
	LoaUploadResponseSuccessTrue LoaUploadResponseSuccess = true
)

type AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams struct {
	// LOA document to upload.
	LoaDocument param.Field[string] `json:"loa_document,required"`
}

func (r AccountAddressLoaDocumentIPAddressManagementPrefixesUploadLoaDocumentParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
