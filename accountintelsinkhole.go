// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelSinkholeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelSinkholeService]
// method instead.
type AccountIntelSinkholeService struct {
	Options []option.RequestOption
}

// NewAccountIntelSinkholeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelSinkholeService(opts ...option.RequestOption) (r *AccountIntelSinkholeService) {
	r = &AccountIntelSinkholeService{}
	r.Options = opts
	return
}

// List sinkholes owned by this account
func (r *AccountIntelSinkholeService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountIntelSinkholeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/sinkholes", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountIntelSinkholeListResponse struct {
	Errors   []AccountIntelSinkholeListResponseError   `json:"errors"`
	Messages []AccountIntelSinkholeListResponseMessage `json:"messages"`
	Result   []AccountIntelSinkholeListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountIntelSinkholeListResponseSuccess `json:"success"`
	JSON    accountIntelSinkholeListResponseJSON    `json:"-"`
}

// accountIntelSinkholeListResponseJSON contains the JSON metadata for the struct
// [AccountIntelSinkholeListResponse]
type accountIntelSinkholeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelSinkholeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelSinkholeListResponseError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountIntelSinkholeListResponseErrorJSON `json:"-"`
}

// accountIntelSinkholeListResponseErrorJSON contains the JSON metadata for the
// struct [AccountIntelSinkholeListResponseError]
type accountIntelSinkholeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelSinkholeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelSinkholeListResponseMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accountIntelSinkholeListResponseMessageJSON `json:"-"`
}

// accountIntelSinkholeListResponseMessageJSON contains the JSON metadata for the
// struct [AccountIntelSinkholeListResponseMessage]
type accountIntelSinkholeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelSinkholeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelSinkholeListResponseResult struct {
	// The unique identifier for the sinkhole
	ID int64 `json:"id"`
	// The account tag that owns this sinkhole
	AccountTag string `json:"account_tag"`
	// The date and time when the sinkhole was created
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The date and time when the sinkhole was last modified
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The name of the sinkhole
	Name string `json:"name"`
	// The name of the R2 bucket to store results
	R2Bucket string `json:"r2_bucket"`
	// The id of the R2 instance
	R2ID string                                     `json:"r2_id"`
	JSON accountIntelSinkholeListResponseResultJSON `json:"-"`
}

// accountIntelSinkholeListResponseResultJSON contains the JSON metadata for the
// struct [AccountIntelSinkholeListResponseResult]
type accountIntelSinkholeListResponseResultJSON struct {
	ID          apijson.Field
	AccountTag  apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	R2Bucket    apijson.Field
	R2ID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelSinkholeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelSinkholeListResponseSuccess bool

const (
	AccountIntelSinkholeListResponseSuccessTrue AccountIntelSinkholeListResponseSuccess = true
)
