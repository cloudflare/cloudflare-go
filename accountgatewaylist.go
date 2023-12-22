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

// Fetch a single Zero Trust List.
func (r *AccountGatewayListService) Get(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *AccountGatewayListGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a configured Zero Trust List.
func (r *AccountGatewayListService) Update(ctx context.Context, identifier interface{}, uuid string, body AccountGatewayListUpdateParams, opts ...option.RequestOption) (res *AccountGatewayListUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Zero Trust List.
func (r *AccountGatewayListService) Delete(ctx context.Context, identifier interface{}, uuid string, opts ...option.RequestOption) (res *EmptyResponseArJnNcMr, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create a new Zero Trust List.
func (r *AccountGatewayListService) ZeroTrustListsNewZeroTrustList(ctx context.Context, identifier interface{}, body AccountGatewayListZeroTrustListsNewZeroTrustListParams, opts ...option.RequestOption) (res *SingleResponseWithListItem, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List Zero Trust Lists for an account.
func (r *AccountGatewayListService) ZeroTrustListsListZeroTrustLists(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseCollection0kL9nhn2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/gateway/lists", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmptyResponseArJnNcMr struct {
	Errors   []EmptyResponseArJnNcMrError   `json:"errors"`
	Messages []EmptyResponseArJnNcMrMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success EmptyResponseArJnNcMrSuccess `json:"success"`
	JSON    emptyResponseArJnNcMrJSON    `json:"-"`
}

// emptyResponseArJnNcMrJSON contains the JSON metadata for the struct
// [EmptyResponseArJnNcMr]
type emptyResponseArJnNcMrJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseArJnNcMr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmptyResponseArJnNcMrError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    emptyResponseArJnNcMrErrorJSON `json:"-"`
}

// emptyResponseArJnNcMrErrorJSON contains the JSON metadata for the struct
// [EmptyResponseArJnNcMrError]
type emptyResponseArJnNcMrErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseArJnNcMrError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EmptyResponseArJnNcMrMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    emptyResponseArJnNcMrMessageJSON `json:"-"`
}

// emptyResponseArJnNcMrMessageJSON contains the JSON metadata for the struct
// [EmptyResponseArJnNcMrMessage]
type emptyResponseArJnNcMrMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmptyResponseArJnNcMrMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type EmptyResponseArJnNcMrSuccess bool

const (
	EmptyResponseArJnNcMrSuccessTrue EmptyResponseArJnNcMrSuccess = true
)

type ResponseCollection0kL9nhn2 struct {
	Errors     []ResponseCollection0kL9nhn2Error    `json:"errors"`
	Messages   []ResponseCollection0kL9nhn2Message  `json:"messages"`
	Result     []ResponseCollection0kL9nhn2Result   `json:"result"`
	ResultInfo ResponseCollection0kL9nhn2ResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollection0kL9nhn2Success `json:"success"`
	JSON    responseCollection0kL9nhn2JSON    `json:"-"`
}

// responseCollection0kL9nhn2JSON contains the JSON metadata for the struct
// [ResponseCollection0kL9nhn2]
type responseCollection0kL9nhn2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollection0kL9nhn2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollection0kL9nhn2Error struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollection0kL9nhn2ErrorJSON `json:"-"`
}

// responseCollection0kL9nhn2ErrorJSON contains the JSON metadata for the struct
// [ResponseCollection0kL9nhn2Error]
type responseCollection0kL9nhn2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollection0kL9nhn2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollection0kL9nhn2Message struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollection0kL9nhn2MessageJSON `json:"-"`
}

