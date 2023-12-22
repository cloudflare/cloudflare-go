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

// Gets a Worker Domain.
func (r *AccountWorkerDomainService) Get(ctx context.Context, accountIdentifier interface{}, domainIdentifier interface{}, opts ...option.RequestOption) (res *DomainResponseSingle9PyIq1Vp, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountIdentifier, domainIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Detaches a worker from a zone and hostname.
func (r *AccountWorkerDomainService) Delete(ctx context.Context, accountIdentifier interface{}, domainIdentifier interface{}, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountIdentifier, domainIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Attaches a worker to a zone and hostname.
func (r *AccountWorkerDomainService) WorkerDomainAttachToDomain(ctx context.Context, accountIdentifier interface{}, body AccountWorkerDomainWorkerDomainAttachToDomainParams, opts ...option.RequestOption) (res *DomainResponseSingle9PyIq1Vp, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Lists all Worker Domains.
func (r *AccountWorkerDomainService) WorkerDomainListDomains(ctx context.Context, accountIdentifier interface{}, query AccountWorkerDomainWorkerDomainListDomainsParams, opts ...option.RequestOption) (res *DomainResponseCollection1E2A3hYu, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DomainResponseCollection1E2A3hYu struct {
	Errors   []DomainResponseCollection1E2A3hYuError   `json:"errors"`
	Messages []DomainResponseCollection1E2A3hYuMessage `json:"messages"`
	Result   []DomainResponseCollection1E2A3hYuResult  `json:"result"`
	// Whether the API call was successful
	Success DomainResponseCollection1E2A3hYuSuccess `json:"success"`
	JSON    domainResponseCollection1E2A3hYuJSON    `json:"-"`
}

// domainResponseCollection1E2A3hYuJSON contains the JSON metadata for the struct
// [DomainResponseCollection1E2A3hYu]
type domainResponseCollection1E2A3hYuJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollection1E2A3hYu) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollection1E2A3hYuError struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    domainResponseCollection1E2A3hYuErrorJSON `json:"-"`
}

// domainResponseCollection1E2A3hYuErrorJSON contains the JSON metadata for the
// struct [DomainResponseCollection1E2A3hYuError]
type domainResponseCollection1E2A3hYuErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollection1E2A3hYuError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollection1E2A3hYuMessage struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    domainResponseCollection1E2A3hYuMessageJSON `json:"-"`
}

// domainResponseCollection1E2A3hYuMessageJSON contains the JSON metadata for the
// struct [DomainResponseCollection1E2A3hYuMessage]
type domainResponseCollection1E2A3hYuMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollection1E2A3hYuMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseCollection1E2A3hYuResult struct {
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
	ZoneName string                                     `json:"zone_name"`
	JSON     domainResponseCollection1E2A3hYuResultJSON `json:"-"`
}

// domainResponseCollection1E2A3hYuResultJSON contains the JSON metadata for the
// struct [DomainResponseCollection1E2A3hYuResult]
type domainResponseCollection1E2A3hYuResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseCollection1E2A3hYuResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainResponseCollection1E2A3hYuSuccess bool

const (
	DomainResponseCollection1E2A3hYuSuccessTrue DomainResponseCollection1E2A3hYuSuccess = true
)

type DomainResponseSingle9PyIq1Vp struct {
	Errors   []DomainResponseSingle9PyIq1VpError   `json:"errors"`
	Messages []DomainResponseSingle9PyIq1VpMessage `json:"messages"`
	Result   DomainResponseSingle9PyIq1VpResult    `json:"result"`
	// Whether the API call was successful
	Success DomainResponseSingle9PyIq1VpSuccess `json:"success"`
	JSON    domainResponseSingle9PyIq1VpJSON    `json:"-"`
}

// domainResponseSingle9PyIq1VpJSON contains the JSON metadata for the struct
// [DomainResponseSingle9PyIq1Vp]
type domainResponseSingle9PyIq1VpJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingle9PyIq1Vp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseSingle9PyIq1VpError struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    domainResponseSingle9PyIq1VpErrorJSON `json:"-"`
}

// domainResponseSingle9PyIq1VpErrorJSON contains the JSON metadata for the struct
// [DomainResponseSingle9PyIq1VpError]
type domainResponseSingle9PyIq1VpErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingle9PyIq1VpError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseSingle9PyIq1VpMessage struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    domainResponseSingle9PyIq1VpMessageJSON `json:"-"`
}

// domainResponseSingle9PyIq1VpMessageJSON contains the JSON metadata for the
// struct [DomainResponseSingle9PyIq1VpMessage]
type domainResponseSingle9PyIq1VpMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingle9PyIq1VpMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DomainResponseSingle9PyIq1VpResult struct {
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
	ZoneName string                                 `json:"zone_name"`
	JSON     domainResponseSingle9PyIq1VpResultJSON `json:"-"`
}

// domainResponseSingle9PyIq1VpResultJSON contains the JSON metadata for the struct
// [DomainResponseSingle9PyIq1VpResult]
type domainResponseSingle9PyIq1VpResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DomainResponseSingle9PyIq1VpResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DomainResponseSingle9PyIq1VpSuccess bool

const (
	DomainResponseSingle9PyIq1VpSuccessTrue DomainResponseSingle9PyIq1VpSuccess = true
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
