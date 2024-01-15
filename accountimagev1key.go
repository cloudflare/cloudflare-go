// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountImageV1KeyService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountImageV1KeyService] method
// instead.
type AccountImageV1KeyService struct {
	Options []option.RequestOption
}

// NewAccountImageV1KeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountImageV1KeyService(opts ...option.RequestOption) (r *AccountImageV1KeyService) {
	r = &AccountImageV1KeyService{}
	r.Options = opts
	return
}

// Lists your signing keys. These can be found on your Cloudflare Images dashboard.
func (r *AccountImageV1KeyService) CloudflareImagesKeysListSigningKeys(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/images/v1/keys", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponse struct {
	Errors   []AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseError   `json:"errors"`
	Messages []AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessage `json:"messages"`
	Result   AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseSuccess `json:"success"`
	JSON    accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON    `json:"-"`
}

// accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON contains the
// JSON metadata for the struct
// [AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponse]
type accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseError struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseErrorJSON `json:"-"`
}

// accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseErrorJSON contains
// the JSON metadata for the struct
// [AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseError]
type accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessage struct {
	Code    int64                                                                   `json:"code,required"`
	Message string                                                                  `json:"message,required"`
	JSON    accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessageJSON `json:"-"`
}

// accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessage]
type accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResult struct {
	Keys []AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKey `json:"keys"`
	JSON accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultJSON  `json:"-"`
}

// accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultJSON contains
// the JSON metadata for the struct
// [AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResult]
type accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultJSON struct {
	Keys        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKey struct {
	// Key name.
	Name string `json:"name"`
	// Key value.
	Value string                                                                    `json:"value"`
	JSON  accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKeyJSON `json:"-"`
}

// accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKeyJSON
// contains the JSON metadata for the struct
// [AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKey]
type accountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKeyJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseResultKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseSuccess bool

const (
	AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseSuccessTrue AccountImageV1KeyCloudflareImagesKeysListSigningKeysResponseSuccess = true
)
