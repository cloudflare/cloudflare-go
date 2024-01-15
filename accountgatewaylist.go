// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountGatewayListService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountGatewayListService] method
// instead.
type AccountGatewayListService struct {
	Options []option.RequestOption
	Items   *AccountGatewayListItemService
}

// NewAccountGatewayListService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountGatewayListService(opts ...option.RequestOption) (r *AccountGatewayListService) {
	r = &AccountGatewayListService{}
	r.Options = opts
	r.Items = NewAccountGatewayListItemService(opts...)
	return
}

// Fetches a single Zero Trust list.
func (r *AccountGatewayListService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a configured Zero Trust list.
func (r *AccountGatewayListService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountGatewayListUpdateParams, opts ...option.RequestOption) (res *AccountGatewayListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a Zero Trust list.
func (r *AccountGatewayListService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayListDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Appends or removes an item from a configured Zero Trust list.
func (r *AccountGatewayListService) Patch(ctx context.Context, identifier interface{}, uuid string, body AccountGatewayListPatchParams, opts ...option.RequestOption) (res *AccountGatewayListPatchResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Creates a new Zero Trust list.
func (r *AccountGatewayListService) ZeroTrustListsNewZeroTrustList(ctx context.Context, identifier interface{}, body AccountGatewayListZeroTrustListsNewZeroTrustListParams, opts ...option.RequestOption) (res *AccountGatewayListZeroTrustListsNewZeroTrustListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all Zero Trust lists for an account.
func (r *AccountGatewayListService) ZeroTrustListsListZeroTrustLists(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGatewayListZeroTrustListsListZeroTrustListsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountGatewayListGetResponse struct {
	Errors   []AccountGatewayListGetResponseError   `json:"errors"`
	Messages []AccountGatewayListGetResponseMessage `json:"messages"`
	Result   AccountGatewayListGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayListGetResponseSuccess `json:"success"`
	JSON    accountGatewayListGetResponseJSON    `json:"-"`
}

// accountGatewayListGetResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListGetResponse]
type accountGatewayListGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListGetResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    accountGatewayListGetResponseErrorJSON `json:"-"`
}

// accountGatewayListGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountGatewayListGetResponseError]
type accountGatewayListGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListGetResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountGatewayListGetResponseMessageJSON `json:"-"`
}

// accountGatewayListGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListGetResponseMessage]
type accountGatewayListGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListGetResponseResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      AccountGatewayListGetResponseResultType `json:"type"`
	UpdatedAt time.Time                               `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListGetResponseResultJSON `json:"-"`
}

// accountGatewayListGetResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayListGetResponseResult]
type accountGatewayListGetResponseResultJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type AccountGatewayListGetResponseResultType string

const (
	AccountGatewayListGetResponseResultTypeSerial AccountGatewayListGetResponseResultType = "SERIAL"
	AccountGatewayListGetResponseResultTypeURL    AccountGatewayListGetResponseResultType = "URL"
	AccountGatewayListGetResponseResultTypeDomain AccountGatewayListGetResponseResultType = "DOMAIN"
	AccountGatewayListGetResponseResultTypeEmail  AccountGatewayListGetResponseResultType = "EMAIL"
	AccountGatewayListGetResponseResultTypeIP     AccountGatewayListGetResponseResultType = "IP"
)

// Whether the API call was successful
type AccountGatewayListGetResponseSuccess bool

const (
	AccountGatewayListGetResponseSuccessTrue AccountGatewayListGetResponseSuccess = true
)

type AccountGatewayListUpdateResponse struct {
	Errors   []AccountGatewayListUpdateResponseError   `json:"errors"`
	Messages []AccountGatewayListUpdateResponseMessage `json:"messages"`
	Result   AccountGatewayListUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayListUpdateResponseSuccess `json:"success"`
	JSON    accountGatewayListUpdateResponseJSON    `json:"-"`
}

// accountGatewayListUpdateResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListUpdateResponse]
type accountGatewayListUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListUpdateResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountGatewayListUpdateResponseErrorJSON `json:"-"`
}

// accountGatewayListUpdateResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListUpdateResponseError]
type accountGatewayListUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListUpdateResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountGatewayListUpdateResponseMessageJSON `json:"-"`
}

// accountGatewayListUpdateResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListUpdateResponseMessage]
type accountGatewayListUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListUpdateResponseResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      AccountGatewayListUpdateResponseResultType `json:"type"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListUpdateResponseResultJSON `json:"-"`
}

// accountGatewayListUpdateResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayListUpdateResponseResult]
type accountGatewayListUpdateResponseResultJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type AccountGatewayListUpdateResponseResultType string

const (
	AccountGatewayListUpdateResponseResultTypeSerial AccountGatewayListUpdateResponseResultType = "SERIAL"
	AccountGatewayListUpdateResponseResultTypeURL    AccountGatewayListUpdateResponseResultType = "URL"
	AccountGatewayListUpdateResponseResultTypeDomain AccountGatewayListUpdateResponseResultType = "DOMAIN"
	AccountGatewayListUpdateResponseResultTypeEmail  AccountGatewayListUpdateResponseResultType = "EMAIL"
	AccountGatewayListUpdateResponseResultTypeIP     AccountGatewayListUpdateResponseResultType = "IP"
)

// Whether the API call was successful
type AccountGatewayListUpdateResponseSuccess bool

const (
	AccountGatewayListUpdateResponseSuccessTrue AccountGatewayListUpdateResponseSuccess = true
)

type AccountGatewayListDeleteResponse struct {
	Errors   []AccountGatewayListDeleteResponseError   `json:"errors"`
	Messages []AccountGatewayListDeleteResponseMessage `json:"messages"`
	Result   interface{}                               `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayListDeleteResponseSuccess `json:"success"`
	JSON    accountGatewayListDeleteResponseJSON    `json:"-"`
}

// accountGatewayListDeleteResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListDeleteResponse]
type accountGatewayListDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListDeleteResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountGatewayListDeleteResponseErrorJSON `json:"-"`
}

// accountGatewayListDeleteResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListDeleteResponseError]
type accountGatewayListDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListDeleteResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountGatewayListDeleteResponseMessageJSON `json:"-"`
}

// accountGatewayListDeleteResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListDeleteResponseMessage]
type accountGatewayListDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayListDeleteResponseSuccess bool

const (
	AccountGatewayListDeleteResponseSuccessTrue AccountGatewayListDeleteResponseSuccess = true
)

type AccountGatewayListPatchResponse struct {
	Errors   []AccountGatewayListPatchResponseError   `json:"errors"`
	Messages []AccountGatewayListPatchResponseMessage `json:"messages"`
	Result   AccountGatewayListPatchResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayListPatchResponseSuccess `json:"success"`
	JSON    accountGatewayListPatchResponseJSON    `json:"-"`
}

// accountGatewayListPatchResponseJSON contains the JSON metadata for the struct
// [AccountGatewayListPatchResponse]
type accountGatewayListPatchResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListPatchResponseError struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accountGatewayListPatchResponseErrorJSON `json:"-"`
}

// accountGatewayListPatchResponseErrorJSON contains the JSON metadata for the
// struct [AccountGatewayListPatchResponseError]
type accountGatewayListPatchResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListPatchResponseMessage struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accountGatewayListPatchResponseMessageJSON `json:"-"`
}

// accountGatewayListPatchResponseMessageJSON contains the JSON metadata for the
// struct [AccountGatewayListPatchResponseMessage]
type accountGatewayListPatchResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListPatchResponseResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      AccountGatewayListPatchResponseResultType `json:"type"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListPatchResponseResultJSON `json:"-"`
}

// accountGatewayListPatchResponseResultJSON contains the JSON metadata for the
// struct [AccountGatewayListPatchResponseResult]
type accountGatewayListPatchResponseResultJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListPatchResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type AccountGatewayListPatchResponseResultType string

const (
	AccountGatewayListPatchResponseResultTypeSerial AccountGatewayListPatchResponseResultType = "SERIAL"
	AccountGatewayListPatchResponseResultTypeURL    AccountGatewayListPatchResponseResultType = "URL"
	AccountGatewayListPatchResponseResultTypeDomain AccountGatewayListPatchResponseResultType = "DOMAIN"
	AccountGatewayListPatchResponseResultTypeEmail  AccountGatewayListPatchResponseResultType = "EMAIL"
	AccountGatewayListPatchResponseResultTypeIP     AccountGatewayListPatchResponseResultType = "IP"
)

// Whether the API call was successful
type AccountGatewayListPatchResponseSuccess bool

const (
	AccountGatewayListPatchResponseSuccessTrue AccountGatewayListPatchResponseSuccess = true
)

type AccountGatewayListZeroTrustListsNewZeroTrustListResponse struct {
	Errors   []AccountGatewayListZeroTrustListsNewZeroTrustListResponseError   `json:"errors"`
	Messages []AccountGatewayListZeroTrustListsNewZeroTrustListResponseMessage `json:"messages"`
	Result   AccountGatewayListZeroTrustListsNewZeroTrustListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountGatewayListZeroTrustListsNewZeroTrustListResponseSuccess `json:"success"`
	JSON    accountGatewayListZeroTrustListsNewZeroTrustListResponseJSON    `json:"-"`
}

// accountGatewayListZeroTrustListsNewZeroTrustListResponseJSON contains the JSON
// metadata for the struct
// [AccountGatewayListZeroTrustListsNewZeroTrustListResponse]
type accountGatewayListZeroTrustListsNewZeroTrustListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsNewZeroTrustListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListResponseError struct {
	Code    int64                                                             `json:"code,required"`
	Message string                                                            `json:"message,required"`
	JSON    accountGatewayListZeroTrustListsNewZeroTrustListResponseErrorJSON `json:"-"`
}

// accountGatewayListZeroTrustListsNewZeroTrustListResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountGatewayListZeroTrustListsNewZeroTrustListResponseError]
type accountGatewayListZeroTrustListsNewZeroTrustListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsNewZeroTrustListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListResponseMessage struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountGatewayListZeroTrustListsNewZeroTrustListResponseMessageJSON `json:"-"`
}

// accountGatewayListZeroTrustListsNewZeroTrustListResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountGatewayListZeroTrustListsNewZeroTrustListResponseMessage]
type accountGatewayListZeroTrustListsNewZeroTrustListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsNewZeroTrustListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListResponseResult struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The items in the list.
	Items []AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultItem `json:"items"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType `json:"type"`
	UpdatedAt time.Time                                                          `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListZeroTrustListsNewZeroTrustListResponseResultJSON `json:"-"`
}

// accountGatewayListZeroTrustListsNewZeroTrustListResponseResultJSON contains the
// JSON metadata for the struct
// [AccountGatewayListZeroTrustListsNewZeroTrustListResponseResult]
type accountGatewayListZeroTrustListsNewZeroTrustListResponseResultJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Items       apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsNewZeroTrustListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a list.
	Value string                                                                 `json:"value"`
	JSON  accountGatewayListZeroTrustListsNewZeroTrustListResponseResultItemJSON `json:"-"`
}

