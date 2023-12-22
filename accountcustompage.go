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

// AccountCustomPageService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountCustomPageService] method
// instead.
type AccountCustomPageService struct {
	Options []option.RequestOption
}

// NewAccountCustomPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountCustomPageService(opts ...option.RequestOption) (r *AccountCustomPageService) {
	r = &AccountCustomPageService{}
	r.Options = opts
	return
}

// Fetches the details of a custom page.
func (r *AccountCustomPageService) Get(ctx context.Context, accountIdentifier string, identifier AccountCustomPageGetParamsIdentifier, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the configuration of an existing custom page.
func (r *AccountCustomPageService) Update(ctx context.Context, accountIdentifier string, identifier AccountCustomPageUpdateParamsIdentifier, body AccountCustomPageUpdateParams, opts ...option.RequestOption) (res *CustomPagesResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages/%v", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Fetches all the custom pages at the account level.
func (r *AccountCustomPageService) CustomPagesForAnAccountListCustomPages(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *CustomPagesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/custom_pages", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CustomPagesResponseCollection struct {
	Errors     []CustomPagesResponseCollectionError    `json:"errors"`
	Messages   []CustomPagesResponseCollectionMessage  `json:"messages"`
	Result     []interface{}                           `json:"result"`
	ResultInfo CustomPagesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success CustomPagesResponseCollectionSuccess `json:"success"`
	JSON    customPagesResponseCollectionJSON    `json:"-"`
}

// customPagesResponseCollectionJSON contains the JSON metadata for the struct
// [CustomPagesResponseCollection]
type customPagesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPagesResponseCollectionError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    customPagesResponseCollectionErrorJSON `json:"-"`
}

// customPagesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [CustomPagesResponseCollectionError]
type customPagesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPagesResponseCollectionMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    customPagesResponseCollectionMessageJSON `json:"-"`
}

// customPagesResponseCollectionMessageJSON contains the JSON metadata for the
// struct [CustomPagesResponseCollectionMessage]
type customPagesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPagesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                     `json:"total_count"`
	JSON       customPagesResponseCollectionResultInfoJSON `json:"-"`
}

// customPagesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [CustomPagesResponseCollectionResultInfo]
type customPagesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomPagesResponseCollectionSuccess bool

const (
	CustomPagesResponseCollectionSuccessTrue CustomPagesResponseCollectionSuccess = true
)

type CustomPagesResponseSingle struct {
	Errors   []CustomPagesResponseSingleError   `json:"errors"`
	Messages []CustomPagesResponseSingleMessage `json:"messages"`
	Result   interface{}                        `json:"result"`
	// Whether the API call was successful
	Success CustomPagesResponseSingleSuccess `json:"success"`
	JSON    customPagesResponseSingleJSON    `json:"-"`
}

// customPagesResponseSingleJSON contains the JSON metadata for the struct
// [CustomPagesResponseSingle]
type customPagesResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPagesResponseSingleError struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    customPagesResponseSingleErrorJSON `json:"-"`
}

// customPagesResponseSingleErrorJSON contains the JSON metadata for the struct
// [CustomPagesResponseSingleError]
type customPagesResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomPagesResponseSingleMessage struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    customPagesResponseSingleMessageJSON `json:"-"`
}

// customPagesResponseSingleMessageJSON contains the JSON metadata for the struct
// [CustomPagesResponseSingleMessage]
type customPagesResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomPagesResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type CustomPagesResponseSingleSuccess bool

const (
	CustomPagesResponseSingleSuccessTrue CustomPagesResponseSingleSuccess = true
)

// The name of the custom page type.
type AccountCustomPageGetParamsIdentifier string

const (
	AccountCustomPageGetParamsIdentifierBasicChallenge   AccountCustomPageGetParamsIdentifier = "basic_challenge"
	AccountCustomPageGetParamsIdentifierManagedChallenge AccountCustomPageGetParamsIdentifier = "managed_challenge"
	AccountCustomPageGetParamsIdentifierWafBlock         AccountCustomPageGetParamsIdentifier = "waf_block"
	AccountCustomPageGetParamsIdentifierCountryChallenge AccountCustomPageGetParamsIdentifier = "country_challenge"
	AccountCustomPageGetParamsIdentifierIPBlock          AccountCustomPageGetParamsIdentifier = "ip_block"
	AccountCustomPageGetParamsIdentifierUnderAttack      AccountCustomPageGetParamsIdentifier = "under_attack"
	AccountCustomPageGetParamsIdentifierRatelimitBlock   AccountCustomPageGetParamsIdentifier = "ratelimit_block"
	AccountCustomPageGetParamsIdentifier500Errors        AccountCustomPageGetParamsIdentifier = "500_errors"
	AccountCustomPageGetParamsIdentifier1000Errors       AccountCustomPageGetParamsIdentifier = "1000_errors"
)

type AccountCustomPageUpdateParams struct {
	// The custom page state.
	State param.Field[AccountCustomPageUpdateParamsState] `json:"state,required"`
	// The URL associated with the custom page.
	URL param.Field[string] `json:"url,required" format:"uri"`
}

func (r AccountCustomPageUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The name of the custom page type.
type AccountCustomPageUpdateParamsIdentifier string

const (
	AccountCustomPageUpdateParamsIdentifierBasicChallenge   AccountCustomPageUpdateParamsIdentifier = "basic_challenge"
	AccountCustomPageUpdateParamsIdentifierManagedChallenge AccountCustomPageUpdateParamsIdentifier = "managed_challenge"
	AccountCustomPageUpdateParamsIdentifierWafBlock         AccountCustomPageUpdateParamsIdentifier = "waf_block"
	AccountCustomPageUpdateParamsIdentifierCountryChallenge AccountCustomPageUpdateParamsIdentifier = "country_challenge"
	AccountCustomPageUpdateParamsIdentifierIPBlock          AccountCustomPageUpdateParamsIdentifier = "ip_block"
	AccountCustomPageUpdateParamsIdentifierUnderAttack      AccountCustomPageUpdateParamsIdentifier = "under_attack"
	AccountCustomPageUpdateParamsIdentifierRatelimitBlock   AccountCustomPageUpdateParamsIdentifier = "ratelimit_block"
	AccountCustomPageUpdateParamsIdentifier500Errors        AccountCustomPageUpdateParamsIdentifier = "500_errors"
	AccountCustomPageUpdateParamsIdentifier1000Errors       AccountCustomPageUpdateParamsIdentifier = "1000_errors"
)

// The custom page state.
type AccountCustomPageUpdateParamsState string

const (
	AccountCustomPageUpdateParamsStateDefault    AccountCustomPageUpdateParamsState = "default"
	AccountCustomPageUpdateParamsStateCustomized AccountCustomPageUpdateParamsState = "customized"
)