// responseCollection0kL9nhn2MessageJSON contains the JSON metadata for the struct
// [ResponseCollection0kL9nhn2Message]
type responseCollection0kL9nhn2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollection0kL9nhn2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollection0kL9nhn2Result struct {
	// API Resource UUID tag.
	ID string `json:"id"`
	// The number of items in the List.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the List.
	Description string `json:"description"`
	// The name of the List.
	Name string `json:"name"`
	// The type of List.
	Type      ResponseCollection0kL9nhn2ResultType `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      responseCollection0kL9nhn2ResultJSON `json:"-"`
}

// responseCollection0kL9nhn2ResultJSON contains the JSON metadata for the struct
// [ResponseCollection0kL9nhn2Result]
type responseCollection0kL9nhn2ResultJSON struct {
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

func (r *ResponseCollection0kL9nhn2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of List.
type ResponseCollection0kL9nhn2ResultType string

const (
	ResponseCollection0kL9nhn2ResultTypeSerial ResponseCollection0kL9nhn2ResultType = "SERIAL"
	ResponseCollection0kL9nhn2ResultTypeURL    ResponseCollection0kL9nhn2ResultType = "URL"
	ResponseCollection0kL9nhn2ResultTypeDomain ResponseCollection0kL9nhn2ResultType = "DOMAIN"
	ResponseCollection0kL9nhn2ResultTypeEmail  ResponseCollection0kL9nhn2ResultType = "EMAIL"
	ResponseCollection0kL9nhn2ResultTypeIP     ResponseCollection0kL9nhn2ResultType = "IP"
)

type ResponseCollection0kL9nhn2ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollection0kL9nhn2ResultInfoJSON `json:"-"`
}

// responseCollection0kL9nhn2ResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollection0kL9nhn2ResultInfo]
type responseCollection0kL9nhn2ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollection0kL9nhn2ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollection0kL9nhn2Success bool

const (
	ResponseCollection0kL9nhn2SuccessTrue ResponseCollection0kL9nhn2Success = true
)

type SingleResponseWithListItem struct {
	Errors   []SingleResponseWithListItemError   `json:"errors"`
	Messages []SingleResponseWithListItemMessage `json:"messages"`
	Result   SingleResponseWithListItemResult    `json:"result"`
	// Whether the API call was successful
	Success SingleResponseWithListItemSuccess `json:"success"`
	JSON    singleResponseWithListItemJSON    `json:"-"`
}

// singleResponseWithListItemJSON contains the JSON metadata for the struct
// [SingleResponseWithListItem]
type singleResponseWithListItemJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWithListItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWithListItemError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    singleResponseWithListItemErrorJSON `json:"-"`
}

// singleResponseWithListItemErrorJSON contains the JSON metadata for the struct
// [SingleResponseWithListItemError]
type singleResponseWithListItemErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWithListItemError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWithListItemMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    singleResponseWithListItemMessageJSON `json:"-"`
}

// singleResponseWithListItemMessageJSON contains the JSON metadata for the struct
// [SingleResponseWithListItemMessage]
type singleResponseWithListItemMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWithListItemMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWithListItemResult struct {
	// API Resource UUID tag.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the List.
	Description string `json:"description"`
	// The items in the List.
	Items []SingleResponseWithListItemResultItem `json:"items"`
	// The name of the List.
	Name string `json:"name"`
	// The type of List.
	Type      SingleResponseWithListItemResultType `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      singleResponseWithListItemResultJSON `json:"-"`
}

// singleResponseWithListItemResultJSON contains the JSON metadata for the struct
// [SingleResponseWithListItemResult]
type singleResponseWithListItemResultJSON struct {
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

func (r *SingleResponseWithListItemResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SingleResponseWithListItemResultItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The value of the item in a List.
	Value string                                   `json:"value"`
	JSON  singleResponseWithListItemResultItemJSON `json:"-"`
}

// singleResponseWithListItemResultItemJSON contains the JSON metadata for the
// struct [SingleResponseWithListItemResultItem]
type singleResponseWithListItemResultItemJSON struct {
	CreatedAt   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SingleResponseWithListItemResultItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of List.
type SingleResponseWithListItemResultType string

const (
	SingleResponseWithListItemResultTypeSerial SingleResponseWithListItemResultType = "SERIAL"
	SingleResponseWithListItemResultTypeURL    SingleResponseWithListItemResultType = "URL"
	SingleResponseWithListItemResultTypeDomain SingleResponseWithListItemResultType = "DOMAIN"
	SingleResponseWithListItemResultTypeEmail  SingleResponseWithListItemResultType = "EMAIL"
	SingleResponseWithListItemResultTypeIP     SingleResponseWithListItemResultType = "IP"
)

// Whether the API call was successful
type SingleResponseWithListItemSuccess bool

const (
	SingleResponseWithListItemSuccessTrue SingleResponseWithListItemSuccess = true
)

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
	// The number of items in the List.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the List.
	Description string `json:"description"`
	// The name of the List.
	Name string `json:"name"`
	// The type of List.
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

// The type of List.
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
	// The number of items in the List.
	Count     float64   `json:"count"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the List.
	Description string `json:"description"`
	// The name of the List.
	Name string `json:"name"`
	// The type of List.
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

// The type of List.
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

type AccountGatewayListUpdateParams struct {
	// The name of the List.
	Name param.Field[string] `json:"name,required"`
	// The description of the List.
	Description param.Field[string] `json:"description"`
}

func (r AccountGatewayListUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountGatewayListZeroTrustListsNewZeroTrustListParams struct {
	// The name of the List.
	Name param.Field[string] `json:"name,required"`
	// The type of List.
	Type param.Field[AccountGatewayListZeroTrustListsNewZeroTrustListParamsType] `json:"type,required"`
	// The description of the List.
	Description param.Field[string] `json:"description"`
	// The items in the List.
	Items param.Field[[]AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem] `json:"items"`
}

func (r AccountGatewayListZeroTrustListsNewZeroTrustListParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of List.
type AccountGatewayListZeroTrustListsNewZeroTrustListParamsType string

const (
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeSerial AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "SERIAL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeURL    AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "URL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeDomain AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "DOMAIN"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeEmail  AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "EMAIL"
	AccountGatewayListZeroTrustListsNewZeroTrustListParamsTypeIP     AccountGatewayListZeroTrustListsNewZeroTrustListParamsType = "IP"
)

type AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem struct {
	// The value of the item in a List.
	Value param.Field[string] `json:"value"`
}

func (r AccountGatewayListZeroTrustListsNewZeroTrustListParamsItem) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
