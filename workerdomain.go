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

// Detaches a Worker from a zone and hostname.
func (r *WorkerDomainService) Delete(ctx context.Context, accountID interface{}, domainID interface{}, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
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
