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

// WorkerDomainService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerDomainService] method
// instead.
type WorkerDomainService struct {
	Options []option.RequestOption
}

// NewWorkerDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerDomainService(opts ...option.RequestOption) (r *WorkerDomainService) {
	r = &WorkerDomainService{}
	r.Options = opts
	return
}

// Detaches a Worker from a zone and hostname.
func (r *WorkerDomainService) Delete(ctx context.Context, accountID interface{}, domainID interface{}, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Gets a Worker domain.
func (r *WorkerDomainService) Get(ctx context.Context, accountID interface{}, domainID interface{}, opts ...option.RequestOption) (res *WorkerDomainGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDomainGetResponseEnvelope
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Attaches a Worker to a zone and hostname.
func (r *WorkerDomainService) WorkerDomainAttachToDomain(ctx context.Context, accountID interface{}, body WorkerDomainWorkerDomainAttachToDomainParams, opts ...option.RequestOption) (res *WorkerDomainWorkerDomainAttachToDomainResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDomainWorkerDomainAttachToDomainResponseEnvelope
	path := fmt.Sprintf("accounts/%v/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Worker Domains for an account.
func (r *WorkerDomainService) WorkerDomainListDomains(ctx context.Context, accountID interface{}, query WorkerDomainWorkerDomainListDomainsParams, opts ...option.RequestOption) (res *[]WorkerDomainWorkerDomainListDomainsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDomainWorkerDomainListDomainsResponseEnvelope
	path := fmt.Sprintf("accounts/%v/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkerDomainGetResponse struct {
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
	ZoneName string                      `json:"zone_name"`
	JSON     workerDomainGetResponseJSON `json:"-"`
}

// workerDomainGetResponseJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponse]
type workerDomainGetResponseJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainAttachToDomainResponse struct {
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
	ZoneName string                                             `json:"zone_name"`
	JSON     workerDomainWorkerDomainAttachToDomainResponseJSON `json:"-"`
}

// workerDomainWorkerDomainAttachToDomainResponseJSON contains the JSON metadata
// for the struct [WorkerDomainWorkerDomainAttachToDomainResponse]
type workerDomainWorkerDomainAttachToDomainResponseJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainAttachToDomainResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainListDomainsResponse struct {
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
	ZoneName string                                          `json:"zone_name"`
	JSON     workerDomainWorkerDomainListDomainsResponseJSON `json:"-"`
}

// workerDomainWorkerDomainListDomainsResponseJSON contains the JSON metadata for
// the struct [WorkerDomainWorkerDomainListDomainsResponse]
type workerDomainWorkerDomainListDomainsResponseJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainListDomainsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseEnvelope struct {
	Errors   []WorkerDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDomainGetResponseEnvelopeJSON    `json:"-"`
}

// workerDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponseEnvelope]
type workerDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    workerDomainGetResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerDomainGetResponseEnvelopeErrors]
type workerDomainGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    workerDomainGetResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerDomainGetResponseEnvelopeMessages]
type workerDomainGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainGetResponseEnvelopeSuccess bool

const (
	WorkerDomainGetResponseEnvelopeSuccessTrue WorkerDomainGetResponseEnvelopeSuccess = true
)

type WorkerDomainWorkerDomainAttachToDomainParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[interface{}] `json:"zone_id,required"`
}

func (r WorkerDomainWorkerDomainAttachToDomainParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDomainWorkerDomainAttachToDomainResponseEnvelope struct {
	Errors   []WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDomainWorkerDomainAttachToDomainResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDomainWorkerDomainAttachToDomainResponseEnvelopeJSON    `json:"-"`
}

// workerDomainWorkerDomainAttachToDomainResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerDomainWorkerDomainAttachToDomainResponseEnvelope]
type workerDomainWorkerDomainAttachToDomainResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainAttachToDomainResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    workerDomainWorkerDomainAttachToDomainResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDomainWorkerDomainAttachToDomainResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeErrors]
type workerDomainWorkerDomainAttachToDomainResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    workerDomainWorkerDomainAttachToDomainResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDomainWorkerDomainAttachToDomainResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeMessages]
type workerDomainWorkerDomainAttachToDomainResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeSuccess bool

const (
	WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeSuccessTrue WorkerDomainWorkerDomainAttachToDomainResponseEnvelopeSuccess = true
)

type WorkerDomainWorkerDomainListDomainsParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `query:"environment"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `query:"hostname"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `query:"service"`
	// Identifier of the zone.
	ZoneID param.Field[interface{}] `query:"zone_id"`
	// Name of the zone.
	ZoneName param.Field[string] `query:"zone_name"`
}

// URLQuery serializes [WorkerDomainWorkerDomainListDomainsParams]'s query
// parameters as `url.Values`.
func (r WorkerDomainWorkerDomainListDomainsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerDomainWorkerDomainListDomainsResponseEnvelope struct {
	Errors   []WorkerDomainWorkerDomainListDomainsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDomainWorkerDomainListDomainsResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerDomainWorkerDomainListDomainsResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDomainWorkerDomainListDomainsResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDomainWorkerDomainListDomainsResponseEnvelopeJSON    `json:"-"`
}

// workerDomainWorkerDomainListDomainsResponseEnvelopeJSON contains the JSON
// metadata for the struct [WorkerDomainWorkerDomainListDomainsResponseEnvelope]
type workerDomainWorkerDomainListDomainsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainListDomainsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainListDomainsResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    workerDomainWorkerDomainListDomainsResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDomainWorkerDomainListDomainsResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [WorkerDomainWorkerDomainListDomainsResponseEnvelopeErrors]
type workerDomainWorkerDomainListDomainsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainListDomainsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainWorkerDomainListDomainsResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    workerDomainWorkerDomainListDomainsResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDomainWorkerDomainListDomainsResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [WorkerDomainWorkerDomainListDomainsResponseEnvelopeMessages]
type workerDomainWorkerDomainListDomainsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainWorkerDomainListDomainsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainWorkerDomainListDomainsResponseEnvelopeSuccess bool

const (
	WorkerDomainWorkerDomainListDomainsResponseEnvelopeSuccessTrue WorkerDomainWorkerDomainListDomainsResponseEnvelopeSuccess = true
)
