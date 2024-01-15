// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountWorkerDomainService]
// method instead.
type AccountWorkerDomainService struct {
	Options []option.RequestOption
}

// NewAccountWorkerDomainService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountWorkerDomainService(opts ...option.RequestOption) (r *AccountWorkerDomainService) {
	r = &AccountWorkerDomainService{}
	r.Options = opts
	return
}

// Gets a Worker domain.
func (r *AccountWorkerDomainService) Get(ctx context.Context, accountIdentifier interface{}, domainIdentifier interface{}, opts ...option.RequestOption) (res *AccountWorkerDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountIdentifier, domainIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Detaches a Worker from a zone and hostname.
func (r *AccountWorkerDomainService) Delete(ctx context.Context, accountIdentifier interface{}, domainIdentifier interface{}, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountIdentifier, domainIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Attaches a Worker to a zone and hostname.
func (r *AccountWorkerDomainService) WorkerDomainAttachToDomain(ctx context.Context, accountIdentifier interface{}, body AccountWorkerDomainWorkerDomainAttachToDomainParams, opts ...option.RequestOption) (res *AccountWorkerDomainWorkerDomainAttachToDomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Worker Domains for an account.
func (r *AccountWorkerDomainService) WorkerDomainListDomains(ctx context.Context, accountIdentifier interface{}, query AccountWorkerDomainWorkerDomainListDomainsParams, opts ...option.RequestOption) (res *AccountWorkerDomainWorkerDomainListDomainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountWorkerDomainGetResponse struct {
	Errors   []AccountWorkerDomainGetResponseError   `json:"errors"`
	Messages []AccountWorkerDomainGetResponseMessage `json:"messages"`
	Result   AccountWorkerDomainGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDomainGetResponseSuccess `json:"success"`
	JSON    accountWorkerDomainGetResponseJSON    `json:"-"`
}

// accountWorkerDomainGetResponseJSON contains the JSON metadata for the struct
// [AccountWorkerDomainGetResponse]
type accountWorkerDomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    accountWorkerDomainGetResponseErrorJSON `json:"-"`
}

// accountWorkerDomainGetResponseErrorJSON contains the JSON metadata for the
// struct [AccountWorkerDomainGetResponseError]
type accountWorkerDomainGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accountWorkerDomainGetResponseMessageJSON `json:"-"`
}

// accountWorkerDomainGetResponseMessageJSON contains the JSON metadata for the
// struct [AccountWorkerDomainGetResponseMessage]
type accountWorkerDomainGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainGetResponseResult struct {
	// Identifer of the Worker Domain.
	ID interface{} `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID interface{} `json:"zone_id"`
	// Name of the zone.
	ZoneName string                                   `json:"zone_name"`
	JSON     accountWorkerDomainGetResponseResultJSON `json:"-"`
}

// accountWorkerDomainGetResponseResultJSON contains the JSON metadata for the
// struct [AccountWorkerDomainGetResponseResult]
type accountWorkerDomainGetResponseResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDomainGetResponseSuccess bool

const (
	AccountWorkerDomainGetResponseSuccessTrue AccountWorkerDomainGetResponseSuccess = true
)

type AccountWorkerDomainWorkerDomainAttachToDomainResponse struct {
	Errors   []AccountWorkerDomainWorkerDomainAttachToDomainResponseError   `json:"errors"`
	Messages []AccountWorkerDomainWorkerDomainAttachToDomainResponseMessage `json:"messages"`
	Result   AccountWorkerDomainWorkerDomainAttachToDomainResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDomainWorkerDomainAttachToDomainResponseSuccess `json:"success"`
	JSON    accountWorkerDomainWorkerDomainAttachToDomainResponseJSON    `json:"-"`
}

// accountWorkerDomainWorkerDomainAttachToDomainResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDomainWorkerDomainAttachToDomainResponse]
type accountWorkerDomainWorkerDomainAttachToDomainResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainAttachToDomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainAttachToDomainResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountWorkerDomainWorkerDomainAttachToDomainResponseErrorJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainAttachToDomainResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountWorkerDomainWorkerDomainAttachToDomainResponseError]
type accountWorkerDomainWorkerDomainAttachToDomainResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainAttachToDomainResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainAttachToDomainResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountWorkerDomainWorkerDomainAttachToDomainResponseMessageJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainAttachToDomainResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountWorkerDomainWorkerDomainAttachToDomainResponseMessage]
type accountWorkerDomainWorkerDomainAttachToDomainResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainAttachToDomainResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainAttachToDomainResponseResult struct {
	// Identifer of the Worker Domain.
	ID interface{} `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID interface{} `json:"zone_id"`
	// Name of the zone.
	ZoneName string                                                          `json:"zone_name"`
	JSON     accountWorkerDomainWorkerDomainAttachToDomainResponseResultJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainAttachToDomainResponseResultJSON contains the
// JSON metadata for the struct
// [AccountWorkerDomainWorkerDomainAttachToDomainResponseResult]
type accountWorkerDomainWorkerDomainAttachToDomainResponseResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainAttachToDomainResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDomainWorkerDomainAttachToDomainResponseSuccess bool

const (
	AccountWorkerDomainWorkerDomainAttachToDomainResponseSuccessTrue AccountWorkerDomainWorkerDomainAttachToDomainResponseSuccess = true
)

type AccountWorkerDomainWorkerDomainListDomainsResponse struct {
	Errors   []AccountWorkerDomainWorkerDomainListDomainsResponseError   `json:"errors"`
	Messages []AccountWorkerDomainWorkerDomainListDomainsResponseMessage `json:"messages"`
	Result   []AccountWorkerDomainWorkerDomainListDomainsResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountWorkerDomainWorkerDomainListDomainsResponseSuccess `json:"success"`
	JSON    accountWorkerDomainWorkerDomainListDomainsResponseJSON    `json:"-"`
}

// accountWorkerDomainWorkerDomainListDomainsResponseJSON contains the JSON
// metadata for the struct [AccountWorkerDomainWorkerDomainListDomainsResponse]
type accountWorkerDomainWorkerDomainListDomainsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainListDomainsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainListDomainsResponseError struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    accountWorkerDomainWorkerDomainListDomainsResponseErrorJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainListDomainsResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountWorkerDomainWorkerDomainListDomainsResponseError]
type accountWorkerDomainWorkerDomainListDomainsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainListDomainsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainListDomainsResponseMessage struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    accountWorkerDomainWorkerDomainListDomainsResponseMessageJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainListDomainsResponseMessageJSON contains the JSON
// metadata for the struct
// [AccountWorkerDomainWorkerDomainListDomainsResponseMessage]
type accountWorkerDomainWorkerDomainListDomainsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainListDomainsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountWorkerDomainWorkerDomainListDomainsResponseResult struct {
	// Identifer of the Worker Domain.
	ID interface{} `json:"id"`
	// Worker environment associated with the zone and hostname.
	Environment string `json:"environment"`
	// Hostname of the Worker Domain.
	Hostname string `json:"hostname"`
	// Worker service associated with the zone and hostname.
	Service string `json:"service"`
	// Identifier of the zone.
	ZoneID interface{} `json:"zone_id"`
	// Name of the zone.
	ZoneName string                                                       `json:"zone_name"`
	JSON     accountWorkerDomainWorkerDomainListDomainsResponseResultJSON `json:"-"`
}

// accountWorkerDomainWorkerDomainListDomainsResponseResultJSON contains the JSON
// metadata for the struct
// [AccountWorkerDomainWorkerDomainListDomainsResponseResult]
type accountWorkerDomainWorkerDomainListDomainsResponseResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountWorkerDomainWorkerDomainListDomainsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountWorkerDomainWorkerDomainListDomainsResponseSuccess bool

const (
	AccountWorkerDomainWorkerDomainListDomainsResponseSuccessTrue AccountWorkerDomainWorkerDomainListDomainsResponseSuccess = true
)

type AccountWorkerDomainWorkerDomainAttachToDomainParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[interface{}] `json:"zone_id,required"`
}

func (r AccountWorkerDomainWorkerDomainAttachToDomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountWorkerDomainWorkerDomainListDomainsParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `query:"hostname"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `query:"service"`
	// Identifier of the zone.
	ZoneIdentifier param.Field[interface{}] `query:"zone_identifier"`
	// Name of the zone.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [AccountWorkerDomainWorkerDomainListDomainsParams]'s query
// parameters as `url.Values`.
func (r AccountWorkerDomainWorkerDomainListDomainsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
