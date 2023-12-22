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

// AccountLogpushOwnershipService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLogpushOwnershipService] method instead.
type AccountLogpushOwnershipService struct {
	Options   []option.RequestOption
	Validates *AccountLogpushOwnershipValidateService
}

// NewAccountLogpushOwnershipService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLogpushOwnershipService(opts ...option.RequestOption) (r *AccountLogpushOwnershipService) {
	r = &AccountLogpushOwnershipService{}
	r.Options = opts
	r.Validates = NewAccountLogpushOwnershipValidateService(opts...)
	return
}

// Gets a new ownership challenge sent to your destination.
func (r *AccountLogpushOwnershipService) PostAccountsAccountIdentifierLogpushOwnership(ctx context.Context, accountIdentifier string, body AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams, opts ...option.RequestOption) (res *GetOwnershipResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/logpush/ownership", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type GetOwnershipResponse struct {
	Errors   []GetOwnershipResponseError   `json:"errors"`
	Messages []GetOwnershipResponseMessage `json:"messages"`
	Result   GetOwnershipResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success GetOwnershipResponseSuccess `json:"success"`
	JSON    getOwnershipResponseJSON    `json:"-"`
}

// getOwnershipResponseJSON contains the JSON metadata for the struct
// [GetOwnershipResponse]
type getOwnershipResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetOwnershipResponseError struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    getOwnershipResponseErrorJSON `json:"-"`
}

// getOwnershipResponseErrorJSON contains the JSON metadata for the struct
// [GetOwnershipResponseError]
type getOwnershipResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetOwnershipResponseMessage struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    getOwnershipResponseMessageJSON `json:"-"`
}

// getOwnershipResponseMessageJSON contains the JSON metadata for the struct
// [GetOwnershipResponseMessage]
type getOwnershipResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type GetOwnershipResponseResult struct {
	Filename string                         `json:"filename"`
	Message  string                         `json:"message"`
	Valid    bool                           `json:"valid"`
	JSON     getOwnershipResponseResultJSON `json:"-"`
}

// getOwnershipResponseResultJSON contains the JSON metadata for the struct
// [GetOwnershipResponseResult]
type getOwnershipResponseResultJSON struct {
	Filename    apijson.Field
	Message     apijson.Field
	Valid       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GetOwnershipResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type GetOwnershipResponseSuccess bool

const (
	GetOwnershipResponseSuccessTrue GetOwnershipResponseSuccess = true
)

type AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams struct {
	// Uniquely identifies a resource (such as an s3 bucket) where data will be pushed.
	// Additional configuration parameters supported by the destination may be
	// included.
	DestinationConf param.Field[string] `json:"destination_conf,required" format:"uri"`
}

func (r AccountLogpushOwnershipPostAccountsAccountIdentifierLogpushOwnershipParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
