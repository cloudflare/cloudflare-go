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
func (r *AccountWorkerSubdomainService) WorkerSubdomainNewSubdomain(ctx context.Context, accountIdentifier string, body AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams, opts ...option.RequestOption) (res *SubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a Workers subdomain for an account.
func (r *AccountWorkerSubdomainService) WorkerSubdomainGetSubdomain(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *SubdomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/subdomain", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SubdomainResponse struct {
	Errors   []SubdomainResponseError   `json:"errors"`
	Messages []SubdomainResponseMessage `json:"messages"`
	Result   SubdomainResponseResult    `json:"result"`
	// Whether the API call was successful
	Success SubdomainResponseSuccess `json:"success"`
	JSON    subdomainResponseJSON    `json:"-"`
}

// subdomainResponseJSON contains the JSON metadata for the struct
// [SubdomainResponse]
type subdomainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubdomainResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    subdomainResponseErrorJSON `json:"-"`
}

// subdomainResponseErrorJSON contains the JSON metadata for the struct
// [SubdomainResponseError]
type subdomainResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubdomainResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    subdomainResponseMessageJSON `json:"-"`
}

// subdomainResponseMessageJSON contains the JSON metadata for the struct
// [SubdomainResponseMessage]
type subdomainResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SubdomainResponseResult struct {
	Name interface{}                 `json:"name"`
	JSON subdomainResponseResultJSON `json:"-"`
}

// subdomainResponseResultJSON contains the JSON metadata for the struct
// [SubdomainResponseResult]
type subdomainResponseResultJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SubdomainResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SubdomainResponseSuccess bool

const (
	SubdomainResponseSuccessTrue SubdomainResponseSuccess = true
)

type AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams struct {
	Body param.Field[interface{}] `json:"body,required"`
}

func (r AccountWorkerSubdomainWorkerSubdomainNewSubdomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
