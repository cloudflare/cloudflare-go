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

// Attaches a Worker to a zone and hostname.
func (r *WorkerDomainService) Update(ctx context.Context, accountID interface{}, body WorkerDomainUpdateParams, opts ...option.RequestOption) (res *WorkerDomainUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDomainUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("accounts/%v/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all Worker Domains for an account.
func (r *WorkerDomainService) List(ctx context.Context, accountID interface{}, query WorkerDomainListParams, opts ...option.RequestOption) (res *[]WorkerDomainListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env WorkerDomainListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("accounts/%v/workers/domains", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
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

type WorkerDomainUpdateResponse struct {
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
	ZoneName string                         `json:"zone_name"`
	JSON     workerDomainUpdateResponseJSON `json:"-"`
}

// workerDomainUpdateResponseJSON contains the JSON metadata for the struct
// [WorkerDomainUpdateResponse]
type workerDomainUpdateResponseJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainListResponse struct {
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
	ZoneName string                       `json:"zone_name"`
	JSON     workerDomainListResponseJSON `json:"-"`
}

// workerDomainListResponseJSON contains the JSON metadata for the struct
// [WorkerDomainListResponse]
type workerDomainListResponseJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
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

type WorkerDomainUpdateParams struct {
	// Worker environment associated with the zone and hostname.
	Environment param.Field[string] `json:"environment,required"`
	// Hostname of the Worker Domain.
	Hostname param.Field[string] `json:"hostname,required"`
	// Worker service associated with the zone and hostname.
	Service param.Field[string] `json:"service,required"`
	// Identifier of the zone.
	ZoneID param.Field[interface{}] `json:"zone_id,required"`
}

func (r WorkerDomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WorkerDomainUpdateResponseEnvelope struct {
	Errors   []WorkerDomainUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDomainUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkerDomainUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDomainUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDomainUpdateResponseEnvelopeJSON    `json:"-"`
}

// workerDomainUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerDomainUpdateResponseEnvelope]
type workerDomainUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainUpdateResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerDomainUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDomainUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerDomainUpdateResponseEnvelopeErrors]
type workerDomainUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainUpdateResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    workerDomainUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDomainUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [WorkerDomainUpdateResponseEnvelopeMessages]
type workerDomainUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainUpdateResponseEnvelopeSuccess bool

const (
	WorkerDomainUpdateResponseEnvelopeSuccessTrue WorkerDomainUpdateResponseEnvelopeSuccess = true
)

type WorkerDomainListParams struct {
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

// URLQuery serializes [WorkerDomainListParams]'s query parameters as `url.Values`.
func (r WorkerDomainListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkerDomainListResponseEnvelope struct {
	Errors   []WorkerDomainListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WorkerDomainListResponseEnvelopeMessages `json:"messages,required"`
	Result   []WorkerDomainListResponse                 `json:"result,required"`
	// Whether the API call was successful
	Success WorkerDomainListResponseEnvelopeSuccess `json:"success,required"`
	JSON    workerDomainListResponseEnvelopeJSON    `json:"-"`
}

// workerDomainListResponseEnvelopeJSON contains the JSON metadata for the struct
// [WorkerDomainListResponseEnvelope]
type workerDomainListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainListResponseEnvelopeErrors struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    workerDomainListResponseEnvelopeErrorsJSON `json:"-"`
}

// workerDomainListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [WorkerDomainListResponseEnvelopeErrors]
type workerDomainListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainListResponseEnvelopeMessages struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    workerDomainListResponseEnvelopeMessagesJSON `json:"-"`
}

// workerDomainListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [WorkerDomainListResponseEnvelopeMessages]
type workerDomainListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainListResponseEnvelopeSuccess bool

const (
	WorkerDomainListResponseEnvelopeSuccessTrue WorkerDomainListResponseEnvelopeSuccess = true
)

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