// accountGatewayListZeroTrustListsNewZeroTrustListResponseResultItemJSON contains
// the JSON metadata for the struct
// [AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultItem]
type accountGatewayListZeroTrustListsNewZeroTrustListResponseResultItemJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType string

const (
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultTypeSerial AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType = "SERIAL"
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultTypeURL    AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType = "URL"
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultTypeDomain AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType = "DOMAIN"
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultTypeEmail  AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType = "EMAIL"
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultTypeIP     AccountGatewayListZeroTrustListsNewZeroTrustListResponseResultType = "IP"
)

// Whether the API call was successful
type AccountGatewayListZeroTrustListsNewZeroTrustListResponseSuccess bool

const (
	AccountGatewayListZeroTrustListsNewZeroTrustListResponseSuccessTrue AccountGatewayListZeroTrustListsNewZeroTrustListResponseSuccess = true
)

type AccountGatewayListZeroTrustListsListZeroTrustListsResponse struct {
	Errors     []AccountGatewayListZeroTrustListsListZeroTrustListsResponseError    `json:"errors"`
	Messages   []AccountGatewayListZeroTrustListsListZeroTrustListsResponseMessage  `json:"messages"`
	Result     []AccountGatewayListZeroTrustListsListZeroTrustListsResponseResult   `json:"result"`
	ResultInfo AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccountGatewayListZeroTrustListsListZeroTrustListsResponseSuccess `json:"success"`
	JSON    accountGatewayListZeroTrustListsListZeroTrustListsResponseJSON    `json:"-"`
}

