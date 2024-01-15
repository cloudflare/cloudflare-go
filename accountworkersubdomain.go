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

// AccountWorkerSubdomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWorkerSubdomainService]
// method instead.
type AccountWorkerSubdomainService struct {
	Options []option.RequestOption
}

// NewAccountWorkerSubdomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerSubdomainService(opts ...option.RequestOption) (r *AccountWorkerSubdomainService) {
	r = &AccountWorkerSubdomainService{}
	r.Options = opts
	return
}

// Creates a Workers subdomain for an account.
func (r *AccountWorkerSubdomainService) WorkerSubdomainNewSubdomain(ctx context.Context, accountIdentifier string, body AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams, opts ...option.RequestOption) (res *AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a Workers subdomain for an account.
func (r *AccountWorkerSubdomainService) WorkerSubdomainGetSubdomain(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponse struct {
	Errors   []AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseError   `json:"errors"`
	Messages []AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessage `json:"messages"`
	Result   AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseSuccess `json:"success"`
	JSON    accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseJSON    `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponse]
type accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseErrorJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseError]
type accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessageJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessage]
type accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResult struct {
	Name interface{}                                                         `json:"name"`
	JSON accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResultJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResult]
type accountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResultJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseSuccess bool

const (
	AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseSuccessTrue AccountWorkerSubdomainWorkerSubdomainNewSubdomainResponseSuccess = true
)

type AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponse struct {
	Errors   []AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseError   `json:"errors"`
	Messages []AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessage `json:"messages"`
	Result   AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseSuccess `json:"success"`
	JSON    accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseJSON    `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseJSON contains the JSON
// metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponse]
type accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseError struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseErrorJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseErrorJSON contains the
// JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseError]
type accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessage struct {
	Code    int64                                                                `json:"code,required"`
	Message string                                                               `json:"message,required"`
	JSON    accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessageJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessageJSON contains
// the JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessage]
type accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResult struct {
	Name interface{}                                                         `json:"name"`
	JSON accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResultJSON `json:"-"`
}

// accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResult]
type accountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResultJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseSuccess bool

const (
	AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseSuccessTrue AccountWorkerSubdomainWorkerSubdomainGetSubdomainResponseSuccess = true
)

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
