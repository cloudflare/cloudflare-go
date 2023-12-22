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

// AccountBillingProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountBillingProfileService]
// method instead.
type AccountBillingProfileService struct {
	Options []option.RequestOption
}

// NewAccountBillingProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountBillingProfileService(opts ...option.RequestOption) (r *AccountBillingProfileService) {
	r = &AccountBillingProfileService{}
	r.Options = opts
	return
}

// Gets the current billing profile for the account.
func (r *AccountBillingProfileService) AccountBillingProfileBillingProfileDetails(ctx context.Context, accountIdentifier interface{}, opts ...option.RequestOption) (res *BillingResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/billing/profile", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BillingResponseSingle struct {
	Errors   []BillingResponseSingleError   `json:"errors"`
	Messages []BillingResponseSingleMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success BillingResponseSingleSuccess `json:"success"`
	JSON    billingResponseSingleJSON    `json:"-"`
}

// billingResponseSingleJSON contains the JSON metadata for the struct
// [BillingResponseSingle]
type billingResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingResponseSingleError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    billingResponseSingleErrorJSON `json:"-"`
}

// billingResponseSingleErrorJSON contains the JSON metadata for the struct
// [BillingResponseSingleError]
type billingResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type BillingResponseSingleMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    billingResponseSingleMessageJSON `json:"-"`
}

// billingResponseSingleMessageJSON contains the JSON metadata for the struct
// [BillingResponseSingleMessage]
type billingResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BillingResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type BillingResponseSingleSuccess bool

const (
	BillingResponseSingleSuccessTrue BillingResponseSingleSuccess = true
)
