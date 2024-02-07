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
	path := fmt.Sprintf("accounts/%v/workers/domains/%v", accountID, domainID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
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
	Errors   []WorkerDomainGetResponseError   `json:"errors"`
	Messages []WorkerDomainGetResponseMessage `json:"messages"`
	Result   WorkerDomainGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WorkerDomainGetResponseSuccess `json:"success"`
	JSON    workerDomainGetResponseJSON    `json:"-"`
}

// workerDomainGetResponseJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponse]
type workerDomainGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseError struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    workerDomainGetResponseErrorJSON `json:"-"`
}

// workerDomainGetResponseErrorJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponseError]
type workerDomainGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseMessage struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    workerDomainGetResponseMessageJSON `json:"-"`
}

// workerDomainGetResponseMessageJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponseMessage]
type workerDomainGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WorkerDomainGetResponseResult struct {
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
	ZoneName string                            `json:"zone_name"`
	JSON     workerDomainGetResponseResultJSON `json:"-"`
}

// workerDomainGetResponseResultJSON contains the JSON metadata for the struct
// [WorkerDomainGetResponseResult]
type workerDomainGetResponseResultJSON struct {
	ID          apijson.Field
	Environment apijson.Field
	Hostname    apijson.Field
	Service     apijson.Field
	ZoneID      apijson.Field
	ZoneName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkerDomainGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WorkerDomainGetResponseSuccess bool

const (
	WorkerDomainGetResponseSuccessTrue WorkerDomainGetResponseSuccess = true
)