// accountGatewayListZeroTrustListsListZeroTrustListsResponseJSON contains the JSON
// metadata for the struct
// [AccountGatewayListZeroTrustListsListZeroTrustListsResponse]
type accountGatewayListZeroTrustListsListZeroTrustListsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsListZeroTrustListsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsListZeroTrustListsResponseError struct {
	Code    int64                                                               `json:"code,required"`
	Message string                                                              `json:"message,required"`
	JSON    accountGatewayListZeroTrustListsListZeroTrustListsResponseErrorJSON `json:"-"`
}

// accountGatewayListZeroTrustListsListZeroTrustListsResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountGatewayListZeroTrustListsListZeroTrustListsResponseError]
type accountGatewayListZeroTrustListsListZeroTrustListsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsListZeroTrustListsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsListZeroTrustListsResponseMessage struct {
	Code    int64                                                                 `json:"code,required"`
	Message string                                                                `json:"message,required"`
	JSON    accountGatewayListZeroTrustListsListZeroTrustListsResponseMessageJSON `json:"-"`
}

// accountGatewayListZeroTrustListsListZeroTrustListsResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountGatewayListZeroTrustListsListZeroTrustListsResponseMessage]
type accountGatewayListZeroTrustListsListZeroTrustListsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsListZeroTrustListsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGatewayListZeroTrustListsListZeroTrustListsResponseResult struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the list.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the list.
	Description string `json:"description"`
	// The name of the list.
	Name string `json:"name"`
	// The type of list.
	Type      AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType `json:"type"`
	UpdatedAt time.Time                                                            `json:"updated_at" format:"date-time"`
	JSON      accountGatewayListZeroTrustListsListZeroTrustListsResponseResultJSON `json:"-"`
}

