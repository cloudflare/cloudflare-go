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

// AccountLogControlCmbConfigService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogControlCmbConfigService] method instead.
type AccountLogControlCmbConfigService struct {
	Options []option.RequestOption
}

// NewAccountLogControlCmbConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLogControlCmbConfigService(opts ...option.RequestOption) (r *AccountLogControlCmbConfigService) {
	r = &AccountLogControlCmbConfigService{}
	r.Options = opts
	return
}

// Deletes CMB config.
func (r *AccountLogControlCmbConfigService) Delete(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountLogControlCmbConfigDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets CMB config.
func (r *AccountLogControlCmbConfigService) GetAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CmbConfigResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates CMB config.
func (r *AccountLogControlCmbConfigService) PutAccountsAccountIdentifierLogsControlCmbConfig(ctx context.Context, accountIdentifier string, body AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams, opts ...option.RequestOption) (res *CmbConfigResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logs/control/cmb/config", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type CmbConfigResponseSingle struct {
	Errors   []CmbConfigResponseSingleError   `json:"errors"`
	Messages []CmbConfigResponseSingleMessage `json:"messages"`
	Result   CmbConfigResponseSingleResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success CmbConfigResponseSingleSuccess `json:"success"`
	JSON    cmbConfigResponseSingleJSON    `json:"-"`
}

// cmbConfigResponseSingleJSON contains the JSON metadata for the struct
// [CmbConfigResponseSingle]
type cmbConfigResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfigResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CmbConfigResponseSingleError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    cmbConfigResponseSingleErrorJSON `json:"-"`
}

// cmbConfigResponseSingleErrorJSON contains the JSON metadata for the struct
// [CmbConfigResponseSingleError]
type cmbConfigResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfigResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CmbConfigResponseSingleMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    cmbConfigResponseSingleMessageJSON `json:"-"`
}

// cmbConfigResponseSingleMessageJSON contains the JSON metadata for the struct
// [CmbConfigResponseSingleMessage]
type cmbConfigResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfigResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CmbConfigResponseSingleResult struct {
	// Comma-separated list of regions.
	Regions string                            `json:"regions"`
	JSON    cmbConfigResponseSingleResultJSON `json:"-"`
}

// cmbConfigResponseSingleResultJSON contains the JSON metadata for the struct
// [CmbConfigResponseSingleResult]
type cmbConfigResponseSingleResultJSON struct {
	Regions     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CmbConfigResponseSingleResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CmbConfigResponseSingleSuccess bool

const (
	CmbConfigResponseSingleSuccessTrue CmbConfigResponseSingleSuccess = true
)

type AccountLogControlCmbConfigDeleteResponse struct {
	Errors   []AccountLogControlCmbConfigDeleteResponseError   `json:"errors"`
	Messages []AccountLogControlCmbConfigDeleteResponseMessage `json:"messages"`
	Result   interface{}                                       `json:"result,nullable"`
	// Whether the API call was successful
	Success AccountLogControlCmbConfigDeleteResponseSuccess `json:"success"`
	JSON    accountLogControlCmbConfigDeleteResponseJSON    `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseJSON contains the JSON metadata for the
// struct [AccountLogControlCmbConfigDeleteResponse]
type accountLogControlCmbConfigDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigDeleteResponseError struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accountLogControlCmbConfigDeleteResponseErrorJSON `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseErrorJSON contains the JSON metadata for
// the struct [AccountLogControlCmbConfigDeleteResponseError]
type accountLogControlCmbConfigDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLogControlCmbConfigDeleteResponseMessage struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accountLogControlCmbConfigDeleteResponseMessageJSON `json:"-"`
}

// accountLogControlCmbConfigDeleteResponseMessageJSON contains the JSON metadata
// for the struct [AccountLogControlCmbConfigDeleteResponseMessage]
type accountLogControlCmbConfigDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLogControlCmbConfigDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLogControlCmbConfigDeleteResponseSuccess bool

const (
	AccountLogControlCmbConfigDeleteResponseSuccessTrue AccountLogControlCmbConfigDeleteResponseSuccess = true
)

type AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams struct {
	// Comma-separated list of regions.
	Regions param.Field[string] `json:"regions"`
}

func (r AccountLogControlCmbConfigPutAccountsAccountIdentifierLogsControlCmbConfigParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