// accountGatewayListZeroTrustListsListZeroTrustListsResponseResultJSON contains
// the JSON metadata for the struct
// [AccountGatewayListZeroTrustListsListZeroTrustListsResponseResult]
type accountGatewayListZeroTrustListsListZeroTrustListsResponseResultJSON struct {
	ID          apijson.Field
	Count       apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsListZeroTrustListsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of list.
type AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType string

const (
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultTypeSerial AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType = "SERIAL"
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultTypeURL    AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType = "URL"
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultTypeDomain AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType = "DOMAIN"
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultTypeEmail  AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType = "EMAIL"
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultTypeIP     AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultType = "IP"
)

type AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                  `json:"total_count"`
	JSON       accountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfoJSON `json:"-"`
}

// accountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfo]
type accountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGatewayListZeroTrustListsListZeroTrustListsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGatewayListZeroTrustListsListZeroTrustListsResponseSuccess bool

const (
	AccountGatewayListZeroTrustListsListZeroTrustListsResponseSuccessTrue AccountGatewayListZeroTrustListsListZeroTrustListsResponseSuccess = true
)

type AccountGatewayListUpdateParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
}

func (r AccountGatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListPatchParams struct {
	// The items in the list.
	Append param.Field[[]AccountGatewayListPatchParamsAppend] `json:"append"`
	// A list of the item values you want to remove.
	Remove param.Field[[]string] `json:"remove"`
}

func (r AccountGatewayListPatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListPatchParamsAppend struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListPatchParamsAppend) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListParams struct {
	// The name of the list.
	Name param.Field[string] `json:"name,required"`
	// The type of list.
	Type param.Field[AccountGatewayListZeroTrustListsNewZeroTrustListParamsType] `json:"type,required"`
	// The description of the list.
	Description param.Field[string] `json:"description"`
	// The items in the list.
	Items param.Field[[]AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem] `json:"items"`
}

func (r AccountGatewayListZeroTrustListsNewZeroTrustListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of list.
type AccountGatewayListZeroTrustListsNewZeroTrustListParamsType string

const (
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeSerial AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "SERIAL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeURL    AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "URL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeDomain AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "DOMAIN"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeEmail  AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "EMAIL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeIP     AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "IP"
)

type AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem struct {
	// The value of the item in a list.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
